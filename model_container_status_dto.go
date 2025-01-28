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

// checks if the ContainerStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerStatusDto{}

// ContainerStatusDto struct for ContainerStatusDto
type ContainerStatusDto struct {
	CurrentState NullableContainerStateDto `json:"current_state,omitempty"`
	Image string `json:"image"`
	LastTerminatedState NullableContainerStateTerminatedDto `json:"last_terminated_state,omitempty"`
	Name string `json:"name"`
	RestartCount int32 `json:"restart_count"`
}

type _ContainerStatusDto ContainerStatusDto

// NewContainerStatusDto instantiates a new ContainerStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerStatusDto(image string, name string, restartCount int32) *ContainerStatusDto {
	this := ContainerStatusDto{}
	this.Image = image
	this.Name = name
	this.RestartCount = restartCount
	return &this
}

// NewContainerStatusDtoWithDefaults instantiates a new ContainerStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerStatusDtoWithDefaults() *ContainerStatusDto {
	this := ContainerStatusDto{}
	return &this
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStatusDto) GetCurrentState() ContainerStateDto {
	if o == nil || IsNil(o.CurrentState.Get()) {
		var ret ContainerStateDto
		return ret
	}
	return *o.CurrentState.Get()
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStatusDto) GetCurrentStateOk() (*ContainerStateDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentState.Get(), o.CurrentState.IsSet()
}

// HasCurrentState returns a boolean if a field has been set.
func (o *ContainerStatusDto) HasCurrentState() bool {
	if o != nil && o.CurrentState.IsSet() {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given NullableContainerStateDto and assigns it to the CurrentState field.
func (o *ContainerStatusDto) SetCurrentState(v ContainerStateDto) {
	o.CurrentState.Set(&v)
}
// SetCurrentStateNil sets the value for CurrentState to be an explicit nil
func (o *ContainerStatusDto) SetCurrentStateNil() {
	o.CurrentState.Set(nil)
}

// UnsetCurrentState ensures that no value is present for CurrentState, not even an explicit nil
func (o *ContainerStatusDto) UnsetCurrentState() {
	o.CurrentState.Unset()
}

// GetImage returns the Image field value
func (o *ContainerStatusDto) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ContainerStatusDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ContainerStatusDto) SetImage(v string) {
	o.Image = v
}

// GetLastTerminatedState returns the LastTerminatedState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerStatusDto) GetLastTerminatedState() ContainerStateTerminatedDto {
	if o == nil || IsNil(o.LastTerminatedState.Get()) {
		var ret ContainerStateTerminatedDto
		return ret
	}
	return *o.LastTerminatedState.Get()
}

// GetLastTerminatedStateOk returns a tuple with the LastTerminatedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerStatusDto) GetLastTerminatedStateOk() (*ContainerStateTerminatedDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTerminatedState.Get(), o.LastTerminatedState.IsSet()
}

// HasLastTerminatedState returns a boolean if a field has been set.
func (o *ContainerStatusDto) HasLastTerminatedState() bool {
	if o != nil && o.LastTerminatedState.IsSet() {
		return true
	}

	return false
}

// SetLastTerminatedState gets a reference to the given NullableContainerStateTerminatedDto and assigns it to the LastTerminatedState field.
func (o *ContainerStatusDto) SetLastTerminatedState(v ContainerStateTerminatedDto) {
	o.LastTerminatedState.Set(&v)
}
// SetLastTerminatedStateNil sets the value for LastTerminatedState to be an explicit nil
func (o *ContainerStatusDto) SetLastTerminatedStateNil() {
	o.LastTerminatedState.Set(nil)
}

// UnsetLastTerminatedState ensures that no value is present for LastTerminatedState, not even an explicit nil
func (o *ContainerStatusDto) UnsetLastTerminatedState() {
	o.LastTerminatedState.Unset()
}

// GetName returns the Name field value
func (o *ContainerStatusDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerStatusDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerStatusDto) SetName(v string) {
	o.Name = v
}

// GetRestartCount returns the RestartCount field value
func (o *ContainerStatusDto) GetRestartCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value
// and a boolean to check if the value has been set.
func (o *ContainerStatusDto) GetRestartCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartCount, true
}

// SetRestartCount sets field value
func (o *ContainerStatusDto) SetRestartCount(v int32) {
	o.RestartCount = v
}

func (o ContainerStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentState.IsSet() {
		toSerialize["current_state"] = o.CurrentState.Get()
	}
	toSerialize["image"] = o.Image
	if o.LastTerminatedState.IsSet() {
		toSerialize["last_terminated_state"] = o.LastTerminatedState.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["restart_count"] = o.RestartCount
	return toSerialize, nil
}

func (o *ContainerStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
		"name",
		"restart_count",
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

	varContainerStatusDto := _ContainerStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerStatusDto)

	if err != nil {
		return err
	}

	*o = ContainerStatusDto(varContainerStatusDto)

	return err
}

type NullableContainerStatusDto struct {
	value *ContainerStatusDto
	isSet bool
}

func (v NullableContainerStatusDto) Get() *ContainerStatusDto {
	return v.value
}

func (v *NullableContainerStatusDto) Set(val *ContainerStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerStatusDto(val *ContainerStatusDto) *NullableContainerStatusDto {
	return &NullableContainerStatusDto{value: val, isSet: true}
}

func (v NullableContainerStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


