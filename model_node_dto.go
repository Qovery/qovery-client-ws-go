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

// checks if the NodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeDto{}

// NodeDto struct for NodeDto
type NodeDto struct {
	Name string `json:"name"`
}

type _NodeDto NodeDto

// NewNodeDto instantiates a new NodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeDto(name string) *NodeDto {
	this := NodeDto{}
	this.Name = name
	return &this
}

// NewNodeDtoWithDefaults instantiates a new NodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeDtoWithDefaults() *NodeDto {
	this := NodeDto{}
	return &this
}

// GetName returns the Name field value
func (o *NodeDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodeDto) SetName(v string) {
	o.Name = v
}

func (o NodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *NodeDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varNodeDto := _NodeDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeDto)

	if err != nil {
		return err
	}

	*o = NodeDto(varNodeDto)

	return err
}

type NullableNodeDto struct {
	value *NodeDto
	isSet bool
}

func (v NullableNodeDto) Get() *NodeDto {
	return v.value
}

func (v *NullableNodeDto) Set(val *NodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeDto(val *NodeDto) *NullableNodeDto {
	return &NullableNodeDto{value: val, isSet: true}
}

func (v NullableNodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


