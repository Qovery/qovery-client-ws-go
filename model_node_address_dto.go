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

// checks if the NodeAddressDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeAddressDto{}

// NodeAddressDto struct for NodeAddressDto
type NodeAddressDto struct {
	Address string `json:"address"`
	Type string `json:"type"`
}

// NewNodeAddressDto instantiates a new NodeAddressDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeAddressDto(address string, type_ string) *NodeAddressDto {
	this := NodeAddressDto{}
	this.Address = address
	this.Type = type_
	return &this
}

// NewNodeAddressDtoWithDefaults instantiates a new NodeAddressDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeAddressDtoWithDefaults() *NodeAddressDto {
	this := NodeAddressDto{}
	return &this
}

// GetAddress returns the Address field value
func (o *NodeAddressDto) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NodeAddressDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NodeAddressDto) SetAddress(v string) {
	o.Address = v
}

// GetType returns the Type field value
func (o *NodeAddressDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NodeAddressDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NodeAddressDto) SetType(v string) {
	o.Type = v
}

func (o NodeAddressDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeAddressDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNodeAddressDto struct {
	value *NodeAddressDto
	isSet bool
}

func (v NullableNodeAddressDto) Get() *NodeAddressDto {
	return v.value
}

func (v *NullableNodeAddressDto) Set(val *NodeAddressDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeAddressDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeAddressDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeAddressDto(val *NodeAddressDto) *NullableNodeAddressDto {
	return &NullableNodeAddressDto{value: val, isSet: true}
}

func (v NullableNodeAddressDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeAddressDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


