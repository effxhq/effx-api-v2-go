# ServiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Kind** | **string** |  | 
**Spec** | [**ServiceConfigurationSpec**](ServiceConfiguration_spec.md) |  | 

## Methods

### NewServiceConfiguration

`func NewServiceConfiguration(version string, kind string, spec ServiceConfigurationSpec, ) *ServiceConfiguration`

NewServiceConfiguration instantiates a new ServiceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceConfigurationWithDefaults

`func NewServiceConfigurationWithDefaults() *ServiceConfiguration`

NewServiceConfigurationWithDefaults instantiates a new ServiceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ServiceConfiguration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceConfiguration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceConfiguration) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetKind

`func (o *ServiceConfiguration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceConfiguration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceConfiguration) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSpec

`func (o *ServiceConfiguration) GetSpec() ServiceConfigurationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ServiceConfiguration) GetSpecOk() (*ServiceConfigurationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ServiceConfiguration) SetSpec(v ServiceConfigurationSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


