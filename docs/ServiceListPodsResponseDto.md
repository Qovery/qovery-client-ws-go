# ServiceListPodsResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pods** | [**[]PodDto**](PodDto.md) |  | 

## Methods

### NewServiceListPodsResponseDto

`func NewServiceListPodsResponseDto(pods []PodDto, ) *ServiceListPodsResponseDto`

NewServiceListPodsResponseDto instantiates a new ServiceListPodsResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListPodsResponseDtoWithDefaults

`func NewServiceListPodsResponseDtoWithDefaults() *ServiceListPodsResponseDto`

NewServiceListPodsResponseDtoWithDefaults instantiates a new ServiceListPodsResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPods

`func (o *ServiceListPodsResponseDto) GetPods() []PodDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ServiceListPodsResponseDto) GetPodsOk() (*[]PodDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ServiceListPodsResponseDto) SetPods(v []PodDto)`

SetPods sets Pods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


