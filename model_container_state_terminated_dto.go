/*
websocket-gateway

Describe the weboscket endpoints

API version: 0.1.0
Contact: erebe@erebe.eu
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ContainerStateTerminatedDto struct for ContainerStateTerminatedDto
type ContainerStateTerminatedDto struct {
	ExitCode int32 `json:"exit_code"`
	ExitCodeMessage string `json:"exit_code_message"`
	FinishedAt NullableInt32 `json:"finished_at,omitempty"`
	Message string `json:"message"`
	Reason string `json:"reason"`
	Signal int32 `json:"signal"`
	StartedAt NullableInt32 `json:"started_at,omitempty"`
}

// NewContainerStateTerminatedDto instantiates a new ContainerStateTerminatedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStateTerminatedDto(exitCode int32, exitCodeMessage string, message string, reason string, signal int32) *ContainerStateTerminatedDto {
	this := ContainerStateTerminatedDto{}
	this.ExitCode = exitCode
	this.ExitCodeMessage = exitCodeMessage
	this.Message = message
	this.Reason = reason
	this.Signal = signal
	return &this
}

// NewContainerStateTerminatedDtoWithDefaults instantiates a new ContainerStateTerminatedDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerStateTerminatedDtoWithDefaults() *ContainerStateTerminatedDto {
	this := ContainerStateTerminatedDto{}
	return &this
}

// GetExitCode returns the ExitCode field value
func (o *ContainerStateTerminatedDto) GetExitCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetExitCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitCode, true
}

// SetExitCode sets field value
func (o *ContainerStateTerminatedDto) SetExitCode(v int32) {
	o.ExitCode = v
}

// GetExitCodeMessage returns the ExitCodeMessage field value
func (o *ContainerStateTerminatedDto) GetExitCodeMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExitCodeMessage
}

// GetExitCodeMessageOk returns a tuple with the ExitCodeMessage field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetExitCodeMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitCodeMessage, true
}

// SetExitCodeMessage sets field value
func (o *ContainerStateTerminatedDto) SetExitCodeMessage(v string) {
	o.ExitCodeMessage = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStateTerminatedDto) GetFinishedAt() int32 {
	if o == nil || o.FinishedAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStateTerminatedDto) GetFinishedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ContainerStateTerminatedDto) HasFinishedAt() bool {
	if o != nil && o.FinishedAt.IsSet() {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given NullableInt32 and assigns it to the FinishedAt field.
func (o *ContainerStateTerminatedDto) SetFinishedAt(v int32) {
	o.FinishedAt.Set(&v)
}
// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil
func (o *ContainerStateTerminatedDto) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
func (o *ContainerStateTerminatedDto) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// GetMessage returns the Message field value
func (o *ContainerStateTerminatedDto) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ContainerStateTerminatedDto) SetMessage(v string) {
	o.Message = v
}

// GetReason returns the Reason field value
func (o *ContainerStateTerminatedDto) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ContainerStateTerminatedDto) SetReason(v string) {
	o.Reason = v
}

// GetSignal returns the Signal field value
func (o *ContainerStateTerminatedDto) GetSignal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Signal
}

// GetSignalOk returns a tuple with the Signal field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetSignalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signal, true
}

// SetSignal sets field value
func (o *ContainerStateTerminatedDto) SetSignal(v int32) {
	o.Signal = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStateTerminatedDto) GetStartedAt() int32 {
	if o == nil || o.StartedAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStateTerminatedDto) GetStartedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ContainerStateTerminatedDto) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableInt32 and assigns it to the StartedAt field.
func (o *ContainerStateTerminatedDto) SetStartedAt(v int32) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *ContainerStateTerminatedDto) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *ContainerStateTerminatedDto) UnsetStartedAt() {
	o.StartedAt.Unset()
}

func (o ContainerStateTerminatedDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exit_code"] = o.ExitCode
	}
	if true {
		toSerialize["exit_code_message"] = o.ExitCodeMessage
	}
	if o.FinishedAt.IsSet() {
		toSerialize["finished_at"] = o.FinishedAt.Get()
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["signal"] = o.Signal
	}
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContainerStateTerminatedDto struct {
	value *ContainerStateTerminatedDto
	isSet bool
}

func (v NullableContainerStateTerminatedDto) Get() *ContainerStateTerminatedDto {
	return v.value
}

func (v *NullableContainerStateTerminatedDto) Set(val *ContainerStateTerminatedDto) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerStateTerminatedDto) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerStateTerminatedDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerStateTerminatedDto(val *ContainerStateTerminatedDto) *NullableContainerStateTerminatedDto {
	return &NullableContainerStateTerminatedDto{value: val, isSet: true}
}

func (v NullableContainerStateTerminatedDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerStateTerminatedDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


