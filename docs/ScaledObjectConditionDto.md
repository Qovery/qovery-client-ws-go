# ScaledObjectConditionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewScaledObjectConditionDto

`func NewScaledObjectConditionDto(status string, type_ string, ) *ScaledObjectConditionDto`

NewScaledObjectConditionDto instantiates a new ScaledObjectConditionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaledObjectConditionDtoWithDefaults

`func NewScaledObjectConditionDtoWithDefaults() *ScaledObjectConditionDto`

NewScaledObjectConditionDtoWithDefaults instantiates a new ScaledObjectConditionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ScaledObjectConditionDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScaledObjectConditionDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScaledObjectConditionDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScaledObjectConditionDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ScaledObjectConditionDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ScaledObjectConditionDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetReason

`func (o *ScaledObjectConditionDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ScaledObjectConditionDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ScaledObjectConditionDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ScaledObjectConditionDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ScaledObjectConditionDto) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ScaledObjectConditionDto) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetStatus

`func (o *ScaledObjectConditionDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScaledObjectConditionDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScaledObjectConditionDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ScaledObjectConditionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScaledObjectConditionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScaledObjectConditionDto) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


