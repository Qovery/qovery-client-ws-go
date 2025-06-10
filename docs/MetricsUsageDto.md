# MetricsUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMilliUsage** | Pointer to **NullableInt32** |  | [optional] 
**CpuPercentUsage** | Pointer to **NullableInt32** |  | [optional] 
**EphemeralStoragePercentUsage** | Pointer to **NullableInt32** |  | [optional] 
**EphemeralStorageUsage** | Pointer to **NullableInt32** |  | [optional] 
**MemoryMibRssUsage** | Pointer to **NullableInt32** |  | [optional] 
**MemoryMibWorkingSetUsage** | Pointer to **NullableInt32** |  | [optional] 
**MemoryPercentRssUsage** | Pointer to **NullableInt32** |  | [optional] 
**MemoryPercentWorkingSetUsage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMetricsUsageDto

`func NewMetricsUsageDto() *MetricsUsageDto`

NewMetricsUsageDto instantiates a new MetricsUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsUsageDtoWithDefaults

`func NewMetricsUsageDtoWithDefaults() *MetricsUsageDto`

NewMetricsUsageDtoWithDefaults instantiates a new MetricsUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMilliUsage

`func (o *MetricsUsageDto) GetCpuMilliUsage() int32`

GetCpuMilliUsage returns the CpuMilliUsage field if non-nil, zero value otherwise.

### GetCpuMilliUsageOk

`func (o *MetricsUsageDto) GetCpuMilliUsageOk() (*int32, bool)`

GetCpuMilliUsageOk returns a tuple with the CpuMilliUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMilliUsage

`func (o *MetricsUsageDto) SetCpuMilliUsage(v int32)`

SetCpuMilliUsage sets CpuMilliUsage field to given value.

### HasCpuMilliUsage

`func (o *MetricsUsageDto) HasCpuMilliUsage() bool`

HasCpuMilliUsage returns a boolean if a field has been set.

### SetCpuMilliUsageNil

`func (o *MetricsUsageDto) SetCpuMilliUsageNil(b bool)`

 SetCpuMilliUsageNil sets the value for CpuMilliUsage to be an explicit nil

### UnsetCpuMilliUsage
`func (o *MetricsUsageDto) UnsetCpuMilliUsage()`

UnsetCpuMilliUsage ensures that no value is present for CpuMilliUsage, not even an explicit nil
### GetCpuPercentUsage

`func (o *MetricsUsageDto) GetCpuPercentUsage() int32`

GetCpuPercentUsage returns the CpuPercentUsage field if non-nil, zero value otherwise.

### GetCpuPercentUsageOk

`func (o *MetricsUsageDto) GetCpuPercentUsageOk() (*int32, bool)`

GetCpuPercentUsageOk returns a tuple with the CpuPercentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercentUsage

`func (o *MetricsUsageDto) SetCpuPercentUsage(v int32)`

SetCpuPercentUsage sets CpuPercentUsage field to given value.

### HasCpuPercentUsage

`func (o *MetricsUsageDto) HasCpuPercentUsage() bool`

HasCpuPercentUsage returns a boolean if a field has been set.

### SetCpuPercentUsageNil

`func (o *MetricsUsageDto) SetCpuPercentUsageNil(b bool)`

 SetCpuPercentUsageNil sets the value for CpuPercentUsage to be an explicit nil

### UnsetCpuPercentUsage
`func (o *MetricsUsageDto) UnsetCpuPercentUsage()`

UnsetCpuPercentUsage ensures that no value is present for CpuPercentUsage, not even an explicit nil
### GetEphemeralStoragePercentUsage

`func (o *MetricsUsageDto) GetEphemeralStoragePercentUsage() int32`

GetEphemeralStoragePercentUsage returns the EphemeralStoragePercentUsage field if non-nil, zero value otherwise.

### GetEphemeralStoragePercentUsageOk

`func (o *MetricsUsageDto) GetEphemeralStoragePercentUsageOk() (*int32, bool)`

GetEphemeralStoragePercentUsageOk returns a tuple with the EphemeralStoragePercentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStoragePercentUsage

`func (o *MetricsUsageDto) SetEphemeralStoragePercentUsage(v int32)`

SetEphemeralStoragePercentUsage sets EphemeralStoragePercentUsage field to given value.

### HasEphemeralStoragePercentUsage

`func (o *MetricsUsageDto) HasEphemeralStoragePercentUsage() bool`

HasEphemeralStoragePercentUsage returns a boolean if a field has been set.

### SetEphemeralStoragePercentUsageNil

`func (o *MetricsUsageDto) SetEphemeralStoragePercentUsageNil(b bool)`

 SetEphemeralStoragePercentUsageNil sets the value for EphemeralStoragePercentUsage to be an explicit nil

### UnsetEphemeralStoragePercentUsage
`func (o *MetricsUsageDto) UnsetEphemeralStoragePercentUsage()`

UnsetEphemeralStoragePercentUsage ensures that no value is present for EphemeralStoragePercentUsage, not even an explicit nil
### GetEphemeralStorageUsage

`func (o *MetricsUsageDto) GetEphemeralStorageUsage() int32`

GetEphemeralStorageUsage returns the EphemeralStorageUsage field if non-nil, zero value otherwise.

### GetEphemeralStorageUsageOk

`func (o *MetricsUsageDto) GetEphemeralStorageUsageOk() (*int32, bool)`

GetEphemeralStorageUsageOk returns a tuple with the EphemeralStorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageUsage

`func (o *MetricsUsageDto) SetEphemeralStorageUsage(v int32)`

SetEphemeralStorageUsage sets EphemeralStorageUsage field to given value.

### HasEphemeralStorageUsage

`func (o *MetricsUsageDto) HasEphemeralStorageUsage() bool`

HasEphemeralStorageUsage returns a boolean if a field has been set.

### SetEphemeralStorageUsageNil

`func (o *MetricsUsageDto) SetEphemeralStorageUsageNil(b bool)`

 SetEphemeralStorageUsageNil sets the value for EphemeralStorageUsage to be an explicit nil

### UnsetEphemeralStorageUsage
`func (o *MetricsUsageDto) UnsetEphemeralStorageUsage()`

UnsetEphemeralStorageUsage ensures that no value is present for EphemeralStorageUsage, not even an explicit nil
### GetMemoryMibRssUsage

`func (o *MetricsUsageDto) GetMemoryMibRssUsage() int32`

GetMemoryMibRssUsage returns the MemoryMibRssUsage field if non-nil, zero value otherwise.

### GetMemoryMibRssUsageOk

`func (o *MetricsUsageDto) GetMemoryMibRssUsageOk() (*int32, bool)`

GetMemoryMibRssUsageOk returns a tuple with the MemoryMibRssUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibRssUsage

`func (o *MetricsUsageDto) SetMemoryMibRssUsage(v int32)`

SetMemoryMibRssUsage sets MemoryMibRssUsage field to given value.

### HasMemoryMibRssUsage

`func (o *MetricsUsageDto) HasMemoryMibRssUsage() bool`

HasMemoryMibRssUsage returns a boolean if a field has been set.

### SetMemoryMibRssUsageNil

`func (o *MetricsUsageDto) SetMemoryMibRssUsageNil(b bool)`

 SetMemoryMibRssUsageNil sets the value for MemoryMibRssUsage to be an explicit nil

### UnsetMemoryMibRssUsage
`func (o *MetricsUsageDto) UnsetMemoryMibRssUsage()`

UnsetMemoryMibRssUsage ensures that no value is present for MemoryMibRssUsage, not even an explicit nil
### GetMemoryMibWorkingSetUsage

`func (o *MetricsUsageDto) GetMemoryMibWorkingSetUsage() int32`

GetMemoryMibWorkingSetUsage returns the MemoryMibWorkingSetUsage field if non-nil, zero value otherwise.

### GetMemoryMibWorkingSetUsageOk

`func (o *MetricsUsageDto) GetMemoryMibWorkingSetUsageOk() (*int32, bool)`

GetMemoryMibWorkingSetUsageOk returns a tuple with the MemoryMibWorkingSetUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMibWorkingSetUsage

`func (o *MetricsUsageDto) SetMemoryMibWorkingSetUsage(v int32)`

SetMemoryMibWorkingSetUsage sets MemoryMibWorkingSetUsage field to given value.

### HasMemoryMibWorkingSetUsage

`func (o *MetricsUsageDto) HasMemoryMibWorkingSetUsage() bool`

HasMemoryMibWorkingSetUsage returns a boolean if a field has been set.

### SetMemoryMibWorkingSetUsageNil

`func (o *MetricsUsageDto) SetMemoryMibWorkingSetUsageNil(b bool)`

 SetMemoryMibWorkingSetUsageNil sets the value for MemoryMibWorkingSetUsage to be an explicit nil

### UnsetMemoryMibWorkingSetUsage
`func (o *MetricsUsageDto) UnsetMemoryMibWorkingSetUsage()`

UnsetMemoryMibWorkingSetUsage ensures that no value is present for MemoryMibWorkingSetUsage, not even an explicit nil
### GetMemoryPercentRssUsage

`func (o *MetricsUsageDto) GetMemoryPercentRssUsage() int32`

GetMemoryPercentRssUsage returns the MemoryPercentRssUsage field if non-nil, zero value otherwise.

### GetMemoryPercentRssUsageOk

`func (o *MetricsUsageDto) GetMemoryPercentRssUsageOk() (*int32, bool)`

GetMemoryPercentRssUsageOk returns a tuple with the MemoryPercentRssUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPercentRssUsage

`func (o *MetricsUsageDto) SetMemoryPercentRssUsage(v int32)`

SetMemoryPercentRssUsage sets MemoryPercentRssUsage field to given value.

### HasMemoryPercentRssUsage

`func (o *MetricsUsageDto) HasMemoryPercentRssUsage() bool`

HasMemoryPercentRssUsage returns a boolean if a field has been set.

### SetMemoryPercentRssUsageNil

`func (o *MetricsUsageDto) SetMemoryPercentRssUsageNil(b bool)`

 SetMemoryPercentRssUsageNil sets the value for MemoryPercentRssUsage to be an explicit nil

### UnsetMemoryPercentRssUsage
`func (o *MetricsUsageDto) UnsetMemoryPercentRssUsage()`

UnsetMemoryPercentRssUsage ensures that no value is present for MemoryPercentRssUsage, not even an explicit nil
### GetMemoryPercentWorkingSetUsage

`func (o *MetricsUsageDto) GetMemoryPercentWorkingSetUsage() int32`

GetMemoryPercentWorkingSetUsage returns the MemoryPercentWorkingSetUsage field if non-nil, zero value otherwise.

### GetMemoryPercentWorkingSetUsageOk

`func (o *MetricsUsageDto) GetMemoryPercentWorkingSetUsageOk() (*int32, bool)`

GetMemoryPercentWorkingSetUsageOk returns a tuple with the MemoryPercentWorkingSetUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPercentWorkingSetUsage

`func (o *MetricsUsageDto) SetMemoryPercentWorkingSetUsage(v int32)`

SetMemoryPercentWorkingSetUsage sets MemoryPercentWorkingSetUsage field to given value.

### HasMemoryPercentWorkingSetUsage

`func (o *MetricsUsageDto) HasMemoryPercentWorkingSetUsage() bool`

HasMemoryPercentWorkingSetUsage returns a boolean if a field has been set.

### SetMemoryPercentWorkingSetUsageNil

`func (o *MetricsUsageDto) SetMemoryPercentWorkingSetUsageNil(b bool)`

 SetMemoryPercentWorkingSetUsageNil sets the value for MemoryPercentWorkingSetUsage to be an explicit nil

### UnsetMemoryPercentWorkingSetUsage
`func (o *MetricsUsageDto) UnsetMemoryPercentWorkingSetUsage()`

UnsetMemoryPercentWorkingSetUsage ensures that no value is present for MemoryPercentWorkingSetUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


