/*
websocket-gateway

Describe the weboscket endpoints

API version: 0.1.0
Contact: erebe@erebe.eu
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery-ws

import (
	"encoding/json"
)

// ContainerStatusDtoCurrentState struct for ContainerStatusDtoCurrentState
type ContainerStatusDtoCurrentState struct {
	StartedAt NullableInt32 `json:"started_at,omitempty"`
	State ServiceStateDto `json:"state"`
	StateMessage NullableString `json:"state_message,omitempty"`
	StateReason NullableString `json:"state_reason,omitempty"`
}

// NewContainerStatusDtoCurrentState instantiates a new ContainerStatusDtoCurrentState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStatusDtoCurrentState(state ServiceStateDto) *ContainerStatusDtoCurrentState {
	this := ContainerStatusDtoCurrentState{}
	this.State = state
	return &this
}

// NewContainerStatusDtoCurrentStateWithDefaults instantiates a new ContainerStatusDtoCurrentState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerStatusDtoCurrentStateWithDefaults() *ContainerStatusDtoCurrentState {
	this := ContainerStatusDtoCurrentState{}
	return &this
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStatusDtoCurrentState) GetStartedAt() int32 {
	if o == nil || o.StartedAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStatusDtoCurrentState) GetStartedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ContainerStatusDtoCurrentState) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableInt32 and assigns it to the StartedAt field.
func (o *ContainerStatusDtoCurrentState) SetStartedAt(v int32) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *ContainerStatusDtoCurrentState) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *ContainerStatusDtoCurrentState) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetState returns the State field value
func (o *ContainerStatusDtoCurrentState) GetState() ServiceStateDto {
	if o == nil {
		var ret ServiceStateDto
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ContainerStatusDtoCurrentState) GetStateOk() (*ServiceStateDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ContainerStatusDtoCurrentState) SetState(v ServiceStateDto) {
	o.State = v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStatusDtoCurrentState) GetStateMessage() string {
	if o == nil || o.StateMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateMessage.Get()
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStatusDtoCurrentState) GetStateMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateMessage.Get(), o.StateMessage.IsSet()
}

// HasStateMessage returns a boolean if a field has been set.
func (o *ContainerStatusDtoCurrentState) HasStateMessage() bool {
	if o != nil && o.StateMessage.IsSet() {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given NullableString and assigns it to the StateMessage field.
func (o *ContainerStatusDtoCurrentState) SetStateMessage(v string) {
	o.StateMessage.Set(&v)
}
// SetStateMessageNil sets the value for StateMessage to be an explicit nil
func (o *ContainerStatusDtoCurrentState) SetStateMessageNil() {
	o.StateMessage.Set(nil)
}

// UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil
func (o *ContainerStatusDtoCurrentState) UnsetStateMessage() {
	o.StateMessage.Unset()
}

// GetStateReason returns the StateReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStatusDtoCurrentState) GetStateReason() string {
	if o == nil || o.StateReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateReason.Get()
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStatusDtoCurrentState) GetStateReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateReason.Get(), o.StateReason.IsSet()
}

// HasStateReason returns a boolean if a field has been set.
func (o *ContainerStatusDtoCurrentState) HasStateReason() bool {
	if o != nil && o.StateReason.IsSet() {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given NullableString and assigns it to the StateReason field.
func (o *ContainerStatusDtoCurrentState) SetStateReason(v string) {
	o.StateReason.Set(&v)
}
// SetStateReasonNil sets the value for StateReason to be an explicit nil
func (o *ContainerStatusDtoCurrentState) SetStateReasonNil() {
	o.StateReason.Set(nil)
}

// UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
func (o *ContainerStatusDtoCurrentState) UnsetStateReason() {
	o.StateReason.Unset()
}

func (o ContainerStatusDtoCurrentState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.StateMessage.IsSet() {
		toSerialize["state_message"] = o.StateMessage.Get()
	}
	if o.StateReason.IsSet() {
		toSerialize["state_reason"] = o.StateReason.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContainerStatusDtoCurrentState struct {
	value *ContainerStatusDtoCurrentState
	isSet bool
}

func (v NullableContainerStatusDtoCurrentState) Get() *ContainerStatusDtoCurrentState {
	return v.value
}

func (v *NullableContainerStatusDtoCurrentState) Set(val *ContainerStatusDtoCurrentState) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerStatusDtoCurrentState) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerStatusDtoCurrentState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerStatusDtoCurrentState(val *ContainerStatusDtoCurrentState) *NullableContainerStatusDtoCurrentState {
	return &NullableContainerStatusDtoCurrentState{value: val, isSet: true}
}

func (v NullableContainerStatusDtoCurrentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerStatusDtoCurrentState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


