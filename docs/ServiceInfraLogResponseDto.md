# ServiceInfraLogResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int32** | Unix timestamp with millisecond precision | 
**Message** | **string** |  | 

## Methods

### NewServiceInfraLogResponseDto

`func NewServiceInfraLogResponseDto(createdAt int32, message string, ) *ServiceInfraLogResponseDto`

NewServiceInfraLogResponseDto instantiates a new ServiceInfraLogResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInfraLogResponseDtoWithDefaults

`func NewServiceInfraLogResponseDtoWithDefaults() *ServiceInfraLogResponseDto`

NewServiceInfraLogResponseDtoWithDefaults instantiates a new ServiceInfraLogResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceInfraLogResponseDto) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceInfraLogResponseDto) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceInfraLogResponseDto) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetMessage

`func (o *ServiceInfraLogResponseDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceInfraLogResponseDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceInfraLogResponseDto) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


