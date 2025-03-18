# ServiceMetricsV2Dto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pods** | [**[]PodInfoDto**](PodInfoDto.md) |  | 
**Pvcs** | [**[]PvcInfoDto**](PvcInfoDto.md) |  | 

## Methods

### NewServiceMetricsV2Dto

`func NewServiceMetricsV2Dto(pods []PodInfoDto, pvcs []PvcInfoDto, ) *ServiceMetricsV2Dto`

NewServiceMetricsV2Dto instantiates a new ServiceMetricsV2Dto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetricsV2DtoWithDefaults

`func NewServiceMetricsV2DtoWithDefaults() *ServiceMetricsV2Dto`

NewServiceMetricsV2DtoWithDefaults instantiates a new ServiceMetricsV2Dto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPods

`func (o *ServiceMetricsV2Dto) GetPods() []PodInfoDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ServiceMetricsV2Dto) GetPodsOk() (*[]PodInfoDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ServiceMetricsV2Dto) SetPods(v []PodInfoDto)`

SetPods sets Pods field to given value.


### GetPvcs

`func (o *ServiceMetricsV2Dto) GetPvcs() []PvcInfoDto`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *ServiceMetricsV2Dto) GetPvcsOk() (*[]PvcInfoDto, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *ServiceMetricsV2Dto) SetPvcs(v []PvcInfoDto)`

SetPvcs sets Pvcs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


