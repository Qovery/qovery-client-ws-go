# ServiceActionDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ServiceActionDto**](ServiceActionDto.md) |  | 
**SubAction** | [**ServiceSubActionDto**](ServiceSubActionDto.md) |  | 

## Methods

### NewServiceActionDetailsDto

`func NewServiceActionDetailsDto(action ServiceActionDto, subAction ServiceSubActionDto, ) *ServiceActionDetailsDto`

NewServiceActionDetailsDto instantiates a new ServiceActionDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceActionDetailsDtoWithDefaults

`func NewServiceActionDetailsDtoWithDefaults() *ServiceActionDetailsDto`

NewServiceActionDetailsDtoWithDefaults instantiates a new ServiceActionDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ServiceActionDetailsDto) GetAction() ServiceActionDto`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServiceActionDetailsDto) GetActionOk() (*ServiceActionDto, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServiceActionDetailsDto) SetAction(v ServiceActionDto)`

SetAction sets Action field to given value.


### GetSubAction

`func (o *ServiceActionDetailsDto) GetSubAction() ServiceSubActionDto`

GetSubAction returns the SubAction field if non-nil, zero value otherwise.

### GetSubActionOk

`func (o *ServiceActionDetailsDto) GetSubActionOk() (*ServiceSubActionDto, bool)`

GetSubActionOk returns a tuple with the SubAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAction

`func (o *ServiceActionDetailsDto) SetSubAction(v ServiceSubActionDto)`

SetSubAction sets SubAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


