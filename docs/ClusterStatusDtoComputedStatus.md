# ClusterStatusDtoComputedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalStatus** | Pointer to **NullableString** |  | [optional] 
**QoveryComponentsInFailure** | Pointer to [**[]ClusterStatusDtoComputedStatusQoveryComponentsInFailureInner**](ClusterStatusDtoComputedStatusQoveryComponentsInFailureInner.md) |  | [optional] 
**NodeWarnings** | Pointer to [**map[string]NodeInWarning**](NodeInWarning.md) |  | [optional] 
**IsMaxNodesSizeReached** | Pointer to **bool** |  | [optional] 
**KubeVersionStatus** | Pointer to [**ClusterStatusDtoComputedStatusKubeVersionStatus**](ClusterStatusDtoComputedStatusKubeVersionStatus.md) |  | [optional] 

## Methods

### NewClusterStatusDtoComputedStatus

`func NewClusterStatusDtoComputedStatus() *ClusterStatusDtoComputedStatus`

NewClusterStatusDtoComputedStatus instantiates a new ClusterStatusDtoComputedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusDtoComputedStatusWithDefaults

`func NewClusterStatusDtoComputedStatusWithDefaults() *ClusterStatusDtoComputedStatus`

NewClusterStatusDtoComputedStatusWithDefaults instantiates a new ClusterStatusDtoComputedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalStatus

`func (o *ClusterStatusDtoComputedStatus) GetGlobalStatus() string`

GetGlobalStatus returns the GlobalStatus field if non-nil, zero value otherwise.

### GetGlobalStatusOk

`func (o *ClusterStatusDtoComputedStatus) GetGlobalStatusOk() (*string, bool)`

GetGlobalStatusOk returns a tuple with the GlobalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStatus

`func (o *ClusterStatusDtoComputedStatus) SetGlobalStatus(v string)`

SetGlobalStatus sets GlobalStatus field to given value.

### HasGlobalStatus

`func (o *ClusterStatusDtoComputedStatus) HasGlobalStatus() bool`

HasGlobalStatus returns a boolean if a field has been set.

### SetGlobalStatusNil

`func (o *ClusterStatusDtoComputedStatus) SetGlobalStatusNil(b bool)`

 SetGlobalStatusNil sets the value for GlobalStatus to be an explicit nil

### UnsetGlobalStatus
`func (o *ClusterStatusDtoComputedStatus) UnsetGlobalStatus()`

UnsetGlobalStatus ensures that no value is present for GlobalStatus, not even an explicit nil
### GetQoveryComponentsInFailure

`func (o *ClusterStatusDtoComputedStatus) GetQoveryComponentsInFailure() []ClusterStatusDtoComputedStatusQoveryComponentsInFailureInner`

GetQoveryComponentsInFailure returns the QoveryComponentsInFailure field if non-nil, zero value otherwise.

### GetQoveryComponentsInFailureOk

`func (o *ClusterStatusDtoComputedStatus) GetQoveryComponentsInFailureOk() (*[]ClusterStatusDtoComputedStatusQoveryComponentsInFailureInner, bool)`

GetQoveryComponentsInFailureOk returns a tuple with the QoveryComponentsInFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryComponentsInFailure

`func (o *ClusterStatusDtoComputedStatus) SetQoveryComponentsInFailure(v []ClusterStatusDtoComputedStatusQoveryComponentsInFailureInner)`

SetQoveryComponentsInFailure sets QoveryComponentsInFailure field to given value.

### HasQoveryComponentsInFailure

`func (o *ClusterStatusDtoComputedStatus) HasQoveryComponentsInFailure() bool`

HasQoveryComponentsInFailure returns a boolean if a field has been set.

### GetNodeWarnings

`func (o *ClusterStatusDtoComputedStatus) GetNodeWarnings() map[string]NodeInWarning`

GetNodeWarnings returns the NodeWarnings field if non-nil, zero value otherwise.

### GetNodeWarningsOk

`func (o *ClusterStatusDtoComputedStatus) GetNodeWarningsOk() (*map[string]NodeInWarning, bool)`

GetNodeWarningsOk returns a tuple with the NodeWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWarnings

`func (o *ClusterStatusDtoComputedStatus) SetNodeWarnings(v map[string]NodeInWarning)`

SetNodeWarnings sets NodeWarnings field to given value.

### HasNodeWarnings

`func (o *ClusterStatusDtoComputedStatus) HasNodeWarnings() bool`

HasNodeWarnings returns a boolean if a field has been set.

### GetIsMaxNodesSizeReached

`func (o *ClusterStatusDtoComputedStatus) GetIsMaxNodesSizeReached() bool`

GetIsMaxNodesSizeReached returns the IsMaxNodesSizeReached field if non-nil, zero value otherwise.

### GetIsMaxNodesSizeReachedOk

`func (o *ClusterStatusDtoComputedStatus) GetIsMaxNodesSizeReachedOk() (*bool, bool)`

GetIsMaxNodesSizeReachedOk returns a tuple with the IsMaxNodesSizeReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaxNodesSizeReached

`func (o *ClusterStatusDtoComputedStatus) SetIsMaxNodesSizeReached(v bool)`

SetIsMaxNodesSizeReached sets IsMaxNodesSizeReached field to given value.

### HasIsMaxNodesSizeReached

`func (o *ClusterStatusDtoComputedStatus) HasIsMaxNodesSizeReached() bool`

HasIsMaxNodesSizeReached returns a boolean if a field has been set.

### GetKubeVersionStatus

`func (o *ClusterStatusDtoComputedStatus) GetKubeVersionStatus() ClusterStatusDtoComputedStatusKubeVersionStatus`

GetKubeVersionStatus returns the KubeVersionStatus field if non-nil, zero value otherwise.

### GetKubeVersionStatusOk

`func (o *ClusterStatusDtoComputedStatus) GetKubeVersionStatusOk() (*ClusterStatusDtoComputedStatusKubeVersionStatus, bool)`

GetKubeVersionStatusOk returns a tuple with the KubeVersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersionStatus

`func (o *ClusterStatusDtoComputedStatus) SetKubeVersionStatus(v ClusterStatusDtoComputedStatusKubeVersionStatus)`

SetKubeVersionStatus sets KubeVersionStatus field to given value.

### HasKubeVersionStatus

`func (o *ClusterStatusDtoComputedStatus) HasKubeVersionStatus() bool`

HasKubeVersionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


