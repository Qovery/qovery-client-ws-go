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

// checks if the QoveryComponentContainerStatusIssue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QoveryComponentContainerStatusIssue{}

// QoveryComponentContainerStatusIssue struct for QoveryComponentContainerStatusIssue
type QoveryComponentContainerStatusIssue struct {
	Level QoveryComponentContainerStatusLevel `json:"level"`
	Message NullableString `json:"message,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
}

type _QoveryComponentContainerStatusIssue QoveryComponentContainerStatusIssue

// NewQoveryComponentContainerStatusIssue instantiates a new QoveryComponentContainerStatusIssue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQoveryComponentContainerStatusIssue(level QoveryComponentContainerStatusLevel) *QoveryComponentContainerStatusIssue {
	this := QoveryComponentContainerStatusIssue{}
	this.Level = level
	return &this
}

// NewQoveryComponentContainerStatusIssueWithDefaults instantiates a new QoveryComponentContainerStatusIssue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQoveryComponentContainerStatusIssueWithDefaults() *QoveryComponentContainerStatusIssue {
	this := QoveryComponentContainerStatusIssue{}
	return &this
}

// GetLevel returns the Level field value
func (o *QoveryComponentContainerStatusIssue) GetLevel() QoveryComponentContainerStatusLevel {
	if o == nil {
		var ret QoveryComponentContainerStatusLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *QoveryComponentContainerStatusIssue) GetLevelOk() (*QoveryComponentContainerStatusLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *QoveryComponentContainerStatusIssue) SetLevel(v QoveryComponentContainerStatusLevel) {
	o.Level = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QoveryComponentContainerStatusIssue) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QoveryComponentContainerStatusIssue) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *QoveryComponentContainerStatusIssue) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *QoveryComponentContainerStatusIssue) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *QoveryComponentContainerStatusIssue) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *QoveryComponentContainerStatusIssue) UnsetMessage() {
	o.Message.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QoveryComponentContainerStatusIssue) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QoveryComponentContainerStatusIssue) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *QoveryComponentContainerStatusIssue) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *QoveryComponentContainerStatusIssue) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *QoveryComponentContainerStatusIssue) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *QoveryComponentContainerStatusIssue) UnsetReason() {
	o.Reason.Unset()
}

func (o QoveryComponentContainerStatusIssue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QoveryComponentContainerStatusIssue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	return toSerialize, nil
}

func (o *QoveryComponentContainerStatusIssue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
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

	varQoveryComponentContainerStatusIssue := _QoveryComponentContainerStatusIssue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQoveryComponentContainerStatusIssue)

	if err != nil {
		return err
	}

	*o = QoveryComponentContainerStatusIssue(varQoveryComponentContainerStatusIssue)

	return err
}

type NullableQoveryComponentContainerStatusIssue struct {
	value *QoveryComponentContainerStatusIssue
	isSet bool
}

func (v NullableQoveryComponentContainerStatusIssue) Get() *QoveryComponentContainerStatusIssue {
	return v.value
}

func (v *NullableQoveryComponentContainerStatusIssue) Set(val *QoveryComponentContainerStatusIssue) {
	v.value = val
	v.isSet = true
}

func (v NullableQoveryComponentContainerStatusIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableQoveryComponentContainerStatusIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQoveryComponentContainerStatusIssue(val *QoveryComponentContainerStatusIssue) *NullableQoveryComponentContainerStatusIssue {
	return &NullableQoveryComponentContainerStatusIssue{value: val, isSet: true}
}

func (v NullableQoveryComponentContainerStatusIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQoveryComponentContainerStatusIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


