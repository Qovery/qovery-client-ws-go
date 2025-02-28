# QoveryComponentInFailure

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

### NewQoveryComponentInFailure

`func NewQoveryComponentInFailure(componentName string, containerName string, level QoveryComponentContainerStatusLevel, podName string, type_ string, ) *QoveryComponentInFailure`

NewQoveryComponentInFailure instantiates a new QoveryComponentInFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoveryComponentInFailureWithDefaults

`func NewQoveryComponentInFailureWithDefaults() *QoveryComponentInFailure`

NewQoveryComponentInFailureWithDefaults instantiates a new QoveryComponentInFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *QoveryComponentInFailure) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *QoveryComponentInFailure) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *QoveryComponentInFailure) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetContainerName

`func (o *QoveryComponentInFailure) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *QoveryComponentInFailure) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *QoveryComponentInFailure) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetLevel

`func (o *QoveryComponentInFailure) GetLevel() QoveryComponentContainerStatusLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *QoveryComponentInFailure) GetLevelOk() (*QoveryComponentContainerStatusLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *QoveryComponentInFailure) SetLevel(v QoveryComponentContainerStatusLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *QoveryComponentInFailure) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QoveryComponentInFailure) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QoveryComponentInFailure) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QoveryComponentInFailure) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QoveryComponentInFailure) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *QoveryComponentInFailure) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPodName

`func (o *QoveryComponentInFailure) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *QoveryComponentInFailure) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *QoveryComponentInFailure) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetReason

`func (o *QoveryComponentInFailure) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QoveryComponentInFailure) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QoveryComponentInFailure) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *QoveryComponentInFailure) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *QoveryComponentInFailure) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *QoveryComponentInFailure) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetType

`func (o *QoveryComponentInFailure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QoveryComponentInFailure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QoveryComponentInFailure) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


