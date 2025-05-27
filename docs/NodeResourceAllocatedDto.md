# NodeResourceAllocatedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitCpuMilli** | Pointer to **NullableInt32** |  | [optional] 
**LimitMemoryMib** | Pointer to **NullableInt32** |  | [optional] 
**RequestCpuMilli** | **int32** |  | 
**RequestMemoryMib** | **int32** |  | 

## Methods

### NewNodeResourceAllocatedDto

`func NewNodeResourceAllocatedDto(requestCpuMilli int32, requestMemoryMib int32, ) *NodeResourceAllocatedDto`

NewNodeResourceAllocatedDto instantiates a new NodeResourceAllocatedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeResourceAllocatedDtoWithDefaults

`func NewNodeResourceAllocatedDtoWithDefaults() *NodeResourceAllocatedDto`

NewNodeResourceAllocatedDtoWithDefaults instantiates a new NodeResourceAllocatedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitCpuMilli

`func (o *NodeResourceAllocatedDto) GetLimitCpuMilli() int32`

GetLimitCpuMilli returns the LimitCpuMilli field if non-nil, zero value otherwise.

### GetLimitCpuMilliOk

`func (o *NodeResourceAllocatedDto) GetLimitCpuMilliOk() (*int32, bool)`

GetLimitCpuMilliOk returns a tuple with the LimitCpuMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCpuMilli

`func (o *NodeResourceAllocatedDto) SetLimitCpuMilli(v int32)`

SetLimitCpuMilli sets LimitCpuMilli field to given value.

### HasLimitCpuMilli

`func (o *NodeResourceAllocatedDto) HasLimitCpuMilli() bool`

HasLimitCpuMilli returns a boolean if a field has been set.

### SetLimitCpuMilliNil

`func (o *NodeResourceAllocatedDto) SetLimitCpuMilliNil(b bool)`

 SetLimitCpuMilliNil sets the value for LimitCpuMilli to be an explicit nil

### UnsetLimitCpuMilli
`func (o *NodeResourceAllocatedDto) UnsetLimitCpuMilli()`

UnsetLimitCpuMilli ensures that no value is present for LimitCpuMilli, not even an explicit nil
### GetLimitMemoryMib

`func (o *NodeResourceAllocatedDto) GetLimitMemoryMib() int32`

GetLimitMemoryMib returns the LimitMemoryMib field if non-nil, zero value otherwise.

### GetLimitMemoryMibOk

`func (o *NodeResourceAllocatedDto) GetLimitMemoryMibOk() (*int32, bool)`

GetLimitMemoryMibOk returns a tuple with the LimitMemoryMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMemoryMib

`func (o *NodeResourceAllocatedDto) SetLimitMemoryMib(v int32)`

SetLimitMemoryMib sets LimitMemoryMib field to given value.

### HasLimitMemoryMib

`func (o *NodeResourceAllocatedDto) HasLimitMemoryMib() bool`

HasLimitMemoryMib returns a boolean if a field has been set.

### SetLimitMemoryMibNil

`func (o *NodeResourceAllocatedDto) SetLimitMemoryMibNil(b bool)`

 SetLimitMemoryMibNil sets the value for LimitMemoryMib to be an explicit nil

### UnsetLimitMemoryMib
`func (o *NodeResourceAllocatedDto) UnsetLimitMemoryMib()`

UnsetLimitMemoryMib ensures that no value is present for LimitMemoryMib, not even an explicit nil
### GetRequestCpuMilli

`func (o *NodeResourceAllocatedDto) GetRequestCpuMilli() int32`

GetRequestCpuMilli returns the RequestCpuMilli field if non-nil, zero value otherwise.

### GetRequestCpuMilliOk

`func (o *NodeResourceAllocatedDto) GetRequestCpuMilliOk() (*int32, bool)`

GetRequestCpuMilliOk returns a tuple with the RequestCpuMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCpuMilli

`func (o *NodeResourceAllocatedDto) SetRequestCpuMilli(v int32)`

SetRequestCpuMilli sets RequestCpuMilli field to given value.


### GetRequestMemoryMib

`func (o *NodeResourceAllocatedDto) GetRequestMemoryMib() int32`

GetRequestMemoryMib returns the RequestMemoryMib field if non-nil, zero value otherwise.

### GetRequestMemoryMibOk

`func (o *NodeResourceAllocatedDto) GetRequestMemoryMibOk() (*int32, bool)`

GetRequestMemoryMibOk returns a tuple with the RequestMemoryMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMemoryMib

`func (o *NodeResourceAllocatedDto) SetRequestMemoryMib(v int32)`

SetRequestMemoryMib sets RequestMemoryMib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


