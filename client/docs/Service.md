# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | uuid of the service object | [optional] 
**Slug** | Pointer to **string** | url safe string representation of the service&#39;s name | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to [**ContactInformation**](ContactInformation.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Attach metadata to resources like teams and services | [optional] 
**LinkGroups** | Pointer to [**[]LinkGroup**](LinkGroup.md) |  | [optional] 
**Dependencies** | Pointer to [**Dependencies**](Dependencies.md) |  | [optional] 
**OwnerId** | Pointer to **NullableString** | team id that owns this service. | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *Service) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Service) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Service) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Service) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContact

`func (o *Service) GetContact() ContactInformation`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Service) GetContactOk() (*ContactInformation, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Service) SetContact(v ContactInformation)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Service) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetTags

`func (o *Service) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Service) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAnnotations

`func (o *Service) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Service) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Service) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Service) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLinkGroups

`func (o *Service) GetLinkGroups() []LinkGroup`

GetLinkGroups returns the LinkGroups field if non-nil, zero value otherwise.

### GetLinkGroupsOk

`func (o *Service) GetLinkGroupsOk() (*[]LinkGroup, bool)`

GetLinkGroupsOk returns a tuple with the LinkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkGroups

`func (o *Service) SetLinkGroups(v []LinkGroup)`

SetLinkGroups sets LinkGroups field to given value.

### HasLinkGroups

`func (o *Service) HasLinkGroups() bool`

HasLinkGroups returns a boolean if a field has been set.

### GetDependencies

`func (o *Service) GetDependencies() Dependencies`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *Service) GetDependenciesOk() (*Dependencies, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *Service) SetDependencies(v Dependencies)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *Service) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetOwnerId

`func (o *Service) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Service) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Service) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Service) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *Service) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *Service) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


