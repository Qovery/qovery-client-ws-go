# QoveryNodeFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**NodeAttributes** | [**QoveryNodeAttributes**](QoveryNodeAttributes.md) |  | 
**Reason** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewQoveryNodeFailure

`func NewQoveryNodeFailure(message string, nodeAttributes QoveryNodeAttributes, reason string, type_ string, ) *QoveryNodeFailure`

NewQoveryNodeFailure instantiates a new QoveryNodeFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoveryNodeFailureWithDefaults

`func NewQoveryNodeFailureWithDefaults() *QoveryNodeFailure`

NewQoveryNodeFailureWithDefaults instantiates a new QoveryNodeFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *QoveryNodeFailure) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QoveryNodeFailure) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QoveryNodeFailure) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNodeAttributes

`func (o *QoveryNodeFailure) GetNodeAttributes() QoveryNodeAttributes`

GetNodeAttributes returns the NodeAttributes field if non-nil, zero value otherwise.

### GetNodeAttributesOk

`func (o *QoveryNodeFailure) GetNodeAttributesOk() (*QoveryNodeAttributes, bool)`

GetNodeAttributesOk returns a tuple with the NodeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAttributes

`func (o *QoveryNodeFailure) SetNodeAttributes(v QoveryNodeAttributes)`

SetNodeAttributes sets NodeAttributes field to given value.


### GetReason

`func (o *QoveryNodeFailure) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QoveryNodeFailure) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QoveryNodeFailure) SetReason(v string)`

SetReason sets Reason field to given value.


### GetType

`func (o *QoveryNodeFailure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QoveryNodeFailure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QoveryNodeFailure) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


