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
	"fmt"
)

// checks if the ContainerStateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerStateDto{}

// ContainerStateDto struct for ContainerStateDto
type ContainerStateDto struct {
	// Unix timestamp with millisecond precision
	StartedAt NullableInt32 `json:"started_at,omitempty"`
	State ServiceStateDto `json:"state"`
	StateMessage NullableString `json:"state_message,omitempty"`
	StateReason NullableString `json:"state_reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContainerStateDto ContainerStateDto

// NewContainerStateDto instantiates a new ContainerStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStateDto(state ServiceStateDto) *ContainerStateDto {
	this := ContainerStateDto{}
	this.State = state
	return &this
}

// NewContainerStateDtoWithDefaults instantiates a new ContainerStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerStateDtoWithDefaults() *ContainerStateDto {
	this := ContainerStateDto{}
	return &this
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStateDto) GetStartedAt() int32 {
	if o == nil || IsNil(o.StartedAt.Get()) {
		var ret int32
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStateDto) GetStartedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ContainerStateDto) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableInt32 and assigns it to the StartedAt field.
func (o *ContainerStateDto) SetStartedAt(v int32) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *ContainerStateDto) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *ContainerStateDto) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetState returns the State field value
func (o *ContainerStateDto) GetState() ServiceStateDto {
	if o == nil {
		var ret ServiceStateDto
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ContainerStateDto) GetStateOk() (*ServiceStateDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ContainerStateDto) SetState(v ServiceStateDto) {
	o.State = v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStateDto) GetStateMessage() string {
	if o == nil || IsNil(o.StateMessage.Get()) {
		var ret string
		return ret
	}
	return *o.StateMessage.Get()
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStateDto) GetStateMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateMessage.Get(), o.StateMessage.IsSet()
}

// HasStateMessage returns a boolean if a field has been set.
func (o *ContainerStateDto) HasStateMessage() bool {
	if o != nil && o.StateMessage.IsSet() {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given NullableString and assigns it to the StateMessage field.
func (o *ContainerStateDto) SetStateMessage(v string) {
	o.StateMessage.Set(&v)
}
// SetStateMessageNil sets the value for StateMessage to be an explicit nil
func (o *ContainerStateDto) SetStateMessageNil() {
	o.StateMessage.Set(nil)
}

// UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil
func (o *ContainerStateDto) UnsetStateMessage() {
	o.StateMessage.Unset()
}

// GetStateReason returns the StateReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStateDto) GetStateReason() string {
	if o == nil || IsNil(o.StateReason.Get()) {
		var ret string
		return ret
	}
	return *o.StateReason.Get()
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStateDto) GetStateReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateReason.Get(), o.StateReason.IsSet()
}

// HasStateReason returns a boolean if a field has been set.
func (o *ContainerStateDto) HasStateReason() bool {
	if o != nil && o.StateReason.IsSet() {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given NullableString and assigns it to the StateReason field.
func (o *ContainerStateDto) SetStateReason(v string) {
	o.StateReason.Set(&v)
}
// SetStateReasonNil sets the value for StateReason to be an explicit nil
func (o *ContainerStateDto) SetStateReasonNil() {
	o.StateReason.Set(nil)
}

// UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
func (o *ContainerStateDto) UnsetStateReason() {
	o.StateReason.Unset()
}

func (o ContainerStateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerStateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	toSerialize["state"] = o.State
	if o.StateMessage.IsSet() {
		toSerialize["state_message"] = o.StateMessage.Get()
	}
	if o.StateReason.IsSet() {
		toSerialize["state_reason"] = o.StateReason.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContainerStateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varContainerStateDto := _ContainerStateDto{}

	err = json.Unmarshal(data, &varContainerStateDto)

	if err != nil {
		return err
	}

	*o = ContainerStateDto(varContainerStateDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "state")
		delete(additionalProperties, "state_message")
		delete(additionalProperties, "state_reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContainerStateDto struct {
	value *ContainerStateDto
	isSet bool
}

func (v NullableContainerStateDto) Get() *ContainerStateDto {
	return v.value
}

func (v *NullableContainerStateDto) Set(val *ContainerStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerStateDto(val *ContainerStateDto) *NullableContainerStateDto {
	return &NullableContainerStateDto{value: val, isSet: true}
}

func (v NullableContainerStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


