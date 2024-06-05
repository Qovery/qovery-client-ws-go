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
	"bytes"
	"fmt"
)

// checks if the PodStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PodStatusDto{}

// PodStatusDto struct for PodStatusDto
type PodStatusDto struct {
	Containers []ContainerStatusDto `json:"containers"`
	Name string `json:"name"`
	RestartCount int32 `json:"restart_count"`
	ServiceVersion string `json:"service_version"`
	// Unix timestamp with millisecond precision
	StartedAt NullableInt32 `json:"started_at,omitempty"`
	State ServiceStateDto `json:"state"`
	StateMessage string `json:"state_message"`
	StateReason string `json:"state_reason"`
}

type _PodStatusDto PodStatusDto

// NewPodStatusDto instantiates a new PodStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodStatusDto(containers []ContainerStatusDto, name string, restartCount int32, serviceVersion string, state ServiceStateDto, stateMessage string, stateReason string) *PodStatusDto {
	this := PodStatusDto{}
	this.Containers = containers
	this.Name = name
	this.RestartCount = restartCount
	this.ServiceVersion = serviceVersion
	this.State = state
	this.StateMessage = stateMessage
	this.StateReason = stateReason
	return &this
}

// NewPodStatusDtoWithDefaults instantiates a new PodStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodStatusDtoWithDefaults() *PodStatusDto {
	this := PodStatusDto{}
	return &this
}

// GetContainers returns the Containers field value
func (o *PodStatusDto) GetContainers() []ContainerStatusDto {
	if o == nil {
		var ret []ContainerStatusDto
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetContainersOk() ([]ContainerStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *PodStatusDto) SetContainers(v []ContainerStatusDto) {
	o.Containers = v
}

// GetName returns the Name field value
func (o *PodStatusDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PodStatusDto) SetName(v string) {
	o.Name = v
}

// GetRestartCount returns the RestartCount field value
func (o *PodStatusDto) GetRestartCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetRestartCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartCount, true
}

// SetRestartCount sets field value
func (o *PodStatusDto) SetRestartCount(v int32) {
	o.RestartCount = v
}

// GetServiceVersion returns the ServiceVersion field value
func (o *PodStatusDto) GetServiceVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetServiceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceVersion, true
}

// SetServiceVersion sets field value
func (o *PodStatusDto) SetServiceVersion(v string) {
	o.ServiceVersion = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodStatusDto) GetStartedAt() int32 {
	if o == nil || IsNil(o.StartedAt.Get()) {
		var ret int32
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodStatusDto) GetStartedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *PodStatusDto) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableInt32 and assigns it to the StartedAt field.
func (o *PodStatusDto) SetStartedAt(v int32) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *PodStatusDto) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *PodStatusDto) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetState returns the State field value
func (o *PodStatusDto) GetState() ServiceStateDto {
	if o == nil {
		var ret ServiceStateDto
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetStateOk() (*ServiceStateDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PodStatusDto) SetState(v ServiceStateDto) {
	o.State = v
}

// GetStateMessage returns the StateMessage field value
func (o *PodStatusDto) GetStateMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateMessage
}

// GetStateMessageOk returns a tuple with the StateMessage field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetStateMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateMessage, true
}

// SetStateMessage sets field value
func (o *PodStatusDto) SetStateMessage(v string) {
	o.StateMessage = v
}

// GetStateReason returns the StateReason field value
func (o *PodStatusDto) GetStateReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value
// and a boolean to check if the value has been set.
func (o *PodStatusDto) GetStateReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateReason, true
}

// SetStateReason sets field value
func (o *PodStatusDto) SetStateReason(v string) {
	o.StateReason = v
}

func (o PodStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containers"] = o.Containers
	toSerialize["name"] = o.Name
	toSerialize["restart_count"] = o.RestartCount
	toSerialize["service_version"] = o.ServiceVersion
	if o.StartedAt.IsSet() {
		toSerialize["started_at"] = o.StartedAt.Get()
	}
	toSerialize["state"] = o.State
	toSerialize["state_message"] = o.StateMessage
	toSerialize["state_reason"] = o.StateReason
	return toSerialize, nil
}

func (o *PodStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"containers",
		"name",
		"restart_count",
		"service_version",
		"state",
		"state_message",
		"state_reason",
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

	varPodStatusDto := _PodStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPodStatusDto)

	if err != nil {
		return err
	}

	*o = PodStatusDto(varPodStatusDto)

	return err
}

type NullablePodStatusDto struct {
	value *PodStatusDto
	isSet bool
}

func (v NullablePodStatusDto) Get() *PodStatusDto {
	return v.value
}

func (v *NullablePodStatusDto) Set(val *PodStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePodStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePodStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodStatusDto(val *PodStatusDto) *NullablePodStatusDto {
	return &NullablePodStatusDto{value: val, isSet: true}
}

func (v NullablePodStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


