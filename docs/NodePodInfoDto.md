# NodePodInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilliLimit** | Pointer to **NullableInt32** |  | [optional] 
**CpuMilliRequest** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **int64** |  | 
**EphemeralStorageMibLimit** | Pointer to **NullableInt32** |  | [optional] 
**EphemeralStorageMibRequest** | Pointer to **NullableInt32** |  | [optional] 
**ErrorContainerStatuses** | [**[]NodePodErrorStatusDto**](NodePodErrorStatusDto.md) |  | 
**ImagesVersion** | **map[string]string** |  | 
**MemoryMibLimit** | Pointer to **NullableInt32** |  | [optional] 
**MemoryMibRequest** | Pointer to **NullableInt32** |  | [optional] 
**MetricsUsage** | [**MetricsUsageDto**](MetricsUsageDto.md) |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**QoveryServiceInfo** | Pointer to [**NullablePodQoveryServiceInfoDto**](PodQoveryServiceInfoDto.md) |  | [optional] 
**RestartCount** | **int32** |  | 

## Methods

### NewNodePodInfoDto

`func NewNodePodInfoDto(createdAt int64, errorContainerStatuses []NodePodErrorStatusDto, imagesVersion map[string]string, metricsUsage MetricsUsageDto, name string, namespace string, restartCount int32, ) *NodePodInfoDto`

NewNodePodInfoDto instantiates a new NodePodInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePodInfoDtoWithDefaults

`func NewNodePodInfoDtoWithDefaults() *NodePodInfoDto`

NewNodePodInfoDtoWithDefaults instantiates a new NodePodInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilliLimit

`func (o *NodePodInfoDto) GetCpuMilliLimit() int32`

GetCpuMilliLimit returns the CpuMilliLimit field if non-nil, zero value otherwise.

### GetCpuMilliLimitOk

`func (o *NodePodInfoDto) GetCpuMilliLimitOk() (*int32, bool)`

GetCpuMilliLimitOk returns a tuple with the CpuMilliLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliLimit

`func (o *NodePodInfoDto) SetCpuMilliLimit(v int32)`

SetCpuMilliLimit sets CpuMilliLimit field to given value.

### HasCpuMilliLimit

`func (o *NodePodInfoDto) HasCpuMilliLimit() bool`

HasCpuMilliLimit returns a boolean if a field has been set.

### SetCpuMilliLimitNil

`func (o *NodePodInfoDto) SetCpuMilliLimitNil(b bool)`

 SetCpuMilliLimitNil sets the value for CpuMilliLimit to be an explicit nil

### UnsetCpuMilliLimit
`func (o *NodePodInfoDto) UnsetCpuMilliLimit()`

UnsetCpuMilliLimit ensures that no value is present for CpuMilliLimit, not even an explicit nil
### GetCpuMilliRequest

`func (o *NodePodInfoDto) GetCpuMilliRequest() int32`

GetCpuMilliRequest returns the CpuMilliRequest field if non-nil, zero value otherwise.

### GetCpuMilliRequestOk

`func (o *NodePodInfoDto) GetCpuMilliRequestOk() (*int32, bool)`

GetCpuMilliRequestOk returns a tuple with the CpuMilliRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliRequest

`func (o *NodePodInfoDto) SetCpuMilliRequest(v int32)`

SetCpuMilliRequest sets CpuMilliRequest field to given value.

### HasCpuMilliRequest

`func (o *NodePodInfoDto) HasCpuMilliRequest() bool`

HasCpuMilliRequest returns a boolean if a field has been set.

### SetCpuMilliRequestNil

`func (o *NodePodInfoDto) SetCpuMilliRequestNil(b bool)`

 SetCpuMilliRequestNil sets the value for CpuMilliRequest to be an explicit nil

### UnsetCpuMilliRequest
`func (o *NodePodInfoDto) UnsetCpuMilliRequest()`

UnsetCpuMilliRequest ensures that no value is present for CpuMilliRequest, not even an explicit nil
### GetCreatedAt

`func (o *NodePodInfoDto) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodePodInfoDto) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodePodInfoDto) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetEphemeralStorageMibLimit

`func (o *NodePodInfoDto) GetEphemeralStorageMibLimit() int32`

GetEphemeralStorageMibLimit returns the EphemeralStorageMibLimit field if non-nil, zero value otherwise.

### GetEphemeralStorageMibLimitOk

`func (o *NodePodInfoDto) GetEphemeralStorageMibLimitOk() (*int32, bool)`

GetEphemeralStorageMibLimitOk returns a tuple with the EphemeralStorageMibLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageMibLimit

`func (o *NodePodInfoDto) SetEphemeralStorageMibLimit(v int32)`

SetEphemeralStorageMibLimit sets EphemeralStorageMibLimit field to given value.

### HasEphemeralStorageMibLimit

`func (o *NodePodInfoDto) HasEphemeralStorageMibLimit() bool`

HasEphemeralStorageMibLimit returns a boolean if a field has been set.

### SetEphemeralStorageMibLimitNil

`func (o *NodePodInfoDto) SetEphemeralStorageMibLimitNil(b bool)`

 SetEphemeralStorageMibLimitNil sets the value for EphemeralStorageMibLimit to be an explicit nil

### UnsetEphemeralStorageMibLimit
`func (o *NodePodInfoDto) UnsetEphemeralStorageMibLimit()`

UnsetEphemeralStorageMibLimit ensures that no value is present for EphemeralStorageMibLimit, not even an explicit nil
### GetEphemeralStorageMibRequest

`func (o *NodePodInfoDto) GetEphemeralStorageMibRequest() int32`

GetEphemeralStorageMibRequest returns the EphemeralStorageMibRequest field if non-nil, zero value otherwise.

### GetEphemeralStorageMibRequestOk

`func (o *NodePodInfoDto) GetEphemeralStorageMibRequestOk() (*int32, bool)`

GetEphemeralStorageMibRequestOk returns a tuple with the EphemeralStorageMibRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageMibRequest

`func (o *NodePodInfoDto) SetEphemeralStorageMibRequest(v int32)`

SetEphemeralStorageMibRequest sets EphemeralStorageMibRequest field to given value.

### HasEphemeralStorageMibRequest

`func (o *NodePodInfoDto) HasEphemeralStorageMibRequest() bool`

HasEphemeralStorageMibRequest returns a boolean if a field has been set.

### SetEphemeralStorageMibRequestNil

`func (o *NodePodInfoDto) SetEphemeralStorageMibRequestNil(b bool)`

 SetEphemeralStorageMibRequestNil sets the value for EphemeralStorageMibRequest to be an explicit nil

### UnsetEphemeralStorageMibRequest
`func (o *NodePodInfoDto) UnsetEphemeralStorageMibRequest()`

UnsetEphemeralStorageMibRequest ensures that no value is present for EphemeralStorageMibRequest, not even an explicit nil
### GetErrorContainerStatuses

`func (o *NodePodInfoDto) GetErrorContainerStatuses() []NodePodErrorStatusDto`

GetErrorContainerStatuses returns the ErrorContainerStatuses field if non-nil, zero value otherwise.

### GetErrorContainerStatusesOk

`func (o *NodePodInfoDto) GetErrorContainerStatusesOk() (*[]NodePodErrorStatusDto, bool)`

GetErrorContainerStatusesOk returns a tuple with the ErrorContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorContainerStatuses

`func (o *NodePodInfoDto) SetErrorContainerStatuses(v []NodePodErrorStatusDto)`

SetErrorContainerStatuses sets ErrorContainerStatuses field to given value.


### GetImagesVersion

`func (o *NodePodInfoDto) GetImagesVersion() map[string]string`

GetImagesVersion returns the ImagesVersion field if non-nil, zero value otherwise.

### GetImagesVersionOk

`func (o *NodePodInfoDto) GetImagesVersionOk() (*map[string]string, bool)`

GetImagesVersionOk returns a tuple with the ImagesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesVersion

`func (o *NodePodInfoDto) SetImagesVersion(v map[string]string)`

SetImagesVersion sets ImagesVersion field to given value.


### GetMemoryMibLimit

`func (o *NodePodInfoDto) GetMemoryMibLimit() int32`

GetMemoryMibLimit returns the MemoryMibLimit field if non-nil, zero value otherwise.

### GetMemoryMibLimitOk

`func (o *NodePodInfoDto) GetMemoryMibLimitOk() (*int32, bool)`

GetMemoryMibLimitOk returns a tuple with the MemoryMibLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibLimit

`func (o *NodePodInfoDto) SetMemoryMibLimit(v int32)`

SetMemoryMibLimit sets MemoryMibLimit field to given value.

### HasMemoryMibLimit

`func (o *NodePodInfoDto) HasMemoryMibLimit() bool`

HasMemoryMibLimit returns a boolean if a field has been set.

### SetMemoryMibLimitNil

`func (o *NodePodInfoDto) SetMemoryMibLimitNil(b bool)`

 SetMemoryMibLimitNil sets the value for MemoryMibLimit to be an explicit nil

### UnsetMemoryMibLimit
`func (o *NodePodInfoDto) UnsetMemoryMibLimit()`

UnsetMemoryMibLimit ensures that no value is present for MemoryMibLimit, not even an explicit nil
### GetMemoryMibRequest

`func (o *NodePodInfoDto) GetMemoryMibRequest() int32`

GetMemoryMibRequest returns the MemoryMibRequest field if non-nil, zero value otherwise.

### GetMemoryMibRequestOk

`func (o *NodePodInfoDto) GetMemoryMibRequestOk() (*int32, bool)`

GetMemoryMibRequestOk returns a tuple with the MemoryMibRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibRequest

`func (o *NodePodInfoDto) SetMemoryMibRequest(v int32)`

SetMemoryMibRequest sets MemoryMibRequest field to given value.

### HasMemoryMibRequest

`func (o *NodePodInfoDto) HasMemoryMibRequest() bool`

HasMemoryMibRequest returns a boolean if a field has been set.

### SetMemoryMibRequestNil

`func (o *NodePodInfoDto) SetMemoryMibRequestNil(b bool)`

 SetMemoryMibRequestNil sets the value for MemoryMibRequest to be an explicit nil

### UnsetMemoryMibRequest
`func (o *NodePodInfoDto) UnsetMemoryMibRequest()`

UnsetMemoryMibRequest ensures that no value is present for MemoryMibRequest, not even an explicit nil
### GetMetricsUsage

`func (o *NodePodInfoDto) GetMetricsUsage() MetricsUsageDto`

GetMetricsUsage returns the MetricsUsage field if non-nil, zero value otherwise.

### GetMetricsUsageOk

`func (o *NodePodInfoDto) GetMetricsUsageOk() (*MetricsUsageDto, bool)`

GetMetricsUsageOk returns a tuple with the MetricsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUsage

`func (o *NodePodInfoDto) SetMetricsUsage(v MetricsUsageDto)`

SetMetricsUsage sets MetricsUsage field to given value.


### GetName

`func (o *NodePodInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePodInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePodInfoDto) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *NodePodInfoDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NodePodInfoDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NodePodInfoDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetQoveryServiceInfo

`func (o *NodePodInfoDto) GetQoveryServiceInfo() PodQoveryServiceInfoDto`

GetQoveryServiceInfo returns the QoveryServiceInfo field if non-nil, zero value otherwise.

### GetQoveryServiceInfoOk

`func (o *NodePodInfoDto) GetQoveryServiceInfoOk() (*PodQoveryServiceInfoDto, bool)`

GetQoveryServiceInfoOk returns a tuple with the QoveryServiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryServiceInfo

`func (o *NodePodInfoDto) SetQoveryServiceInfo(v PodQoveryServiceInfoDto)`

SetQoveryServiceInfo sets QoveryServiceInfo field to given value.

### HasQoveryServiceInfo

`func (o *NodePodInfoDto) HasQoveryServiceInfo() bool`

HasQoveryServiceInfo returns a boolean if a field has been set.

### SetQoveryServiceInfoNil

`func (o *NodePodInfoDto) SetQoveryServiceInfoNil(b bool)`

 SetQoveryServiceInfoNil sets the value for QoveryServiceInfo to be an explicit nil

### UnsetQoveryServiceInfo
`func (o *NodePodInfoDto) UnsetQoveryServiceInfo()`

UnsetQoveryServiceInfo ensures that no value is present for QoveryServiceInfo, not even an explicit nil
### GetRestartCount

`func (o *NodePodInfoDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *NodePodInfoDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *NodePodInfoDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


