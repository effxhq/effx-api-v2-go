/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ManualDependency - An explicit dependency on another entity
type ManualDependency struct {

	// The kind of entity to depend on. Defaults to 'service'
	Kind string `json:"kind,omitempty"`

	// Matches a dependency by the kind's name
	Name string `json:"name,omitempty"`
}
