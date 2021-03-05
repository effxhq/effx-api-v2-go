# ServiceConfigurationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Contact** | [**ContactInformation**](ContactInformation.md) |  | 
**LinkGroups** | Pointer to [**[]LinkGroup**](LinkGroup.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Attach metadata to resources like teams and services | [optional] 
**Dependencies** | Pointer to [**Dependencies**](Dependencies.md) |  | [optional] 

## Methods

### NewServiceConfigurationSpec

`func NewServiceConfigurationSpec(name string, contact ContactInformation, ) *ServiceConfigurationSpec`

NewServiceConfigurationSpec instantiates a new ServiceConfigurationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConfigurationSpecWithDefaults

`func NewServiceConfigurationSpecWithDefaults() *ServiceConfigurationSpec`

NewServiceConfigurationSpecWithDefaults instantiates a new ServiceConfigurationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceConfigurationSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceConfigurationSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceConfigurationSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceConfigurationSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceConfigurationSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceConfigurationSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceConfigurationSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContact

`func (o *ServiceConfigurationSpec) GetContact() ContactInformation`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ServiceConfigurationSpec) GetContactOk() (*ContactInformation, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ServiceConfigurationSpec) SetContact(v ContactInformation)`

SetContact sets Contact field to given value.


### GetLinkGroups

`func (o *ServiceConfigurationSpec) GetLinkGroups() []LinkGroup`

GetLinkGroups returns the LinkGroups field if non-nil, zero value otherwise.

### GetLinkGroupsOk

`func (o *ServiceConfigurationSpec) GetLinkGroupsOk() (*[]LinkGroup, bool)`

GetLinkGroupsOk returns a tuple with the LinkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkGroups

`func (o *ServiceConfigurationSpec) SetLinkGroups(v []LinkGroup)`

SetLinkGroups sets LinkGroups field to given value.

### HasLinkGroups

`func (o *ServiceConfigurationSpec) HasLinkGroups() bool`

HasLinkGroups returns a boolean if a field has been set.

### GetTags

`func (o *ServiceConfigurationSpec) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceConfigurationSpec) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceConfigurationSpec) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceConfigurationSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAnnotations

`func (o *ServiceConfigurationSpec) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ServiceConfigurationSpec) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ServiceConfigurationSpec) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ServiceConfigurationSpec) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDependencies

`func (o *ServiceConfigurationSpec) GetDependencies() Dependencies`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ServiceConfigurationSpec) GetDependenciesOk() (*Dependencies, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ServiceConfigurationSpec) SetDependencies(v Dependencies)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *ServiceConfigurationSpec) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


