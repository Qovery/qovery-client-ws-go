# \LogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleInfraLogsRequest**](LogsAPI.md#HandleInfraLogsRequest) | **Get** /infra/logs | 
[**HandleServiceLogsRequest**](LogsAPI.md#HandleServiceLogsRequest) | **Get** /service/logs | 



## HandleInfraLogsRequest

> ServiceInfraLogResponseDto HandleInfraLogsRequest(ctx, organization, cluster, project, environment, service, infraComponentType).Execute()



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
	infraComponentType := "infraComponentType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.HandleInfraLogsRequest(context.Background(), organization, cluster, project, environment, service, infraComponentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.HandleInfraLogsRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleInfraLogsRequest`: ServiceInfraLogResponseDto
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.HandleInfraLogsRequest`: %v\n", resp)
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
**infraComponentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleInfraLogsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







### Return type

[**ServiceInfraLogResponseDto**](ServiceInfraLogResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleServiceLogsRequest

> ServiceLogResponseDto HandleServiceLogsRequest(ctx, organization, cluster, project, environment, service, podName, deploymentId, query, start).Execute()



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
	deploymentId := "deploymentId_example" // string | 
	query := "query_example" // string | 
	start := "start_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.HandleServiceLogsRequest(context.Background(), organization, cluster, project, environment, service, podName, deploymentId, query, start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.HandleServiceLogsRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleServiceLogsRequest`: ServiceLogResponseDto
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.HandleServiceLogsRequest`: %v\n", resp)
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
**deploymentId** | **string** |  | 
**query** | **string** |  | 
**start** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleServiceLogsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------










### Return type

[**ServiceLogResponseDto**](ServiceLogResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

