# \TeamsApi

All URIs are relative to *https://api.effx.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTeamById**](TeamsApi.md#GetTeamById) | **Get** /teams/{teamId} | Info for a specific team
[**TeamsGet**](TeamsApi.md#TeamsGet) | **Get** /teams | 
[**TeamsPut**](TeamsApi.md#TeamsPut) | **Put** /teams | 



## GetTeamById

> Team GetTeamById(ctx, teamId).XEffxApiKey(xEffxApiKey).Execute()

Info for a specific team

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xEffxApiKey := TODO // string | Your API Key
    teamId := "teamId_example" // string | The id or slug of the team to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.GetTeamById(context.Background(), teamId).XEffxApiKey(xEffxApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeamById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamById`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetTeamById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The id or slug of the team to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 


### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGet

> InlineResponse2001 TeamsGet(ctx).XEffxApiKey(xEffxApiKey).Limit(limit).Offset(offset).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xEffxApiKey := TODO // string | Your API Key
    limit := int32(56) // int32 | How many records to retrieve per request. (optional) (default to 50)
    offset := int32(56) // int32 | The number of records to skip. Used for pagination through a list of services. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.TeamsGet(context.Background()).XEffxApiKey(xEffxApiKey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 
 **limit** | **int32** | How many records to retrieve per request. | [default to 50]
 **offset** | **int32** | The number of records to skip. Used for pagination through a list of services. | [default to 0]

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPut

> TeamsPut(ctx).XEffxApiKey(xEffxApiKey).TeamConfiguration(teamConfiguration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xEffxApiKey := TODO // string | Your API Key
    teamConfiguration := *openapiclient.NewTeamConfiguration("effx/v1", "Team", *openapiclient.NewTeamConfigurationSpec("authentication team", *openapiclient.NewContactInformation())) // TeamConfiguration | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.TeamsPut(context.Background()).XEffxApiKey(xEffxApiKey).TeamConfiguration(teamConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 
 **teamConfiguration** | [**TeamConfiguration**](TeamConfiguration.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

