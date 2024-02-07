# \ServiceStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleServiceStatusRequest**](ServiceStatusAPI.md#HandleServiceStatusRequest) | **Get** /service/status | 



## HandleServiceStatusRequest

> ServiceStatusDto HandleServiceStatusRequest(ctx, organization, cluster, project, environment).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceStatusAPI.HandleServiceStatusRequest(context.Background(), organization, cluster, project, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceStatusAPI.HandleServiceStatusRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleServiceStatusRequest`: ServiceStatusDto
	fmt.Fprintf(os.Stdout, "Response from `ServiceStatusAPI.HandleServiceStatusRequest`: %v\n", resp)
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

### Other Parameters

Other parameters are passed through a pointer to a apiHandleServiceStatusRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ServiceStatusDto**](ServiceStatusDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

