# ServiceStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | [**[]EnvironmentStatusDto**](EnvironmentStatusDto.md) |  | 

## Methods

### NewServiceStatusDto

`func NewServiceStatusDto(environments []EnvironmentStatusDto, ) *ServiceStatusDto`

NewServiceStatusDto instantiates a new ServiceStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusDtoWithDefaults

`func NewServiceStatusDtoWithDefaults() *ServiceStatusDto`

NewServiceStatusDtoWithDefaults instantiates a new ServiceStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *ServiceStatusDto) GetEnvironments() []EnvironmentStatusDto`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ServiceStatusDto) GetEnvironmentsOk() (*[]EnvironmentStatusDto, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ServiceStatusDto) SetEnvironments(v []EnvironmentStatusDto)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


