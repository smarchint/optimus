"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[223],{3905:function(e,t,a){a.d(t,{Zo:function(){return p},kt:function(){return c}});var n=a(7294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function l(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function i(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?l(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):l(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function o(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},l=Object.keys(e);for(n=0;n<l.length;n++)a=l[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)a=l[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var d=n.createContext({}),s=function(e){var t=n.useContext(d),a=t;return e&&(a="function"==typeof e?e(t):i(i({},t),e)),a},p=function(e){var t=s(e.components);return n.createElement(d.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,l=e.originalType,d=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),m=s(a),c=r,k=m["".concat(d,".").concat(c)]||m[c]||u[c]||l;return a?n.createElement(k,i(i({ref:t},p),{},{components:a})):n.createElement(k,i({ref:t},p))}));function c(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=a.length,i=new Array(l);i[0]=m;var o={};for(var d in t)hasOwnProperty.call(t,d)&&(o[d]=t[d]);o.originalType=e,o.mdxType="string"==typeof e?e:r,i[1]=o;for(var s=2;s<l;s++)i[s]=a[s];return n.createElement.apply(null,i)}return n.createElement.apply(null,a)}m.displayName="MDXCreateElement"},1937:function(e,t,a){a.r(t),a.d(t,{frontMatter:function(){return o},contentTitle:function(){return d},metadata:function(){return s},toc:function(){return p},default:function(){return m}});var n=a(7462),r=a(3366),l=(a(7294),a(3905)),i=["components"],o={id:"task-bq2bq",title:"Bigquery to bigquery transformation task"},d=void 0,s={unversionedId:"guides/task-bq2bq",id:"guides/task-bq2bq",isDocsHomePage:!1,title:"Bigquery to bigquery transformation task",description:"Creating Task",source:"@site/docs/guides/task-bq2bq.md",sourceDirName:"guides",slug:"/guides/task-bq2bq",permalink:"/optimus/docs/guides/task-bq2bq",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/guides/task-bq2bq.md",tags:[],version:"current",lastUpdatedBy:"Anwar Hidayat",lastUpdatedAt:1650264863,formattedLastUpdatedAt:"4/18/2022",frontMatter:{id:"task-bq2bq",title:"Bigquery to bigquery transformation task"},sidebar:"docsSidebar",previous:{title:"Starting Optimus Server",permalink:"/optimus/docs/guides/optimus-serve"},next:{title:"Backup Resources",permalink:"/optimus/docs/guides/backup"}},p=[{value:"Creating Task",id:"creating-task",children:[]},{value:"Load Method",id:"load-method",children:[]},{value:"query.sql file",id:"querysql-file",children:[{value:"SQL macros",id:"sql-macros",children:[]}]},{value:"SQL Helpers",id:"sql-helpers",children:[]}],u={toc:p};function m(e){var t=e.components,a=(0,r.Z)(e,i);return(0,l.kt)("wrapper",(0,n.Z)({},u,a,{components:t,mdxType:"MDXLayout"}),(0,l.kt)("h3",{id:"creating-task"},"Creating Task"),(0,l.kt)("p",null,"Command to create a task :"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre"},"optimus job create\n")),(0,l.kt)("p",null,"This command will invoke an interactive cli that contains configurations that\nneed to be filled for the task. The tasks files will be generated at\n",(0,l.kt)("inlineCode",{parentName:"p"},"{PWD}/jobs/{JOB_NAME}/assets")," folder. "),(0,l.kt)("p",null,"Inside the assets folder there could be several files, one that is\nneeded to configure this task is :"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"query.sql - file that contains the transformation query")),(0,l.kt)("p",null,"This will also configure the ",(0,l.kt)("inlineCode",{parentName:"p"},"job.yaml")," with few defaults and few inputs requested at the time\nof creation. User still able to change the config values after the file is generated."),(0,l.kt)("p",null,"For example ",(0,l.kt)("inlineCode",{parentName:"p"},"job.yaml")," config :"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 1\nname: example_job\nowner: example@example.com\nschedule:\n  start_date: "2021-02-18"\n  interval: 0 3 * * *\nbehavior:\n  depends_on_past: false\n  catch_up: true\ntask:\n  name: bq2bq\n  config:\n    DATASET: data\n    LOAD_METHOD: APPEND\n    PROJECT: example\n    SQL_TYPE: STANDARD\n    TABLE: hello_table\n  window:\n    size: 24h\n    offset: "0"\n    truncate_to: d\n')),(0,l.kt)("p",null,"Here are the details of each configuration and the allowed values :"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Config Name"),(0,l.kt)("th",{parentName:"tr",align:null},"Description"),(0,l.kt)("th",{parentName:"tr",align:null},"Values"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"PROJECT")),(0,l.kt)("td",{parentName:"tr",align:null},"google cloud platform project id of the destination bigquery table"),(0,l.kt)("td",{parentName:"tr",align:null},"...")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"DATASET")),(0,l.kt)("td",{parentName:"tr",align:null},"bigquery dataset name of the destination table"),(0,l.kt)("td",{parentName:"tr",align:null},"...")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"TABLE")),(0,l.kt)("td",{parentName:"tr",align:null},"the table name of the destination table"),(0,l.kt)("td",{parentName:"tr",align:null},"...")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"LOAD_METHOD")),(0,l.kt)("td",{parentName:"tr",align:null},"method to load data to the destination tables"),(0,l.kt)("td",{parentName:"tr",align:null},"APPEND, REPLACE, MERGE")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},(0,l.kt)("inlineCode",{parentName:"td"},"PARTITION_FILTER")),(0,l.kt)("td",{parentName:"tr",align:null},"Used to identify target partitions to replace in a REPLACE query. This can be left empty and optimus will figure the target partitions automatically but its cheaper and faster to specify the condition. This filter will be used as a where clause in a merge statement to delete the partitions from the destination table."),(0,l.kt)("td",{parentName:"tr",align:null},'event_timestamp >= "{{.DSTART}}" AND event_timestamp < "{{.DEND}}"')))),(0,l.kt)("h3",{id:"load-method"},"Load Method"),(0,l.kt)("p",null,"The way data loaded to destination table depends on the partition configuration of the destination tables"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Load Method"),(0,l.kt)("th",{parentName:"tr",align:null},"No Partition"),(0,l.kt)("th",{parentName:"tr",align:null},"Partitioned Table"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"APPEND"),(0,l.kt)("td",{parentName:"tr",align:null},"Append new records to destination table"),(0,l.kt)("td",{parentName:"tr",align:null},"Append new records to destination table per partition based on localised start_time")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"REPLACE"),(0,l.kt)("td",{parentName:"tr",align:null},"Truncate/Clean the table before insert new records"),(0,l.kt)("td",{parentName:"tr",align:null},"Clean records in destination partition before insert new record to new partition")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"MERGE"),(0,l.kt)("td",{parentName:"tr",align:null},"Load the data using DML Merge statement, all of the load logic lies on DML merge statement"),(0,l.kt)("td",{parentName:"tr",align:null},"Load the data using DML Merge statement, all of the load logic lies on DML merge statement")))),(0,l.kt)("h2",{id:"querysql-file"},"query.sql file"),(0,l.kt)("p",null,"The ",(0,l.kt)("em",{parentName:"p"},"query.sql")," file contains transformation logic"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sql"},"select count(1) as count, date(created_time) as dt\nfrom `project.dataset.tablename`\nwhere date(created_time) >= '{{.DSTART|Date}}' and date(booking_creation_time) < '{{.DEND|Date}}'\ngroup by dt\n")),(0,l.kt)("h3",{id:"sql-macros"},"SQL macros"),(0,l.kt)("p",null,"Macros is special variables in SQL that will be replaced by actual values when transformation executed"),(0,l.kt)("p",null,"There are several SQL macros available"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},"{{.DSTART}} - start date/datetime of the window as ",(0,l.kt)("inlineCode",{parentName:"li"},"2021-02-10T10:00:00+00:00"),"\nthat is, RFC3339"),(0,l.kt)("li",{parentName:"ul"},"{{.DEND}} - end date/datetime of the window, as RFC3339"),(0,l.kt)("li",{parentName:"ul"},"{{.JOB_DESTINATION}} - full qualified table name used in DML statement"),(0,l.kt)("li",{parentName:"ul"},"{{.EXECUTION_TIME}} - full qualified table name used in DML statement")),(0,l.kt)("p",null,"The value of ",(0,l.kt)("inlineCode",{parentName:"p"},"DSTART")," and ",(0,l.kt)("inlineCode",{parentName:"p"},"DEND")," depends on ",(0,l.kt)("inlineCode",{parentName:"p"},"window")," config in ",(0,l.kt)("inlineCode",{parentName:"p"},"job.yaml"),". This is very similar to Optimus v1"),(0,l.kt)("table",null,(0,l.kt)("thead",{parentName:"table"},(0,l.kt)("tr",{parentName:"thead"},(0,l.kt)("th",{parentName:"tr",align:null},"Window config"),(0,l.kt)("th",{parentName:"tr",align:null},"DSTART"),(0,l.kt)("th",{parentName:"tr",align:null},"DEND"))),(0,l.kt)("tbody",{parentName:"table"},(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"size:24h, offset:0, truncate_to:d"),(0,l.kt)("td",{parentName:"tr",align:null},"The current date taken from input, for example 2019-01-01"),(0,l.kt)("td",{parentName:"tr",align:null},"The next day after DSTART date 2019-01-02")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"size:168h, offset:0, truncate_to:w"),(0,l.kt)("td",{parentName:"tr",align:null},"Start of the week date for example : 2019-04-01"),(0,l.kt)("td",{parentName:"tr",align:null},"End date of the week , for example : 2019-04-07")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"size:1M, offset:0, truncate_to:M"),(0,l.kt)("td",{parentName:"tr",align:null},"Start of the month date, example : 2019-01-01"),(0,l.kt)("td",{parentName:"tr",align:null},"End date of the month, for example : 2019-01-31")),(0,l.kt)("tr",{parentName:"tbody"},(0,l.kt)("td",{parentName:"tr",align:null},"size:2h, offset:0, truncate_to:h"),(0,l.kt)("td",{parentName:"tr",align:null},"Datetime of the start of the hour, for example 2019-01-01 01:00:00"),(0,l.kt)("td",{parentName:"tr",align:null},"Datetime the start of the next hour, for example 2019-01-01 02:00:00")))),(0,l.kt)("p",null,"Please find more details under ",(0,l.kt)("a",{parentName:"p",href:"/optimus/docs/concepts/intervals-and-windows"},"concepts")," section."),(0,l.kt)("p",null,"Macros in SQL transformation example :"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sql"},"select count(1) as count, date(created_time) as dt\nfrom `project.dataset.tablename`\nwhere date(created_time) >= '{{.DSTART|Date}}' and date(booking_creation_time) < '{{.DEND|Date}}'\ngroup by dt\n")),(0,l.kt)("p",null,"Rendered SQL for DAILY window example :"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sql"},"select count(1) as count, date(created_time) as dt\nfrom `project.dataset.tablename`\nwhere date(created_time) >= '2019-01-01' and date(booking_creation_time) < '2019-01-02'\ngroup by dt\n")),(0,l.kt)("p",null,"Rendered SQL for HOURLY window example :\nthe value of ",(0,l.kt)("inlineCode",{parentName:"p"},"DSTART")," and ",(0,l.kt)("inlineCode",{parentName:"p"},"DEND")," is YYYY-mm-dd HH:MM:SS formatted datetime "),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sql"},"select count(1) as count, date(created_time) as dt\nfrom `project.dataset.tablename`\nwhere date(created_time) >= '2019-01-01 06:00:00' and date(booking_creation_time) < '2019-01-01 07:00:00'\ngroup by dt\n")),(0,l.kt)("p",null,"destination_table macros example :"),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sql"},"MERGE `{{.JOB_DESTINATION}}` S\nusing\n(\nselect count(1) as count, date(created_time) as dt\nfrom `project.dataset.tablename`\nwhere date(created_time) >= '{{.DSTART}}' and date(created_time) < '{{.DEND}}'\ngroup by dt\n) N\non S.date = N.date\nWHEN MATCHED then\nUPDATE SET `count` = N.count\nwhen not matched then\nINSERT (`date`, `count`) VALUES(N.date, N.count)\n")),(0,l.kt)("h2",{id:"sql-helpers"},"SQL Helpers"),(0,l.kt)("p",null,"Sometimes default behaviour of how tasks are being understood by optimus is not ideal. You can change this using helpers inside the query.sql file. To use, simply add them inside sql multiline comments where it\u2019s required.\nAt the moment there is only one sql helper:"),(0,l.kt)("ul",null,(0,l.kt)("li",{parentName:"ul"},(0,l.kt)("inlineCode",{parentName:"li"},"@ignoreupstream"),": By default, Optimus adds all the external tables used inside the query file as its upstream\ndependency. This helper can help ignore unwanted waits for upstream dependency to finish before the current transformation can be executed.\nHelper needs to be added just before the external table name. For example:")),(0,l.kt)("pre",null,(0,l.kt)("code",{parentName:"pre",className:"language-sql"},"select\nhakai,\nrasengan,\n`over`,\nload_timestamp as `event_timestamp`\nfrom /* @ignoreupstream */\n`g-project.playground.sample_select`\nWHERE\nDATE(`load_timestamp`) >= DATE('{{.DSTART}}')\nAND DATE(`load_timestamp`) < DATE('{{.DEND}}')\n")))}m.isMDXComponent=!0}}]);