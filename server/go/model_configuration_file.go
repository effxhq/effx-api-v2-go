/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ConfigurationFile struct {

	// An Effx Yaml file serialized into a string
	FileContents string `json:"fileContents"`

	Tags map[string]string `json:"tags,omitempty"`

	// Attach metadata to resources like teams and services
	Annotations map[string]string `json:"annotations,omitempty"`
}
