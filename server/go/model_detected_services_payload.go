/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DetectedServicesPayload struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	SourceName string `json:"sourceName,omitempty"`

	Email string `json:"email,omitempty"`

	OnCall Link `json:"onCall,omitempty"`

	Chat Link `json:"chat,omitempty"`

	IssueTracker Link `json:"issueTracker,omitempty"`

	LinkGroups []LinkGroup `json:"linkGroups,omitempty"`

	Tags map[string]string `json:"tags,omitempty"`

	// Attach metadata to resources like teams and services
	Annotations map[string]string `json:"annotations,omitempty"`
}
