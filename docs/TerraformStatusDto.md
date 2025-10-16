# TerraformStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Pods** | [**[]PodStatusDto**](PodStatusDto.md) |  | 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 
**TriggeredAction** | [**ServiceActionDetailsDto**](ServiceActionDetailsDto.md) |  | 

## Methods

### NewTerraformStatusDto

`func NewTerraformStatusDto(id string, pods []PodStatusDto, state ServiceStateDto, triggeredAction ServiceActionDetailsDto, ) *TerraformStatusDto`

NewTerraformStatusDto instantiates a new TerraformStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformStatusDtoWithDefaults

`func NewTerraformStatusDtoWithDefaults() *TerraformStatusDto`

NewTerraformStatusDtoWithDefaults instantiates a new TerraformStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerraformStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerraformStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerraformStatusDto) SetId(v string)`

SetId sets Id field to given value.


### GetPods

`func (o *TerraformStatusDto) GetPods() []PodStatusDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *TerraformStatusDto) GetPodsOk() (*[]PodStatusDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *TerraformStatusDto) SetPods(v []PodStatusDto)`

SetPods sets Pods field to given value.


### GetState

`func (o *TerraformStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TerraformStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TerraformStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.


### GetTriggeredAction

`func (o *TerraformStatusDto) GetTriggeredAction() ServiceActionDetailsDto`

GetTriggeredAction returns the TriggeredAction field if non-nil, zero value otherwise.

### GetTriggeredActionOk

`func (o *TerraformStatusDto) GetTriggeredActionOk() (*ServiceActionDetailsDto, bool)`

GetTriggeredActionOk returns a tuple with the TriggeredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAction

`func (o *TerraformStatusDto) SetTriggeredAction(v ServiceActionDetailsDto)`

SetTriggeredAction sets TriggeredAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


