# NodeResourceAllocatedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilli** | **int32** |  | 
**MemoryMib** | **int32** |  | 

## Methods

### NewNodeResourceAllocatedDto

`func NewNodeResourceAllocatedDto(cpuMilli int32, memoryMib int32, ) *NodeResourceAllocatedDto`

NewNodeResourceAllocatedDto instantiates a new NodeResourceAllocatedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeResourceAllocatedDtoWithDefaults

`func NewNodeResourceAllocatedDtoWithDefaults() *NodeResourceAllocatedDto`

NewNodeResourceAllocatedDtoWithDefaults instantiates a new NodeResourceAllocatedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilli

`func (o *NodeResourceAllocatedDto) GetCpuMilli() int32`

GetCpuMilli returns the CpuMilli field if non-nil, zero value otherwise.

### GetCpuMilliOk

`func (o *NodeResourceAllocatedDto) GetCpuMilliOk() (*int32, bool)`

GetCpuMilliOk returns a tuple with the CpuMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilli

`func (o *NodeResourceAllocatedDto) SetCpuMilli(v int32)`

SetCpuMilli sets CpuMilli field to given value.


### GetMemoryMib

`func (o *NodeResourceAllocatedDto) GetMemoryMib() int32`

GetMemoryMib returns the MemoryMib field if non-nil, zero value otherwise.

### GetMemoryMibOk

`func (o *NodeResourceAllocatedDto) GetMemoryMibOk() (*int32, bool)`

GetMemoryMibOk returns a tuple with the MemoryMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMib

`func (o *NodeResourceAllocatedDto) SetMemoryMib(v int32)`

SetMemoryMib sets MemoryMib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


