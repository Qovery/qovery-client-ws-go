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

// checks if the NodeResourceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeResourceDto{}

// NodeResourceDto struct for NodeResourceDto
type NodeResourceDto struct {
	CpuMilli int64 `json:"cpu_milli"`
	EphemeralStorageGib int64 `json:"ephemeral_storage_gib"`
	MemoryMib int64 `json:"memory_mib"`
	Pods int64 `json:"pods"`
}

type _NodeResourceDto NodeResourceDto

// NewNodeResourceDto instantiates a new NodeResourceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeResourceDto(cpuMilli int64, ephemeralStorageGib int64, memoryMib int64, pods int64) *NodeResourceDto {
	this := NodeResourceDto{}
	this.CpuMilli = cpuMilli
	this.EphemeralStorageGib = ephemeralStorageGib
	this.MemoryMib = memoryMib
	this.Pods = pods
	return &this
}

// NewNodeResourceDtoWithDefaults instantiates a new NodeResourceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeResourceDtoWithDefaults() *NodeResourceDto {
	this := NodeResourceDto{}
	return &this
}

// GetCpuMilli returns the CpuMilli field value
func (o *NodeResourceDto) GetCpuMilli() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CpuMilli
}

// GetCpuMilliOk returns a tuple with the CpuMilli field value
// and a boolean to check if the value has been set.
func (o *NodeResourceDto) GetCpuMilliOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMilli, true
}

// SetCpuMilli sets field value
func (o *NodeResourceDto) SetCpuMilli(v int64) {
	o.CpuMilli = v
}

// GetEphemeralStorageGib returns the EphemeralStorageGib field value
func (o *NodeResourceDto) GetEphemeralStorageGib() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EphemeralStorageGib
}

// GetEphemeralStorageGibOk returns a tuple with the EphemeralStorageGib field value
// and a boolean to check if the value has been set.
func (o *NodeResourceDto) GetEphemeralStorageGibOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EphemeralStorageGib, true
}

// SetEphemeralStorageGib sets field value
func (o *NodeResourceDto) SetEphemeralStorageGib(v int64) {
	o.EphemeralStorageGib = v
}

// GetMemoryMib returns the MemoryMib field value
func (o *NodeResourceDto) GetMemoryMib() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MemoryMib
}

// GetMemoryMibOk returns a tuple with the MemoryMib field value
// and a boolean to check if the value has been set.
func (o *NodeResourceDto) GetMemoryMibOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryMib, true
}

// SetMemoryMib sets field value
func (o *NodeResourceDto) SetMemoryMib(v int64) {
	o.MemoryMib = v
}

// GetPods returns the Pods field value
func (o *NodeResourceDto) GetPods() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value
// and a boolean to check if the value has been set.
func (o *NodeResourceDto) GetPodsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pods, true
}

// SetPods sets field value
func (o *NodeResourceDto) SetPods(v int64) {
	o.Pods = v
}

func (o NodeResourceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeResourceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu_milli"] = o.CpuMilli
	toSerialize["ephemeral_storage_gib"] = o.EphemeralStorageGib
	toSerialize["memory_mib"] = o.MemoryMib
	toSerialize["pods"] = o.Pods
	return toSerialize, nil
}

func (o *NodeResourceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cpu_milli",
		"ephemeral_storage_gib",
		"memory_mib",
		"pods",
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

	varNodeResourceDto := _NodeResourceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeResourceDto)

	if err != nil {
		return err
	}

	*o = NodeResourceDto(varNodeResourceDto)

	return err
}

type NullableNodeResourceDto struct {
	value *NodeResourceDto
	isSet bool
}

func (v NullableNodeResourceDto) Get() *NodeResourceDto {
	return v.value
}

func (v *NullableNodeResourceDto) Set(val *NodeResourceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeResourceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeResourceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeResourceDto(val *NodeResourceDto) *NullableNodeResourceDto {
	return &NullableNodeResourceDto{value: val, isSet: true}
}

func (v NullableNodeResourceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeResourceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


