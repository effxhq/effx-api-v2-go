# \ServicesApi

All URIs are relative to *https://api.effx.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetectedServicesPut**](ServicesApi.md#DetectedServicesPut) | **Put** /detected_services | 
[**GetServiceById**](ServicesApi.md#GetServiceById) | **Get** /services/{serviceId} | Info for a specific service
[**ServicesGet**](ServicesApi.md#ServicesGet) | **Get** /services | 
[**ServicesPut**](ServicesApi.md#ServicesPut) | **Put** /services | 



## DetectedServicesPut

> DetectedServicesPut(ctx).XEffxApiKey(xEffxApiKey).DetectedServicesPayload(detectedServicesPayload).Execute()





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
    detectedServicesPayload := *openapiclient.NewDetectedServicesPayload("Name_example") // DetectedServicesPayload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DetectedServicesPut(context.Background()).XEffxApiKey(xEffxApiKey).DetectedServicesPayload(detectedServicesPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DetectedServicesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetectedServicesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 
 **detectedServicesPayload** | [**DetectedServicesPayload**](DetectedServicesPayload.md) |  | 

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


## GetServiceById

> Service GetServiceById(ctx, serviceId).XEffxApiKey(xEffxApiKey).Execute()

Info for a specific service

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
    serviceId := "serviceId_example" // string | The id or slug of the service to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetServiceById(context.Background(), serviceId).XEffxApiKey(xEffxApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceById`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetServiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The id or slug of the service to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 


### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesGet

> InlineResponse200 ServicesGet(ctx).XEffxApiKey(xEffxApiKey).Limit(limit).Offset(offset).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesGet(context.Background()).XEffxApiKey(xEffxApiKey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 
 **limit** | **int32** | How many records to retrieve per request. | [default to 50]
 **offset** | **int32** | The number of records to skip. Used for pagination through a list of services. | [default to 0]

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesPut

> ServicesPut(ctx).XEffxApiKey(xEffxApiKey).ServiceConfiguration(serviceConfiguration).Execute()





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
    serviceConfiguration := *openapiclient.NewServiceConfiguration("effx/v1", "Service", *openapiclient.NewServiceConfigurationSpec("the canonical name of the service", *openapiclient.NewContactInformation())) // ServiceConfiguration | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesPut(context.Background()).XEffxApiKey(xEffxApiKey).ServiceConfiguration(serviceConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 
 **serviceConfiguration** | [**ServiceConfiguration**](ServiceConfiguration.md) |  | 

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

