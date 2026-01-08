# ScaledObjectHealthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfFailures** | Pointer to **NullableInt32** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewScaledObjectHealthDto

`func NewScaledObjectHealthDto(status string, ) *ScaledObjectHealthDto`

NewScaledObjectHealthDto instantiates a new ScaledObjectHealthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaledObjectHealthDtoWithDefaults

`func NewScaledObjectHealthDtoWithDefaults() *ScaledObjectHealthDto`

NewScaledObjectHealthDtoWithDefaults instantiates a new ScaledObjectHealthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfFailures

`func (o *ScaledObjectHealthDto) GetNumberOfFailures() int32`

GetNumberOfFailures returns the NumberOfFailures field if non-nil, zero value otherwise.

### GetNumberOfFailuresOk

`func (o *ScaledObjectHealthDto) GetNumberOfFailuresOk() (*int32, bool)`

GetNumberOfFailuresOk returns a tuple with the NumberOfFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFailures

`func (o *ScaledObjectHealthDto) SetNumberOfFailures(v int32)`

SetNumberOfFailures sets NumberOfFailures field to given value.

### HasNumberOfFailures

`func (o *ScaledObjectHealthDto) HasNumberOfFailures() bool`

HasNumberOfFailures returns a boolean if a field has been set.

### SetNumberOfFailuresNil

`func (o *ScaledObjectHealthDto) SetNumberOfFailuresNil(b bool)`

 SetNumberOfFailuresNil sets the value for NumberOfFailures to be an explicit nil

### UnsetNumberOfFailures
`func (o *ScaledObjectHealthDto) UnsetNumberOfFailures()`

UnsetNumberOfFailures ensures that no value is present for NumberOfFailures, not even an explicit nil
### GetStatus

`func (o *ScaledObjectHealthDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScaledObjectHealthDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScaledObjectHealthDto) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


