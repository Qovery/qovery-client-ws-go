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

// checks if the ClusterComputedStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterComputedStatusDto{}

// ClusterComputedStatusDto struct for ClusterComputedStatusDto
type ClusterComputedStatusDto struct {
	GlobalStatus ClusterStatusGlobalStatus `json:"global_status"`
	IsMaxNodesSizeReached bool `json:"is_max_nodes_size_reached"`
	KubeVersionStatus QoveryClusterKubeVersionStatus `json:"kube_version_status"`
	NodeWarnings map[string][]QoveryNodeFailure `json:"node_warnings"`
	QoveryComponentsInFailure []QoveryComponentInFailure `json:"qovery_components_in_failure"`
}

type _ClusterComputedStatusDto ClusterComputedStatusDto

// NewClusterComputedStatusDto instantiates a new ClusterComputedStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterComputedStatusDto(globalStatus ClusterStatusGlobalStatus, isMaxNodesSizeReached bool, kubeVersionStatus QoveryClusterKubeVersionStatus, nodeWarnings map[string][]QoveryNodeFailure, qoveryComponentsInFailure []QoveryComponentInFailure) *ClusterComputedStatusDto {
	this := ClusterComputedStatusDto{}
	this.GlobalStatus = globalStatus
	this.IsMaxNodesSizeReached = isMaxNodesSizeReached
	this.KubeVersionStatus = kubeVersionStatus
	this.NodeWarnings = nodeWarnings
	this.QoveryComponentsInFailure = qoveryComponentsInFailure
	return &this
}

// NewClusterComputedStatusDtoWithDefaults instantiates a new ClusterComputedStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterComputedStatusDtoWithDefaults() *ClusterComputedStatusDto {
	this := ClusterComputedStatusDto{}
	return &this
}

// GetGlobalStatus returns the GlobalStatus field value
func (o *ClusterComputedStatusDto) GetGlobalStatus() ClusterStatusGlobalStatus {
	if o == nil {
		var ret ClusterStatusGlobalStatus
		return ret
	}

	return o.GlobalStatus
}

// GetGlobalStatusOk returns a tuple with the GlobalStatus field value
// and a boolean to check if the value has been set.
func (o *ClusterComputedStatusDto) GetGlobalStatusOk() (*ClusterStatusGlobalStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalStatus, true
}

// SetGlobalStatus sets field value
func (o *ClusterComputedStatusDto) SetGlobalStatus(v ClusterStatusGlobalStatus) {
	o.GlobalStatus = v
}

// GetIsMaxNodesSizeReached returns the IsMaxNodesSizeReached field value
func (o *ClusterComputedStatusDto) GetIsMaxNodesSizeReached() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMaxNodesSizeReached
}

// GetIsMaxNodesSizeReachedOk returns a tuple with the IsMaxNodesSizeReached field value
// and a boolean to check if the value has been set.
func (o *ClusterComputedStatusDto) GetIsMaxNodesSizeReachedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMaxNodesSizeReached, true
}

// SetIsMaxNodesSizeReached sets field value
func (o *ClusterComputedStatusDto) SetIsMaxNodesSizeReached(v bool) {
	o.IsMaxNodesSizeReached = v
}

// GetKubeVersionStatus returns the KubeVersionStatus field value
func (o *ClusterComputedStatusDto) GetKubeVersionStatus() QoveryClusterKubeVersionStatus {
	if o == nil {
		var ret QoveryClusterKubeVersionStatus
		return ret
	}

	return o.KubeVersionStatus
}

// GetKubeVersionStatusOk returns a tuple with the KubeVersionStatus field value
// and a boolean to check if the value has been set.
func (o *ClusterComputedStatusDto) GetKubeVersionStatusOk() (*QoveryClusterKubeVersionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeVersionStatus, true
}

// SetKubeVersionStatus sets field value
func (o *ClusterComputedStatusDto) SetKubeVersionStatus(v QoveryClusterKubeVersionStatus) {
	o.KubeVersionStatus = v
}

// GetNodeWarnings returns the NodeWarnings field value
func (o *ClusterComputedStatusDto) GetNodeWarnings() map[string][]QoveryNodeFailure {
	if o == nil {
		var ret map[string][]QoveryNodeFailure
		return ret
	}

	return o.NodeWarnings
}

// GetNodeWarningsOk returns a tuple with the NodeWarnings field value
// and a boolean to check if the value has been set.
func (o *ClusterComputedStatusDto) GetNodeWarningsOk() (map[string][]QoveryNodeFailure, bool) {
	if o == nil {
		return map[string][]QoveryNodeFailure{}, false
	}
	return o.NodeWarnings, true
}

// SetNodeWarnings sets field value
func (o *ClusterComputedStatusDto) SetNodeWarnings(v map[string][]QoveryNodeFailure) {
	o.NodeWarnings = v
}

// GetQoveryComponentsInFailure returns the QoveryComponentsInFailure field value
func (o *ClusterComputedStatusDto) GetQoveryComponentsInFailure() []QoveryComponentInFailure {
	if o == nil {
		var ret []QoveryComponentInFailure
		return ret
	}

	return o.QoveryComponentsInFailure
}

// GetQoveryComponentsInFailureOk returns a tuple with the QoveryComponentsInFailure field value
// and a boolean to check if the value has been set.
func (o *ClusterComputedStatusDto) GetQoveryComponentsInFailureOk() ([]QoveryComponentInFailure, bool) {
	if o == nil {
		return nil, false
	}
	return o.QoveryComponentsInFailure, true
}

// SetQoveryComponentsInFailure sets field value
func (o *ClusterComputedStatusDto) SetQoveryComponentsInFailure(v []QoveryComponentInFailure) {
	o.QoveryComponentsInFailure = v
}

func (o ClusterComputedStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterComputedStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["global_status"] = o.GlobalStatus
	toSerialize["is_max_nodes_size_reached"] = o.IsMaxNodesSizeReached
	toSerialize["kube_version_status"] = o.KubeVersionStatus
	toSerialize["node_warnings"] = o.NodeWarnings
	toSerialize["qovery_components_in_failure"] = o.QoveryComponentsInFailure
	return toSerialize, nil
}

func (o *ClusterComputedStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"global_status",
		"is_max_nodes_size_reached",
		"kube_version_status",
		"node_warnings",
		"qovery_components_in_failure",
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

	varClusterComputedStatusDto := _ClusterComputedStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterComputedStatusDto)

	if err != nil {
		return err
	}

	*o = ClusterComputedStatusDto(varClusterComputedStatusDto)

	return err
}

type NullableClusterComputedStatusDto struct {
	value *ClusterComputedStatusDto
	isSet bool
}

func (v NullableClusterComputedStatusDto) Get() *ClusterComputedStatusDto {
	return v.value
}

func (v *NullableClusterComputedStatusDto) Set(val *ClusterComputedStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterComputedStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterComputedStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterComputedStatusDto(val *ClusterComputedStatusDto) *NullableClusterComputedStatusDto {
	return &NullableClusterComputedStatusDto{value: val, isSet: true}
}

func (v NullableClusterComputedStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterComputedStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


