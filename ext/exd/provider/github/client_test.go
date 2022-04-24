package github_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	tMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/odpf/optimus/ext/exd"
	"github.com/odpf/optimus/ext/exd/provider/github"
	"github.com/odpf/optimus/mock"
)

type ClientTestSuite struct {
	suite.Suite
}

func (c *ClientTestSuite) TestDownload() {
	var ctx = context.Background()
	var httpDoer = &mock.HTTPDoer{}
	client, err := github.NewClient(ctx, httpDoer)
	if err != nil {
		panic(err)
	}

	c.Run("should return nil and error if metadata is nil", func() {
		var metadata *exd.Metadata

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if metadata provider is not recognized", func() {
		metadata := &exd.Metadata{
			ProviderName: "unrecognized",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if error when creating request to API path", func() {
		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: ":invalid-url",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if encountered error when doing request", func() {
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(nil, errors.New("random error"))

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if encountered error when decoding response", func() {
		response := &http.Response{
			Body: io.NopCloser(strings.NewReader("invalid-body")),
		}
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(response, nil)

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if cannot find asset with the specified suffix", func() {
		release := github.RepositoryRelease{}
		marshalled, _ := json.Marshal(release)
		response := &http.Response{
			Body: io.NopCloser(bytes.NewReader(marshalled)),
		}
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(response, nil)

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if error when creating request to download url", func() {
		release := github.RepositoryRelease{
			Assets: []*github.ReleaseAsset{
				{
					Name:               "asset" + runtime.GOOS + "-" + runtime.GOARCH,
					BrowserDownloadURL: ":invalid-url",
				},
			},
		}
		marshalled, _ := json.Marshal(release)
		response := &http.Response{
			Body: io.NopCloser(bytes.NewReader(marshalled)),
		}
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(response, nil)

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if error when sending download request", func() {
		release := github.RepositoryRelease{
			Assets: []*github.ReleaseAsset{
				{
					Name:               "asset" + runtime.GOOS + "-" + runtime.GOARCH,
					BrowserDownloadURL: "http://github.com/odpf/optimus",
				},
			},
		}
		marshalled, _ := json.Marshal(release)
		response := &http.Response{
			Body: io.NopCloser(bytes.NewReader(marshalled)),
		}
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(response, nil).Once()
		httpDoer.On("Do", tMock.Anything).Return(nil, errors.New("random error")).Once()

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if error when decoding response", func() {
		marshalled := []byte("unknown message")
		response := &http.Response{
			Body: io.NopCloser(bytes.NewReader(marshalled)),
		}
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(response, nil).Once()
		httpDoer.On("Do", tMock.Anything).Return(response, nil).Once()

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return nil and error if the specified release is either draft or pre-release", func() {
		release := github.RepositoryRelease{
			Draft:      true,
			Prerelease: true,
			Assets: []*github.ReleaseAsset{
				{
					Name:               "asset" + runtime.GOOS + "-" + runtime.GOARCH,
					BrowserDownloadURL: "http://github.com/odpf/optimus",
				},
			},
		}
		marshalled, _ := json.Marshal(release)
		response := &http.Response{
			Body: io.NopCloser(bytes.NewReader(marshalled)),
		}
		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(response, nil).Once()
		httpDoer.On("Do", tMock.Anything).Return(response, nil).Once()

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.Nil(actualAsset)
		c.Error(actualErr)
	})

	c.Run("should return bytes and nil if no error is encountered", func() {
		release := github.RepositoryRelease{
			Assets: []*github.ReleaseAsset{
				{
					Name:               "asset" + runtime.GOOS + "-" + runtime.GOARCH,
					BrowserDownloadURL: "http://github.com/odpf/optimus",
				},
			},
		}
		marshalled, _ := json.Marshal(release)
		releaseResponse := &http.Response{
			Body: io.NopCloser(bytes.NewReader(marshalled)),
		}
		downloadResponse := &http.Response{
			Body: io.NopCloser(strings.NewReader("random payload")),
		}

		httpDoer := &mock.HTTPDoer{}
		httpDoer.On("Do", tMock.Anything).Return(releaseResponse, nil).Once()
		httpDoer.On("Do", tMock.Anything).Return(downloadResponse, nil).Once()

		client, err := github.NewClient(ctx, httpDoer)
		if err != nil {
			panic(err)
		}

		metadata := &exd.Metadata{
			ProviderName: "github",
			AssetAPIPath: "http://github.com/odpf/optimus",
		}

		actualAsset, actualErr := client.Download(metadata)

		c.NotNil(actualAsset)
		c.NoError(actualErr)
	})
}

func TestNewClient(t *testing.T) {
	t.Run("should return nil and error if context is nil", func(t *testing.T) {
		var ctx context.Context
		httpDoer := &mock.HTTPDoer{}

		actualGithub, actualErr := github.NewClient(ctx, httpDoer)

		assert.Nil(t, actualGithub)
		assert.Error(t, actualErr)
	})

	t.Run("should return nil and error if http doer is nil", func(t *testing.T) {
		ctx := context.Background()
		var httpDoer exd.HTTPDoer

		actualGithub, actualErr := github.NewClient(ctx, httpDoer)

		assert.Nil(t, actualGithub)
		assert.Error(t, actualErr)
	})

	t.Run("should return github and nil if no error encountered", func(t *testing.T) {
		ctx := context.Background()
		httpDoer := &mock.HTTPDoer{}

		actualGithub, actualErr := github.NewClient(ctx, httpDoer)

		assert.NotNil(t, actualGithub)
		assert.NoError(t, actualErr)
	})
}

func TestGithub(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}