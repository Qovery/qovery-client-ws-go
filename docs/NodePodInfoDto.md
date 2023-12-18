# NodePodInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilliLimit** | Pointer to **NullableInt32** |  | [optional] 
**CpuMilliRequest** | Pointer to **NullableInt32** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**MemoryMibLimit** | Pointer to **NullableInt32** |  | [optional] 
**MemoryMibRequest** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ServiceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNodePodInfoDto

`func NewNodePodInfoDto(name string, namespace string, ) *NodePodInfoDto`

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
### GetEnvironmentId

`func (o *NodePodInfoDto) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *NodePodInfoDto) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *NodePodInfoDto) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *NodePodInfoDto) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *NodePodInfoDto) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *NodePodInfoDto) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
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


### GetProjectId

`func (o *NodePodInfoDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NodePodInfoDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NodePodInfoDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NodePodInfoDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *NodePodInfoDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NodePodInfoDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetServiceId

`func (o *NodePodInfoDto) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *NodePodInfoDto) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *NodePodInfoDto) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *NodePodInfoDto) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *NodePodInfoDto) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *NodePodInfoDto) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


