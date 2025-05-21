# QoveryComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDetail** | [**[]QoveryComponentInFailure**](QoveryComponentInFailure.md) |  | 
**ImagesVersion** | **[]string** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Status** | [**QoveryComponentStatus**](QoveryComponentStatus.md) |  | 

## Methods

### NewQoveryComponentDto

`func NewQoveryComponentDto(errorDetail []QoveryComponentInFailure, imagesVersion []string, name string, namespace string, status QoveryComponentStatus, ) *QoveryComponentDto`

NewQoveryComponentDto instantiates a new QoveryComponentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoveryComponentDtoWithDefaults

`func NewQoveryComponentDtoWithDefaults() *QoveryComponentDto`

NewQoveryComponentDtoWithDefaults instantiates a new QoveryComponentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDetail

`func (o *QoveryComponentDto) GetErrorDetail() []QoveryComponentInFailure`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *QoveryComponentDto) GetErrorDetailOk() (*[]QoveryComponentInFailure, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *QoveryComponentDto) SetErrorDetail(v []QoveryComponentInFailure)`

SetErrorDetail sets ErrorDetail field to given value.


### GetImagesVersion

`func (o *QoveryComponentDto) GetImagesVersion() []string`

GetImagesVersion returns the ImagesVersion field if non-nil, zero value otherwise.

### GetImagesVersionOk

`func (o *QoveryComponentDto) GetImagesVersionOk() (*[]string, bool)`

GetImagesVersionOk returns a tuple with the ImagesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesVersion

`func (o *QoveryComponentDto) SetImagesVersion(v []string)`

SetImagesVersion sets ImagesVersion field to given value.


### GetName

`func (o *QoveryComponentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QoveryComponentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QoveryComponentDto) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *QoveryComponentDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *QoveryComponentDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *QoveryComponentDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetStatus

`func (o *QoveryComponentDto) GetStatus() QoveryComponentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QoveryComponentDto) GetStatusOk() (*QoveryComponentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QoveryComponentDto) SetStatus(v QoveryComponentStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


