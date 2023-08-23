# \DeploymentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleDeploymentLogsRequest**](DeploymentApi.md#HandleDeploymentLogsRequest) | **Get** /deployment/logs | 
[**HandleDeploymentStatusRequest**](DeploymentApi.md#HandleDeploymentStatusRequest) | **Get** /deployment/status | 



## HandleDeploymentLogsRequest

> string HandleDeploymentLogsRequest(ctx, organization, cluster, project, environment, version).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    environment := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.HandleDeploymentLogsRequest(context.Background(), organization, cluster, project, environment, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.HandleDeploymentLogsRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleDeploymentLogsRequest`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.HandleDeploymentLogsRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**cluster** | **string** |  | 
**project** | **string** |  | 
**environment** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleDeploymentLogsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleDeploymentStatusRequest

> string HandleDeploymentStatusRequest(ctx, organization, cluster, project, environment, version).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    environment := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.HandleDeploymentStatusRequest(context.Background(), organization, cluster, project, environment, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.HandleDeploymentStatusRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleDeploymentStatusRequest`: string
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.HandleDeploymentStatusRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**cluster** | **string** |  | 
**project** | **string** |  | 
**environment** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleDeploymentStatusRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

