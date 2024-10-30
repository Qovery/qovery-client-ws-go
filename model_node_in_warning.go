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

// checks if the NodeInWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeInWarning{}

// NodeInWarning struct for NodeInWarning
type NodeInWarning struct {
	Reason string `json:"reason"`
	Message string `json:"message"`
}

type _NodeInWarning NodeInWarning

// NewNodeInWarning instantiates a new NodeInWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInWarning(reason string, message string) *NodeInWarning {
	this := NodeInWarning{}
	this.Reason = reason
	this.Message = message
	return &this
}

// NewNodeInWarningWithDefaults instantiates a new NodeInWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInWarningWithDefaults() *NodeInWarning {
	this := NodeInWarning{}
	return &this
}

// GetReason returns the Reason field value
func (o *NodeInWarning) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *NodeInWarning) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *NodeInWarning) SetReason(v string) {
	o.Reason = v
}

// GetMessage returns the Message field value
func (o *NodeInWarning) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NodeInWarning) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NodeInWarning) SetMessage(v string) {
	o.Message = v
}

func (o NodeInWarning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeInWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *NodeInWarning) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reason",
		"message",
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

	varNodeInWarning := _NodeInWarning{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeInWarning)

	if err != nil {
		return err
	}

	*o = NodeInWarning(varNodeInWarning)

	return err
}

type NullableNodeInWarning struct {
	value *NodeInWarning
	isSet bool
}

func (v NullableNodeInWarning) Get() *NodeInWarning {
	return v.value
}

func (v *NullableNodeInWarning) Set(val *NodeInWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInWarning(val *NodeInWarning) *NullableNodeInWarning {
	return &NullableNodeInWarning{value: val, isSet: true}
}

func (v NullableNodeInWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

