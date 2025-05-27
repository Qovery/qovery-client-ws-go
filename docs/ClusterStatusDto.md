# ClusterStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedStatus** | [**ClusterComputedStatusDto**](ClusterComputedStatusDto.md) |  | 
**TlsCertificate** | Pointer to [**NullableCertificateStatusDto**](CertificateStatusDto.md) |  | [optional] 

## Methods

### NewClusterStatusDto

`func NewClusterStatusDto(computedStatus ClusterComputedStatusDto, ) *ClusterStatusDto`

NewClusterStatusDto instantiates a new ClusterStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusDtoWithDefaults

`func NewClusterStatusDtoWithDefaults() *ClusterStatusDto`

NewClusterStatusDtoWithDefaults instantiates a new ClusterStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedStatus

`func (o *ClusterStatusDto) GetComputedStatus() ClusterComputedStatusDto`

GetComputedStatus returns the ComputedStatus field if non-nil, zero value otherwise.

### GetComputedStatusOk

`func (o *ClusterStatusDto) GetComputedStatusOk() (*ClusterComputedStatusDto, bool)`

GetComputedStatusOk returns a tuple with the ComputedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedStatus

`func (o *ClusterStatusDto) SetComputedStatus(v ClusterComputedStatusDto)`

SetComputedStatus sets ComputedStatus field to given value.


### GetTlsCertificate

`func (o *ClusterStatusDto) GetTlsCertificate() CertificateStatusDto`

GetTlsCertificate returns the TlsCertificate field if non-nil, zero value otherwise.

### GetTlsCertificateOk

`func (o *ClusterStatusDto) GetTlsCertificateOk() (*CertificateStatusDto, bool)`

GetTlsCertificateOk returns a tuple with the TlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificate

`func (o *ClusterStatusDto) SetTlsCertificate(v CertificateStatusDto)`

SetTlsCertificate sets TlsCertificate field to given value.

### HasTlsCertificate

`func (o *ClusterStatusDto) HasTlsCertificate() bool`

HasTlsCertificate returns a boolean if a field has been set.

### SetTlsCertificateNil

`func (o *ClusterStatusDto) SetTlsCertificateNil(b bool)`

 SetTlsCertificateNil sets the value for TlsCertificate to be an explicit nil

### UnsetTlsCertificate
`func (o *ClusterStatusDto) UnsetTlsCertificate()`

UnsetTlsCertificate ensures that no value is present for TlsCertificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


