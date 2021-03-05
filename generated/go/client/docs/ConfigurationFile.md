# ConfigurationFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileContents** | **string** | An Effx Yaml file serialized into a string | 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Attach metadata to resources like teams and services | [optional] 

## Methods

### NewConfigurationFile

`func NewConfigurationFile(fileContents string, ) *ConfigurationFile`

NewConfigurationFile instantiates a new ConfigurationFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationFileWithDefaults

`func NewConfigurationFileWithDefaults() *ConfigurationFile`

NewConfigurationFileWithDefaults instantiates a new ConfigurationFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileContents

`func (o *ConfigurationFile) GetFileContents() string`

GetFileContents returns the FileContents field if non-nil, zero value otherwise.

### GetFileContentsOk

`func (o *ConfigurationFile) GetFileContentsOk() (*string, bool)`

GetFileContentsOk returns a tuple with the FileContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContents

`func (o *ConfigurationFile) SetFileContents(v string)`

SetFileContents sets FileContents field to given value.


### GetTags

`func (o *ConfigurationFile) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigurationFile) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigurationFile) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigurationFile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConfigurationFile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConfigurationFile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConfigurationFile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConfigurationFile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


