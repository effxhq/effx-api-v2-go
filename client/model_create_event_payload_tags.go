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

// CreateEventPayloadTags struct for CreateEventPayloadTags
type CreateEventPayloadTags struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// NewCreateEventPayloadTags instantiates a new CreateEventPayloadTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEventPayloadTags(key string, value string, ) *CreateEventPayloadTags {
	this := CreateEventPayloadTags{}
	this.Key = key
	this.Value = value
	return &this
}

// NewCreateEventPayloadTagsWithDefaults instantiates a new CreateEventPayloadTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEventPayloadTagsWithDefaults() *CreateEventPayloadTags {
	this := CreateEventPayloadTags{}
	return &this
}

// GetKey returns the Key field value
func (o *CreateEventPayloadTags) GetKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateEventPayloadTags) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateEventPayloadTags) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *CreateEventPayloadTags) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateEventPayloadTags) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateEventPayloadTags) SetValue(v string) {
	o.Value = v
}

func (o CreateEventPayloadTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEventPayloadTags struct {
	value *CreateEventPayloadTags
	isSet bool
}

func (v NullableCreateEventPayloadTags) Get() *CreateEventPayloadTags {
	return v.value
}

func (v *NullableCreateEventPayloadTags) Set(val *CreateEventPayloadTags) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEventPayloadTags) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEventPayloadTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEventPayloadTags(val *CreateEventPayloadTags) *NullableCreateEventPayloadTags {
	return &NullableCreateEventPayloadTags{value: val, isSet: true}
}

func (v NullableCreateEventPayloadTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEventPayloadTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


