# NodePoolInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilli** | **int64** |  | 
**CpuMilliLimit** | Pointer to **NullableInt64** |  | [optional] 
**MemoryMib** | **int64** |  | 
**MemoryMibLimit** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**NodesCount** | **int32** |  | 

## Methods

### NewNodePoolInfoDto

`func NewNodePoolInfoDto(cpuMilli int64, memoryMib int64, name string, nodesCount int32, ) *NodePoolInfoDto`

NewNodePoolInfoDto instantiates a new NodePoolInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolInfoDtoWithDefaults

`func NewNodePoolInfoDtoWithDefaults() *NodePoolInfoDto`

NewNodePoolInfoDtoWithDefaults instantiates a new NodePoolInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilli

`func (o *NodePoolInfoDto) GetCpuMilli() int64`

GetCpuMilli returns the CpuMilli field if non-nil, zero value otherwise.

### GetCpuMilliOk

`func (o *NodePoolInfoDto) GetCpuMilliOk() (*int64, bool)`

GetCpuMilliOk returns a tuple with the CpuMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilli

`func (o *NodePoolInfoDto) SetCpuMilli(v int64)`

SetCpuMilli sets CpuMilli field to given value.


### GetCpuMilliLimit

`func (o *NodePoolInfoDto) GetCpuMilliLimit() int64`

GetCpuMilliLimit returns the CpuMilliLimit field if non-nil, zero value otherwise.

### GetCpuMilliLimitOk

`func (o *NodePoolInfoDto) GetCpuMilliLimitOk() (*int64, bool)`

GetCpuMilliLimitOk returns a tuple with the CpuMilliLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliLimit

`func (o *NodePoolInfoDto) SetCpuMilliLimit(v int64)`

SetCpuMilliLimit sets CpuMilliLimit field to given value.

### HasCpuMilliLimit

`func (o *NodePoolInfoDto) HasCpuMilliLimit() bool`

HasCpuMilliLimit returns a boolean if a field has been set.

### SetCpuMilliLimitNil

`func (o *NodePoolInfoDto) SetCpuMilliLimitNil(b bool)`

 SetCpuMilliLimitNil sets the value for CpuMilliLimit to be an explicit nil

### UnsetCpuMilliLimit
`func (o *NodePoolInfoDto) UnsetCpuMilliLimit()`

UnsetCpuMilliLimit ensures that no value is present for CpuMilliLimit, not even an explicit nil
### GetMemoryMib

`func (o *NodePoolInfoDto) GetMemoryMib() int64`

GetMemoryMib returns the MemoryMib field if non-nil, zero value otherwise.

### GetMemoryMibOk

`func (o *NodePoolInfoDto) GetMemoryMibOk() (*int64, bool)`

GetMemoryMibOk returns a tuple with the MemoryMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMib

`func (o *NodePoolInfoDto) SetMemoryMib(v int64)`

SetMemoryMib sets MemoryMib field to given value.


### GetMemoryMibLimit

`func (o *NodePoolInfoDto) GetMemoryMibLimit() int64`

GetMemoryMibLimit returns the MemoryMibLimit field if non-nil, zero value otherwise.

### GetMemoryMibLimitOk

`func (o *NodePoolInfoDto) GetMemoryMibLimitOk() (*int64, bool)`

GetMemoryMibLimitOk returns a tuple with the MemoryMibLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibLimit

`func (o *NodePoolInfoDto) SetMemoryMibLimit(v int64)`

SetMemoryMibLimit sets MemoryMibLimit field to given value.

### HasMemoryMibLimit

`func (o *NodePoolInfoDto) HasMemoryMibLimit() bool`

HasMemoryMibLimit returns a boolean if a field has been set.

### SetMemoryMibLimitNil

`func (o *NodePoolInfoDto) SetMemoryMibLimitNil(b bool)`

 SetMemoryMibLimitNil sets the value for MemoryMibLimit to be an explicit nil

### UnsetMemoryMibLimit
`func (o *NodePoolInfoDto) UnsetMemoryMibLimit()`

UnsetMemoryMibLimit ensures that no value is present for MemoryMibLimit, not even an explicit nil
### GetName

`func (o *NodePoolInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePoolInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePoolInfoDto) SetName(v string)`

SetName sets Name field to given value.


### GetNodesCount

`func (o *NodePoolInfoDto) GetNodesCount() int32`

GetNodesCount returns the NodesCount field if non-nil, zero value otherwise.

### GetNodesCountOk

`func (o *NodePoolInfoDto) GetNodesCountOk() (*int32, bool)`

GetNodesCountOk returns a tuple with the NodesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesCount

`func (o *NodePoolInfoDto) SetNodesCount(v int32)`

SetNodesCount sets NodesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


