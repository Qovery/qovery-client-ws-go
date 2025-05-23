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
	"fmt"
)

// QoveryComponentInFailure - struct for QoveryComponentInFailure
type QoveryComponentInFailure struct {
	MissingComponentValue *MissingComponentValue
	PodInErrorValue *PodInErrorValue
}

// MissingComponentValueAsQoveryComponentInFailure is a convenience function that returns MissingComponentValue wrapped in QoveryComponentInFailure
func MissingComponentValueAsQoveryComponentInFailure(v *MissingComponentValue) QoveryComponentInFailure {
	return QoveryComponentInFailure{
		MissingComponentValue: v,
	}
}

// PodInErrorValueAsQoveryComponentInFailure is a convenience function that returns PodInErrorValue wrapped in QoveryComponentInFailure
func PodInErrorValueAsQoveryComponentInFailure(v *PodInErrorValue) QoveryComponentInFailure {
	return QoveryComponentInFailure{
		PodInErrorValue: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QoveryComponentInFailure) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MissingComponentValue
	err = newStrictDecoder(data).Decode(&dst.MissingComponentValue)
	if err == nil {
		jsonMissingComponentValue, _ := json.Marshal(dst.MissingComponentValue)
		if string(jsonMissingComponentValue) == "{}" { // empty struct
			dst.MissingComponentValue = nil
		} else {
			match++
		}
	} else {
		dst.MissingComponentValue = nil
	}

	// try to unmarshal data into PodInErrorValue
	err = newStrictDecoder(data).Decode(&dst.PodInErrorValue)
	if err == nil {
		jsonPodInErrorValue, _ := json.Marshal(dst.PodInErrorValue)
		if string(jsonPodInErrorValue) == "{}" { // empty struct
			dst.PodInErrorValue = nil
		} else {
			match++
		}
	} else {
		dst.PodInErrorValue = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MissingComponentValue = nil
		dst.PodInErrorValue = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QoveryComponentInFailure)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QoveryComponentInFailure)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QoveryComponentInFailure) MarshalJSON() ([]byte, error) {
	if src.MissingComponentValue != nil {
		return json.Marshal(&src.MissingComponentValue)
	}

	if src.PodInErrorValue != nil {
		return json.Marshal(&src.PodInErrorValue)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QoveryComponentInFailure) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MissingComponentValue != nil {
		return obj.MissingComponentValue
	}

	if obj.PodInErrorValue != nil {
		return obj.PodInErrorValue
	}

	// all schemas are nil
	return nil
}

type NullableQoveryComponentInFailure struct {
	value *QoveryComponentInFailure
	isSet bool
}

func (v NullableQoveryComponentInFailure) Get() *QoveryComponentInFailure {
	return v.value
}

func (v *NullableQoveryComponentInFailure) Set(val *QoveryComponentInFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableQoveryComponentInFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableQoveryComponentInFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoveryComponentInFailure(val *QoveryComponentInFailure) *NullableQoveryComponentInFailure {
	return &NullableQoveryComponentInFailure{value: val, isSet: true}
}

func (v NullableQoveryComponentInFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoveryComponentInFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


