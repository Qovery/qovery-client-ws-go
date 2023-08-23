# ServiceMetricsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | [**MetricDto**](MetricDto.md) |  | 
**Memory** | [**MetricDto**](MetricDto.md) |  | 
**PodName** | **string** |  | 
**Storages** | [**[]MetricDto**](MetricDto.md) |  | 

## Methods

### NewServiceMetricsDto

`func NewServiceMetricsDto(cpu MetricDto, memory MetricDto, podName string, storages []MetricDto, ) *ServiceMetricsDto`

NewServiceMetricsDto instantiates a new ServiceMetricsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetricsDtoWithDefaults

`func NewServiceMetricsDtoWithDefaults() *ServiceMetricsDto`

NewServiceMetricsDtoWithDefaults instantiates a new ServiceMetricsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ServiceMetricsDto) GetCpu() MetricDto`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServiceMetricsDto) GetCpuOk() (*MetricDto, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServiceMetricsDto) SetCpu(v MetricDto)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ServiceMetricsDto) GetMemory() MetricDto`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ServiceMetricsDto) GetMemoryOk() (*MetricDto, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ServiceMetricsDto) SetMemory(v MetricDto)`

SetMemory sets Memory field to given value.


### GetPodName

`func (o *ServiceMetricsDto) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ServiceMetricsDto) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ServiceMetricsDto) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetStorages

`func (o *ServiceMetricsDto) GetStorages() []MetricDto`

GetStorages returns the Storages field if non-nil, zero value otherwise.

### GetStoragesOk

`func (o *ServiceMetricsDto) GetStoragesOk() (*[]MetricDto, bool)`

GetStoragesOk returns a tuple with the Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorages

`func (o *ServiceMetricsDto) SetStorages(v []MetricDto)`

SetStorages sets Storages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


