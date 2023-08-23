# ContainerStatusDtoLastTerminatedState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | **int32** |  | 
**ExitCodeMessage** | **string** |  | 
**FinishedAt** | Pointer to **NullableInt32** |  | [optional] 
**Message** | **string** |  | 
**Reason** | **string** |  | 
**Signal** | **int32** |  | 
**StartedAt** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewContainerStatusDtoLastTerminatedState

`func NewContainerStatusDtoLastTerminatedState(exitCode int32, exitCodeMessage string, message string, reason string, signal int32, ) *ContainerStatusDtoLastTerminatedState`

NewContainerStatusDtoLastTerminatedState instantiates a new ContainerStatusDtoLastTerminatedState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStatusDtoLastTerminatedStateWithDefaults

`func NewContainerStatusDtoLastTerminatedStateWithDefaults() *ContainerStatusDtoLastTerminatedState`

NewContainerStatusDtoLastTerminatedStateWithDefaults instantiates a new ContainerStatusDtoLastTerminatedState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *ContainerStatusDtoLastTerminatedState) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ContainerStatusDtoLastTerminatedState) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ContainerStatusDtoLastTerminatedState) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetExitCodeMessage

`func (o *ContainerStatusDtoLastTerminatedState) GetExitCodeMessage() string`

GetExitCodeMessage returns the ExitCodeMessage field if non-nil, zero value otherwise.

### GetExitCodeMessageOk

`func (o *ContainerStatusDtoLastTerminatedState) GetExitCodeMessageOk() (*string, bool)`

GetExitCodeMessageOk returns a tuple with the ExitCodeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCodeMessage

`func (o *ContainerStatusDtoLastTerminatedState) SetExitCodeMessage(v string)`

SetExitCodeMessage sets ExitCodeMessage field to given value.


### GetFinishedAt

`func (o *ContainerStatusDtoLastTerminatedState) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ContainerStatusDtoLastTerminatedState) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ContainerStatusDtoLastTerminatedState) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ContainerStatusDtoLastTerminatedState) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ContainerStatusDtoLastTerminatedState) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ContainerStatusDtoLastTerminatedState) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetMessage

`func (o *ContainerStatusDtoLastTerminatedState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContainerStatusDtoLastTerminatedState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContainerStatusDtoLastTerminatedState) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *ContainerStatusDtoLastTerminatedState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContainerStatusDtoLastTerminatedState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContainerStatusDtoLastTerminatedState) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSignal

`func (o *ContainerStatusDtoLastTerminatedState) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *ContainerStatusDtoLastTerminatedState) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *ContainerStatusDtoLastTerminatedState) SetSignal(v int32)`

SetSignal sets Signal field to given value.


### GetStartedAt

`func (o *ContainerStatusDtoLastTerminatedState) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerStatusDtoLastTerminatedState) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerStatusDtoLastTerminatedState) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ContainerStatusDtoLastTerminatedState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ContainerStatusDtoLastTerminatedState) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ContainerStatusDtoLastTerminatedState) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


