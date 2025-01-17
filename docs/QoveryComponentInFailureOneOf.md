# QoveryComponentInFailureOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | **string** |  | 
**ContainerName** | **string** |  | 
**Level** | [**QoveryComponentContainerStatusLevel**](QoveryComponentContainerStatusLevel.md) |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**PodName** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewQoveryComponentInFailureOneOf

`func NewQoveryComponentInFailureOneOf(componentName string, containerName string, level QoveryComponentContainerStatusLevel, podName string, type_ string, ) *QoveryComponentInFailureOneOf`

NewQoveryComponentInFailureOneOf instantiates a new QoveryComponentInFailureOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoveryComponentInFailureOneOfWithDefaults

`func NewQoveryComponentInFailureOneOfWithDefaults() *QoveryComponentInFailureOneOf`

NewQoveryComponentInFailureOneOfWithDefaults instantiates a new QoveryComponentInFailureOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *QoveryComponentInFailureOneOf) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *QoveryComponentInFailureOneOf) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *QoveryComponentInFailureOneOf) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetContainerName

`func (o *QoveryComponentInFailureOneOf) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *QoveryComponentInFailureOneOf) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *QoveryComponentInFailureOneOf) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetLevel

`func (o *QoveryComponentInFailureOneOf) GetLevel() QoveryComponentContainerStatusLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *QoveryComponentInFailureOneOf) GetLevelOk() (*QoveryComponentContainerStatusLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *QoveryComponentInFailureOneOf) SetLevel(v QoveryComponentContainerStatusLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *QoveryComponentInFailureOneOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QoveryComponentInFailureOneOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QoveryComponentInFailureOneOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QoveryComponentInFailureOneOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QoveryComponentInFailureOneOf) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *QoveryComponentInFailureOneOf) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPodName

`func (o *QoveryComponentInFailureOneOf) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *QoveryComponentInFailureOneOf) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *QoveryComponentInFailureOneOf) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetReason

`func (o *QoveryComponentInFailureOneOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QoveryComponentInFailureOneOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QoveryComponentInFailureOneOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *QoveryComponentInFailureOneOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *QoveryComponentInFailureOneOf) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *QoveryComponentInFailureOneOf) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetType

`func (o *QoveryComponentInFailureOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QoveryComponentInFailureOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QoveryComponentInFailureOneOf) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


