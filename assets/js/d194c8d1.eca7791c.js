"use strict";(self.webpackChunkoptimus=self.webpackChunkoptimus||[]).push([[8462],{3905:function(e,t,n){n.d(t,{Zo:function(){return l},kt:function(){return f}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var c=r.createContext({}),u=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},l=function(e){var t=u(e.components);return r.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,c=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),d=u(n),f=o,m=d["".concat(c,".").concat(f)]||d[f]||p[f]||a;return n?r.createElement(m,s(s({ref:t},l),{},{components:n})):r.createElement(m,s({ref:t},l))}));function f(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,s=new Array(a);s[0]=d;var i={};for(var c in t)hasOwnProperty.call(t,c)&&(i[c]=t[c]);i.originalType=e,i.mdxType="string"==typeof e?e:o,s[1]=i;for(var u=2;u<a;u++)s[u]=n[u];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},1647:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return i},contentTitle:function(){return c},metadata:function(){return u},toc:function(){return l},default:function(){return d}});var r=n(7462),o=n(3366),a=(n(7294),n(3905)),s=["components"],i={},c=void 0,u={unversionedId:"rfcs/support_for_depending_on_external_sensors",id:"rfcs/support_for_depending_on_external_sensors",isDocsHomePage:!1,title:"support_for_depending_on_external_sensors",description:"- Feature Name: Support For Depndening on External Sources",source:"@site/docs/rfcs/20220123_support_for_depending_on_external_sensors.md",sourceDirName:"rfcs",slug:"/rfcs/support_for_depending_on_external_sensors",permalink:"/optimus/docs/rfcs/support_for_depending_on_external_sensors",editUrl:"https://github.com/odpf/optimus/edit/master/docs/docs/rfcs/20220123_support_for_depending_on_external_sensors.md",tags:[],version:"current",lastUpdatedBy:"Dery Rahman Ahaddienata",lastUpdatedAt:1649667084,formattedLastUpdatedAt:"4/11/2022",sidebarPosition:20220123,frontMatter:{}},l=[{value:"GCS Sensor",id:"gcs-sensor",children:[]}],p={toc:l};function d(e){var t=e.components,n=(0,o.Z)(e,s);return(0,a.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Feature Name: Support For Depndening on External Sources"),(0,a.kt)("li",{parentName:"ul"},"Status: Draft"),(0,a.kt)("li",{parentName:"ul"},"Start Date: 2022-01-23"),(0,a.kt)("li",{parentName:"ul"},"Authors: ")),(0,a.kt)("h1",{id:"summary"},"Summary"),(0,a.kt)("p",null,"Optimus supports job dependencies, but there is a need for optimus jobs to depend on external sources which are not managed by the optimus server. For example, depending the BQ or GCS data availability or data being managed by another optimus server. Whatever data sources optimus is managing lets have sensors for basic data availability check, in GCS checking for file exists & in BQ taking a select query & returning success when rowcount > 0. For other requirements let's have a http sensor."),(0,a.kt)("h1",{id:"technical-design"},"Technical Design"),(0,a.kt)("p",null,"Optimus can add support for all the sensors as libraries, which will be evaulated within the execution envrionment of the user, all variables will be returned for a given scheduled date through the api call which will be used by the actual sensor execution. "),(0,a.kt)("p",null,"Optimus provides libraries needed for the above operations which can be used in the respective execution environment of the scheduler, currently the library will be offered in python."),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"/intance")," api call can accept params to filter what to return to just reduce the unnecessary payload & return only the needed variables, as sensors execute a lot."),(0,a.kt)("h4",{id:"http-sensor"},(0,a.kt)("strong",{parentName:"h4"},"Http Sensor")),(0,a.kt)("p",null,"If the call returns 200 then the sensor succeeds"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"dependencies : \n type : http\n endpoint : url\n headers :\n body :\n  \n")),(0,a.kt)("h4",{id:"bq-sensor"},"BQ Sensor"),(0,a.kt)("p",null,"If the query results in rows then the sensor succeeds"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"dependencies : \n type : bq\n query : \n service_account :\n  \n")),(0,a.kt)("h3",{id:"gcs-sensor"},"GCS Sensor"),(0,a.kt)("p",null,"If the path exists then the sensor succeeds"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"dependencies : \n type : gcs\n path : \n service_account :  \n")))}d.isMDXComponent=!0}}]);