# ApplicationStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]CertificateStatusDto**](CertificateStatusDto.md) |  | 
**Id** | **string** |  | 
**Pods** | [**[]PodStatusDto**](PodStatusDto.md) |  | 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 

## Methods

### NewApplicationStatusDto

`func NewApplicationStatusDto(certificates []CertificateStatusDto, id string, pods []PodStatusDto, state ServiceStateDto, ) *ApplicationStatusDto`

NewApplicationStatusDto instantiates a new ApplicationStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStatusDtoWithDefaults

`func NewApplicationStatusDtoWithDefaults() *ApplicationStatusDto`

NewApplicationStatusDtoWithDefaults instantiates a new ApplicationStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *ApplicationStatusDto) GetCertificates() []CertificateStatusDto`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ApplicationStatusDto) GetCertificatesOk() (*[]CertificateStatusDto, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ApplicationStatusDto) SetCertificates(v []CertificateStatusDto)`

SetCertificates sets Certificates field to given value.


### GetId

`func (o *ApplicationStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationStatusDto) SetId(v string)`

SetId sets Id field to given value.


### GetPods

`func (o *ApplicationStatusDto) GetPods() []PodStatusDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ApplicationStatusDto) GetPodsOk() (*[]PodStatusDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ApplicationStatusDto) SetPods(v []PodStatusDto)`

SetPods sets Pods field to given value.


### GetState

`func (o *ApplicationStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApplicationStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


