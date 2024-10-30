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

// checks if the QoveryMissingComponentInFailure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QoveryMissingComponentInFailure{}

// QoveryMissingComponentInFailure struct for QoveryMissingComponentInFailure
type QoveryMissingComponentInFailure struct {
	ComponentName string `json:"component_name"`
}

type _QoveryMissingComponentInFailure QoveryMissingComponentInFailure

// NewQoveryMissingComponentInFailure instantiates a new QoveryMissingComponentInFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQoveryMissingComponentInFailure(componentName string) *QoveryMissingComponentInFailure {
	this := QoveryMissingComponentInFailure{}
	this.ComponentName = componentName
	return &this
}

// NewQoveryMissingComponentInFailureWithDefaults instantiates a new QoveryMissingComponentInFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQoveryMissingComponentInFailureWithDefaults() *QoveryMissingComponentInFailure {
	this := QoveryMissingComponentInFailure{}
	return &this
}

// GetComponentName returns the ComponentName field value
func (o *QoveryMissingComponentInFailure) GetComponentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value
// and a boolean to check if the value has been set.
func (o *QoveryMissingComponentInFailure) GetComponentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentName, true
}

// SetComponentName sets field value
func (o *QoveryMissingComponentInFailure) SetComponentName(v string) {
	o.ComponentName = v
}

func (o QoveryMissingComponentInFailure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QoveryMissingComponentInFailure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["component_name"] = o.ComponentName
	return toSerialize, nil
}

func (o *QoveryMissingComponentInFailure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"component_name",
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

	varQoveryMissingComponentInFailure := _QoveryMissingComponentInFailure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQoveryMissingComponentInFailure)

	if err != nil {
		return err
	}

	*o = QoveryMissingComponentInFailure(varQoveryMissingComponentInFailure)

	return err
}

type NullableQoveryMissingComponentInFailure struct {
	value *QoveryMissingComponentInFailure
	isSet bool
}

func (v NullableQoveryMissingComponentInFailure) Get() *QoveryMissingComponentInFailure {
	return v.value
}

func (v *NullableQoveryMissingComponentInFailure) Set(val *QoveryMissingComponentInFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableQoveryMissingComponentInFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableQoveryMissingComponentInFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoveryMissingComponentInFailure(val *QoveryMissingComponentInFailure) *NullableQoveryMissingComponentInFailure {
	return &NullableQoveryMissingComponentInFailure{value: val, isSet: true}
}

func (v NullableQoveryMissingComponentInFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoveryMissingComponentInFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

