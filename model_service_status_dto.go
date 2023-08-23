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
)

// ServiceStatusDto struct for ServiceStatusDto
type ServiceStatusDto struct {
	Environments []EnvironmentStatusDto `json:"environments"`
}

// NewServiceStatusDto instantiates a new ServiceStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStatusDto(environments []EnvironmentStatusDto) *ServiceStatusDto {
	this := ServiceStatusDto{}
	this.Environments = environments
	return &this
}

// NewServiceStatusDtoWithDefaults instantiates a new ServiceStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStatusDtoWithDefaults() *ServiceStatusDto {
	this := ServiceStatusDto{}
	return &this
}

// GetEnvironments returns the Environments field value
func (o *ServiceStatusDto) GetEnvironments() []EnvironmentStatusDto {
	if o == nil {
		var ret []EnvironmentStatusDto
		return ret
	}

	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *ServiceStatusDto) GetEnvironmentsOk() ([]EnvironmentStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environments, true
}

// SetEnvironments sets field value
func (o *ServiceStatusDto) SetEnvironments(v []EnvironmentStatusDto) {
	o.Environments = v
}

func (o ServiceStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["environments"] = o.Environments
	}
	return json.Marshal(toSerialize)
}

type NullableServiceStatusDto struct {
	value *ServiceStatusDto
	isSet bool
}

func (v NullableServiceStatusDto) Get() *ServiceStatusDto {
	return v.value
}

func (v *NullableServiceStatusDto) Set(val *ServiceStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStatusDto(val *ServiceStatusDto) *NullableServiceStatusDto {
	return &NullableServiceStatusDto{value: val, isSet: true}
}

func (v NullableServiceStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


