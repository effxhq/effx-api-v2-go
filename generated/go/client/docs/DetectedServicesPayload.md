# DetectedServicesPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OnCall** | Pointer to [**Link**](Link.md) |  | [optional] 
**Chat** | Pointer to [**Link**](Link.md) |  | [optional] 
**IssueTracker** | Pointer to [**Link**](Link.md) |  | [optional] 
**LinkGroups** | Pointer to [**[]LinkGroup**](LinkGroup.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** | Attach metadata to resources like teams and services | [optional] 

## Methods

### NewDetectedServicesPayload

`func NewDetectedServicesPayload(name string, ) *DetectedServicesPayload`

NewDetectedServicesPayload instantiates a new DetectedServicesPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectedServicesPayloadWithDefaults

`func NewDetectedServicesPayloadWithDefaults() *DetectedServicesPayload`

NewDetectedServicesPayloadWithDefaults instantiates a new DetectedServicesPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DetectedServicesPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetectedServicesPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetectedServicesPayload) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DetectedServicesPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DetectedServicesPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DetectedServicesPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DetectedServicesPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceName

`func (o *DetectedServicesPayload) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DetectedServicesPayload) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DetectedServicesPayload) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *DetectedServicesPayload) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetEmail

`func (o *DetectedServicesPayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DetectedServicesPayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DetectedServicesPayload) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DetectedServicesPayload) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOnCall

`func (o *DetectedServicesPayload) GetOnCall() Link`

GetOnCall returns the OnCall field if non-nil, zero value otherwise.

### GetOnCallOk

`func (o *DetectedServicesPayload) GetOnCallOk() (*Link, bool)`

GetOnCallOk returns a tuple with the OnCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCall

`func (o *DetectedServicesPayload) SetOnCall(v Link)`

SetOnCall sets OnCall field to given value.

### HasOnCall

`func (o *DetectedServicesPayload) HasOnCall() bool`

HasOnCall returns a boolean if a field has been set.

### GetChat

`func (o *DetectedServicesPayload) GetChat() Link`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *DetectedServicesPayload) GetChatOk() (*Link, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *DetectedServicesPayload) SetChat(v Link)`

SetChat sets Chat field to given value.

### HasChat

`func (o *DetectedServicesPayload) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetIssueTracker

`func (o *DetectedServicesPayload) GetIssueTracker() Link`

GetIssueTracker returns the IssueTracker field if non-nil, zero value otherwise.

### GetIssueTrackerOk

`func (o *DetectedServicesPayload) GetIssueTrackerOk() (*Link, bool)`

GetIssueTrackerOk returns a tuple with the IssueTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTracker

`func (o *DetectedServicesPayload) SetIssueTracker(v Link)`

SetIssueTracker sets IssueTracker field to given value.

### HasIssueTracker

`func (o *DetectedServicesPayload) HasIssueTracker() bool`

HasIssueTracker returns a boolean if a field has been set.

### GetLinkGroups

`func (o *DetectedServicesPayload) GetLinkGroups() []LinkGroup`

GetLinkGroups returns the LinkGroups field if non-nil, zero value otherwise.

### GetLinkGroupsOk

`func (o *DetectedServicesPayload) GetLinkGroupsOk() (*[]LinkGroup, bool)`

GetLinkGroupsOk returns a tuple with the LinkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkGroups

`func (o *DetectedServicesPayload) SetLinkGroups(v []LinkGroup)`

SetLinkGroups sets LinkGroups field to given value.

### HasLinkGroups

`func (o *DetectedServicesPayload) HasLinkGroups() bool`

HasLinkGroups returns a boolean if a field has been set.

### GetTags

`func (o *DetectedServicesPayload) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetectedServicesPayload) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetectedServicesPayload) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetectedServicesPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAnnotations

`func (o *DetectedServicesPayload) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DetectedServicesPayload) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DetectedServicesPayload) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DetectedServicesPayload) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


