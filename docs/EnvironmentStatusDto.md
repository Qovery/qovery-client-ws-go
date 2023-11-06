# EnvironmentStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | [**[]ApplicationStatusDto**](ApplicationStatusDto.md) |  | 
**Containers** | [**[]ApplicationStatusDto**](ApplicationStatusDto.md) |  | 
**Databases** | [**[]DatabaseStatusDto**](DatabaseStatusDto.md) |  | 
**Helms** | [**[]ApplicationStatusDto**](ApplicationStatusDto.md) |  | 
**Id** | **string** |  | 
**Jobs** | [**[]ApplicationStatusDto**](ApplicationStatusDto.md) |  | 
**ProjectId** | **string** |  | 
**State** | [**ServiceStateDto**](ServiceStateDto.md) |  | 

## Methods

### NewEnvironmentStatusDto

`func NewEnvironmentStatusDto(applications []ApplicationStatusDto, containers []ApplicationStatusDto, databases []DatabaseStatusDto, helms []ApplicationStatusDto, id string, jobs []ApplicationStatusDto, projectId string, state ServiceStateDto, ) *EnvironmentStatusDto`

NewEnvironmentStatusDto instantiates a new EnvironmentStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatusDtoWithDefaults

`func NewEnvironmentStatusDtoWithDefaults() *EnvironmentStatusDto`

NewEnvironmentStatusDtoWithDefaults instantiates a new EnvironmentStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *EnvironmentStatusDto) GetApplications() []ApplicationStatusDto`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EnvironmentStatusDto) GetApplicationsOk() (*[]ApplicationStatusDto, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EnvironmentStatusDto) SetApplications(v []ApplicationStatusDto)`

SetApplications sets Applications field to given value.


### GetContainers

`func (o *EnvironmentStatusDto) GetContainers() []ApplicationStatusDto`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *EnvironmentStatusDto) GetContainersOk() (*[]ApplicationStatusDto, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *EnvironmentStatusDto) SetContainers(v []ApplicationStatusDto)`

SetContainers sets Containers field to given value.


### GetDatabases

`func (o *EnvironmentStatusDto) GetDatabases() []DatabaseStatusDto`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *EnvironmentStatusDto) GetDatabasesOk() (*[]DatabaseStatusDto, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *EnvironmentStatusDto) SetDatabases(v []DatabaseStatusDto)`

SetDatabases sets Databases field to given value.


### GetHelms

`func (o *EnvironmentStatusDto) GetHelms() []ApplicationStatusDto`

GetHelms returns the Helms field if non-nil, zero value otherwise.

### GetHelmsOk

`func (o *EnvironmentStatusDto) GetHelmsOk() (*[]ApplicationStatusDto, bool)`

GetHelmsOk returns a tuple with the Helms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelms

`func (o *EnvironmentStatusDto) SetHelms(v []ApplicationStatusDto)`

SetHelms sets Helms field to given value.


### GetId

`func (o *EnvironmentStatusDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentStatusDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentStatusDto) SetId(v string)`

SetId sets Id field to given value.


### GetJobs

`func (o *EnvironmentStatusDto) GetJobs() []ApplicationStatusDto`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *EnvironmentStatusDto) GetJobsOk() (*[]ApplicationStatusDto, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *EnvironmentStatusDto) SetJobs(v []ApplicationStatusDto)`

SetJobs sets Jobs field to given value.


### GetProjectId

`func (o *EnvironmentStatusDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnvironmentStatusDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnvironmentStatusDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetState

`func (o *EnvironmentStatusDto) GetState() ServiceStateDto`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentStatusDto) GetStateOk() (*ServiceStateDto, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentStatusDto) SetState(v ServiceStateDto)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


