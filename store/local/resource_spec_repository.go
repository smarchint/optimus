package local

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"github.com/odpf/optimus/core/fs"
	"github.com/odpf/optimus/models"
)

const (
	ResourceSpecFileName = "resource.yaml"
)

type Resource struct {
	Version int
	Name    string
	Type    string
	Spec    interface{}
	Labels  map[string]string
}

type resourceRepository struct {
	fs    fs.FileSystem
	cache struct {
		dirty bool

		// cache is mapped with spec name -> resource
		data map[string]models.ResourceSpec
	}

	ds models.Datastorer
}

func (repo *resourceRepository) Save(resourceSpec models.ResourceSpec) error {
	if len(resourceSpec.Name) == 0 || len(resourceSpec.Type) == 0 {
		return fmt.Errorf("resource is missing required fields")
	}

	typeController, _ := resourceSpec.Datastore.Types()[resourceSpec.Type]
	specBytes, err := typeController.Adapter().ToYaml(resourceSpec)
	if err != nil {
		return err
	}

	// save assets
	for assetName, assetValue := range resourceSpec.Assets {
		assetFd, err := repo.fs.OpenForWrite(repo.assetFilePath(resourceSpec.Name, assetName))
		if err != nil {
			return err
		}
		_, err = assetFd.Write([]byte(assetValue))
		if err != nil {
			return err
		}
		assetFd.Close()
	}

	// save resource
	fileName := repo.resourceFilePath(resourceSpec.Name)
	fd, err := repo.fs.OpenForWrite(fileName)
	if err != nil {
		return err
	}

	if _, err := fd.Write(specBytes); err != nil {
		return err
	}
	fd.Close()

	repo.cache.dirty = true
	return nil
}

// GetAll finds all the resources recursively in current and sub directory
func (repo *resourceRepository) GetAll() ([]models.ResourceSpec, error) {
	resourceSpecs := []models.ResourceSpec{}
	if repo.cache.dirty {
		if err := repo.refreshCache(); err != nil {
			return resourceSpecs, err
		}
	}

	for _, j := range repo.cache.data {
		resourceSpecs = append(resourceSpecs, j)
	}
	return resourceSpecs, nil
}

// GetByName returns a job requested by the name
func (repo *resourceRepository) GetByName(jobName string) (models.ResourceSpec, error) {
	resourceSpec := models.ResourceSpec{}
	if strings.TrimSpace(jobName) == "" {
		return resourceSpec, errors.Errorf("resource name cannot be an empty string")
	}

	// refresh local cache if needed
	if repo.cache.dirty {
		if err := repo.refreshCache(); err != nil {
			return resourceSpec, err
		}
	}

	// check if available in cache
	resourceSpec, ok := repo.cache.data[jobName]
	if !ok {
		return resourceSpec, models.ErrNoSuchSpec
	}
	return resourceSpec, nil
}

// Delete deletes a requested job by name
func (repo *resourceRepository) Delete(jobName string) error {
	panic("unimplemented")
}

func (repo *resourceRepository) refreshCache() error {
	repo.cache.dirty = true
	repo.cache.data = make(map[string]models.ResourceSpec)

	resourceSpecs, err := repo.scanDirs(".")
	if err != nil && err != fs.ErrNoSuchFile {
		return err
	}
	if len(resourceSpecs) < 1 {
		return models.ErrNoResources
	}

	repo.cache.dirty = false
	return nil
}

func (repo *resourceRepository) findInDir(dirName string) (models.ResourceSpec, error) {
	resourceSpec := models.ResourceSpec{}
	if strings.TrimSpace(dirName) == "" {
		return resourceSpec, fmt.Errorf("dir name cannot be an empty string")
	}

	resourceFD, err := repo.fs.Open(repo.resourceFilePath(dirName))
	if err != nil {
		if err == fs.ErrNoSuchFile {
			err = models.ErrNoSuchSpec
		}
		return resourceSpec, err
	}

	// need to parse type of the resource before it can proceed and pass on to datastore
	resourceBytes, err := ioutil.ReadAll(resourceFD)
	if err != nil {
		return models.ResourceSpec{}, err
	}
	resourceFD.Close()

	var rawResource Resource
	if err := yaml.Unmarshal(resourceBytes, &rawResource); err != nil {
		return resourceSpec, errors.Wrapf(err, "error parsing resource spec in %s", dirName)
	}
	typeController, ok := repo.ds.Types()[models.ResourceType(rawResource.Type)]
	if !ok {
		return models.ResourceSpec{}, fmt.Errorf("unsupported type %s for datastore %s", rawResource.Type, repo.ds.Name())
	}

	// convert to internal model
	resourceSpec, err = typeController.Adapter().FromYaml(resourceBytes)
	if err != nil {
		return resourceSpec, errors.Wrapf(err, "failed to read spec in: %s", dirName)
	}

	assets := map[string]string{}
	assetFolderFD, err := repo.fs.Open(repo.assetFolderPath(dirName))
	if err == nil {
		fileNames, err := assetFolderFD.Readdirnames(-1)
		if err != nil {
			return resourceSpec, err
		}
		assetFolderFD.Close()

		for _, fileName := range fileNames {
			// don't include base resource file as asset
			if fileName == ResourceSpecFileName {
				continue
			}

			assetFd, err := repo.fs.Open(repo.assetFilePath(dirName, fileName))
			if err != nil {
				return resourceSpec, err
			}

			if isDir, err := assetFd.IsDir(); err == nil && isDir {
				assetFd.Close()
				continue
			} else if err != nil {
				return models.ResourceSpec{}, err
			}

			fileContent, err := ioutil.ReadAll(assetFd)
			if err != nil {
				return resourceSpec, err
			}
			assets[fileName] = string(fileContent)
			assetFd.Close()
		}
	}
	resourceSpec.Assets = assets

	if _, ok := repo.cache.data[resourceSpec.Name]; ok {
		return resourceSpec, errors.Errorf("job name should be unique across directories: %s", resourceSpec.Name)
	}
	repo.cache.data[resourceSpec.Name] = resourceSpec
	return resourceSpec, nil
}

func (repo *resourceRepository) scanDirs(path string) ([]models.ResourceSpec, error) {
	specs := []models.ResourceSpec{}

	// filter folders & scan recursively
	folders, err := repo.getDirs(path)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		s, err := repo.scanDirs(filepath.Join(path, folder))
		if err != nil && err != fs.ErrNoSuchFile {
			return s, err
		}
		specs = append(specs, s...)
	}

	// find resources in this folder
	spec, err := repo.findInDir(path)
	if err != nil {
		if err != fs.ErrNoSuchFile && err != models.ErrNoSuchSpec {
			return nil, err
		}
	} else {
		specs = append(specs, spec)
	}
	return specs, nil
}

// getDirs return names of all the folders in provided path
func (repo *resourceRepository) getDirs(dirPath string) ([]string, error) {
	currentDir, err := repo.fs.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames, err := currentDir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	currentDir.Close()

	folderPath := []string{}
	for _, fileName := range fileNames {
		if strings.HasPrefix(fileName, ".") {
			continue
		}
		if specSuffixRegex.FindString(fileName) != "" || fileName == AssetFolderName {
			continue
		}

		fd, err := repo.fs.Open(filepath.Join(dirPath, fileName))
		if err != nil {
			return nil, err
		}
		if isDir, err := fd.IsDir(); err == nil && !isDir {
			fd.Close()
			continue
		} else if err != nil {
			fd.Close()
			return nil, err
		}
		fd.Close()

		folderPath = append(folderPath, fileName)
	}
	return folderPath, nil
}

// resourceFilePath generates the filename for a given job
func (repo *resourceRepository) resourceFilePath(name string) string {
	return filepath.Join(name, ResourceSpecFileName)
}

// assetFolderPath generates the path to asset directory folder
func (repo *resourceRepository) assetFolderPath(name string) string {
	return name
}

// assetFilePath generates the path to asset directory files
func (repo *resourceRepository) assetFilePath(job string, file string) string {
	return filepath.Join(repo.assetFolderPath(job), file)
}

func NewResourceSpecRepository(fs fs.FileSystem, ds models.Datastorer) *resourceRepository {
	repo := new(resourceRepository)
	repo.fs = fs
	repo.cache.dirty = true
	repo.ds = ds
	return repo
}
