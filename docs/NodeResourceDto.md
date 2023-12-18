# NodeResourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilli** | **int64** |  | 
**EphemeralStorageGib** | **int64** |  | 
**MemoryMib** | **int64** |  | 
**Pods** | **int64** |  | 

## Methods

### NewNodeResourceDto

`func NewNodeResourceDto(cpuMilli int64, ephemeralStorageGib int64, memoryMib int64, pods int64, ) *NodeResourceDto`

NewNodeResourceDto instantiates a new NodeResourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeResourceDtoWithDefaults

`func NewNodeResourceDtoWithDefaults() *NodeResourceDto`

NewNodeResourceDtoWithDefaults instantiates a new NodeResourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilli

`func (o *NodeResourceDto) GetCpuMilli() int64`

GetCpuMilli returns the CpuMilli field if non-nil, zero value otherwise.

### GetCpuMilliOk

`func (o *NodeResourceDto) GetCpuMilliOk() (*int64, bool)`

GetCpuMilliOk returns a tuple with the CpuMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilli

`func (o *NodeResourceDto) SetCpuMilli(v int64)`

SetCpuMilli sets CpuMilli field to given value.


### GetEphemeralStorageGib

`func (o *NodeResourceDto) GetEphemeralStorageGib() int64`

GetEphemeralStorageGib returns the EphemeralStorageGib field if non-nil, zero value otherwise.

### GetEphemeralStorageGibOk

`func (o *NodeResourceDto) GetEphemeralStorageGibOk() (*int64, bool)`

GetEphemeralStorageGibOk returns a tuple with the EphemeralStorageGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageGib

`func (o *NodeResourceDto) SetEphemeralStorageGib(v int64)`

SetEphemeralStorageGib sets EphemeralStorageGib field to given value.


### GetMemoryMib

`func (o *NodeResourceDto) GetMemoryMib() int64`

GetMemoryMib returns the MemoryMib field if non-nil, zero value otherwise.

### GetMemoryMibOk

`func (o *NodeResourceDto) GetMemoryMibOk() (*int64, bool)`

GetMemoryMibOk returns a tuple with the MemoryMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMib

`func (o *NodeResourceDto) SetMemoryMib(v int64)`

SetMemoryMib sets MemoryMib field to given value.


### GetPods

`func (o *NodeResourceDto) GetPods() int64`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *NodeResourceDto) GetPodsOk() (*int64, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *NodeResourceDto) SetPods(v int64)`

SetPods sets Pods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


