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

// checks if the QoveryClusterKubeVersionStatusOneOfOk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QoveryClusterKubeVersionStatusOneOfOk{}

// QoveryClusterKubeVersionStatusOneOfOk struct for QoveryClusterKubeVersionStatusOneOfOk
type QoveryClusterKubeVersionStatusOneOfOk struct {
	KubeVersion string `json:"kube_version"`
}

type _QoveryClusterKubeVersionStatusOneOfOk QoveryClusterKubeVersionStatusOneOfOk

// NewQoveryClusterKubeVersionStatusOneOfOk instantiates a new QoveryClusterKubeVersionStatusOneOfOk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQoveryClusterKubeVersionStatusOneOfOk(kubeVersion string) *QoveryClusterKubeVersionStatusOneOfOk {
	this := QoveryClusterKubeVersionStatusOneOfOk{}
	this.KubeVersion = kubeVersion
	return &this
}

// NewQoveryClusterKubeVersionStatusOneOfOkWithDefaults instantiates a new QoveryClusterKubeVersionStatusOneOfOk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQoveryClusterKubeVersionStatusOneOfOkWithDefaults() *QoveryClusterKubeVersionStatusOneOfOk {
	this := QoveryClusterKubeVersionStatusOneOfOk{}
	return &this
}

// GetKubeVersion returns the KubeVersion field value
func (o *QoveryClusterKubeVersionStatusOneOfOk) GetKubeVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubeVersion
}

// GetKubeVersionOk returns a tuple with the KubeVersion field value
// and a boolean to check if the value has been set.
func (o *QoveryClusterKubeVersionStatusOneOfOk) GetKubeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeVersion, true
}

// SetKubeVersion sets field value
func (o *QoveryClusterKubeVersionStatusOneOfOk) SetKubeVersion(v string) {
	o.KubeVersion = v
}

func (o QoveryClusterKubeVersionStatusOneOfOk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QoveryClusterKubeVersionStatusOneOfOk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kube_version"] = o.KubeVersion
	return toSerialize, nil
}

func (o *QoveryClusterKubeVersionStatusOneOfOk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kube_version",
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

	varQoveryClusterKubeVersionStatusOneOfOk := _QoveryClusterKubeVersionStatusOneOfOk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQoveryClusterKubeVersionStatusOneOfOk)

	if err != nil {
		return err
	}

	*o = QoveryClusterKubeVersionStatusOneOfOk(varQoveryClusterKubeVersionStatusOneOfOk)

	return err
}

type NullableQoveryClusterKubeVersionStatusOneOfOk struct {
	value *QoveryClusterKubeVersionStatusOneOfOk
	isSet bool
}

func (v NullableQoveryClusterKubeVersionStatusOneOfOk) Get() *QoveryClusterKubeVersionStatusOneOfOk {
	return v.value
}

func (v *NullableQoveryClusterKubeVersionStatusOneOfOk) Set(val *QoveryClusterKubeVersionStatusOneOfOk) {
	v.value = val
	v.isSet = true
}

func (v NullableQoveryClusterKubeVersionStatusOneOfOk) IsSet() bool {
	return v.isSet
}

func (v *NullableQoveryClusterKubeVersionStatusOneOfOk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoveryClusterKubeVersionStatusOneOfOk(val *QoveryClusterKubeVersionStatusOneOfOk) *NullableQoveryClusterKubeVersionStatusOneOfOk {
	return &NullableQoveryClusterKubeVersionStatusOneOfOk{value: val, isSet: true}
}

func (v NullableQoveryClusterKubeVersionStatusOneOfOk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoveryClusterKubeVersionStatusOneOfOk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


