# ScaledObjectStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationsTypes** | Pointer to **NullableString** |  | [optional] 
**CompositeScalerName** | Pointer to **NullableString** |  | [optional] 
**Conditions** | [**[]ScaledObjectConditionDto**](ScaledObjectConditionDto.md) |  | 
**ExternalMetricNames** | **[]string** |  | 
**Health** | [**map[string]ScaledObjectHealthDto**](ScaledObjectHealthDto.md) |  | 
**HpaName** | Pointer to **NullableString** |  | [optional] 
**LastActiveTime** | **int64** |  | 
**Name** | **string** |  | 
**OriginalReplicaCount** | Pointer to **NullableInt32** |  | [optional] 
**PausedReplicaCount** | Pointer to **NullableInt32** |  | [optional] 
**ResourceMetricNames** | **[]string** |  | 
**ScaleTargetGvkr** | Pointer to [**NullableScaleTargetGvkrDto**](ScaleTargetGvkrDto.md) |  | [optional] 
**ScaleTargetKind** | Pointer to **NullableString** |  | [optional] 
**TriggersTypes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScaledObjectStatusDto

`func NewScaledObjectStatusDto(conditions []ScaledObjectConditionDto, externalMetricNames []string, health map[string]ScaledObjectHealthDto, lastActiveTime int64, name string, resourceMetricNames []string, ) *ScaledObjectStatusDto`

NewScaledObjectStatusDto instantiates a new ScaledObjectStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaledObjectStatusDtoWithDefaults

`func NewScaledObjectStatusDtoWithDefaults() *ScaledObjectStatusDto`

NewScaledObjectStatusDtoWithDefaults instantiates a new ScaledObjectStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationsTypes

`func (o *ScaledObjectStatusDto) GetAuthenticationsTypes() string`

GetAuthenticationsTypes returns the AuthenticationsTypes field if non-nil, zero value otherwise.

### GetAuthenticationsTypesOk

`func (o *ScaledObjectStatusDto) GetAuthenticationsTypesOk() (*string, bool)`

GetAuthenticationsTypesOk returns a tuple with the AuthenticationsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationsTypes

`func (o *ScaledObjectStatusDto) SetAuthenticationsTypes(v string)`

SetAuthenticationsTypes sets AuthenticationsTypes field to given value.

### HasAuthenticationsTypes

`func (o *ScaledObjectStatusDto) HasAuthenticationsTypes() bool`

HasAuthenticationsTypes returns a boolean if a field has been set.

### SetAuthenticationsTypesNil

`func (o *ScaledObjectStatusDto) SetAuthenticationsTypesNil(b bool)`

 SetAuthenticationsTypesNil sets the value for AuthenticationsTypes to be an explicit nil

### UnsetAuthenticationsTypes
`func (o *ScaledObjectStatusDto) UnsetAuthenticationsTypes()`

UnsetAuthenticationsTypes ensures that no value is present for AuthenticationsTypes, not even an explicit nil
### GetCompositeScalerName

`func (o *ScaledObjectStatusDto) GetCompositeScalerName() string`

GetCompositeScalerName returns the CompositeScalerName field if non-nil, zero value otherwise.

### GetCompositeScalerNameOk

`func (o *ScaledObjectStatusDto) GetCompositeScalerNameOk() (*string, bool)`

GetCompositeScalerNameOk returns a tuple with the CompositeScalerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeScalerName

`func (o *ScaledObjectStatusDto) SetCompositeScalerName(v string)`

SetCompositeScalerName sets CompositeScalerName field to given value.

### HasCompositeScalerName

`func (o *ScaledObjectStatusDto) HasCompositeScalerName() bool`

HasCompositeScalerName returns a boolean if a field has been set.

### SetCompositeScalerNameNil

`func (o *ScaledObjectStatusDto) SetCompositeScalerNameNil(b bool)`

 SetCompositeScalerNameNil sets the value for CompositeScalerName to be an explicit nil

### UnsetCompositeScalerName
`func (o *ScaledObjectStatusDto) UnsetCompositeScalerName()`

UnsetCompositeScalerName ensures that no value is present for CompositeScalerName, not even an explicit nil
### GetConditions

`func (o *ScaledObjectStatusDto) GetConditions() []ScaledObjectConditionDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ScaledObjectStatusDto) GetConditionsOk() (*[]ScaledObjectConditionDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ScaledObjectStatusDto) SetConditions(v []ScaledObjectConditionDto)`

SetConditions sets Conditions field to given value.


### GetExternalMetricNames

`func (o *ScaledObjectStatusDto) GetExternalMetricNames() []string`

GetExternalMetricNames returns the ExternalMetricNames field if non-nil, zero value otherwise.

### GetExternalMetricNamesOk

`func (o *ScaledObjectStatusDto) GetExternalMetricNamesOk() (*[]string, bool)`

GetExternalMetricNamesOk returns a tuple with the ExternalMetricNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMetricNames

`func (o *ScaledObjectStatusDto) SetExternalMetricNames(v []string)`

SetExternalMetricNames sets ExternalMetricNames field to given value.


### GetHealth

`func (o *ScaledObjectStatusDto) GetHealth() map[string]ScaledObjectHealthDto`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ScaledObjectStatusDto) GetHealthOk() (*map[string]ScaledObjectHealthDto, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ScaledObjectStatusDto) SetHealth(v map[string]ScaledObjectHealthDto)`

SetHealth sets Health field to given value.


### GetHpaName

`func (o *ScaledObjectStatusDto) GetHpaName() string`

GetHpaName returns the HpaName field if non-nil, zero value otherwise.

### GetHpaNameOk

`func (o *ScaledObjectStatusDto) GetHpaNameOk() (*string, bool)`

GetHpaNameOk returns a tuple with the HpaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpaName

`func (o *ScaledObjectStatusDto) SetHpaName(v string)`

SetHpaName sets HpaName field to given value.

### HasHpaName

`func (o *ScaledObjectStatusDto) HasHpaName() bool`

HasHpaName returns a boolean if a field has been set.

### SetHpaNameNil

`func (o *ScaledObjectStatusDto) SetHpaNameNil(b bool)`

 SetHpaNameNil sets the value for HpaName to be an explicit nil

### UnsetHpaName
`func (o *ScaledObjectStatusDto) UnsetHpaName()`

UnsetHpaName ensures that no value is present for HpaName, not even an explicit nil
### GetLastActiveTime

`func (o *ScaledObjectStatusDto) GetLastActiveTime() int64`

GetLastActiveTime returns the LastActiveTime field if non-nil, zero value otherwise.

### GetLastActiveTimeOk

`func (o *ScaledObjectStatusDto) GetLastActiveTimeOk() (*int64, bool)`

GetLastActiveTimeOk returns a tuple with the LastActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveTime

`func (o *ScaledObjectStatusDto) SetLastActiveTime(v int64)`

SetLastActiveTime sets LastActiveTime field to given value.


### GetName

`func (o *ScaledObjectStatusDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScaledObjectStatusDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScaledObjectStatusDto) SetName(v string)`

SetName sets Name field to given value.


### GetOriginalReplicaCount

`func (o *ScaledObjectStatusDto) GetOriginalReplicaCount() int32`

GetOriginalReplicaCount returns the OriginalReplicaCount field if non-nil, zero value otherwise.

### GetOriginalReplicaCountOk

`func (o *ScaledObjectStatusDto) GetOriginalReplicaCountOk() (*int32, bool)`

GetOriginalReplicaCountOk returns a tuple with the OriginalReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReplicaCount

`func (o *ScaledObjectStatusDto) SetOriginalReplicaCount(v int32)`

SetOriginalReplicaCount sets OriginalReplicaCount field to given value.

### HasOriginalReplicaCount

`func (o *ScaledObjectStatusDto) HasOriginalReplicaCount() bool`

HasOriginalReplicaCount returns a boolean if a field has been set.

### SetOriginalReplicaCountNil

`func (o *ScaledObjectStatusDto) SetOriginalReplicaCountNil(b bool)`

 SetOriginalReplicaCountNil sets the value for OriginalReplicaCount to be an explicit nil

### UnsetOriginalReplicaCount
`func (o *ScaledObjectStatusDto) UnsetOriginalReplicaCount()`

UnsetOriginalReplicaCount ensures that no value is present for OriginalReplicaCount, not even an explicit nil
### GetPausedReplicaCount

`func (o *ScaledObjectStatusDto) GetPausedReplicaCount() int32`

GetPausedReplicaCount returns the PausedReplicaCount field if non-nil, zero value otherwise.

### GetPausedReplicaCountOk

`func (o *ScaledObjectStatusDto) GetPausedReplicaCountOk() (*int32, bool)`

GetPausedReplicaCountOk returns a tuple with the PausedReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedReplicaCount

`func (o *ScaledObjectStatusDto) SetPausedReplicaCount(v int32)`

SetPausedReplicaCount sets PausedReplicaCount field to given value.

### HasPausedReplicaCount

`func (o *ScaledObjectStatusDto) HasPausedReplicaCount() bool`

HasPausedReplicaCount returns a boolean if a field has been set.

### SetPausedReplicaCountNil

`func (o *ScaledObjectStatusDto) SetPausedReplicaCountNil(b bool)`

 SetPausedReplicaCountNil sets the value for PausedReplicaCount to be an explicit nil

### UnsetPausedReplicaCount
`func (o *ScaledObjectStatusDto) UnsetPausedReplicaCount()`

UnsetPausedReplicaCount ensures that no value is present for PausedReplicaCount, not even an explicit nil
### GetResourceMetricNames

`func (o *ScaledObjectStatusDto) GetResourceMetricNames() []string`

GetResourceMetricNames returns the ResourceMetricNames field if non-nil, zero value otherwise.

### GetResourceMetricNamesOk

`func (o *ScaledObjectStatusDto) GetResourceMetricNamesOk() (*[]string, bool)`

GetResourceMetricNamesOk returns a tuple with the ResourceMetricNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMetricNames

`func (o *ScaledObjectStatusDto) SetResourceMetricNames(v []string)`

SetResourceMetricNames sets ResourceMetricNames field to given value.


### GetScaleTargetGvkr

`func (o *ScaledObjectStatusDto) GetScaleTargetGvkr() ScaleTargetGvkrDto`

GetScaleTargetGvkr returns the ScaleTargetGvkr field if non-nil, zero value otherwise.

### GetScaleTargetGvkrOk

`func (o *ScaledObjectStatusDto) GetScaleTargetGvkrOk() (*ScaleTargetGvkrDto, bool)`

GetScaleTargetGvkrOk returns a tuple with the ScaleTargetGvkr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleTargetGvkr

`func (o *ScaledObjectStatusDto) SetScaleTargetGvkr(v ScaleTargetGvkrDto)`

SetScaleTargetGvkr sets ScaleTargetGvkr field to given value.

### HasScaleTargetGvkr

`func (o *ScaledObjectStatusDto) HasScaleTargetGvkr() bool`

HasScaleTargetGvkr returns a boolean if a field has been set.

### SetScaleTargetGvkrNil

`func (o *ScaledObjectStatusDto) SetScaleTargetGvkrNil(b bool)`

 SetScaleTargetGvkrNil sets the value for ScaleTargetGvkr to be an explicit nil

### UnsetScaleTargetGvkr
`func (o *ScaledObjectStatusDto) UnsetScaleTargetGvkr()`

UnsetScaleTargetGvkr ensures that no value is present for ScaleTargetGvkr, not even an explicit nil
### GetScaleTargetKind

`func (o *ScaledObjectStatusDto) GetScaleTargetKind() string`

GetScaleTargetKind returns the ScaleTargetKind field if non-nil, zero value otherwise.

### GetScaleTargetKindOk

`func (o *ScaledObjectStatusDto) GetScaleTargetKindOk() (*string, bool)`

GetScaleTargetKindOk returns a tuple with the ScaleTargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleTargetKind

`func (o *ScaledObjectStatusDto) SetScaleTargetKind(v string)`

SetScaleTargetKind sets ScaleTargetKind field to given value.

### HasScaleTargetKind

`func (o *ScaledObjectStatusDto) HasScaleTargetKind() bool`

HasScaleTargetKind returns a boolean if a field has been set.

### SetScaleTargetKindNil

`func (o *ScaledObjectStatusDto) SetScaleTargetKindNil(b bool)`

 SetScaleTargetKindNil sets the value for ScaleTargetKind to be an explicit nil

### UnsetScaleTargetKind
`func (o *ScaledObjectStatusDto) UnsetScaleTargetKind()`

UnsetScaleTargetKind ensures that no value is present for ScaleTargetKind, not even an explicit nil
### GetTriggersTypes

`func (o *ScaledObjectStatusDto) GetTriggersTypes() string`

GetTriggersTypes returns the TriggersTypes field if non-nil, zero value otherwise.

### GetTriggersTypesOk

`func (o *ScaledObjectStatusDto) GetTriggersTypesOk() (*string, bool)`

GetTriggersTypesOk returns a tuple with the TriggersTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggersTypes

`func (o *ScaledObjectStatusDto) SetTriggersTypes(v string)`

SetTriggersTypes sets TriggersTypes field to given value.

### HasTriggersTypes

`func (o *ScaledObjectStatusDto) HasTriggersTypes() bool`

HasTriggersTypes returns a boolean if a field has been set.

### SetTriggersTypesNil

`func (o *ScaledObjectStatusDto) SetTriggersTypesNil(b bool)`

 SetTriggersTypesNil sets the value for TriggersTypes to be an explicit nil

### UnsetTriggersTypes
`func (o *ScaledObjectStatusDto) UnsetTriggersTypes()`

UnsetTriggersTypes ensures that no value is present for TriggersTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


