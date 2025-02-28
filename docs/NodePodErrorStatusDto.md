# NodePodErrorStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNodePodErrorStatusDto

`func NewNodePodErrorStatusDto(containerName string, ) *NodePodErrorStatusDto`

NewNodePodErrorStatusDto instantiates a new NodePodErrorStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePodErrorStatusDtoWithDefaults

`func NewNodePodErrorStatusDtoWithDefaults() *NodePodErrorStatusDto`

NewNodePodErrorStatusDtoWithDefaults instantiates a new NodePodErrorStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *NodePodErrorStatusDto) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *NodePodErrorStatusDto) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *NodePodErrorStatusDto) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetMessage

`func (o *NodePodErrorStatusDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodePodErrorStatusDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodePodErrorStatusDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NodePodErrorStatusDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NodePodErrorStatusDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NodePodErrorStatusDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetReason

`func (o *NodePodErrorStatusDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NodePodErrorStatusDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NodePodErrorStatusDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NodePodErrorStatusDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *NodePodErrorStatusDto) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *NodePodErrorStatusDto) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


