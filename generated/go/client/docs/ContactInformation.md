# ContactInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**OnCall** | Pointer to [**Link**](Link.md) |  | [optional] 
**Chat** | Pointer to [**Link**](Link.md) |  | [optional] 
**IssueTracker** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewContactInformation

`func NewContactInformation() *ContactInformation`

NewContactInformation instantiates a new ContactInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactInformationWithDefaults

`func NewContactInformationWithDefaults() *ContactInformation`

NewContactInformationWithDefaults instantiates a new ContactInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ContactInformation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactInformation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactInformation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactInformation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOnCall

`func (o *ContactInformation) GetOnCall() Link`

GetOnCall returns the OnCall field if non-nil, zero value otherwise.

### GetOnCallOk

`func (o *ContactInformation) GetOnCallOk() (*Link, bool)`

GetOnCallOk returns a tuple with the OnCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCall

`func (o *ContactInformation) SetOnCall(v Link)`

SetOnCall sets OnCall field to given value.

### HasOnCall

`func (o *ContactInformation) HasOnCall() bool`

HasOnCall returns a boolean if a field has been set.

### GetChat

`func (o *ContactInformation) GetChat() Link`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ContactInformation) GetChatOk() (*Link, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ContactInformation) SetChat(v Link)`

SetChat sets Chat field to given value.

### HasChat

`func (o *ContactInformation) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetIssueTracker

`func (o *ContactInformation) GetIssueTracker() Link`

GetIssueTracker returns the IssueTracker field if non-nil, zero value otherwise.

### GetIssueTrackerOk

`func (o *ContactInformation) GetIssueTrackerOk() (*Link, bool)`

GetIssueTrackerOk returns a tuple with the IssueTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTracker

`func (o *ContactInformation) SetIssueTracker(v Link)`

SetIssueTracker sets IssueTracker field to given value.

### HasIssueTracker

`func (o *ContactInformation) HasIssueTracker() bool`

HasIssueTracker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


