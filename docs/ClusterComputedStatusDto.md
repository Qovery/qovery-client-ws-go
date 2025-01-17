# ClusterComputedStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalStatus** | [**ClusterStatusGlobalStatus**](ClusterStatusGlobalStatus.md) |  | 
**IsMaxNodesSizeReached** | **bool** |  | 
**KubeVersionStatus** | [**QoveryClusterKubeVersionStatus**](QoveryClusterKubeVersionStatus.md) |  | 
**NodeWarnings** | [**map[string][]QoveryNodeFailure**](array.md) |  | 
**QoveryComponentsInFailure** | [**[]QoveryComponentInFailure**](QoveryComponentInFailure.md) |  | 

## Methods

### NewClusterComputedStatusDto

`func NewClusterComputedStatusDto(globalStatus ClusterStatusGlobalStatus, isMaxNodesSizeReached bool, kubeVersionStatus QoveryClusterKubeVersionStatus, nodeWarnings map[string][]QoveryNodeFailure, qoveryComponentsInFailure []QoveryComponentInFailure, ) *ClusterComputedStatusDto`

NewClusterComputedStatusDto instantiates a new ClusterComputedStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterComputedStatusDtoWithDefaults

`func NewClusterComputedStatusDtoWithDefaults() *ClusterComputedStatusDto`

NewClusterComputedStatusDtoWithDefaults instantiates a new ClusterComputedStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalStatus

`func (o *ClusterComputedStatusDto) GetGlobalStatus() ClusterStatusGlobalStatus`

GetGlobalStatus returns the GlobalStatus field if non-nil, zero value otherwise.

### GetGlobalStatusOk

`func (o *ClusterComputedStatusDto) GetGlobalStatusOk() (*ClusterStatusGlobalStatus, bool)`

GetGlobalStatusOk returns a tuple with the GlobalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStatus

`func (o *ClusterComputedStatusDto) SetGlobalStatus(v ClusterStatusGlobalStatus)`

SetGlobalStatus sets GlobalStatus field to given value.


### GetIsMaxNodesSizeReached

`func (o *ClusterComputedStatusDto) GetIsMaxNodesSizeReached() bool`

GetIsMaxNodesSizeReached returns the IsMaxNodesSizeReached field if non-nil, zero value otherwise.

### GetIsMaxNodesSizeReachedOk

`func (o *ClusterComputedStatusDto) GetIsMaxNodesSizeReachedOk() (*bool, bool)`

GetIsMaxNodesSizeReachedOk returns a tuple with the IsMaxNodesSizeReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaxNodesSizeReached

`func (o *ClusterComputedStatusDto) SetIsMaxNodesSizeReached(v bool)`

SetIsMaxNodesSizeReached sets IsMaxNodesSizeReached field to given value.


### GetKubeVersionStatus

`func (o *ClusterComputedStatusDto) GetKubeVersionStatus() QoveryClusterKubeVersionStatus`

GetKubeVersionStatus returns the KubeVersionStatus field if non-nil, zero value otherwise.

### GetKubeVersionStatusOk

`func (o *ClusterComputedStatusDto) GetKubeVersionStatusOk() (*QoveryClusterKubeVersionStatus, bool)`

GetKubeVersionStatusOk returns a tuple with the KubeVersionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersionStatus

`func (o *ClusterComputedStatusDto) SetKubeVersionStatus(v QoveryClusterKubeVersionStatus)`

SetKubeVersionStatus sets KubeVersionStatus field to given value.


### GetNodeWarnings

`func (o *ClusterComputedStatusDto) GetNodeWarnings() map[string][]QoveryNodeFailure`

GetNodeWarnings returns the NodeWarnings field if non-nil, zero value otherwise.

### GetNodeWarningsOk

`func (o *ClusterComputedStatusDto) GetNodeWarningsOk() (*map[string][]QoveryNodeFailure, bool)`

GetNodeWarningsOk returns a tuple with the NodeWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWarnings

`func (o *ClusterComputedStatusDto) SetNodeWarnings(v map[string][]QoveryNodeFailure)`

SetNodeWarnings sets NodeWarnings field to given value.


### GetQoveryComponentsInFailure

`func (o *ClusterComputedStatusDto) GetQoveryComponentsInFailure() []QoveryComponentInFailure`

GetQoveryComponentsInFailure returns the QoveryComponentsInFailure field if non-nil, zero value otherwise.

### GetQoveryComponentsInFailureOk

`func (o *ClusterComputedStatusDto) GetQoveryComponentsInFailureOk() (*[]QoveryComponentInFailure, bool)`

GetQoveryComponentsInFailureOk returns a tuple with the QoveryComponentsInFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryComponentsInFailure

`func (o *ClusterComputedStatusDto) SetQoveryComponentsInFailure(v []QoveryComponentInFailure)`

SetQoveryComponentsInFailure sets QoveryComponentsInFailure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


