/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateEventPayload struct for CreateEventPayload
type CreateEventPayload struct {
	Message string `json:"message"`
	Title string `json:"title"`
	// the timestamp in milliseconds on which this event occurred. must be greater than 01/01/2020. If not specified, it will default to the timestamp on which the request was received.
	TimestampMilliseconds *int64 `json:"timestamp_milliseconds,omitempty"`
	Tags *[]CreateEventPayloadTags `json:"tags,omitempty"`
	Actions *[]CreateEventPayloadActions `json:"actions,omitempty"`
	// the service's name that this event belongs to
	ServiceName *string `json:"service_name,omitempty"`
	Integration *CreateEventPayloadIntegration `json:"integration,omitempty"`
}

// NewCreateEventPayload instantiates a new CreateEventPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEventPayload(message string, title string, ) *CreateEventPayload {
	this := CreateEventPayload{}
	this.Message = message
	this.Title = title
	return &this
}

// NewCreateEventPayloadWithDefaults instantiates a new CreateEventPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEventPayloadWithDefaults() *CreateEventPayload {
	this := CreateEventPayload{}
	return &this
}

// GetMessage returns the Message field value
func (o *CreateEventPayload) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateEventPayload) SetMessage(v string) {
	o.Message = v
}

// GetTitle returns the Title field value
func (o *CreateEventPayload) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateEventPayload) SetTitle(v string) {
	o.Title = v
}

// GetTimestampMilliseconds returns the TimestampMilliseconds field value if set, zero value otherwise.
func (o *CreateEventPayload) GetTimestampMilliseconds() int64 {
	if o == nil || o.TimestampMilliseconds == nil {
		var ret int64
		return ret
	}
	return *o.TimestampMilliseconds
}

// GetTimestampMillisecondsOk returns a tuple with the TimestampMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetTimestampMillisecondsOk() (*int64, bool) {
	if o == nil || o.TimestampMilliseconds == nil {
		return nil, false
	}
	return o.TimestampMilliseconds, true
}

// HasTimestampMilliseconds returns a boolean if a field has been set.
func (o *CreateEventPayload) HasTimestampMilliseconds() bool {
	if o != nil && o.TimestampMilliseconds != nil {
		return true
	}

	return false
}

// SetTimestampMilliseconds gets a reference to the given int64 and assigns it to the TimestampMilliseconds field.
func (o *CreateEventPayload) SetTimestampMilliseconds(v int64) {
	o.TimestampMilliseconds = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateEventPayload) GetTags() []CreateEventPayloadTags {
	if o == nil || o.Tags == nil {
		var ret []CreateEventPayloadTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetTagsOk() (*[]CreateEventPayloadTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateEventPayload) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []CreateEventPayloadTags and assigns it to the Tags field.
func (o *CreateEventPayload) SetTags(v []CreateEventPayloadTags) {
	o.Tags = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CreateEventPayload) GetActions() []CreateEventPayloadActions {
	if o == nil || o.Actions == nil {
		var ret []CreateEventPayloadActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetActionsOk() (*[]CreateEventPayloadActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CreateEventPayload) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []CreateEventPayloadActions and assigns it to the Actions field.
func (o *CreateEventPayload) SetActions(v []CreateEventPayloadActions) {
	o.Actions = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *CreateEventPayload) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *CreateEventPayload) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *CreateEventPayload) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *CreateEventPayload) GetIntegration() CreateEventPayloadIntegration {
	if o == nil || o.Integration == nil {
		var ret CreateEventPayloadIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEventPayload) GetIntegrationOk() (*CreateEventPayloadIntegration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *CreateEventPayload) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given CreateEventPayloadIntegration and assigns it to the Integration field.
func (o *CreateEventPayload) SetIntegration(v CreateEventPayloadIntegration) {
	o.Integration = &v
}

func (o CreateEventPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.TimestampMilliseconds != nil {
		toSerialize["timestamp_milliseconds"] = o.TimestampMilliseconds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEventPayload struct {
	value *CreateEventPayload
	isSet bool
}

func (v NullableCreateEventPayload) Get() *CreateEventPayload {
	return v.value
}

func (v *NullableCreateEventPayload) Set(val *CreateEventPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEventPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEventPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEventPayload(val *CreateEventPayload) *NullableCreateEventPayload {
	return &NullableCreateEventPayload{value: val, isSet: true}
}

func (v NullableCreateEventPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEventPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


