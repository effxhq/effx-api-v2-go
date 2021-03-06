/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// ConfigApiRouter defines the required methods for binding the api requests to a responses for the ConfigApi
// The ConfigApiRouter implementation should parse necessary information from the http request, 
// pass the data to a ConfigApiServicer to perform the required actions, then write the service results to the http response.
type ConfigApiRouter interface { 
	ConfigLint(http.ResponseWriter, *http.Request)
	ConfigPut(http.ResponseWriter, *http.Request)
}
// EventsApiRouter defines the required methods for binding the api requests to a responses for the EventsApi
// The EventsApiRouter implementation should parse necessary information from the http request, 
// pass the data to a EventsApiServicer to perform the required actions, then write the service results to the http response.
type EventsApiRouter interface { 
	EventsPost(http.ResponseWriter, *http.Request)
}
// ServicesApiRouter defines the required methods for binding the api requests to a responses for the ServicesApi
// The ServicesApiRouter implementation should parse necessary information from the http request, 
// pass the data to a ServicesApiServicer to perform the required actions, then write the service results to the http response.
type ServicesApiRouter interface { 
	DetectedServicesPut(http.ResponseWriter, *http.Request)
	GetServiceById(http.ResponseWriter, *http.Request)
	ServicesGet(http.ResponseWriter, *http.Request)
	ServicesPut(http.ResponseWriter, *http.Request)
}
// TeamsApiRouter defines the required methods for binding the api requests to a responses for the TeamsApi
// The TeamsApiRouter implementation should parse necessary information from the http request, 
// pass the data to a TeamsApiServicer to perform the required actions, then write the service results to the http response.
type TeamsApiRouter interface { 
	GetTeamById(http.ResponseWriter, *http.Request)
	TeamsGet(http.ResponseWriter, *http.Request)
	TeamsPut(http.ResponseWriter, *http.Request)
}


// ConfigApiServicer defines the api actions for the ConfigApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type ConfigApiServicer interface { 
	ConfigLint(context.Context, ConfigurationFile) (ImplResponse, error)
	ConfigPut(context.Context, ConfigurationFile) (ImplResponse, error)
}


// EventsApiServicer defines the api actions for the EventsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type EventsApiServicer interface { 
	EventsPost(context.Context, string, CreateEventPayload) (ImplResponse, error)
}


// ServicesApiServicer defines the api actions for the ServicesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type ServicesApiServicer interface { 
	DetectedServicesPut(context.Context, string, DetectedServicesPayload) (ImplResponse, error)
	GetServiceById(context.Context, string, string) (ImplResponse, error)
	ServicesGet(context.Context, string, int32, int32) (ImplResponse, error)
	ServicesPut(context.Context, string, ServiceConfiguration) (ImplResponse, error)
}


// TeamsApiServicer defines the api actions for the TeamsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type TeamsApiServicer interface { 
	GetTeamById(context.Context, string, string) (ImplResponse, error)
	TeamsGet(context.Context, string, int32, int32) (ImplResponse, error)
	TeamsPut(context.Context, string, TeamConfiguration) (ImplResponse, error)
}
