# QoveryComponentContainerStatusIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**QoveryComponentContainerStatusLevel**](QoveryComponentContainerStatusLevel.md) |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQoveryComponentContainerStatusIssue

`func NewQoveryComponentContainerStatusIssue(level QoveryComponentContainerStatusLevel, ) *QoveryComponentContainerStatusIssue`

NewQoveryComponentContainerStatusIssue instantiates a new QoveryComponentContainerStatusIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoveryComponentContainerStatusIssueWithDefaults

`func NewQoveryComponentContainerStatusIssueWithDefaults() *QoveryComponentContainerStatusIssue`

NewQoveryComponentContainerStatusIssueWithDefaults instantiates a new QoveryComponentContainerStatusIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *QoveryComponentContainerStatusIssue) GetLevel() QoveryComponentContainerStatusLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *QoveryComponentContainerStatusIssue) GetLevelOk() (*QoveryComponentContainerStatusLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *QoveryComponentContainerStatusIssue) SetLevel(v QoveryComponentContainerStatusLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *QoveryComponentContainerStatusIssue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QoveryComponentContainerStatusIssue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QoveryComponentContainerStatusIssue) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QoveryComponentContainerStatusIssue) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QoveryComponentContainerStatusIssue) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *QoveryComponentContainerStatusIssue) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetReason

`func (o *QoveryComponentContainerStatusIssue) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QoveryComponentContainerStatusIssue) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QoveryComponentContainerStatusIssue) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *QoveryComponentContainerStatusIssue) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *QoveryComponentContainerStatusIssue) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *QoveryComponentContainerStatusIssue) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


