/*
 * effx API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/effxhq/effx-api-v2-go/generated/go/client/go"
)

func main() {
	log.Printf("Server started")

	ConfigApiService := openapi.NewConfigApiService()
	ConfigApiController := openapi.NewConfigApiController(ConfigApiService)

	EventsApiService := openapi.NewEventsApiService()
	EventsApiController := openapi.NewEventsApiController(EventsApiService)

	ServicesApiService := openapi.NewServicesApiService()
	ServicesApiController := openapi.NewServicesApiController(ServicesApiService)

	TeamsApiService := openapi.NewTeamsApiService()
	TeamsApiController := openapi.NewTeamsApiController(TeamsApiService)

	router := openapi.NewRouter(ConfigApiController, EventsApiController, ServicesApiController, TeamsApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
