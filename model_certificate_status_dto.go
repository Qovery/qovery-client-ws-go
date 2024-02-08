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
	"fmt"
)

// checks if the CertificateStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateStatusDto{}

// CertificateStatusDto struct for CertificateStatusDto
type CertificateStatusDto struct {
	DnsNames []string `json:"dns_names"`
	FailedIssuanceAttemptCount int64 `json:"failed_issuance_attempt_count"`
	// Unix timestamp with millisecond precision
	LastFailureIssuanceTime NullableInt32 `json:"last_failure_issuance_time,omitempty"`
	// Unix timestamp with millisecond precision
	NotAfter NullableInt32 `json:"not_after,omitempty"`
	// Unix timestamp with millisecond precision
	NotBefore NullableInt32 `json:"not_before,omitempty"`
	// Unix timestamp with millisecond precision
	RenewalTime NullableInt32 `json:"renewal_time,omitempty"`
	State ServiceStateDto `json:"state"`
	StateMessage NullableString `json:"state_message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateStatusDto CertificateStatusDto

// NewCertificateStatusDto instantiates a new CertificateStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateStatusDto(dnsNames []string, failedIssuanceAttemptCount int64, state ServiceStateDto) *CertificateStatusDto {
	this := CertificateStatusDto{}
	this.DnsNames = dnsNames
	this.FailedIssuanceAttemptCount = failedIssuanceAttemptCount
	this.State = state
	return &this
}

// NewCertificateStatusDtoWithDefaults instantiates a new CertificateStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateStatusDtoWithDefaults() *CertificateStatusDto {
	this := CertificateStatusDto{}
	return &this
}

// GetDnsNames returns the DnsNames field value
func (o *CertificateStatusDto) GetDnsNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value
// and a boolean to check if the value has been set.
func (o *CertificateStatusDto) GetDnsNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsNames, true
}

// SetDnsNames sets field value
func (o *CertificateStatusDto) SetDnsNames(v []string) {
	o.DnsNames = v
}

// GetFailedIssuanceAttemptCount returns the FailedIssuanceAttemptCount field value
func (o *CertificateStatusDto) GetFailedIssuanceAttemptCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedIssuanceAttemptCount
}

// GetFailedIssuanceAttemptCountOk returns a tuple with the FailedIssuanceAttemptCount field value
// and a boolean to check if the value has been set.
func (o *CertificateStatusDto) GetFailedIssuanceAttemptCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedIssuanceAttemptCount, true
}

// SetFailedIssuanceAttemptCount sets field value
func (o *CertificateStatusDto) SetFailedIssuanceAttemptCount(v int64) {
	o.FailedIssuanceAttemptCount = v
}

// GetLastFailureIssuanceTime returns the LastFailureIssuanceTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateStatusDto) GetLastFailureIssuanceTime() int32 {
	if o == nil || IsNil(o.LastFailureIssuanceTime.Get()) {
		var ret int32
		return ret
	}
	return *o.LastFailureIssuanceTime.Get()
}

// GetLastFailureIssuanceTimeOk returns a tuple with the LastFailureIssuanceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateStatusDto) GetLastFailureIssuanceTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFailureIssuanceTime.Get(), o.LastFailureIssuanceTime.IsSet()
}

// HasLastFailureIssuanceTime returns a boolean if a field has been set.
func (o *CertificateStatusDto) HasLastFailureIssuanceTime() bool {
	if o != nil && o.LastFailureIssuanceTime.IsSet() {
		return true
	}

	return false
}

// SetLastFailureIssuanceTime gets a reference to the given NullableInt32 and assigns it to the LastFailureIssuanceTime field.
func (o *CertificateStatusDto) SetLastFailureIssuanceTime(v int32) {
	o.LastFailureIssuanceTime.Set(&v)
}
// SetLastFailureIssuanceTimeNil sets the value for LastFailureIssuanceTime to be an explicit nil
func (o *CertificateStatusDto) SetLastFailureIssuanceTimeNil() {
	o.LastFailureIssuanceTime.Set(nil)
}

// UnsetLastFailureIssuanceTime ensures that no value is present for LastFailureIssuanceTime, not even an explicit nil
func (o *CertificateStatusDto) UnsetLastFailureIssuanceTime() {
	o.LastFailureIssuanceTime.Unset()
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateStatusDto) GetNotAfter() int32 {
	if o == nil || IsNil(o.NotAfter.Get()) {
		var ret int32
		return ret
	}
	return *o.NotAfter.Get()
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateStatusDto) GetNotAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotAfter.Get(), o.NotAfter.IsSet()
}

// HasNotAfter returns a boolean if a field has been set.
func (o *CertificateStatusDto) HasNotAfter() bool {
	if o != nil && o.NotAfter.IsSet() {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given NullableInt32 and assigns it to the NotAfter field.
func (o *CertificateStatusDto) SetNotAfter(v int32) {
	o.NotAfter.Set(&v)
}
// SetNotAfterNil sets the value for NotAfter to be an explicit nil
func (o *CertificateStatusDto) SetNotAfterNil() {
	o.NotAfter.Set(nil)
}

// UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
func (o *CertificateStatusDto) UnsetNotAfter() {
	o.NotAfter.Unset()
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateStatusDto) GetNotBefore() int32 {
	if o == nil || IsNil(o.NotBefore.Get()) {
		var ret int32
		return ret
	}
	return *o.NotBefore.Get()
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateStatusDto) GetNotBeforeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotBefore.Get(), o.NotBefore.IsSet()
}

// HasNotBefore returns a boolean if a field has been set.
func (o *CertificateStatusDto) HasNotBefore() bool {
	if o != nil && o.NotBefore.IsSet() {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given NullableInt32 and assigns it to the NotBefore field.
func (o *CertificateStatusDto) SetNotBefore(v int32) {
	o.NotBefore.Set(&v)
}
// SetNotBeforeNil sets the value for NotBefore to be an explicit nil
func (o *CertificateStatusDto) SetNotBeforeNil() {
	o.NotBefore.Set(nil)
}

// UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
func (o *CertificateStatusDto) UnsetNotBefore() {
	o.NotBefore.Unset()
}

// GetRenewalTime returns the RenewalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateStatusDto) GetRenewalTime() int32 {
	if o == nil || IsNil(o.RenewalTime.Get()) {
		var ret int32
		return ret
	}
	return *o.RenewalTime.Get()
}

// GetRenewalTimeOk returns a tuple with the RenewalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateStatusDto) GetRenewalTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalTime.Get(), o.RenewalTime.IsSet()
}

// HasRenewalTime returns a boolean if a field has been set.
func (o *CertificateStatusDto) HasRenewalTime() bool {
	if o != nil && o.RenewalTime.IsSet() {
		return true
	}

	return false
}

// SetRenewalTime gets a reference to the given NullableInt32 and assigns it to the RenewalTime field.
func (o *CertificateStatusDto) SetRenewalTime(v int32) {
	o.RenewalTime.Set(&v)
}
// SetRenewalTimeNil sets the value for RenewalTime to be an explicit nil
func (o *CertificateStatusDto) SetRenewalTimeNil() {
	o.RenewalTime.Set(nil)
}

// UnsetRenewalTime ensures that no value is present for RenewalTime, not even an explicit nil
func (o *CertificateStatusDto) UnsetRenewalTime() {
	o.RenewalTime.Unset()
}

// GetState returns the State field value
func (o *CertificateStatusDto) GetState() ServiceStateDto {
	if o == nil {
		var ret ServiceStateDto
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CertificateStatusDto) GetStateOk() (*ServiceStateDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CertificateStatusDto) SetState(v ServiceStateDto) {
	o.State = v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateStatusDto) GetStateMessage() string {
	if o == nil || IsNil(o.StateMessage.Get()) {
		var ret string
		return ret
	}
	return *o.StateMessage.Get()
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateStatusDto) GetStateMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateMessage.Get(), o.StateMessage.IsSet()
}

// HasStateMessage returns a boolean if a field has been set.
func (o *CertificateStatusDto) HasStateMessage() bool {
	if o != nil && o.StateMessage.IsSet() {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given NullableString and assigns it to the StateMessage field.
func (o *CertificateStatusDto) SetStateMessage(v string) {
	o.StateMessage.Set(&v)
}
// SetStateMessageNil sets the value for StateMessage to be an explicit nil
func (o *CertificateStatusDto) SetStateMessageNil() {
	o.StateMessage.Set(nil)
}

// UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil
func (o *CertificateStatusDto) UnsetStateMessage() {
	o.StateMessage.Unset()
}

func (o CertificateStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dns_names"] = o.DnsNames
	toSerialize["failed_issuance_attempt_count"] = o.FailedIssuanceAttemptCount
	if o.LastFailureIssuanceTime.IsSet() {
		toSerialize["last_failure_issuance_time"] = o.LastFailureIssuanceTime.Get()
	}
	if o.NotAfter.IsSet() {
		toSerialize["not_after"] = o.NotAfter.Get()
	}
	if o.NotBefore.IsSet() {
		toSerialize["not_before"] = o.NotBefore.Get()
	}
	if o.RenewalTime.IsSet() {
		toSerialize["renewal_time"] = o.RenewalTime.Get()
	}
	toSerialize["state"] = o.State
	if o.StateMessage.IsSet() {
		toSerialize["state_message"] = o.StateMessage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dns_names",
		"failed_issuance_attempt_count",
		"state",
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

	varCertificateStatusDto := _CertificateStatusDto{}

	err = json.Unmarshal(data, &varCertificateStatusDto)

	if err != nil {
		return err
	}

	*o = CertificateStatusDto(varCertificateStatusDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_names")
		delete(additionalProperties, "failed_issuance_attempt_count")
		delete(additionalProperties, "last_failure_issuance_time")
		delete(additionalProperties, "not_after")
		delete(additionalProperties, "not_before")
		delete(additionalProperties, "renewal_time")
		delete(additionalProperties, "state")
		delete(additionalProperties, "state_message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateStatusDto struct {
	value *CertificateStatusDto
	isSet bool
}

func (v NullableCertificateStatusDto) Get() *CertificateStatusDto {
	return v.value
}

func (v *NullableCertificateStatusDto) Set(val *CertificateStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateStatusDto(val *CertificateStatusDto) *NullableCertificateStatusDto {
	return &NullableCertificateStatusDto{value: val, isSet: true}
}

func (v NullableCertificateStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


