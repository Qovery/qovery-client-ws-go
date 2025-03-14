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

// checks if the KubeVersionStatusOkValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubeVersionStatusOkValue{}

// KubeVersionStatusOkValue struct for KubeVersionStatusOkValue
type KubeVersionStatusOkValue struct {
	KubeVersion string `json:"kube_version"`
	Type string `json:"type"`
}

type _KubeVersionStatusOkValue KubeVersionStatusOkValue

// NewKubeVersionStatusOkValue instantiates a new KubeVersionStatusOkValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubeVersionStatusOkValue(kubeVersion string, type_ string) *KubeVersionStatusOkValue {
	this := KubeVersionStatusOkValue{}
	this.KubeVersion = kubeVersion
	this.Type = type_
	return &this
}

// NewKubeVersionStatusOkValueWithDefaults instantiates a new KubeVersionStatusOkValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubeVersionStatusOkValueWithDefaults() *KubeVersionStatusOkValue {
	this := KubeVersionStatusOkValue{}
	return &this
}

// GetKubeVersion returns the KubeVersion field value
func (o *KubeVersionStatusOkValue) GetKubeVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubeVersion
}

// GetKubeVersionOk returns a tuple with the KubeVersion field value
// and a boolean to check if the value has been set.
func (o *KubeVersionStatusOkValue) GetKubeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeVersion, true
}

// SetKubeVersion sets field value
func (o *KubeVersionStatusOkValue) SetKubeVersion(v string) {
	o.KubeVersion = v
}

// GetType returns the Type field value
func (o *KubeVersionStatusOkValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KubeVersionStatusOkValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KubeVersionStatusOkValue) SetType(v string) {
	o.Type = v
}

func (o KubeVersionStatusOkValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubeVersionStatusOkValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kube_version"] = o.KubeVersion
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *KubeVersionStatusOkValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kube_version",
		"type",
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

	varKubeVersionStatusOkValue := _KubeVersionStatusOkValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubeVersionStatusOkValue)

	if err != nil {
		return err
	}

	*o = KubeVersionStatusOkValue(varKubeVersionStatusOkValue)

	return err
}

type NullableKubeVersionStatusOkValue struct {
	value *KubeVersionStatusOkValue
	isSet bool
}

func (v NullableKubeVersionStatusOkValue) Get() *KubeVersionStatusOkValue {
	return v.value
}

func (v *NullableKubeVersionStatusOkValue) Set(val *KubeVersionStatusOkValue) {
	v.value = val
	v.isSet = true
}

func (v NullableKubeVersionStatusOkValue) IsSet() bool {
	return v.isSet
}

func (v *NullableKubeVersionStatusOkValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubeVersionStatusOkValue(val *KubeVersionStatusOkValue) *NullableKubeVersionStatusOkValue {
	return &NullableKubeVersionStatusOkValue{value: val, isSet: true}
}

func (v NullableKubeVersionStatusOkValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubeVersionStatusOkValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


