# Go API client for qovery-ws

Describe the weboscket endpoints

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 0.1.0
- Generator version: 7.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import qovery-ws "github.com/qovery/qovery-client-ws-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `qovery-ws.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), qovery-ws.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `qovery-ws.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), qovery-ws.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `qovery-ws.ContextOperationServerIndices` and `qovery-ws.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), qovery-ws.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), qovery-ws.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClusterStatusAPI* | [**HandleClusterStatusRequest**](docs/ClusterStatusAPI.md#handleclusterstatusrequest) | **Get** /cluster/status | 
*DeploymentAPI* | [**HandleDeploymentLogsRequest**](docs/DeploymentAPI.md#handledeploymentlogsrequest) | **Get** /deployment/logs | 
*DeploymentAPI* | [**HandleDeploymentStatusRequest**](docs/DeploymentAPI.md#handledeploymentstatusrequest) | **Get** /deployment/status | 
*LogsAPI* | [**HandleInfraLogsRequest**](docs/LogsAPI.md#handleinfralogsrequest) | **Get** /infra/logs | 
*LogsAPI* | [**HandleServiceLogsRequest**](docs/LogsAPI.md#handleservicelogsrequest) | **Get** /service/logs | 
*ServiceListPodsAPI* | [**HandleServiceListPodsRequest**](docs/ServiceListPodsAPI.md#handleservicelistpodsrequest) | **Get** /service/pods | 
*ServiceMetricsAPI* | [**HandleMetricsRequest**](docs/ServiceMetricsAPI.md#handlemetricsrequest) | **Get** /service/metrics | 
*ServiceStatusAPI* | [**HandleServiceStatusRequest**](docs/ServiceStatusAPI.md#handleservicestatusrequest) | **Get** /service/status | 
*ShellAPI* | [**HandleShellExec**](docs/ShellAPI.md#handleshellexec) | **Get** /shell/exec | 


## Documentation For Models

 - [ApplicationStatusDto](docs/ApplicationStatusDto.md)
 - [CertificateStatusDto](docs/CertificateStatusDto.md)
 - [ClusterNodeDto](docs/ClusterNodeDto.md)
 - [ClusterStatusDto](docs/ClusterStatusDto.md)
 - [ContainerStateDto](docs/ContainerStateDto.md)
 - [ContainerStateTerminatedDto](docs/ContainerStateTerminatedDto.md)
 - [ContainerStatusDto](docs/ContainerStatusDto.md)
 - [DatabaseStatusDto](docs/DatabaseStatusDto.md)
 - [EnvironmentStatusDto](docs/EnvironmentStatusDto.md)
 - [MetricDto](docs/MetricDto.md)
 - [NodeAddressDto](docs/NodeAddressDto.md)
 - [NodeConditionDto](docs/NodeConditionDto.md)
 - [NodePodInfoDto](docs/NodePodInfoDto.md)
 - [NodeResourceDto](docs/NodeResourceDto.md)
 - [NodeTaintDto](docs/NodeTaintDto.md)
 - [PodDto](docs/PodDto.md)
 - [PodStatusDto](docs/PodStatusDto.md)
 - [ResourceStatusDto](docs/ResourceStatusDto.md)
 - [ServiceInfraLogResponseDto](docs/ServiceInfraLogResponseDto.md)
 - [ServiceListPodsResponseDto](docs/ServiceListPodsResponseDto.md)
 - [ServiceLogResponseDto](docs/ServiceLogResponseDto.md)
 - [ServiceMetricsDto](docs/ServiceMetricsDto.md)
 - [ServiceStateDto](docs/ServiceStateDto.md)
 - [ServiceStatusDto](docs/ServiceStatusDto.md)
 - [ServiceType](docs/ServiceType.md)
 - [UnitDto](docs/UnitDto.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

erebe@erebe.eu

