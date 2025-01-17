# \ClusterListNodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleClusterListNodesRequest**](ClusterListNodesAPI.md#HandleClusterListNodesRequest) | **Get** /cluster/node | 



## HandleClusterListNodesRequest

> ClusterListNodesResponseDto HandleClusterListNodesRequest(ctx, organization, cluster).Execute()



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
	resp, r, err := apiClient.ClusterListNodesAPI.HandleClusterListNodesRequest(context.Background(), organization, cluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterListNodesAPI.HandleClusterListNodesRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleClusterListNodesRequest`: ClusterListNodesResponseDto
	fmt.Fprintf(os.Stdout, "Response from `ClusterListNodesAPI.HandleClusterListNodesRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**cluster** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleClusterListNodesRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterListNodesResponseDto**](ClusterListNodesResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

