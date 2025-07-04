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

// PodStatusPhase the model 'PodStatusPhase'
type PodStatusPhase string

// List of PodStatusPhase
const (
	PODSTATUSPHASE_PENDING PodStatusPhase = "PENDING"
	PODSTATUSPHASE_RUNNING PodStatusPhase = "RUNNING"
	PODSTATUSPHASE_SUCCEEDED PodStatusPhase = "SUCCEEDED"
	PODSTATUSPHASE_FAILED PodStatusPhase = "FAILED"
	PODSTATUSPHASE_UNKNOWN PodStatusPhase = "UNKNOWN"
)

// All allowed values of PodStatusPhase enum
var AllowedPodStatusPhaseEnumValues = []PodStatusPhase{
	"PENDING",
	"RUNNING",
	"SUCCEEDED",
	"FAILED",
	"UNKNOWN",
}

func (v *PodStatusPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PodStatusPhase(value)
	for _, existing := range AllowedPodStatusPhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PodStatusPhase", value)
}

// NewPodStatusPhaseFromValue returns a pointer to a valid PodStatusPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPodStatusPhaseFromValue(v string) (*PodStatusPhase, error) {
	ev := PodStatusPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PodStatusPhase: valid values are %v", v, AllowedPodStatusPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PodStatusPhase) IsValid() bool {
	for _, existing := range AllowedPodStatusPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PodStatusPhase value
func (v PodStatusPhase) Ptr() *PodStatusPhase {
	return &v
}

type NullablePodStatusPhase struct {
	value *PodStatusPhase
	isSet bool
}

func (v NullablePodStatusPhase) Get() *PodStatusPhase {
	return v.value
}

func (v *NullablePodStatusPhase) Set(val *PodStatusPhase) {
	v.value = val
	v.isSet = true
}

func (v NullablePodStatusPhase) IsSet() bool {
	return v.isSet
}

func (v *NullablePodStatusPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodStatusPhase(val *PodStatusPhase) *NullablePodStatusPhase {
	return &NullablePodStatusPhase{value: val, isSet: true}
}

func (v NullablePodStatusPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodStatusPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

