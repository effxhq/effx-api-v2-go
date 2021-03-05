# TeamConfigurationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Contact** | [**ContactInformation**](ContactInformation.md) |  | 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Attach metadata to resources like teams and services | [optional] 

## Methods

### NewTeamConfigurationSpec

`func NewTeamConfigurationSpec(name string, contact ContactInformation, ) *TeamConfigurationSpec`

NewTeamConfigurationSpec instantiates a new TeamConfigurationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamConfigurationSpecWithDefaults

`func NewTeamConfigurationSpecWithDefaults() *TeamConfigurationSpec`

NewTeamConfigurationSpecWithDefaults instantiates a new TeamConfigurationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamConfigurationSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamConfigurationSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamConfigurationSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamConfigurationSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamConfigurationSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamConfigurationSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamConfigurationSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContact

`func (o *TeamConfigurationSpec) GetContact() ContactInformation`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TeamConfigurationSpec) GetContactOk() (*ContactInformation, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TeamConfigurationSpec) SetContact(v ContactInformation)`

SetContact sets Contact field to given value.


### GetTags

`func (o *TeamConfigurationSpec) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TeamConfigurationSpec) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TeamConfigurationSpec) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TeamConfigurationSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAnnotations

`func (o *TeamConfigurationSpec) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *TeamConfigurationSpec) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *TeamConfigurationSpec) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *TeamConfigurationSpec) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


