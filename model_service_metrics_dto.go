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

// checks if the ServiceMetricsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceMetricsDto{}

// ServiceMetricsDto struct for ServiceMetricsDto
type ServiceMetricsDto struct {
	Cpu MetricDto `json:"cpu"`
	Memory MetricDto `json:"memory"`
	PodName string `json:"pod_name"`
	Storages []MetricDto `json:"storages"`
}

type _ServiceMetricsDto ServiceMetricsDto

// NewServiceMetricsDto instantiates a new ServiceMetricsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMetricsDto(cpu MetricDto, memory MetricDto, podName string, storages []MetricDto) *ServiceMetricsDto {
	this := ServiceMetricsDto{}
	this.Cpu = cpu
	this.Memory = memory
	this.PodName = podName
	this.Storages = storages
	return &this
}

// NewServiceMetricsDtoWithDefaults instantiates a new ServiceMetricsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMetricsDtoWithDefaults() *ServiceMetricsDto {
	this := ServiceMetricsDto{}
	return &this
}

// GetCpu returns the Cpu field value
func (o *ServiceMetricsDto) GetCpu() MetricDto {
	if o == nil {
		var ret MetricDto
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *ServiceMetricsDto) GetCpuOk() (*MetricDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *ServiceMetricsDto) SetCpu(v MetricDto) {
	o.Cpu = v
}

// GetMemory returns the Memory field value
func (o *ServiceMetricsDto) GetMemory() MetricDto {
	if o == nil {
		var ret MetricDto
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *ServiceMetricsDto) GetMemoryOk() (*MetricDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *ServiceMetricsDto) SetMemory(v MetricDto) {
	o.Memory = v
}

// GetPodName returns the PodName field value
func (o *ServiceMetricsDto) GetPodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value
// and a boolean to check if the value has been set.
func (o *ServiceMetricsDto) GetPodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodName, true
}

// SetPodName sets field value
func (o *ServiceMetricsDto) SetPodName(v string) {
	o.PodName = v
}

// GetStorages returns the Storages field value
func (o *ServiceMetricsDto) GetStorages() []MetricDto {
	if o == nil {
		var ret []MetricDto
		return ret
	}

	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value
// and a boolean to check if the value has been set.
func (o *ServiceMetricsDto) GetStoragesOk() ([]MetricDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storages, true
}

// SetStorages sets field value
func (o *ServiceMetricsDto) SetStorages(v []MetricDto) {
	o.Storages = v
}

func (o ServiceMetricsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceMetricsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	toSerialize["pod_name"] = o.PodName
	toSerialize["storages"] = o.Storages
	return toSerialize, nil
}

func (o *ServiceMetricsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cpu",
		"memory",
		"pod_name",
		"storages",
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

	varServiceMetricsDto := _ServiceMetricsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceMetricsDto)

	if err != nil {
		return err
	}

	*o = ServiceMetricsDto(varServiceMetricsDto)

	return err
}

type NullableServiceMetricsDto struct {
	value *ServiceMetricsDto
	isSet bool
}

func (v NullableServiceMetricsDto) Get() *ServiceMetricsDto {
	return v.value
}

func (v *NullableServiceMetricsDto) Set(val *ServiceMetricsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMetricsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMetricsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMetricsDto(val *ServiceMetricsDto) *NullableServiceMetricsDto {
	return &NullableServiceMetricsDto{value: val, isSet: true}
}

func (v NullableServiceMetricsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMetricsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


