# ContainerStateTerminatedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | **int32** |  | 
**ExitCodeMessage** | **string** |  | 
**FinishedAt** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**Message** | **string** |  | 
**Reason** | **string** |  | 
**Signal** | **int32** |  | 
**StartedAt** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 

## Methods

### NewContainerStateTerminatedDto

`func NewContainerStateTerminatedDto(exitCode int32, exitCodeMessage string, message string, reason string, signal int32, ) *ContainerStateTerminatedDto`

NewContainerStateTerminatedDto instantiates a new ContainerStateTerminatedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStateTerminatedDtoWithDefaults

`func NewContainerStateTerminatedDtoWithDefaults() *ContainerStateTerminatedDto`

NewContainerStateTerminatedDtoWithDefaults instantiates a new ContainerStateTerminatedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *ContainerStateTerminatedDto) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ContainerStateTerminatedDto) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ContainerStateTerminatedDto) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetExitCodeMessage

`func (o *ContainerStateTerminatedDto) GetExitCodeMessage() string`

GetExitCodeMessage returns the ExitCodeMessage field if non-nil, zero value otherwise.

### GetExitCodeMessageOk

`func (o *ContainerStateTerminatedDto) GetExitCodeMessageOk() (*string, bool)`

GetExitCodeMessageOk returns a tuple with the ExitCodeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCodeMessage

`func (o *ContainerStateTerminatedDto) SetExitCodeMessage(v string)`

SetExitCodeMessage sets ExitCodeMessage field to given value.


### GetFinishedAt

`func (o *ContainerStateTerminatedDto) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ContainerStateTerminatedDto) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ContainerStateTerminatedDto) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ContainerStateTerminatedDto) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ContainerStateTerminatedDto) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ContainerStateTerminatedDto) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetMessage

`func (o *ContainerStateTerminatedDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ContainerStateTerminatedDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ContainerStateTerminatedDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *ContainerStateTerminatedDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContainerStateTerminatedDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContainerStateTerminatedDto) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSignal

`func (o *ContainerStateTerminatedDto) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *ContainerStateTerminatedDto) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *ContainerStateTerminatedDto) SetSignal(v int32)`

SetSignal sets Signal field to given value.


### GetStartedAt

`func (o *ContainerStateTerminatedDto) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerStateTerminatedDto) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerStateTerminatedDto) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ContainerStateTerminatedDto) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ContainerStateTerminatedDto) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ContainerStateTerminatedDto) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


