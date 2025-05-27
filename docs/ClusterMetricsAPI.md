# \ClusterMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleClusterMetricsRequest**](ClusterMetricsAPI.md#HandleClusterMetricsRequest) | **Get** /cluster/metrics | 



## HandleClusterMetricsRequest

> ClusterStatusDto HandleClusterMetricsRequest(ctx, organization, cluster).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterMetricsAPI.HandleClusterMetricsRequest(context.Background(), organization, cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterMetricsAPI.HandleClusterMetricsRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterMetricsRequest`: ClusterStatusDto
	fmt.Fprintf(os.Stdout, "Response from `ClusterMetricsAPI.HandleClusterMetricsRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**cluster** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterMetricsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterStatusDto**](ClusterStatusDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

