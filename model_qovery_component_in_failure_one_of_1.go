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

// checks if the QoveryComponentInFailureOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QoveryComponentInFailureOneOf1{}

// QoveryComponentInFailureOneOf1 struct for QoveryComponentInFailureOneOf1
type QoveryComponentInFailureOneOf1 struct {
	MISSING_COMPONENT QoveryComponentInFailureOneOf1MISSINGCOMPONENT `json:"MISSING_COMPONENT"`
}

type _QoveryComponentInFailureOneOf1 QoveryComponentInFailureOneOf1

// NewQoveryComponentInFailureOneOf1 instantiates a new QoveryComponentInFailureOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQoveryComponentInFailureOneOf1(mISSINGCOMPONENT QoveryComponentInFailureOneOf1MISSINGCOMPONENT) *QoveryComponentInFailureOneOf1 {
	this := QoveryComponentInFailureOneOf1{}
	this.MISSING_COMPONENT = mISSINGCOMPONENT
	return &this
}

// NewQoveryComponentInFailureOneOf1WithDefaults instantiates a new QoveryComponentInFailureOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQoveryComponentInFailureOneOf1WithDefaults() *QoveryComponentInFailureOneOf1 {
	this := QoveryComponentInFailureOneOf1{}
	return &this
}

// GetMISSING_COMPONENT returns the MISSING_COMPONENT field value
func (o *QoveryComponentInFailureOneOf1) GetMISSING_COMPONENT() QoveryComponentInFailureOneOf1MISSINGCOMPONENT {
	if o == nil {
		var ret QoveryComponentInFailureOneOf1MISSINGCOMPONENT
		return ret
	}

	return o.MISSING_COMPONENT
}

// GetMISSING_COMPONENTOk returns a tuple with the MISSING_COMPONENT field value
// and a boolean to check if the value has been set.
func (o *QoveryComponentInFailureOneOf1) GetMISSING_COMPONENTOk() (*QoveryComponentInFailureOneOf1MISSINGCOMPONENT, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MISSING_COMPONENT, true
}

// SetMISSING_COMPONENT sets field value
func (o *QoveryComponentInFailureOneOf1) SetMISSING_COMPONENT(v QoveryComponentInFailureOneOf1MISSINGCOMPONENT) {
	o.MISSING_COMPONENT = v
}

func (o QoveryComponentInFailureOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QoveryComponentInFailureOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MISSING_COMPONENT"] = o.MISSING_COMPONENT
	return toSerialize, nil
}

func (o *QoveryComponentInFailureOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"MISSING_COMPONENT",
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

	varQoveryComponentInFailureOneOf1 := _QoveryComponentInFailureOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQoveryComponentInFailureOneOf1)

	if err != nil {
		return err
	}

	*o = QoveryComponentInFailureOneOf1(varQoveryComponentInFailureOneOf1)

	return err
}

type NullableQoveryComponentInFailureOneOf1 struct {
	value *QoveryComponentInFailureOneOf1
	isSet bool
}

func (v NullableQoveryComponentInFailureOneOf1) Get() *QoveryComponentInFailureOneOf1 {
	return v.value
}

func (v *NullableQoveryComponentInFailureOneOf1) Set(val *QoveryComponentInFailureOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableQoveryComponentInFailureOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableQoveryComponentInFailureOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoveryComponentInFailureOneOf1(val *QoveryComponentInFailureOneOf1) *NullableQoveryComponentInFailureOneOf1 {
	return &NullableQoveryComponentInFailureOneOf1{value: val, isSet: true}
}

func (v NullableQoveryComponentInFailureOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoveryComponentInFailureOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


