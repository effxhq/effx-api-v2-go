# Dependencies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manual** | Pointer to [**[]ManualDependency**](ManualDependency.md) |  | [optional] 

## Methods

### NewDependencies

`func NewDependencies() *Dependencies`

NewDependencies instantiates a new Dependencies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependenciesWithDefaults

`func NewDependenciesWithDefaults() *Dependencies`

NewDependenciesWithDefaults instantiates a new Dependencies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManual

`func (o *Dependencies) GetManual() []ManualDependency`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *Dependencies) GetManualOk() (*[]ManualDependency, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *Dependencies) SetManual(v []ManualDependency)`

SetManual sets Manual field to given value.

### HasManual

`func (o *Dependencies) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


