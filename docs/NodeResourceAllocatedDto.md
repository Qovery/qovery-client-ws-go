# NodeResourceAllocatedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitCpuMilli** | **int32** |  | 
**LimitMemoryMib** | **int32** |  | 
**RequestCpuMilli** | **int32** |  | 
**RequestMemoryMib** | **int32** |  | 

## Methods

### NewNodeResourceAllocatedDto

`func NewNodeResourceAllocatedDto(limitCpuMilli int32, limitMemoryMib int32, requestCpuMilli int32, requestMemoryMib int32, ) *NodeResourceAllocatedDto`

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


