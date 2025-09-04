# ServiceLogResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** |  | 
**CreatedAt** | **int64** |  | 
**Labels** | **map[string]string** |  | 
**Message** | **string** |  | 
**PodName** | **string** |  | 
**SeverityText** | **int32** |  | 
**Version** | **string** |  | 

## Methods

### NewServiceLogResponseDto

`func NewServiceLogResponseDto(containerName string, createdAt int64, labels map[string]string, message string, podName string, severityText int32, version string, ) *ServiceLogResponseDto`

NewServiceLogResponseDto instantiates a new ServiceLogResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLogResponseDtoWithDefaults

`func NewServiceLogResponseDtoWithDefaults() *ServiceLogResponseDto`

NewServiceLogResponseDtoWithDefaults instantiates a new ServiceLogResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *ServiceLogResponseDto) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ServiceLogResponseDto) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ServiceLogResponseDto) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetCreatedAt

`func (o *ServiceLogResponseDto) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceLogResponseDto) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceLogResponseDto) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetLabels

`func (o *ServiceLogResponseDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServiceLogResponseDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServiceLogResponseDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetMessage

`func (o *ServiceLogResponseDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceLogResponseDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceLogResponseDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPodName

`func (o *ServiceLogResponseDto) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ServiceLogResponseDto) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ServiceLogResponseDto) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetSeverityText

`func (o *ServiceLogResponseDto) GetSeverityText() int32`

GetSeverityText returns the SeverityText field if non-nil, zero value otherwise.

### GetSeverityTextOk

`func (o *ServiceLogResponseDto) GetSeverityTextOk() (*int32, bool)`

GetSeverityTextOk returns a tuple with the SeverityText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityText

`func (o *ServiceLogResponseDto) SetSeverityText(v int32)`

SetSeverityText sets SeverityText field to given value.


### GetVersion

`func (o *ServiceLogResponseDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceLogResponseDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceLogResponseDto) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


