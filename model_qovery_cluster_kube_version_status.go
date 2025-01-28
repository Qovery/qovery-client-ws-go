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

// QoveryClusterKubeVersionStatus - struct for QoveryClusterKubeVersionStatus
type QoveryClusterKubeVersionStatus struct {
	QoveryClusterKubeVersionStatusOneOf *QoveryClusterKubeVersionStatusOneOf
	QoveryClusterKubeVersionStatusOneOf1 *QoveryClusterKubeVersionStatusOneOf1
	QoveryClusterKubeVersionStatusOneOf2 *QoveryClusterKubeVersionStatusOneOf2
}

// QoveryClusterKubeVersionStatusOneOfAsQoveryClusterKubeVersionStatus is a convenience function that returns QoveryClusterKubeVersionStatusOneOf wrapped in QoveryClusterKubeVersionStatus
func QoveryClusterKubeVersionStatusOneOfAsQoveryClusterKubeVersionStatus(v *QoveryClusterKubeVersionStatusOneOf) QoveryClusterKubeVersionStatus {
	return QoveryClusterKubeVersionStatus{
		QoveryClusterKubeVersionStatusOneOf: v,
	}
}

// QoveryClusterKubeVersionStatusOneOf1AsQoveryClusterKubeVersionStatus is a convenience function that returns QoveryClusterKubeVersionStatusOneOf1 wrapped in QoveryClusterKubeVersionStatus
func QoveryClusterKubeVersionStatusOneOf1AsQoveryClusterKubeVersionStatus(v *QoveryClusterKubeVersionStatusOneOf1) QoveryClusterKubeVersionStatus {
	return QoveryClusterKubeVersionStatus{
		QoveryClusterKubeVersionStatusOneOf1: v,
	}
}

// QoveryClusterKubeVersionStatusOneOf2AsQoveryClusterKubeVersionStatus is a convenience function that returns QoveryClusterKubeVersionStatusOneOf2 wrapped in QoveryClusterKubeVersionStatus
func QoveryClusterKubeVersionStatusOneOf2AsQoveryClusterKubeVersionStatus(v *QoveryClusterKubeVersionStatusOneOf2) QoveryClusterKubeVersionStatus {
	return QoveryClusterKubeVersionStatus{
		QoveryClusterKubeVersionStatusOneOf2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QoveryClusterKubeVersionStatus) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into QoveryClusterKubeVersionStatusOneOf
	err = newStrictDecoder(data).Decode(&dst.QoveryClusterKubeVersionStatusOneOf)
	if err == nil {
		jsonQoveryClusterKubeVersionStatusOneOf, _ := json.Marshal(dst.QoveryClusterKubeVersionStatusOneOf)
		if string(jsonQoveryClusterKubeVersionStatusOneOf) == "{}" { // empty struct
			dst.QoveryClusterKubeVersionStatusOneOf = nil
		} else {
			match++
		}
	} else {
		dst.QoveryClusterKubeVersionStatusOneOf = nil
	}

	// try to unmarshal data into QoveryClusterKubeVersionStatusOneOf1
	err = newStrictDecoder(data).Decode(&dst.QoveryClusterKubeVersionStatusOneOf1)
	if err == nil {
		jsonQoveryClusterKubeVersionStatusOneOf1, _ := json.Marshal(dst.QoveryClusterKubeVersionStatusOneOf1)
		if string(jsonQoveryClusterKubeVersionStatusOneOf1) == "{}" { // empty struct
			dst.QoveryClusterKubeVersionStatusOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.QoveryClusterKubeVersionStatusOneOf1 = nil
	}

	// try to unmarshal data into QoveryClusterKubeVersionStatusOneOf2
	err = newStrictDecoder(data).Decode(&dst.QoveryClusterKubeVersionStatusOneOf2)
	if err == nil {
		jsonQoveryClusterKubeVersionStatusOneOf2, _ := json.Marshal(dst.QoveryClusterKubeVersionStatusOneOf2)
		if string(jsonQoveryClusterKubeVersionStatusOneOf2) == "{}" { // empty struct
			dst.QoveryClusterKubeVersionStatusOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.QoveryClusterKubeVersionStatusOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.QoveryClusterKubeVersionStatusOneOf = nil
		dst.QoveryClusterKubeVersionStatusOneOf1 = nil
		dst.QoveryClusterKubeVersionStatusOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QoveryClusterKubeVersionStatus)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QoveryClusterKubeVersionStatus)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QoveryClusterKubeVersionStatus) MarshalJSON() ([]byte, error) {
	if src.QoveryClusterKubeVersionStatusOneOf != nil {
		return json.Marshal(&src.QoveryClusterKubeVersionStatusOneOf)
	}

	if src.QoveryClusterKubeVersionStatusOneOf1 != nil {
		return json.Marshal(&src.QoveryClusterKubeVersionStatusOneOf1)
	}

	if src.QoveryClusterKubeVersionStatusOneOf2 != nil {
		return json.Marshal(&src.QoveryClusterKubeVersionStatusOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QoveryClusterKubeVersionStatus) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.QoveryClusterKubeVersionStatusOneOf != nil {
		return obj.QoveryClusterKubeVersionStatusOneOf
	}

	if obj.QoveryClusterKubeVersionStatusOneOf1 != nil {
		return obj.QoveryClusterKubeVersionStatusOneOf1
	}

	if obj.QoveryClusterKubeVersionStatusOneOf2 != nil {
		return obj.QoveryClusterKubeVersionStatusOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableQoveryClusterKubeVersionStatus struct {
	value *QoveryClusterKubeVersionStatus
	isSet bool
}

func (v NullableQoveryClusterKubeVersionStatus) Get() *QoveryClusterKubeVersionStatus {
	return v.value
}

func (v *NullableQoveryClusterKubeVersionStatus) Set(val *QoveryClusterKubeVersionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableQoveryClusterKubeVersionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableQoveryClusterKubeVersionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoveryClusterKubeVersionStatus(val *QoveryClusterKubeVersionStatus) *NullableQoveryClusterKubeVersionStatus {
	return &NullableQoveryClusterKubeVersionStatus{value: val, isSet: true}
}

func (v NullableQoveryClusterKubeVersionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoveryClusterKubeVersionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


