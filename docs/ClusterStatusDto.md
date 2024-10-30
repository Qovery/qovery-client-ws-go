# ClusterStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedStatus** | [**ClusterStatusDtoComputedStatus**](ClusterStatusDtoComputedStatus.md) |  | 
**Nodes** | [**[]ClusterNodeDto**](ClusterNodeDto.md) |  | 

## Methods

### NewClusterStatusDto

`func NewClusterStatusDto(computedStatus ClusterStatusDtoComputedStatus, nodes []ClusterNodeDto, ) *ClusterStatusDto`

NewClusterStatusDto instantiates a new ClusterStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusDtoWithDefaults

`func NewClusterStatusDtoWithDefaults() *ClusterStatusDto`

NewClusterStatusDtoWithDefaults instantiates a new ClusterStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedStatus

`func (o *ClusterStatusDto) GetComputedStatus() ClusterStatusDtoComputedStatus`

GetComputedStatus returns the ComputedStatus field if non-nil, zero value otherwise.

### GetComputedStatusOk

`func (o *ClusterStatusDto) GetComputedStatusOk() (*ClusterStatusDtoComputedStatus, bool)`

GetComputedStatusOk returns a tuple with the ComputedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedStatus

`func (o *ClusterStatusDto) SetComputedStatus(v ClusterStatusDtoComputedStatus)`

SetComputedStatus sets ComputedStatus field to given value.


### GetNodes

`func (o *ClusterStatusDto) GetNodes() []ClusterNodeDto`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterStatusDto) GetNodesOk() (*[]ClusterNodeDto, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterStatusDto) SetNodes(v []ClusterNodeDto)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


