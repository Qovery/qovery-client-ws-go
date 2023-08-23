# ContainerStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to [**NullableContainerStatusDtoCurrentState**](ContainerStatusDtoCurrentState.md) |  | [optional] 
**Image** | **string** |  | 
**LastTerminatedState** | Pointer to [**NullableContainerStatusDtoLastTerminatedState**](ContainerStatusDtoLastTerminatedState.md) |  | [optional] 
**Name** | **string** |  | 
**RestartCount** | **int32** |  | 

## Methods

### NewContainerStatusDto

`func NewContainerStatusDto(image string, name string, restartCount int32, ) *ContainerStatusDto`

NewContainerStatusDto instantiates a new ContainerStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStatusDtoWithDefaults

`func NewContainerStatusDtoWithDefaults() *ContainerStatusDto`

NewContainerStatusDtoWithDefaults instantiates a new ContainerStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentState

`func (o *ContainerStatusDto) GetCurrentState() ContainerStatusDtoCurrentState`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *ContainerStatusDto) GetCurrentStateOk() (*ContainerStatusDtoCurrentState, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *ContainerStatusDto) SetCurrentState(v ContainerStatusDtoCurrentState)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *ContainerStatusDto) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### SetCurrentStateNil

`func (o *ContainerStatusDto) SetCurrentStateNil(b bool)`

 SetCurrentStateNil sets the value for CurrentState to be an explicit nil

### UnsetCurrentState
`func (o *ContainerStatusDto) UnsetCurrentState()`

UnsetCurrentState ensures that no value is present for CurrentState, not even an explicit nil
### GetImage

`func (o *ContainerStatusDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerStatusDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerStatusDto) SetImage(v string)`

SetImage sets Image field to given value.


### GetLastTerminatedState

`func (o *ContainerStatusDto) GetLastTerminatedState() ContainerStatusDtoLastTerminatedState`

GetLastTerminatedState returns the LastTerminatedState field if non-nil, zero value otherwise.

### GetLastTerminatedStateOk

`func (o *ContainerStatusDto) GetLastTerminatedStateOk() (*ContainerStatusDtoLastTerminatedState, bool)`

GetLastTerminatedStateOk returns a tuple with the LastTerminatedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTerminatedState

`func (o *ContainerStatusDto) SetLastTerminatedState(v ContainerStatusDtoLastTerminatedState)`

SetLastTerminatedState sets LastTerminatedState field to given value.

### HasLastTerminatedState

`func (o *ContainerStatusDto) HasLastTerminatedState() bool`

HasLastTerminatedState returns a boolean if a field has been set.

### SetLastTerminatedStateNil

`func (o *ContainerStatusDto) SetLastTerminatedStateNil(b bool)`

 SetLastTerminatedStateNil sets the value for LastTerminatedState to be an explicit nil

### UnsetLastTerminatedState
`func (o *ContainerStatusDto) UnsetLastTerminatedState()`

UnsetLastTerminatedState ensures that no value is present for LastTerminatedState, not even an explicit nil
### GetName

`func (o *ContainerStatusDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerStatusDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerStatusDto) SetName(v string)`

SetName sets Name field to given value.


### GetRestartCount

`func (o *ContainerStatusDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *ContainerStatusDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *ContainerStatusDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


