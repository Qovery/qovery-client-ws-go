# BuildAnalysisEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Details about the event, e.g. why the analysis failed. | [optional] 
**Status** | [**BuildAnalysisStatus**](BuildAnalysisStatus.md) |  | 
**Step** | **string** | Human-readable label of the current step. | 
**Terminal** | **bool** | True on the last event of the analysis (&#x60;DONE&#x60; or &#x60;FAILED&#x60;). | 

## Methods

### NewBuildAnalysisEvent

`func NewBuildAnalysisEvent(status BuildAnalysisStatus, step string, terminal bool, ) *BuildAnalysisEvent`

NewBuildAnalysisEvent instantiates a new BuildAnalysisEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildAnalysisEventWithDefaults

`func NewBuildAnalysisEventWithDefaults() *BuildAnalysisEvent`

NewBuildAnalysisEventWithDefaults instantiates a new BuildAnalysisEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BuildAnalysisEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BuildAnalysisEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BuildAnalysisEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BuildAnalysisEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BuildAnalysisEvent) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BuildAnalysisEvent) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStatus

`func (o *BuildAnalysisEvent) GetStatus() BuildAnalysisStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildAnalysisEvent) GetStatusOk() (*BuildAnalysisStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildAnalysisEvent) SetStatus(v BuildAnalysisStatus)`

SetStatus sets Status field to given value.


### GetStep

`func (o *BuildAnalysisEvent) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *BuildAnalysisEvent) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *BuildAnalysisEvent) SetStep(v string)`

SetStep sets Step field to given value.


### GetTerminal

`func (o *BuildAnalysisEvent) GetTerminal() bool`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *BuildAnalysisEvent) GetTerminalOk() (*bool, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *BuildAnalysisEvent) SetTerminal(v bool)`

SetTerminal sets Terminal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


