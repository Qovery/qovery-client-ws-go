# DatabaseStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Pods** | [**[]PodStatusDto**](PodStatusDto.md) |  | 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 

## Methods

### NewDatabaseStatusDto

`func NewDatabaseStatusDto(id string, pods []PodStatusDto, state ServiceStateDto, ) *DatabaseStatusDto`

NewDatabaseStatusDto instantiates a new DatabaseStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseStatusDtoWithDefaults

`func NewDatabaseStatusDtoWithDefaults() *DatabaseStatusDto`

NewDatabaseStatusDtoWithDefaults instantiates a new DatabaseStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabaseStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseStatusDto) SetId(v string)`

SetId sets Id field to given value.


### GetPods

`func (o *DatabaseStatusDto) GetPods() []PodStatusDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *DatabaseStatusDto) GetPodsOk() (*[]PodStatusDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *DatabaseStatusDto) SetPods(v []PodStatusDto)`

SetPods sets Pods field to given value.


### GetState

`func (o *DatabaseStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DatabaseStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DatabaseStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


