# PodInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilliLimit** | Pointer to **NullableInt32** |  | [optional] 
**CpuMilliRequest** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **int64** |  | 
**ErrorContainerStatuses** | [**[]NodePodErrorStatusDto**](NodePodErrorStatusDto.md) |  | 
**ImagesVersion** | **map[string]string** |  | 
**MemoryMibLimit** | Pointer to **NullableInt32** |  | [optional] 
**MemoryMibRequest** | Pointer to **NullableInt32** |  | [optional] 
**MetricsUsage** | [**MetricsUsageDto**](MetricsUsageDto.md) |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**RestartCount** | **int32** |  | 

## Methods

### NewPodInfoDto

`func NewPodInfoDto(createdAt int64, errorContainerStatuses []NodePodErrorStatusDto, imagesVersion map[string]string, metricsUsage MetricsUsageDto, name string, namespace string, restartCount int32, ) *PodInfoDto`

NewPodInfoDto instantiates a new PodInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodInfoDtoWithDefaults

`func NewPodInfoDtoWithDefaults() *PodInfoDto`

NewPodInfoDtoWithDefaults instantiates a new PodInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilliLimit

`func (o *PodInfoDto) GetCpuMilliLimit() int32`

GetCpuMilliLimit returns the CpuMilliLimit field if non-nil, zero value otherwise.

### GetCpuMilliLimitOk

`func (o *PodInfoDto) GetCpuMilliLimitOk() (*int32, bool)`

GetCpuMilliLimitOk returns a tuple with the CpuMilliLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliLimit

`func (o *PodInfoDto) SetCpuMilliLimit(v int32)`

SetCpuMilliLimit sets CpuMilliLimit field to given value.

### HasCpuMilliLimit

`func (o *PodInfoDto) HasCpuMilliLimit() bool`

HasCpuMilliLimit returns a boolean if a field has been set.

### SetCpuMilliLimitNil

`func (o *PodInfoDto) SetCpuMilliLimitNil(b bool)`

 SetCpuMilliLimitNil sets the value for CpuMilliLimit to be an explicit nil

### UnsetCpuMilliLimit
`func (o *PodInfoDto) UnsetCpuMilliLimit()`

UnsetCpuMilliLimit ensures that no value is present for CpuMilliLimit, not even an explicit nil
### GetCpuMilliRequest

`func (o *PodInfoDto) GetCpuMilliRequest() int32`

GetCpuMilliRequest returns the CpuMilliRequest field if non-nil, zero value otherwise.

### GetCpuMilliRequestOk

`func (o *PodInfoDto) GetCpuMilliRequestOk() (*int32, bool)`

GetCpuMilliRequestOk returns a tuple with the CpuMilliRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliRequest

`func (o *PodInfoDto) SetCpuMilliRequest(v int32)`

SetCpuMilliRequest sets CpuMilliRequest field to given value.

### HasCpuMilliRequest

`func (o *PodInfoDto) HasCpuMilliRequest() bool`

HasCpuMilliRequest returns a boolean if a field has been set.

### SetCpuMilliRequestNil

`func (o *PodInfoDto) SetCpuMilliRequestNil(b bool)`

 SetCpuMilliRequestNil sets the value for CpuMilliRequest to be an explicit nil

### UnsetCpuMilliRequest
`func (o *PodInfoDto) UnsetCpuMilliRequest()`

UnsetCpuMilliRequest ensures that no value is present for CpuMilliRequest, not even an explicit nil
### GetCreatedAt

`func (o *PodInfoDto) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PodInfoDto) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PodInfoDto) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorContainerStatuses

`func (o *PodInfoDto) GetErrorContainerStatuses() []NodePodErrorStatusDto`

GetErrorContainerStatuses returns the ErrorContainerStatuses field if non-nil, zero value otherwise.

### GetErrorContainerStatusesOk

`func (o *PodInfoDto) GetErrorContainerStatusesOk() (*[]NodePodErrorStatusDto, bool)`

GetErrorContainerStatusesOk returns a tuple with the ErrorContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorContainerStatuses

`func (o *PodInfoDto) SetErrorContainerStatuses(v []NodePodErrorStatusDto)`

SetErrorContainerStatuses sets ErrorContainerStatuses field to given value.


### GetImagesVersion

`func (o *PodInfoDto) GetImagesVersion() map[string]string`

GetImagesVersion returns the ImagesVersion field if non-nil, zero value otherwise.

### GetImagesVersionOk

`func (o *PodInfoDto) GetImagesVersionOk() (*map[string]string, bool)`

GetImagesVersionOk returns a tuple with the ImagesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesVersion

`func (o *PodInfoDto) SetImagesVersion(v map[string]string)`

SetImagesVersion sets ImagesVersion field to given value.


### GetMemoryMibLimit

`func (o *PodInfoDto) GetMemoryMibLimit() int32`

GetMemoryMibLimit returns the MemoryMibLimit field if non-nil, zero value otherwise.

### GetMemoryMibLimitOk

`func (o *PodInfoDto) GetMemoryMibLimitOk() (*int32, bool)`

GetMemoryMibLimitOk returns a tuple with the MemoryMibLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibLimit

`func (o *PodInfoDto) SetMemoryMibLimit(v int32)`

SetMemoryMibLimit sets MemoryMibLimit field to given value.

### HasMemoryMibLimit

`func (o *PodInfoDto) HasMemoryMibLimit() bool`

HasMemoryMibLimit returns a boolean if a field has been set.

### SetMemoryMibLimitNil

`func (o *PodInfoDto) SetMemoryMibLimitNil(b bool)`

 SetMemoryMibLimitNil sets the value for MemoryMibLimit to be an explicit nil

### UnsetMemoryMibLimit
`func (o *PodInfoDto) UnsetMemoryMibLimit()`

UnsetMemoryMibLimit ensures that no value is present for MemoryMibLimit, not even an explicit nil
### GetMemoryMibRequest

`func (o *PodInfoDto) GetMemoryMibRequest() int32`

GetMemoryMibRequest returns the MemoryMibRequest field if non-nil, zero value otherwise.

### GetMemoryMibRequestOk

`func (o *PodInfoDto) GetMemoryMibRequestOk() (*int32, bool)`

GetMemoryMibRequestOk returns a tuple with the MemoryMibRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibRequest

`func (o *PodInfoDto) SetMemoryMibRequest(v int32)`

SetMemoryMibRequest sets MemoryMibRequest field to given value.

### HasMemoryMibRequest

`func (o *PodInfoDto) HasMemoryMibRequest() bool`

HasMemoryMibRequest returns a boolean if a field has been set.

### SetMemoryMibRequestNil

`func (o *PodInfoDto) SetMemoryMibRequestNil(b bool)`

 SetMemoryMibRequestNil sets the value for MemoryMibRequest to be an explicit nil

### UnsetMemoryMibRequest
`func (o *PodInfoDto) UnsetMemoryMibRequest()`

UnsetMemoryMibRequest ensures that no value is present for MemoryMibRequest, not even an explicit nil
### GetMetricsUsage

`func (o *PodInfoDto) GetMetricsUsage() MetricsUsageDto`

GetMetricsUsage returns the MetricsUsage field if non-nil, zero value otherwise.

### GetMetricsUsageOk

`func (o *PodInfoDto) GetMetricsUsageOk() (*MetricsUsageDto, bool)`

GetMetricsUsageOk returns a tuple with the MetricsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUsage

`func (o *PodInfoDto) SetMetricsUsage(v MetricsUsageDto)`

SetMetricsUsage sets MetricsUsage field to given value.


### GetName

`func (o *PodInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodInfoDto) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *PodInfoDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PodInfoDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PodInfoDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetRestartCount

`func (o *PodInfoDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *PodInfoDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *PodInfoDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


