# ClusterNodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]NodeAddressDto**](NodeAddressDto.md) |  | 
**Annotations** | **map[string]string** |  | 
**Architecture** | **string** |  | 
**Conditions** | [**[]NodeConditionDto**](NodeConditionDto.md) |  | 
**CreatedAt** | **int64** |  | 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
**KernelVersion** | **string** |  | 
**KubeletVersion** | **string** |  | 
**Labels** | **map[string]string** |  | 
**MetricsUsage** | [**MetricsUsageDto**](MetricsUsageDto.md) |  | 
**Name** | **string** |  | 
**OperatingSystem** | **string** |  | 
**OsImage** | **string** |  | 
**Pods** | [**[]NodePodInfoDto**](NodePodInfoDto.md) |  | 
**ResourcesAllocatable** | [**NodeResourceDto**](NodeResourceDto.md) |  | 
**ResourcesAllocated** | [**NodeResourceAllocatedDto**](NodeResourceAllocatedDto.md) |  | 
**ResourcesCapacity** | [**NodeResourceDto**](NodeResourceDto.md) |  | 
**Taints** | [**[]NodeTaintDto**](NodeTaintDto.md) |  | 
**Unschedulable** | **bool** |  | 

## Methods

### NewClusterNodeDto

`func NewClusterNodeDto(addresses []NodeAddressDto, annotations map[string]string, architecture string, conditions []NodeConditionDto, createdAt int64, kernelVersion string, kubeletVersion string, labels map[string]string, metricsUsage MetricsUsageDto, name string, operatingSystem string, osImage string, pods []NodePodInfoDto, resourcesAllocatable NodeResourceDto, resourcesAllocated NodeResourceAllocatedDto, resourcesCapacity NodeResourceDto, taints []NodeTaintDto, unschedulable bool, ) *ClusterNodeDto`

NewClusterNodeDto instantiates a new ClusterNodeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeDtoWithDefaults

`func NewClusterNodeDtoWithDefaults() *ClusterNodeDto`

NewClusterNodeDtoWithDefaults instantiates a new ClusterNodeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ClusterNodeDto) GetAddresses() []NodeAddressDto`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ClusterNodeDto) GetAddressesOk() (*[]NodeAddressDto, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ClusterNodeDto) SetAddresses(v []NodeAddressDto)`

SetAddresses sets Addresses field to given value.


### GetAnnotations

`func (o *ClusterNodeDto) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ClusterNodeDto) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ClusterNodeDto) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetArchitecture

`func (o *ClusterNodeDto) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ClusterNodeDto) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ClusterNodeDto) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetConditions

`func (o *ClusterNodeDto) GetConditions() []NodeConditionDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterNodeDto) GetConditionsOk() (*[]NodeConditionDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterNodeDto) SetConditions(v []NodeConditionDto)`

SetConditions sets Conditions field to given value.


### GetCreatedAt

`func (o *ClusterNodeDto) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterNodeDto) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterNodeDto) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetInstanceType

`func (o *ClusterNodeDto) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterNodeDto) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterNodeDto) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterNodeDto) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *ClusterNodeDto) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *ClusterNodeDto) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetKernelVersion

`func (o *ClusterNodeDto) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ClusterNodeDto) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ClusterNodeDto) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### GetKubeletVersion

`func (o *ClusterNodeDto) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *ClusterNodeDto) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *ClusterNodeDto) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.


### GetLabels

`func (o *ClusterNodeDto) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterNodeDto) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterNodeDto) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetMetricsUsage

`func (o *ClusterNodeDto) GetMetricsUsage() MetricsUsageDto`

GetMetricsUsage returns the MetricsUsage field if non-nil, zero value otherwise.

### GetMetricsUsageOk

`func (o *ClusterNodeDto) GetMetricsUsageOk() (*MetricsUsageDto, bool)`

GetMetricsUsageOk returns a tuple with the MetricsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUsage

`func (o *ClusterNodeDto) SetMetricsUsage(v MetricsUsageDto)`

SetMetricsUsage sets MetricsUsage field to given value.


### GetName

`func (o *ClusterNodeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNodeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNodeDto) SetName(v string)`

SetName sets Name field to given value.


### GetOperatingSystem

`func (o *ClusterNodeDto) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ClusterNodeDto) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ClusterNodeDto) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetOsImage

`func (o *ClusterNodeDto) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *ClusterNodeDto) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *ClusterNodeDto) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.


### GetPods

`func (o *ClusterNodeDto) GetPods() []NodePodInfoDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ClusterNodeDto) GetPodsOk() (*[]NodePodInfoDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ClusterNodeDto) SetPods(v []NodePodInfoDto)`

SetPods sets Pods field to given value.


### GetResourcesAllocatable

`func (o *ClusterNodeDto) GetResourcesAllocatable() NodeResourceDto`

GetResourcesAllocatable returns the ResourcesAllocatable field if non-nil, zero value otherwise.

### GetResourcesAllocatableOk

`func (o *ClusterNodeDto) GetResourcesAllocatableOk() (*NodeResourceDto, bool)`

GetResourcesAllocatableOk returns a tuple with the ResourcesAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesAllocatable

`func (o *ClusterNodeDto) SetResourcesAllocatable(v NodeResourceDto)`

SetResourcesAllocatable sets ResourcesAllocatable field to given value.


### GetResourcesAllocated

`func (o *ClusterNodeDto) GetResourcesAllocated() NodeResourceAllocatedDto`

GetResourcesAllocated returns the ResourcesAllocated field if non-nil, zero value otherwise.

### GetResourcesAllocatedOk

`func (o *ClusterNodeDto) GetResourcesAllocatedOk() (*NodeResourceAllocatedDto, bool)`

GetResourcesAllocatedOk returns a tuple with the ResourcesAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesAllocated

`func (o *ClusterNodeDto) SetResourcesAllocated(v NodeResourceAllocatedDto)`

SetResourcesAllocated sets ResourcesAllocated field to given value.


### GetResourcesCapacity

`func (o *ClusterNodeDto) GetResourcesCapacity() NodeResourceDto`

GetResourcesCapacity returns the ResourcesCapacity field if non-nil, zero value otherwise.

### GetResourcesCapacityOk

`func (o *ClusterNodeDto) GetResourcesCapacityOk() (*NodeResourceDto, bool)`

GetResourcesCapacityOk returns a tuple with the ResourcesCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesCapacity

`func (o *ClusterNodeDto) SetResourcesCapacity(v NodeResourceDto)`

SetResourcesCapacity sets ResourcesCapacity field to given value.


### GetTaints

`func (o *ClusterNodeDto) GetTaints() []NodeTaintDto`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *ClusterNodeDto) GetTaintsOk() (*[]NodeTaintDto, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *ClusterNodeDto) SetTaints(v []NodeTaintDto)`

SetTaints sets Taints field to given value.


### GetUnschedulable

`func (o *ClusterNodeDto) GetUnschedulable() bool`

GetUnschedulable returns the Unschedulable field if non-nil, zero value otherwise.

### GetUnschedulableOk

`func (o *ClusterNodeDto) GetUnschedulableOk() (*bool, bool)`

GetUnschedulableOk returns a tuple with the Unschedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnschedulable

`func (o *ClusterNodeDto) SetUnschedulable(v bool)`

SetUnschedulable sets Unschedulable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


