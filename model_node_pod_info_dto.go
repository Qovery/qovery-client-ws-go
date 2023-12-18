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

// checks if the NodePodInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodePodInfoDto{}

// NodePodInfoDto struct for NodePodInfoDto
type NodePodInfoDto struct {
	CpuMilliLimit NullableInt32 `json:"cpu_milli_limit,omitempty"`
	CpuMilliRequest NullableInt32 `json:"cpu_milli_request,omitempty"`
	EnvironmentId NullableString `json:"environment_id,omitempty"`
	ImagesVersion map[string]string `json:"images_version"`
	MemoryMibLimit NullableInt32 `json:"memory_mib_limit,omitempty"`
	MemoryMibRequest NullableInt32 `json:"memory_mib_request,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	ProjectId NullableString `json:"project_id,omitempty"`
	ServiceId NullableString `json:"service_id,omitempty"`
}

// NewNodePodInfoDto instantiates a new NodePodInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePodInfoDto(imagesVersion map[string]string, name string, namespace string) *NodePodInfoDto {
	this := NodePodInfoDto{}
	this.ImagesVersion = imagesVersion
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewNodePodInfoDtoWithDefaults instantiates a new NodePodInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePodInfoDtoWithDefaults() *NodePodInfoDto {
	this := NodePodInfoDto{}
	return &this
}

// GetCpuMilliLimit returns the CpuMilliLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetCpuMilliLimit() int32 {
	if o == nil || IsNil(o.CpuMilliLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.CpuMilliLimit.Get()
}

// GetCpuMilliLimitOk returns a tuple with the CpuMilliLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetCpuMilliLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuMilliLimit.Get(), o.CpuMilliLimit.IsSet()
}

// HasCpuMilliLimit returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasCpuMilliLimit() bool {
	if o != nil && o.CpuMilliLimit.IsSet() {
		return true
	}

	return false
}

// SetCpuMilliLimit gets a reference to the given NullableInt32 and assigns it to the CpuMilliLimit field.
func (o *NodePodInfoDto) SetCpuMilliLimit(v int32) {
	o.CpuMilliLimit.Set(&v)
}
// SetCpuMilliLimitNil sets the value for CpuMilliLimit to be an explicit nil
func (o *NodePodInfoDto) SetCpuMilliLimitNil() {
	o.CpuMilliLimit.Set(nil)
}

// UnsetCpuMilliLimit ensures that no value is present for CpuMilliLimit, not even an explicit nil
func (o *NodePodInfoDto) UnsetCpuMilliLimit() {
	o.CpuMilliLimit.Unset()
}

// GetCpuMilliRequest returns the CpuMilliRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetCpuMilliRequest() int32 {
	if o == nil || IsNil(o.CpuMilliRequest.Get()) {
		var ret int32
		return ret
	}
	return *o.CpuMilliRequest.Get()
}

// GetCpuMilliRequestOk returns a tuple with the CpuMilliRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetCpuMilliRequestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuMilliRequest.Get(), o.CpuMilliRequest.IsSet()
}

// HasCpuMilliRequest returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasCpuMilliRequest() bool {
	if o != nil && o.CpuMilliRequest.IsSet() {
		return true
	}

	return false
}

// SetCpuMilliRequest gets a reference to the given NullableInt32 and assigns it to the CpuMilliRequest field.
func (o *NodePodInfoDto) SetCpuMilliRequest(v int32) {
	o.CpuMilliRequest.Set(&v)
}
// SetCpuMilliRequestNil sets the value for CpuMilliRequest to be an explicit nil
func (o *NodePodInfoDto) SetCpuMilliRequestNil() {
	o.CpuMilliRequest.Set(nil)
}

// UnsetCpuMilliRequest ensures that no value is present for CpuMilliRequest, not even an explicit nil
func (o *NodePodInfoDto) UnsetCpuMilliRequest() {
	o.CpuMilliRequest.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *NodePodInfoDto) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}
// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *NodePodInfoDto) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *NodePodInfoDto) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetImagesVersion returns the ImagesVersion field value
func (o *NodePodInfoDto) GetImagesVersion() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.ImagesVersion
}

// GetImagesVersionOk returns a tuple with the ImagesVersion field value
// and a boolean to check if the value has been set.
func (o *NodePodInfoDto) GetImagesVersionOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImagesVersion, true
}

// SetImagesVersion sets field value
func (o *NodePodInfoDto) SetImagesVersion(v map[string]string) {
	o.ImagesVersion = v
}

// GetMemoryMibLimit returns the MemoryMibLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetMemoryMibLimit() int32 {
	if o == nil || IsNil(o.MemoryMibLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.MemoryMibLimit.Get()
}

// GetMemoryMibLimitOk returns a tuple with the MemoryMibLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetMemoryMibLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryMibLimit.Get(), o.MemoryMibLimit.IsSet()
}

// HasMemoryMibLimit returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasMemoryMibLimit() bool {
	if o != nil && o.MemoryMibLimit.IsSet() {
		return true
	}

	return false
}

// SetMemoryMibLimit gets a reference to the given NullableInt32 and assigns it to the MemoryMibLimit field.
func (o *NodePodInfoDto) SetMemoryMibLimit(v int32) {
	o.MemoryMibLimit.Set(&v)
}
// SetMemoryMibLimitNil sets the value for MemoryMibLimit to be an explicit nil
func (o *NodePodInfoDto) SetMemoryMibLimitNil() {
	o.MemoryMibLimit.Set(nil)
}

// UnsetMemoryMibLimit ensures that no value is present for MemoryMibLimit, not even an explicit nil
func (o *NodePodInfoDto) UnsetMemoryMibLimit() {
	o.MemoryMibLimit.Unset()
}

// GetMemoryMibRequest returns the MemoryMibRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetMemoryMibRequest() int32 {
	if o == nil || IsNil(o.MemoryMibRequest.Get()) {
		var ret int32
		return ret
	}
	return *o.MemoryMibRequest.Get()
}

// GetMemoryMibRequestOk returns a tuple with the MemoryMibRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetMemoryMibRequestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryMibRequest.Get(), o.MemoryMibRequest.IsSet()
}

// HasMemoryMibRequest returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasMemoryMibRequest() bool {
	if o != nil && o.MemoryMibRequest.IsSet() {
		return true
	}

	return false
}

// SetMemoryMibRequest gets a reference to the given NullableInt32 and assigns it to the MemoryMibRequest field.
func (o *NodePodInfoDto) SetMemoryMibRequest(v int32) {
	o.MemoryMibRequest.Set(&v)
}
// SetMemoryMibRequestNil sets the value for MemoryMibRequest to be an explicit nil
func (o *NodePodInfoDto) SetMemoryMibRequestNil() {
	o.MemoryMibRequest.Set(nil)
}

// UnsetMemoryMibRequest ensures that no value is present for MemoryMibRequest, not even an explicit nil
func (o *NodePodInfoDto) UnsetMemoryMibRequest() {
	o.MemoryMibRequest.Unset()
}

// GetName returns the Name field value
func (o *NodePodInfoDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodePodInfoDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodePodInfoDto) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *NodePodInfoDto) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *NodePodInfoDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *NodePodInfoDto) SetNamespace(v string) {
	o.Namespace = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *NodePodInfoDto) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *NodePodInfoDto) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *NodePodInfoDto) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodePodInfoDto) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceId.Get()
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodePodInfoDto) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceId.Get(), o.ServiceId.IsSet()
}

// HasServiceId returns a boolean if a field has been set.
func (o *NodePodInfoDto) HasServiceId() bool {
	if o != nil && o.ServiceId.IsSet() {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given NullableString and assigns it to the ServiceId field.
func (o *NodePodInfoDto) SetServiceId(v string) {
	o.ServiceId.Set(&v)
}
// SetServiceIdNil sets the value for ServiceId to be an explicit nil
func (o *NodePodInfoDto) SetServiceIdNil() {
	o.ServiceId.Set(nil)
}

// UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
func (o *NodePodInfoDto) UnsetServiceId() {
	o.ServiceId.Unset()
}

func (o NodePodInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodePodInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuMilliLimit.IsSet() {
		toSerialize["cpu_milli_limit"] = o.CpuMilliLimit.Get()
	}
	if o.CpuMilliRequest.IsSet() {
		toSerialize["cpu_milli_request"] = o.CpuMilliRequest.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	toSerialize["images_version"] = o.ImagesVersion
	if o.MemoryMibLimit.IsSet() {
		toSerialize["memory_mib_limit"] = o.MemoryMibLimit.Get()
	}
	if o.MemoryMibRequest.IsSet() {
		toSerialize["memory_mib_request"] = o.MemoryMibRequest.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if o.ServiceId.IsSet() {
		toSerialize["service_id"] = o.ServiceId.Get()
	}
	return toSerialize, nil
}

type NullableNodePodInfoDto struct {
	value *NodePodInfoDto
	isSet bool
}

func (v NullableNodePodInfoDto) Get() *NodePodInfoDto {
	return v.value
}

func (v *NullableNodePodInfoDto) Set(val *NodePodInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePodInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePodInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePodInfoDto(val *NodePodInfoDto) *NullableNodePodInfoDto {
	return &NullableNodePodInfoDto{value: val, isSet: true}
}

func (v NullableNodePodInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePodInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

