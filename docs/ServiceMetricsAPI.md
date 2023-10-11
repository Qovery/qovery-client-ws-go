# \ServiceMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleMetricsRequest**](ServiceMetricsAPI.md#HandleMetricsRequest) | **Get** /service/metrics | 



## HandleMetricsRequest

> ServiceMetricsDto HandleMetricsRequest(ctx, organization, cluster, project, environment, service, serviceType).Execute()



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
    serviceType := openapiclient.ServiceType("APPLICATION") // ServiceType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceMetricsAPI.HandleMetricsRequest(context.Background(), organization, cluster, project, environment, service, serviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceMetricsAPI.HandleMetricsRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleMetricsRequest`: ServiceMetricsDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceMetricsAPI.HandleMetricsRequest`: %v\n", resp)
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
**serviceType** | [**ServiceType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleMetricsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







### Return type

[**ServiceMetricsDto**](ServiceMetricsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

