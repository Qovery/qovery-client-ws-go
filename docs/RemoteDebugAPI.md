# \RemoteDebugAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleShellRemoteDebug**](RemoteDebugAPI.md#HandleShellRemoteDebug) | **Get** /shell/debug | 



## HandleShellRemoteDebug

> string HandleShellRemoteDebug(ctx, organization, cluster, flavor, ttyWidth, ttyHeight, nodeSelector).Execute()



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
	flavor := openapiclient.DebugFlavor("REGULAR_PRIVILEGE") // DebugFlavor | 
	ttyWidth := int32(56) // int32 | 
	ttyHeight := int32(56) // int32 | 
	nodeSelector := "nodeSelector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteDebugAPI.HandleShellRemoteDebug(context.Background(), organization, cluster, flavor, ttyWidth, ttyHeight, nodeSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteDebugAPI.HandleShellRemoteDebug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleShellRemoteDebug`: string
	fmt.Fprintf(os.Stdout, "Response from `RemoteDebugAPI.HandleShellRemoteDebug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**cluster** | **string** |  | 
**flavor** | [**DebugFlavor**](.md) |  | 
**ttyWidth** | **int32** |  | 
**ttyHeight** | **int32** |  | 
**nodeSelector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleShellRemoteDebugRequest struct via the builder pattern


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

