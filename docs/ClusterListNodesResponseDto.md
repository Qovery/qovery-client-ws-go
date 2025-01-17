# ClusterListNodesResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]NodeDto**](NodeDto.md) |  | 

## Methods

### NewClusterListNodesResponseDto

`func NewClusterListNodesResponseDto(nodes []NodeDto, ) *ClusterListNodesResponseDto`

NewClusterListNodesResponseDto instantiates a new ClusterListNodesResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterListNodesResponseDtoWithDefaults

`func NewClusterListNodesResponseDtoWithDefaults() *ClusterListNodesResponseDto`

NewClusterListNodesResponseDtoWithDefaults instantiates a new ClusterListNodesResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ClusterListNodesResponseDto) GetNodes() []NodeDto`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterListNodesResponseDto) GetNodesOk() (*[]NodeDto, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterListNodesResponseDto) SetNodes(v []NodeDto)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


