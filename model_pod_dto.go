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

// checks if the PodDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PodDto{}

// PodDto struct for PodDto
type PodDto struct {
	Name string `json:"name"`
	Ports []int32 `json:"ports"`
}

type _PodDto PodDto

// NewPodDto instantiates a new PodDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodDto(name string, ports []int32) *PodDto {
	this := PodDto{}
	this.Name = name
	this.Ports = ports
	return &this
}

// NewPodDtoWithDefaults instantiates a new PodDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodDtoWithDefaults() *PodDto {
	this := PodDto{}
	return &this
}

// GetName returns the Name field value
func (o *PodDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PodDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PodDto) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value
func (o *PodDto) GetPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *PodDto) GetPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *PodDto) SetPorts(v []int32) {
	o.Ports = v
}

func (o PodDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["ports"] = o.Ports
	return toSerialize, nil
}

func (o *PodDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"ports",
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

	varPodDto := _PodDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPodDto)

	if err != nil {
		return err
	}

	*o = PodDto(varPodDto)

	return err
}

type NullablePodDto struct {
	value *PodDto
	isSet bool
}

func (v NullablePodDto) Get() *PodDto {
	return v.value
}

func (v *NullablePodDto) Set(val *PodDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePodDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePodDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodDto(val *PodDto) *NullablePodDto {
	return &NullablePodDto{value: val, isSet: true}
}

func (v NullablePodDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


