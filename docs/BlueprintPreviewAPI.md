# \BlueprintPreviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleBlueprintPreviewRequest**](BlueprintPreviewAPI.md#HandleBlueprintPreviewRequest) | **Get** /blueprint/preview | 



## HandleBlueprintPreviewRequest

> BlueprintPreviewResult HandleBlueprintPreviewRequest(ctx, organization, cluster, previewId).Execute()



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
	organization := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Organization id — for permission scoping. Frontend already has it.
	cluster := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster id — for permission scoping. Frontend already has it.
	previewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Preview id — the UUIDv7 q-core returned from `POST /api/blueprint/{id}/update/preview`. Doubles as the Pub/Sub channel suffix: `core.blueprint.preview.{preview_id}`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintPreviewAPI.HandleBlueprintPreviewRequest(context.Background(), organization, cluster, previewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintPreviewAPI.HandleBlueprintPreviewRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleBlueprintPreviewRequest`: BlueprintPreviewResult
	fmt.Fprintf(os.Stdout, "Response from `BlueprintPreviewAPI.HandleBlueprintPreviewRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization id — for permission scoping. Frontend already has it. | 
**cluster** | **string** | Cluster id — for permission scoping. Frontend already has it. | 
**previewId** | **string** | Preview id — the UUIDv7 q-core returned from &#x60;POST /api/blueprint/{id}/update/preview&#x60;. Doubles as the Pub/Sub channel suffix: &#x60;core.blueprint.preview.{preview_id}&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleBlueprintPreviewRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BlueprintPreviewResult**](BlueprintPreviewResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

