# CertificateStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsNames** | **[]string** |  | 
**FailedIssuanceAttemptCount** | **int64** |  | 
**LastFailureIssuanceTime** | **int64** |  | 
**NotAfter** | **int64** |  | 
**NotBefore** | **int64** |  | 
**RenewalTime** | **int64** |  | 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 
**StateMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateStatusDto

`func NewCertificateStatusDto(dnsNames []string, failedIssuanceAttemptCount int64, lastFailureIssuanceTime int64, notAfter int64, notBefore int64, renewalTime int64, state ServiceStateDto, ) *CertificateStatusDto`

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

`func (o *CertificateStatusDto) GetLastFailureIssuanceTime() int64`

GetLastFailureIssuanceTime returns the LastFailureIssuanceTime field if non-nil, zero value otherwise.

### GetLastFailureIssuanceTimeOk

`func (o *CertificateStatusDto) GetLastFailureIssuanceTimeOk() (*int64, bool)`

GetLastFailureIssuanceTimeOk returns a tuple with the LastFailureIssuanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureIssuanceTime

`func (o *CertificateStatusDto) SetLastFailureIssuanceTime(v int64)`

SetLastFailureIssuanceTime sets LastFailureIssuanceTime field to given value.


### GetNotAfter

`func (o *CertificateStatusDto) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateStatusDto) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateStatusDto) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.


### GetNotBefore

`func (o *CertificateStatusDto) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateStatusDto) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateStatusDto) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.


### GetRenewalTime

`func (o *CertificateStatusDto) GetRenewalTime() int64`

GetRenewalTime returns the RenewalTime field if non-nil, zero value otherwise.

### GetRenewalTimeOk

`func (o *CertificateStatusDto) GetRenewalTimeOk() (*int64, bool)`

GetRenewalTimeOk returns a tuple with the RenewalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalTime

`func (o *CertificateStatusDto) SetRenewalTime(v int64)`

SetRenewalTime sets RenewalTime field to given value.


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


