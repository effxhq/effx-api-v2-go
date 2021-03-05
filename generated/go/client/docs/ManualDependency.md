# ManualDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind of entity to depend on. Defaults to &#39;service&#39; | [optional] [default to "service"]
**Name** | Pointer to **string** | Matches a dependency by the kind&#39;s name | [optional] 

## Methods

### NewManualDependency

`func NewManualDependency() *ManualDependency`

NewManualDependency instantiates a new ManualDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualDependencyWithDefaults

`func NewManualDependencyWithDefaults() *ManualDependency`

NewManualDependencyWithDefaults instantiates a new ManualDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ManualDependency) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManualDependency) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManualDependency) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ManualDependency) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ManualDependency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManualDependency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManualDependency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManualDependency) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


