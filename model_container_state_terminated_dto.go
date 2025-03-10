/*
websocket-gateway

Describe the websocket endpoints of Qovery

API version: 0.1.0
Contact: erebe@erebe.eu
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery-ws

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerStateTerminatedDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerStateTerminatedDto{}

// ContainerStateTerminatedDto struct for ContainerStateTerminatedDto
type ContainerStateTerminatedDto struct {
	ExitCode int32 `json:"exit_code"`
	ExitCodeMessage string `json:"exit_code_message"`
	FinishedAt int64 `json:"finished_at"`
	Message string `json:"message"`
	Reason string `json:"reason"`
	Signal int32 `json:"signal"`
	StartedAt int64 `json:"started_at"`
}

type _ContainerStateTerminatedDto ContainerStateTerminatedDto

// NewContainerStateTerminatedDto instantiates a new ContainerStateTerminatedDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStateTerminatedDto(exitCode int32, exitCodeMessage string, finishedAt int64, message string, reason string, signal int32, startedAt int64) *ContainerStateTerminatedDto {
	this := ContainerStateTerminatedDto{}
	this.ExitCode = exitCode
	this.ExitCodeMessage = exitCodeMessage
	this.FinishedAt = finishedAt
	this.Message = message
	this.Reason = reason
	this.Signal = signal
	this.StartedAt = startedAt
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

// GetFinishedAt returns the FinishedAt field value
func (o *ContainerStateTerminatedDto) GetFinishedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetFinishedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *ContainerStateTerminatedDto) SetFinishedAt(v int64) {
	o.FinishedAt = v
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

// GetStartedAt returns the StartedAt field value
func (o *ContainerStateTerminatedDto) GetStartedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *ContainerStateTerminatedDto) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *ContainerStateTerminatedDto) SetStartedAt(v int64) {
	o.StartedAt = v
}

func (o ContainerStateTerminatedDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerStateTerminatedDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exit_code"] = o.ExitCode
	toSerialize["exit_code_message"] = o.ExitCodeMessage
	toSerialize["finished_at"] = o.FinishedAt
	toSerialize["message"] = o.Message
	toSerialize["reason"] = o.Reason
	toSerialize["signal"] = o.Signal
	toSerialize["started_at"] = o.StartedAt
	return toSerialize, nil
}

func (o *ContainerStateTerminatedDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exit_code",
		"exit_code_message",
		"finished_at",
		"message",
		"reason",
		"signal",
		"started_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContainerStateTerminatedDto := _ContainerStateTerminatedDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerStateTerminatedDto)

	if err != nil {
		return err
	}

	*o = ContainerStateTerminatedDto(varContainerStateTerminatedDto)

	return err
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


