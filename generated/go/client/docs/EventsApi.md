# \EventsApi

All URIs are relative to *https://api.effx.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsPost**](EventsApi.md#EventsPost) | **Post** /events | 



## EventsPost

> EventsPost(ctx).XEffxApiKey(xEffxApiKey).CreateEventPayload(createEventPayload).Execute()





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
    createEventPayload := *openapiclient.NewCreateEventPayload("Message_example", "Title_example") // CreateEventPayload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventsApi.EventsPost(context.Background()).XEffxApiKey(xEffxApiKey).CreateEventPayload(createEventPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEffxApiKey** | [**string**](string.md) | Your API Key | 
 **createEventPayload** | [**CreateEventPayload**](CreateEventPayload.md) |  | 

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

