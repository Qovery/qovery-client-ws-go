# \ShellAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleShellExec**](ShellAPI.md#HandleShellExec) | **Get** /shell/exec | 



## HandleShellExec

> string HandleShellExec(ctx, organization, cluster, project, environment, service, podName, containerName, command, ttyWidth, ttyHeight).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/qovery/qovery-client-ws-go"
)

func main() {
    organization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    environment := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    service := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    podName := "podName_example" // string | 
    containerName := "containerName_example" // string | 
    command := []string{"Inner_example"} // []string | 
    ttyWidth := int32(56) // int32 | 
    ttyHeight := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShellAPI.HandleShellExec(context.Background(), organization, cluster, project, environment, service, podName, containerName, command, ttyWidth, ttyHeight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShellAPI.HandleShellExec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleShellExec`: string
    fmt.Fprintf(os.Stdout, "Response from `ShellAPI.HandleShellExec`: %v\n", resp)
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
**service** | **string** |  | 
**podName** | **string** |  | 
**containerName** | **string** |  | 
**command** | [**[]string**](string.md) |  | 
**ttyWidth** | **int32** |  | 
**ttyHeight** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleShellExecRequest struct via the builder pattern


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

