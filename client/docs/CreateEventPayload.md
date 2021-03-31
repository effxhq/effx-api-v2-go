# CreateEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Title** | **string** |  | 
**TimestampMilliseconds** | Pointer to **int64** | the timestamp in milliseconds on which this event occurred. must be greater than 01/01/2020. If not specified, it will default to the timestamp on which the request was received. | [optional] 
**Tags** | Pointer to [**[]CreateEventPayloadTags**](CreateEventPayloadTags.md) |  | [optional] 
**Actions** | Pointer to [**[]CreateEventPayloadActions**](CreateEventPayloadActions.md) |  | [optional] 
**ServiceName** | Pointer to **string** | the service&#39;s name that this event belongs to | [optional] 
**Integration** | Pointer to [**CreateEventPayloadIntegration**](CreateEventPayload_integration.md) |  | [optional] 

## Methods

### NewCreateEventPayload

`func NewCreateEventPayload(message string, title string, ) *CreateEventPayload`

NewCreateEventPayload instantiates a new CreateEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventPayloadWithDefaults

`func NewCreateEventPayloadWithDefaults() *CreateEventPayload`

NewCreateEventPayloadWithDefaults instantiates a new CreateEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateEventPayload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateEventPayload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateEventPayload) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *CreateEventPayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateEventPayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateEventPayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTimestampMilliseconds

`func (o *CreateEventPayload) GetTimestampMilliseconds() int64`

GetTimestampMilliseconds returns the TimestampMilliseconds field if non-nil, zero value otherwise.

### GetTimestampMillisecondsOk

`func (o *CreateEventPayload) GetTimestampMillisecondsOk() (*int64, bool)`

GetTimestampMillisecondsOk returns a tuple with the TimestampMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMilliseconds

`func (o *CreateEventPayload) SetTimestampMilliseconds(v int64)`

SetTimestampMilliseconds sets TimestampMilliseconds field to given value.

### HasTimestampMilliseconds

`func (o *CreateEventPayload) HasTimestampMilliseconds() bool`

HasTimestampMilliseconds returns a boolean if a field has been set.

### GetTags

`func (o *CreateEventPayload) GetTags() []CreateEventPayloadTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateEventPayload) GetTagsOk() (*[]CreateEventPayloadTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateEventPayload) SetTags(v []CreateEventPayloadTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateEventPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetActions

`func (o *CreateEventPayload) GetActions() []CreateEventPayloadActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateEventPayload) GetActionsOk() (*[]CreateEventPayloadActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateEventPayload) SetActions(v []CreateEventPayloadActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CreateEventPayload) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetServiceName

`func (o *CreateEventPayload) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CreateEventPayload) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CreateEventPayload) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CreateEventPayload) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetIntegration

`func (o *CreateEventPayload) GetIntegration() CreateEventPayloadIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CreateEventPayload) GetIntegrationOk() (*CreateEventPayloadIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CreateEventPayload) SetIntegration(v CreateEventPayloadIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CreateEventPayload) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


