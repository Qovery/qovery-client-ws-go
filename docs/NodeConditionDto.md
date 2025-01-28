# NodeConditionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHeartbeatTime** | **int64** |  | 
**LastTransitionTime** | **int64** |  | 
**Message** | **string** |  | 
**Reason** | **string** |  | 
**Status** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewNodeConditionDto

`func NewNodeConditionDto(lastHeartbeatTime int64, lastTransitionTime int64, message string, reason string, status string, type_ string, ) *NodeConditionDto`

NewNodeConditionDto instantiates a new NodeConditionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConditionDtoWithDefaults

`func NewNodeConditionDtoWithDefaults() *NodeConditionDto`

NewNodeConditionDtoWithDefaults instantiates a new NodeConditionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastHeartbeatTime

`func (o *NodeConditionDto) GetLastHeartbeatTime() int64`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *NodeConditionDto) GetLastHeartbeatTimeOk() (*int64, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *NodeConditionDto) SetLastHeartbeatTime(v int64)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.


### GetLastTransitionTime

`func (o *NodeConditionDto) GetLastTransitionTime() int64`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *NodeConditionDto) GetLastTransitionTimeOk() (*int64, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *NodeConditionDto) SetLastTransitionTime(v int64)`

SetLastTransitionTime sets LastTransitionTime field to given value.


### GetMessage

`func (o *NodeConditionDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeConditionDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeConditionDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *NodeConditionDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NodeConditionDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NodeConditionDto) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *NodeConditionDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeConditionDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeConditionDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *NodeConditionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeConditionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeConditionDto) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


