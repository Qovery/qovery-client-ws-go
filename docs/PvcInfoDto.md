# PvcInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**DiskMibCapacity** | **int32** |  | 
**DiskMibUsage** | **int32** |  | 
**DiskPercentUsage** | **int32** |  | 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**PodName** | **string** |  | 
**QoveryServiceInfo** | Pointer to [**NullablePodQoveryServiceInfoDto**](PodQoveryServiceInfoDto.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPvcInfoDto

`func NewPvcInfoDto(createdAt int64, diskMibCapacity int32, diskMibUsage int32, diskPercentUsage int32, name string, namespace string, podName string, ) *PvcInfoDto`

NewPvcInfoDto instantiates a new PvcInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvcInfoDtoWithDefaults

`func NewPvcInfoDtoWithDefaults() *PvcInfoDto`

NewPvcInfoDtoWithDefaults instantiates a new PvcInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PvcInfoDto) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PvcInfoDto) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PvcInfoDto) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetDiskMibCapacity

`func (o *PvcInfoDto) GetDiskMibCapacity() int32`

GetDiskMibCapacity returns the DiskMibCapacity field if non-nil, zero value otherwise.

### GetDiskMibCapacityOk

`func (o *PvcInfoDto) GetDiskMibCapacityOk() (*int32, bool)`

GetDiskMibCapacityOk returns a tuple with the DiskMibCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMibCapacity

`func (o *PvcInfoDto) SetDiskMibCapacity(v int32)`

SetDiskMibCapacity sets DiskMibCapacity field to given value.


### GetDiskMibUsage

`func (o *PvcInfoDto) GetDiskMibUsage() int32`

GetDiskMibUsage returns the DiskMibUsage field if non-nil, zero value otherwise.

### GetDiskMibUsageOk

`func (o *PvcInfoDto) GetDiskMibUsageOk() (*int32, bool)`

GetDiskMibUsageOk returns a tuple with the DiskMibUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMibUsage

`func (o *PvcInfoDto) SetDiskMibUsage(v int32)`

SetDiskMibUsage sets DiskMibUsage field to given value.


### GetDiskPercentUsage

`func (o *PvcInfoDto) GetDiskPercentUsage() int32`

GetDiskPercentUsage returns the DiskPercentUsage field if non-nil, zero value otherwise.

### GetDiskPercentUsageOk

`func (o *PvcInfoDto) GetDiskPercentUsageOk() (*int32, bool)`

GetDiskPercentUsageOk returns a tuple with the DiskPercentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPercentUsage

`func (o *PvcInfoDto) SetDiskPercentUsage(v int32)`

SetDiskPercentUsage sets DiskPercentUsage field to given value.


### GetName

`func (o *PvcInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PvcInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PvcInfoDto) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *PvcInfoDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PvcInfoDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PvcInfoDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPodName

`func (o *PvcInfoDto) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *PvcInfoDto) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *PvcInfoDto) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetQoveryServiceInfo

`func (o *PvcInfoDto) GetQoveryServiceInfo() PodQoveryServiceInfoDto`

GetQoveryServiceInfo returns the QoveryServiceInfo field if non-nil, zero value otherwise.

### GetQoveryServiceInfoOk

`func (o *PvcInfoDto) GetQoveryServiceInfoOk() (*PodQoveryServiceInfoDto, bool)`

GetQoveryServiceInfoOk returns a tuple with the QoveryServiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryServiceInfo

`func (o *PvcInfoDto) SetQoveryServiceInfo(v PodQoveryServiceInfoDto)`

SetQoveryServiceInfo sets QoveryServiceInfo field to given value.

### HasQoveryServiceInfo

`func (o *PvcInfoDto) HasQoveryServiceInfo() bool`

HasQoveryServiceInfo returns a boolean if a field has been set.

### SetQoveryServiceInfoNil

`func (o *PvcInfoDto) SetQoveryServiceInfoNil(b bool)`

 SetQoveryServiceInfoNil sets the value for QoveryServiceInfo to be an explicit nil

### UnsetQoveryServiceInfo
`func (o *PvcInfoDto) UnsetQoveryServiceInfo()`

UnsetQoveryServiceInfo ensures that no value is present for QoveryServiceInfo, not even an explicit nil
### GetStatus

`func (o *PvcInfoDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PvcInfoDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PvcInfoDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PvcInfoDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PvcInfoDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PvcInfoDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


