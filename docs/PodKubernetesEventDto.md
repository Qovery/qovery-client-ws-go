# PodKubernetesEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**Message** | **string** |  | 
**Reason** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewPodKubernetesEventDto

`func NewPodKubernetesEventDto(createdAt int64, message string, reason string, type_ string, ) *PodKubernetesEventDto`

NewPodKubernetesEventDto instantiates a new PodKubernetesEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodKubernetesEventDtoWithDefaults

`func NewPodKubernetesEventDtoWithDefaults() *PodKubernetesEventDto`

NewPodKubernetesEventDtoWithDefaults instantiates a new PodKubernetesEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PodKubernetesEventDto) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PodKubernetesEventDto) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PodKubernetesEventDto) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetMessage

`func (o *PodKubernetesEventDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PodKubernetesEventDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PodKubernetesEventDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *PodKubernetesEventDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PodKubernetesEventDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PodKubernetesEventDto) SetReason(v string)`

SetReason sets Reason field to given value.


### GetType

`func (o *PodKubernetesEventDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PodKubernetesEventDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PodKubernetesEventDto) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


