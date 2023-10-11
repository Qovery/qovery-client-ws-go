# CertificateStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsNames** | **[]string** |  | 
**FailedIssuanceAttemptCount** | **int64** |  | 
**LastFailureIssuanceTime** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**NotAfter** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**NotBefore** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**RenewalTime** | Pointer to **NullableInt32** | Unix timestamp with millisecond precision | [optional] 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 
**StateMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateStatusDto

`func NewCertificateStatusDto(dnsNames []string, failedIssuanceAttemptCount int64, state ServiceStateDto, ) *CertificateStatusDto`

NewCertificateStatusDto instantiates a new CertificateStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStatusDtoWithDefaults

`func NewCertificateStatusDtoWithDefaults() *CertificateStatusDto`

NewCertificateStatusDtoWithDefaults instantiates a new CertificateStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsNames

`func (o *CertificateStatusDto) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *CertificateStatusDto) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *CertificateStatusDto) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.


### GetFailedIssuanceAttemptCount

`func (o *CertificateStatusDto) GetFailedIssuanceAttemptCount() int64`

GetFailedIssuanceAttemptCount returns the FailedIssuanceAttemptCount field if non-nil, zero value otherwise.

### GetFailedIssuanceAttemptCountOk

`func (o *CertificateStatusDto) GetFailedIssuanceAttemptCountOk() (*int64, bool)`

GetFailedIssuanceAttemptCountOk returns a tuple with the FailedIssuanceAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIssuanceAttemptCount

`func (o *CertificateStatusDto) SetFailedIssuanceAttemptCount(v int64)`

SetFailedIssuanceAttemptCount sets FailedIssuanceAttemptCount field to given value.


### GetLastFailureIssuanceTime

`func (o *CertificateStatusDto) GetLastFailureIssuanceTime() int32`

GetLastFailureIssuanceTime returns the LastFailureIssuanceTime field if non-nil, zero value otherwise.

### GetLastFailureIssuanceTimeOk

`func (o *CertificateStatusDto) GetLastFailureIssuanceTimeOk() (*int32, bool)`

GetLastFailureIssuanceTimeOk returns a tuple with the LastFailureIssuanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureIssuanceTime

`func (o *CertificateStatusDto) SetLastFailureIssuanceTime(v int32)`

SetLastFailureIssuanceTime sets LastFailureIssuanceTime field to given value.

### HasLastFailureIssuanceTime

`func (o *CertificateStatusDto) HasLastFailureIssuanceTime() bool`

HasLastFailureIssuanceTime returns a boolean if a field has been set.

### SetLastFailureIssuanceTimeNil

`func (o *CertificateStatusDto) SetLastFailureIssuanceTimeNil(b bool)`

 SetLastFailureIssuanceTimeNil sets the value for LastFailureIssuanceTime to be an explicit nil

### UnsetLastFailureIssuanceTime
`func (o *CertificateStatusDto) UnsetLastFailureIssuanceTime()`

UnsetLastFailureIssuanceTime ensures that no value is present for LastFailureIssuanceTime, not even an explicit nil
### GetNotAfter

`func (o *CertificateStatusDto) GetNotAfter() int32`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateStatusDto) GetNotAfterOk() (*int32, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateStatusDto) SetNotAfter(v int32)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateStatusDto) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *CertificateStatusDto) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *CertificateStatusDto) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
### GetNotBefore

`func (o *CertificateStatusDto) GetNotBefore() int32`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateStatusDto) GetNotBeforeOk() (*int32, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateStatusDto) SetNotBefore(v int32)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateStatusDto) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### SetNotBeforeNil

`func (o *CertificateStatusDto) SetNotBeforeNil(b bool)`

 SetNotBeforeNil sets the value for NotBefore to be an explicit nil

### UnsetNotBefore
`func (o *CertificateStatusDto) UnsetNotBefore()`

UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
### GetRenewalTime

`func (o *CertificateStatusDto) GetRenewalTime() int32`

GetRenewalTime returns the RenewalTime field if non-nil, zero value otherwise.

### GetRenewalTimeOk

`func (o *CertificateStatusDto) GetRenewalTimeOk() (*int32, bool)`

GetRenewalTimeOk returns a tuple with the RenewalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalTime

`func (o *CertificateStatusDto) SetRenewalTime(v int32)`

SetRenewalTime sets RenewalTime field to given value.

### HasRenewalTime

`func (o *CertificateStatusDto) HasRenewalTime() bool`

HasRenewalTime returns a boolean if a field has been set.

### SetRenewalTimeNil

`func (o *CertificateStatusDto) SetRenewalTimeNil(b bool)`

 SetRenewalTimeNil sets the value for RenewalTime to be an explicit nil

### UnsetRenewalTime
`func (o *CertificateStatusDto) UnsetRenewalTime()`

UnsetRenewalTime ensures that no value is present for RenewalTime, not even an explicit nil
### GetState

`func (o *CertificateStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificateStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificateStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.


### GetStateMessage

`func (o *CertificateStatusDto) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *CertificateStatusDto) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *CertificateStatusDto) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *CertificateStatusDto) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### SetStateMessageNil

`func (o *CertificateStatusDto) SetStateMessageNil(b bool)`

 SetStateMessageNil sets the value for StateMessage to be an explicit nil

### UnsetStateMessage
`func (o *CertificateStatusDto) UnsetStateMessage()`

UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


