# ArgoCdAppStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Pods** | [**[]PodStatusDto**](PodStatusDto.md) |  | 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 

## Methods

### NewArgoCdAppStatusDto

`func NewArgoCdAppStatusDto(id string, pods []PodStatusDto, state ServiceStateDto, ) *ArgoCdAppStatusDto`

NewArgoCdAppStatusDto instantiates a new ArgoCdAppStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdAppStatusDtoWithDefaults

`func NewArgoCdAppStatusDtoWithDefaults() *ArgoCdAppStatusDto`

NewArgoCdAppStatusDtoWithDefaults instantiates a new ArgoCdAppStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArgoCdAppStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgoCdAppStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgoCdAppStatusDto) SetId(v string)`

SetId sets Id field to given value.


### GetPods

`func (o *ArgoCdAppStatusDto) GetPods() []PodStatusDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ArgoCdAppStatusDto) GetPodsOk() (*[]PodStatusDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ArgoCdAppStatusDto) SetPods(v []PodStatusDto)`

SetPods sets Pods field to given value.


### GetState

`func (o *ArgoCdAppStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ArgoCdAppStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ArgoCdAppStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


