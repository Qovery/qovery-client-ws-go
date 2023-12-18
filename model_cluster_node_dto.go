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

// checks if the ClusterNodeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterNodeDto{}

// ClusterNodeDto struct for ClusterNodeDto
type ClusterNodeDto struct {
	Addresses []NodeAddressDto `json:"addresses"`
	Annotations map[string]string `json:"annotations"`
	Architecture string `json:"architecture"`
	Conditions []NodeConditionDto `json:"conditions"`
	KernelVersion string `json:"kernel_version"`
	KubeletVersion string `json:"kubelet_version"`
	Labels map[string]string `json:"labels"`
	Name string `json:"name"`
	OperatingSystem string `json:"operating_system"`
	OsImage string `json:"os_image"`
	Pods []NodePodInfoDto `json:"pods"`
	ResourcesAllocatable NodeResourceDto `json:"resources_allocatable"`
	Taints []NodeTaintDto `json:"taints"`
	Unschedulable bool `json:"unschedulable"`
}

// NewClusterNodeDto instantiates a new ClusterNodeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNodeDto(addresses []NodeAddressDto, annotations map[string]string, architecture string, conditions []NodeConditionDto, kernelVersion string, kubeletVersion string, labels map[string]string, name string, operatingSystem string, osImage string, pods []NodePodInfoDto, resourcesAllocatable NodeResourceDto, taints []NodeTaintDto, unschedulable bool) *ClusterNodeDto {
	this := ClusterNodeDto{}
	this.Addresses = addresses
	this.Annotations = annotations
	this.Architecture = architecture
	this.Conditions = conditions
	this.KernelVersion = kernelVersion
	this.KubeletVersion = kubeletVersion
	this.Labels = labels
	this.Name = name
	this.OperatingSystem = operatingSystem
	this.OsImage = osImage
	this.Pods = pods
	this.ResourcesAllocatable = resourcesAllocatable
	this.Taints = taints
	this.Unschedulable = unschedulable
	return &this
}

// NewClusterNodeDtoWithDefaults instantiates a new ClusterNodeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeDtoWithDefaults() *ClusterNodeDto {
	this := ClusterNodeDto{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *ClusterNodeDto) GetAddresses() []NodeAddressDto {
	if o == nil {
		var ret []NodeAddressDto
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetAddressesOk() ([]NodeAddressDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ClusterNodeDto) SetAddresses(v []NodeAddressDto) {
	o.Addresses = v
}

// GetAnnotations returns the Annotations field value
func (o *ClusterNodeDto) GetAnnotations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value
func (o *ClusterNodeDto) SetAnnotations(v map[string]string) {
	o.Annotations = v
}

// GetArchitecture returns the Architecture field value
func (o *ClusterNodeDto) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *ClusterNodeDto) SetArchitecture(v string) {
	o.Architecture = v
}

// GetConditions returns the Conditions field value
func (o *ClusterNodeDto) GetConditions() []NodeConditionDto {
	if o == nil {
		var ret []NodeConditionDto
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetConditionsOk() ([]NodeConditionDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *ClusterNodeDto) SetConditions(v []NodeConditionDto) {
	o.Conditions = v
}

// GetKernelVersion returns the KernelVersion field value
func (o *ClusterNodeDto) GetKernelVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetKernelVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KernelVersion, true
}

// SetKernelVersion sets field value
func (o *ClusterNodeDto) SetKernelVersion(v string) {
	o.KernelVersion = v
}

// GetKubeletVersion returns the KubeletVersion field value
func (o *ClusterNodeDto) GetKubeletVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetKubeletVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeletVersion, true
}

// SetKubeletVersion sets field value
func (o *ClusterNodeDto) SetKubeletVersion(v string) {
	o.KubeletVersion = v
}

// GetLabels returns the Labels field value
func (o *ClusterNodeDto) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *ClusterNodeDto) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *ClusterNodeDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterNodeDto) SetName(v string) {
	o.Name = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *ClusterNodeDto) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *ClusterNodeDto) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetOsImage returns the OsImage field value
func (o *ClusterNodeDto) GetOsImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsImage
}

// GetOsImageOk returns a tuple with the OsImage field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetOsImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsImage, true
}

// SetOsImage sets field value
func (o *ClusterNodeDto) SetOsImage(v string) {
	o.OsImage = v
}

// GetPods returns the Pods field value
func (o *ClusterNodeDto) GetPods() []NodePodInfoDto {
	if o == nil {
		var ret []NodePodInfoDto
		return ret
	}

	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetPodsOk() ([]NodePodInfoDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pods, true
}

// SetPods sets field value
func (o *ClusterNodeDto) SetPods(v []NodePodInfoDto) {
	o.Pods = v
}

// GetResourcesAllocatable returns the ResourcesAllocatable field value
func (o *ClusterNodeDto) GetResourcesAllocatable() NodeResourceDto {
	if o == nil {
		var ret NodeResourceDto
		return ret
	}

	return o.ResourcesAllocatable
}

// GetResourcesAllocatableOk returns a tuple with the ResourcesAllocatable field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetResourcesAllocatableOk() (*NodeResourceDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcesAllocatable, true
}

// SetResourcesAllocatable sets field value
func (o *ClusterNodeDto) SetResourcesAllocatable(v NodeResourceDto) {
	o.ResourcesAllocatable = v
}

// GetTaints returns the Taints field value
func (o *ClusterNodeDto) GetTaints() []NodeTaintDto {
	if o == nil {
		var ret []NodeTaintDto
		return ret
	}

	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetTaintsOk() ([]NodeTaintDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Taints, true
}

// SetTaints sets field value
func (o *ClusterNodeDto) SetTaints(v []NodeTaintDto) {
	o.Taints = v
}

// GetUnschedulable returns the Unschedulable field value
func (o *ClusterNodeDto) GetUnschedulable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Unschedulable
}

// GetUnschedulableOk returns a tuple with the Unschedulable field value
// and a boolean to check if the value has been set.
func (o *ClusterNodeDto) GetUnschedulableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unschedulable, true
}

// SetUnschedulable sets field value
func (o *ClusterNodeDto) SetUnschedulable(v bool) {
	o.Unschedulable = v
}

func (o ClusterNodeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterNodeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addresses"] = o.Addresses
	toSerialize["annotations"] = o.Annotations
	toSerialize["architecture"] = o.Architecture
	toSerialize["conditions"] = o.Conditions
	toSerialize["kernel_version"] = o.KernelVersion
	toSerialize["kubelet_version"] = o.KubeletVersion
	toSerialize["labels"] = o.Labels
	toSerialize["name"] = o.Name
	toSerialize["operating_system"] = o.OperatingSystem
	toSerialize["os_image"] = o.OsImage
	toSerialize["pods"] = o.Pods
	toSerialize["resources_allocatable"] = o.ResourcesAllocatable
	toSerialize["taints"] = o.Taints
	toSerialize["unschedulable"] = o.Unschedulable
	return toSerialize, nil
}

type NullableClusterNodeDto struct {
	value *ClusterNodeDto
	isSet bool
}

func (v NullableClusterNodeDto) Get() *ClusterNodeDto {
	return v.value
}

func (v *NullableClusterNodeDto) Set(val *ClusterNodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNodeDto(val *ClusterNodeDto) *NullableClusterNodeDto {
	return &NullableClusterNodeDto{value: val, isSet: true}
}

func (v NullableClusterNodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

