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

// TeamConfigurationSpec struct for TeamConfigurationSpec
type TeamConfigurationSpec struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Contact ContactInformation `json:"contact"`
	Tags *map[string]string `json:"tags,omitempty"`
	// Attach metadata to resources like teams and services
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewTeamConfigurationSpec instantiates a new TeamConfigurationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamConfigurationSpec(name string, contact ContactInformation, ) *TeamConfigurationSpec {
	this := TeamConfigurationSpec{}
	this.Name = name
	this.Contact = contact
	return &this
}

// NewTeamConfigurationSpecWithDefaults instantiates a new TeamConfigurationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamConfigurationSpecWithDefaults() *TeamConfigurationSpec {
	this := TeamConfigurationSpec{}
	return &this
}

// GetName returns the Name field value
func (o *TeamConfigurationSpec) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamConfigurationSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamConfigurationSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamConfigurationSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConfigurationSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamConfigurationSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamConfigurationSpec) SetDescription(v string) {
	o.Description = &v
}

// GetContact returns the Contact field value
func (o *TeamConfigurationSpec) GetContact() ContactInformation {
	if o == nil  {
		var ret ContactInformation
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *TeamConfigurationSpec) GetContactOk() (*ContactInformation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *TeamConfigurationSpec) SetContact(v ContactInformation) {
	o.Contact = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TeamConfigurationSpec) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConfigurationSpec) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TeamConfigurationSpec) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *TeamConfigurationSpec) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *TeamConfigurationSpec) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConfigurationSpec) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *TeamConfigurationSpec) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *TeamConfigurationSpec) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o TeamConfigurationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["contact"] = o.Contact
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableTeamConfigurationSpec struct {
	value *TeamConfigurationSpec
	isSet bool
}

func (v NullableTeamConfigurationSpec) Get() *TeamConfigurationSpec {
	return v.value
}

func (v *NullableTeamConfigurationSpec) Set(val *TeamConfigurationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamConfigurationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamConfigurationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamConfigurationSpec(val *TeamConfigurationSpec) *NullableTeamConfigurationSpec {
	return &NullableTeamConfigurationSpec{value: val, isSet: true}
}

func (v NullableTeamConfigurationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamConfigurationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


