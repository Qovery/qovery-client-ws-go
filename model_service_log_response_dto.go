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

// checks if the ServiceLogResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceLogResponseDto{}

// ServiceLogResponseDto struct for ServiceLogResponseDto
type ServiceLogResponseDto struct {
	// Unix timestamp with millisecond precision
	CreatedAt int32 `json:"created_at"`
	Message string `json:"message"`
	PodName string `json:"pod_name"`
	Version string `json:"version"`
}

// NewServiceLogResponseDto instantiates a new ServiceLogResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLogResponseDto(createdAt int32, message string, podName string, version string) *ServiceLogResponseDto {
	this := ServiceLogResponseDto{}
	this.CreatedAt = createdAt
	this.Message = message
	this.PodName = podName
	this.Version = version
	return &this
}

// NewServiceLogResponseDtoWithDefaults instantiates a new ServiceLogResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLogResponseDtoWithDefaults() *ServiceLogResponseDto {
	this := ServiceLogResponseDto{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceLogResponseDto) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceLogResponseDto) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceLogResponseDto) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetMessage returns the Message field value
func (o *ServiceLogResponseDto) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ServiceLogResponseDto) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ServiceLogResponseDto) SetMessage(v string) {
	o.Message = v
}

// GetPodName returns the PodName field value
func (o *ServiceLogResponseDto) GetPodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value
// and a boolean to check if the value has been set.
func (o *ServiceLogResponseDto) GetPodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodName, true
}

// SetPodName sets field value
func (o *ServiceLogResponseDto) SetPodName(v string) {
	o.PodName = v
}

// GetVersion returns the Version field value
func (o *ServiceLogResponseDto) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ServiceLogResponseDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ServiceLogResponseDto) SetVersion(v string) {
	o.Version = v
}

func (o ServiceLogResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceLogResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["message"] = o.Message
	toSerialize["pod_name"] = o.PodName
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableServiceLogResponseDto struct {
	value *ServiceLogResponseDto
	isSet bool
}

func (v NullableServiceLogResponseDto) Get() *ServiceLogResponseDto {
	return v.value
}

func (v *NullableServiceLogResponseDto) Set(val *ServiceLogResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLogResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLogResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLogResponseDto(val *ServiceLogResponseDto) *NullableServiceLogResponseDto {
	return &NullableServiceLogResponseDto{value: val, isSet: true}
}

func (v NullableServiceLogResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLogResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


