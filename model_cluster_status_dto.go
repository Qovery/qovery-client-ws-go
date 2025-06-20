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

// checks if the ClusterStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterStatusDto{}

// ClusterStatusDto struct for ClusterStatusDto
type ClusterStatusDto struct {
	ComputedStatus ClusterComputedStatusDto `json:"computed_status"`
	TlsCertificate NullableCertificateStatusDto `json:"tls_certificate,omitempty"`
}

type _ClusterStatusDto ClusterStatusDto

// NewClusterStatusDto instantiates a new ClusterStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatusDto(computedStatus ClusterComputedStatusDto) *ClusterStatusDto {
	this := ClusterStatusDto{}
	this.ComputedStatus = computedStatus
	return &this
}

// NewClusterStatusDtoWithDefaults instantiates a new ClusterStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusDtoWithDefaults() *ClusterStatusDto {
	this := ClusterStatusDto{}
	return &this
}

// GetComputedStatus returns the ComputedStatus field value
func (o *ClusterStatusDto) GetComputedStatus() ClusterComputedStatusDto {
	if o == nil {
		var ret ClusterComputedStatusDto
		return ret
	}

	return o.ComputedStatus
}

// GetComputedStatusOk returns a tuple with the ComputedStatus field value
// and a boolean to check if the value has been set.
func (o *ClusterStatusDto) GetComputedStatusOk() (*ClusterComputedStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputedStatus, true
}

// SetComputedStatus sets field value
func (o *ClusterStatusDto) SetComputedStatus(v ClusterComputedStatusDto) {
	o.ComputedStatus = v
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterStatusDto) GetTlsCertificate() CertificateStatusDto {
	if o == nil || IsNil(o.TlsCertificate.Get()) {
		var ret CertificateStatusDto
		return ret
	}
	return *o.TlsCertificate.Get()
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterStatusDto) GetTlsCertificateOk() (*CertificateStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCertificate.Get(), o.TlsCertificate.IsSet()
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *ClusterStatusDto) HasTlsCertificate() bool {
	if o != nil && o.TlsCertificate.IsSet() {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given NullableCertificateStatusDto and assigns it to the TlsCertificate field.
func (o *ClusterStatusDto) SetTlsCertificate(v CertificateStatusDto) {
	o.TlsCertificate.Set(&v)
}
// SetTlsCertificateNil sets the value for TlsCertificate to be an explicit nil
func (o *ClusterStatusDto) SetTlsCertificateNil() {
	o.TlsCertificate.Set(nil)
}

// UnsetTlsCertificate ensures that no value is present for TlsCertificate, not even an explicit nil
func (o *ClusterStatusDto) UnsetTlsCertificate() {
	o.TlsCertificate.Unset()
}

func (o ClusterStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["computed_status"] = o.ComputedStatus
	if o.TlsCertificate.IsSet() {
		toSerialize["tls_certificate"] = o.TlsCertificate.Get()
	}
	return toSerialize, nil
}

func (o *ClusterStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"computed_status",
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

	varClusterStatusDto := _ClusterStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClusterStatusDto)

	if err != nil {
		return err
	}

	*o = ClusterStatusDto(varClusterStatusDto)

	return err
}

type NullableClusterStatusDto struct {
	value *ClusterStatusDto
	isSet bool
}

func (v NullableClusterStatusDto) Get() *ClusterStatusDto {
	return v.value
}

func (v *NullableClusterStatusDto) Set(val *ClusterStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatusDto(val *ClusterStatusDto) *NullableClusterStatusDto {
	return &NullableClusterStatusDto{value: val, isSet: true}
}

func (v NullableClusterStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


