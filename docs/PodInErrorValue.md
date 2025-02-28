# PodInErrorValue

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

### NewPodInErrorValue

`func NewPodInErrorValue(componentName string, containerName string, level QoveryComponentContainerStatusLevel, podName string, type_ string, ) *PodInErrorValue`

NewPodInErrorValue instantiates a new PodInErrorValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodInErrorValueWithDefaults

`func NewPodInErrorValueWithDefaults() *PodInErrorValue`

NewPodInErrorValueWithDefaults instantiates a new PodInErrorValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *PodInErrorValue) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *PodInErrorValue) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *PodInErrorValue) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetContainerName

`func (o *PodInErrorValue) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *PodInErrorValue) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *PodInErrorValue) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetLevel

`func (o *PodInErrorValue) GetLevel() QoveryComponentContainerStatusLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PodInErrorValue) GetLevelOk() (*QoveryComponentContainerStatusLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PodInErrorValue) SetLevel(v QoveryComponentContainerStatusLevel)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *PodInErrorValue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PodInErrorValue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PodInErrorValue) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PodInErrorValue) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PodInErrorValue) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PodInErrorValue) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPodName

`func (o *PodInErrorValue) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *PodInErrorValue) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *PodInErrorValue) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetReason

`func (o *PodInErrorValue) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PodInErrorValue) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PodInErrorValue) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PodInErrorValue) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *PodInErrorValue) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PodInErrorValue) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetType

`func (o *PodInErrorValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PodInErrorValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PodInErrorValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


