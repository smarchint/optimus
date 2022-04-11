"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[2082],{3905:function(e,n,t){t.d(n,{Zo:function(){return c},kt:function(){return d}});var a=t(7294);function r(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,a)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){r(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,a,r=function(e,n){if(null==e)return{};var t,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)t=i[a],n.indexOf(t)>=0||(r[t]=e[t]);return r}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)t=i[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(r[t]=e[t])}return r}var p=a.createContext({}),l=function(e){var n=a.useContext(p),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},c=function(e){var n=l(e.components);return a.createElement(p.Provider,{value:n},e.children)},u={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},m=a.forwardRef((function(e,n){var t=e.components,r=e.mdxType,i=e.originalType,p=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),m=l(t),d=r,f=m["".concat(p,".").concat(d)]||m[d]||u[d]||i;return t?a.createElement(f,o(o({ref:n},c),{},{components:t})):a.createElement(f,o({ref:n},c))}));function d(e,n){var t=arguments,r=n&&n.mdxType;if("string"==typeof e||r){var i=t.length,o=new Array(i);o[0]=m;var s={};for(var p in n)hasOwnProperty.call(n,p)&&(s[p]=n[p]);s.originalType=e,s.mdxType="string"==typeof e?e:r,o[1]=s;for(var l=2;l<i;l++)o[l]=t[l];return a.createElement.apply(null,o)}return a.createElement.apply(null,t)}m.displayName="MDXCreateElement"},9807:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return s},contentTitle:function(){return p},metadata:function(){return l},toc:function(){return c},default:function(){return m}});var a=t(7462),r=t(3366),i=(t(7294),t(3905)),o=["components"],s={id:"organising-specifications",title:"Organising specifications"},p=void 0,l={unversionedId:"guides/organising-specifications",id:"guides/organising-specifications",isDocsHomePage:!1,title:"Organising specifications",description:"Optimus supports two ways to deploy specifications",source:"@site/docs/guides/organising-specifcations.md",sourceDirName:"guides",slug:"/guides/organising-specifications",permalink:"/optimus/docs/guides/organising-specifications",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/guides/organising-specifcations.md",tags:[],version:"current",lastUpdatedBy:"Dery Rahman Ahaddienata",lastUpdatedAt:1649667084,formattedLastUpdatedAt:"4/11/2022",frontMatter:{id:"organising-specifications",title:"Organising specifications"},sidebar:"docsSidebar",previous:{title:"Create bigquery external table",permalink:"/optimus/docs/guides/create-bigquery-external-table"},next:{title:"Starting Optimus Server",permalink:"/optimus/docs/guides/optimus-serve"}},c=[],u={toc:c};function m(e){var n=e.components,t=(0,r.Z)(e,o);return(0,i.kt)("wrapper",(0,a.Z)({},u,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"Optimus supports two ways to deploy specifications"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"REST/GRPC"),(0,i.kt)("li",{parentName:"ul"},"Optimus CLI deploy command")),(0,i.kt)("p",null,"When using Optimus cli to deploy, either manually or from a CI pipeline, it is advised to use version control system like git. Here is a simple directory structure that can be used as a template for jobs and datastore resources."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},".\n\u251c\u2500\u2500 optimus.yaml\n\u251c\u2500\u2500 README.md\n\u251c\u2500\u2500 namespace-1\n\u2502\xa0\xa0 \u251c\u2500\u2500 jobs\n|   \u2502\xa0\xa0 \u251c\u2500\u2500 job1\n|   \u2502\xa0\xa0 \u251c\u2500\u2500 job2\n|   \u2502\xa0\xa0 \u2514\u2500\u2500 this.yaml\n\u2502\xa0\xa0 \u2514\u2500\u2500 resources\n|    \xa0\xa0 \u251c\u2500\u2500 bigquery\n\u2502\xa0\xa0  \xa0\xa0 \u2502   \u251c\u2500\u2500 table1\n\u2502\xa0\xa0  \xa0\xa0 \u2502\xa0\xa0 \u251c\u2500\u2500 table2\n|    \xa0\xa0 |   \u2514\u2500\u2500 this.yaml\n\u2502\xa0\xa0     \u2514\u2500\u2500 postgres\n\u2502\xa0\xa0         \u2514\u2500\u2500 table1\n\u2514\u2500\u2500 namespace-2\n \xa0\xa0 \u2514\u2500\u2500 jobs\n \xa0\xa0 \u2514\u2500\u2500 resources\n")),(0,i.kt)("p",null,"A sample ",(0,i.kt)("inlineCode",{parentName:"p"},"optimus.yaml")," would look like"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"version: 1\nhost: localhost:9100\nproject:\n  name: project-1\nnamespaces:\n- name: namespace-1\n  config: {}\n  job:\n    path: namespace-1/jobs\n  datastore:\n  - type: bigquery\n    path: namespace-1/resources/bigquery\n  - type: postgres\n    path: namespace-1/resources/postgres\n- name: namespace-2\n  config: {}\n  job:\n    path: namespaces-2/jobs\n  datastore:\n  - type: bigquery\n    path: namespace-2/resources/bigquery\n")),(0,i.kt)("p",null,"You might have also noticed there are ",(0,i.kt)("inlineCode",{parentName:"p"},"this.yaml")," files being used in some directories. This file is used to share a single set of configuration across multiple sub directories. For example if I create a file at ",(0,i.kt)("inlineCode",{parentName:"p"},"/namespace-1/jobs/this.yaml"),", then all sub directories inside ",(0,i.kt)("inlineCode",{parentName:"p"},"/namespaces-1/jobs")," will inherit this config as defaults. If same config is specified in sub directory, then sub directory will override the parent defaults. For example a ",(0,i.kt)("inlineCode",{parentName:"p"},"this.yaml")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"/namespace-1/jobs")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"version: 1\nschedule:\n  interval: @daily\nbehavior:\n  depends_on_past: false\n  catch_up: true\n  retry:\n    count: 1\n    delay: 5s\nlabels:\n  owner: overlords\n  transform: sql\n")),(0,i.kt)("p",null,"and a ",(0,i.kt)("inlineCode",{parentName:"p"},"job.yaml")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"/namespace-1/jobs/job1")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},'name: sample_replace\nowner: optimus@example.io\nschedule:\n  start_date: "2020-09-25"\n  interval: 0 10 * * *\nbehavior:\n  depends_on_past: true\ntask:\n  name: bq2bq\n  config:\n    project: project_name\n    dataset: project_dataset\n    table: sample_replace\n    load_method: REPLACE\n  window:\n    size: 48h\n    offset: 24h\nlabels:\n  process: bq\n')),(0,i.kt)("p",null,"will result in final computed ",(0,i.kt)("inlineCode",{parentName:"p"},"job.yaml")," during deployment as"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 1\nname: sample_replace\nowner: optimus@example.io\nschedule:\n  start_date: "2020-10-06"\n  interval: 0 10 * * *\nbehavior:\n  depends_on_past: true\n  catch_up: true\n  retry:\n    count: 1\n    delay: 5s\ntask:\n  name: bq2bq\n  config:\n    project: project_name\n    dataset: project_dataset\n    table: sample_replace\n    load_method: REPLACE\n  window:\n    size: 48h\n    offset: 24h\nlabels:\n  process: bq\n  owner: overlords\n  transform: sql\n')))}m.isMDXComponent=!0}}]);