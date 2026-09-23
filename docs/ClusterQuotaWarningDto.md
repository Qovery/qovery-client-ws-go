# ClusterQuotaWarningDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectedAt** | **int64** |  | 
**LastSeenAt** | **int64** |  | 
**Message** | **string** |  | 
**Provider** | **string** |  | 
**QuotaCode** | **string** |  | 
**QuotaName** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Resource** | Pointer to **NullableString** |  | [optional] 
**Source** | **string** |  | 
**Status** | **string** |  | 
**SuggestedAction** | **string** |  | 

## Methods

### NewClusterQuotaWarningDto

`func NewClusterQuotaWarningDto(detectedAt int64, lastSeenAt int64, message string, provider string, quotaCode string, source string, status string, suggestedAction string, ) *ClusterQuotaWarningDto`

NewClusterQuotaWarningDto instantiates a new ClusterQuotaWarningDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterQuotaWarningDtoWithDefaults

`func NewClusterQuotaWarningDtoWithDefaults() *ClusterQuotaWarningDto`

NewClusterQuotaWarningDtoWithDefaults instantiates a new ClusterQuotaWarningDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectedAt

`func (o *ClusterQuotaWarningDto) GetDetectedAt() int64`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *ClusterQuotaWarningDto) GetDetectedAtOk() (*int64, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *ClusterQuotaWarningDto) SetDetectedAt(v int64)`

SetDetectedAt sets DetectedAt field to given value.


### GetLastSeenAt

`func (o *ClusterQuotaWarningDto) GetLastSeenAt() int64`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ClusterQuotaWarningDto) GetLastSeenAtOk() (*int64, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ClusterQuotaWarningDto) SetLastSeenAt(v int64)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetMessage

`func (o *ClusterQuotaWarningDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterQuotaWarningDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterQuotaWarningDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetProvider

`func (o *ClusterQuotaWarningDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterQuotaWarningDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterQuotaWarningDto) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetQuotaCode

`func (o *ClusterQuotaWarningDto) GetQuotaCode() string`

GetQuotaCode returns the QuotaCode field if non-nil, zero value otherwise.

### GetQuotaCodeOk

`func (o *ClusterQuotaWarningDto) GetQuotaCodeOk() (*string, bool)`

GetQuotaCodeOk returns a tuple with the QuotaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaCode

`func (o *ClusterQuotaWarningDto) SetQuotaCode(v string)`

SetQuotaCode sets QuotaCode field to given value.


### GetQuotaName

`func (o *ClusterQuotaWarningDto) GetQuotaName() string`

GetQuotaName returns the QuotaName field if non-nil, zero value otherwise.

### GetQuotaNameOk

`func (o *ClusterQuotaWarningDto) GetQuotaNameOk() (*string, bool)`

GetQuotaNameOk returns a tuple with the QuotaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaName

`func (o *ClusterQuotaWarningDto) SetQuotaName(v string)`

SetQuotaName sets QuotaName field to given value.

### HasQuotaName

`func (o *ClusterQuotaWarningDto) HasQuotaName() bool`

HasQuotaName returns a boolean if a field has been set.

### SetQuotaNameNil

`func (o *ClusterQuotaWarningDto) SetQuotaNameNil(b bool)`

 SetQuotaNameNil sets the value for QuotaName to be an explicit nil

### UnsetQuotaName
`func (o *ClusterQuotaWarningDto) UnsetQuotaName()`

UnsetQuotaName ensures that no value is present for QuotaName, not even an explicit nil
### GetRegion

`func (o *ClusterQuotaWarningDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterQuotaWarningDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterQuotaWarningDto) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterQuotaWarningDto) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ClusterQuotaWarningDto) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ClusterQuotaWarningDto) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetResource

`func (o *ClusterQuotaWarningDto) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ClusterQuotaWarningDto) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ClusterQuotaWarningDto) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ClusterQuotaWarningDto) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ClusterQuotaWarningDto) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ClusterQuotaWarningDto) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetSource

`func (o *ClusterQuotaWarningDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClusterQuotaWarningDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClusterQuotaWarningDto) SetSource(v string)`

SetSource sets Source field to given value.


### GetStatus

`func (o *ClusterQuotaWarningDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterQuotaWarningDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterQuotaWarningDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuggestedAction

`func (o *ClusterQuotaWarningDto) GetSuggestedAction() string`

GetSuggestedAction returns the SuggestedAction field if non-nil, zero value otherwise.

### GetSuggestedActionOk

`func (o *ClusterQuotaWarningDto) GetSuggestedActionOk() (*string, bool)`

GetSuggestedActionOk returns a tuple with the SuggestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAction

`func (o *ClusterQuotaWarningDto) SetSuggestedAction(v string)`

SetSuggestedAction sets SuggestedAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


