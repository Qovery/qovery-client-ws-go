# MetricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **int32** |  | 
**CurrentPercent** | **int32** |  | 
**Limit** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | [**ResourceStatusDto**](ResourceStatusDto.md) |  | 
**Unit** | [**UnitDto**](UnitDto.md) |  | 

## Methods

### NewMetricDto

`func NewMetricDto(current int32, currentPercent int32, limit int32, status ResourceStatusDto, unit UnitDto, ) *MetricDto`

NewMetricDto instantiates a new MetricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDtoWithDefaults

`func NewMetricDtoWithDefaults() *MetricDto`

NewMetricDtoWithDefaults instantiates a new MetricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *MetricDto) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *MetricDto) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *MetricDto) SetCurrent(v int32)`

SetCurrent sets Current field to given value.


### GetCurrentPercent

`func (o *MetricDto) GetCurrentPercent() int32`

GetCurrentPercent returns the CurrentPercent field if non-nil, zero value otherwise.

### GetCurrentPercentOk

`func (o *MetricDto) GetCurrentPercentOk() (*int32, bool)`

GetCurrentPercentOk returns a tuple with the CurrentPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPercent

`func (o *MetricDto) SetCurrentPercent(v int32)`

SetCurrentPercent sets CurrentPercent field to given value.


### GetLimit

`func (o *MetricDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetName

`func (o *MetricDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetricDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetricDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *MetricDto) GetStatus() ResourceStatusDto`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricDto) GetStatusOk() (*ResourceStatusDto, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricDto) SetStatus(v ResourceStatusDto)`

SetStatus sets Status field to given value.


### GetUnit

`func (o *MetricDto) GetUnit() UnitDto`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricDto) GetUnitOk() (*UnitDto, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricDto) SetUnit(v UnitDto)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


