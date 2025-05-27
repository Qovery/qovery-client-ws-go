# ClusterMetricsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePools** | [**[]NodePoolInfoDto**](NodePoolInfoDto.md) |  | 
**Nodes** | [**[]ClusterNodeDto**](ClusterNodeDto.md) |  | 
**Pvcs** | [**[]PvcInfoDto**](PvcInfoDto.md) |  | 

## Methods

### NewClusterMetricsDto

`func NewClusterMetricsDto(nodePools []NodePoolInfoDto, nodes []ClusterNodeDto, pvcs []PvcInfoDto, ) *ClusterMetricsDto`

NewClusterMetricsDto instantiates a new ClusterMetricsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMetricsDtoWithDefaults

`func NewClusterMetricsDtoWithDefaults() *ClusterMetricsDto`

NewClusterMetricsDtoWithDefaults instantiates a new ClusterMetricsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePools

`func (o *ClusterMetricsDto) GetNodePools() []NodePoolInfoDto`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *ClusterMetricsDto) GetNodePoolsOk() (*[]NodePoolInfoDto, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *ClusterMetricsDto) SetNodePools(v []NodePoolInfoDto)`

SetNodePools sets NodePools field to given value.


### GetNodes

`func (o *ClusterMetricsDto) GetNodes() []ClusterNodeDto`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterMetricsDto) GetNodesOk() (*[]ClusterNodeDto, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterMetricsDto) SetNodes(v []ClusterNodeDto)`

SetNodes sets Nodes field to given value.


### GetPvcs

`func (o *ClusterMetricsDto) GetPvcs() []PvcInfoDto`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *ClusterMetricsDto) GetPvcsOk() (*[]PvcInfoDto, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *ClusterMetricsDto) SetPvcs(v []PvcInfoDto)`

SetPvcs sets Pvcs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


