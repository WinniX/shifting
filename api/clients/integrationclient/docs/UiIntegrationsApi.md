# \UiIntegrationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationUiIntegrationsByTargetByIdDelete**](UiIntegrationsApi.md#IntegrationUiIntegrationsByTargetByIdDelete) | **Delete** /integration/v1/ui-integrations/{target}/{id} | 
[**IntegrationUiIntegrationsByTargetByIdGet**](UiIntegrationsApi.md#IntegrationUiIntegrationsByTargetByIdGet) | **Get** /integration/v1/ui-integrations/{target}/{id} | 
[**IntegrationUiIntegrationsByTargetByIdPut**](UiIntegrationsApi.md#IntegrationUiIntegrationsByTargetByIdPut) | **Put** /integration/v1/ui-integrations/{target}/{id} | 
[**IntegrationUiIntegrationsByTargetByIdtestGet**](UiIntegrationsApi.md#IntegrationUiIntegrationsByTargetByIdtestGet) | **Get** /integration/v1/ui-integrations/{target}/{id}/$test | Tests a private source integration.
[**IntegrationUiIntegrationsByTargetGet**](UiIntegrationsApi.md#IntegrationUiIntegrationsByTargetGet) | **Get** /integration/v1/ui-integrations/{target} | 
[**IntegrationUiIntegrationsByTargetPost**](UiIntegrationsApi.md#IntegrationUiIntegrationsByTargetPost) | **Post** /integration/v1/ui-integrations/{target} | 
[**IntegrationUiIntegrationsGet**](UiIntegrationsApi.md#IntegrationUiIntegrationsGet) | **Get** /integration/v1/ui-integrations | 



## IntegrationUiIntegrationsByTargetByIdDelete

> IntegrationUiIntegrationsByTargetByIdDelete(ctx, target, id).Execute()





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
    target := "target_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdDelete(context.Background(), target, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsByTargetByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUiIntegrationsByTargetByIdGet

> UiIntegrationItemModel IntegrationUiIntegrationsByTargetByIdGet(ctx, target, id).Execute()





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
    target := "target_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdGet(context.Background(), target, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationUiIntegrationsByTargetByIdGet`: UiIntegrationItemModel
    fmt.Fprintf(os.Stdout, "Response from `UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsByTargetByIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UiIntegrationItemModel**](UiIntegrationItemModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUiIntegrationsByTargetByIdPut

> IntegrationUiIntegrationsByTargetByIdPut(ctx, target, id).Body(body).Execute()





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
    target := "target_example" // string | 
    id := "id_example" // string | 
    body := *openapiclient.NewReplaceUiIntegrationModel("Label_example", "SourceUrl_example", "SourceType_example") // ReplaceUiIntegrationModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdPut(context.Background(), target, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsByTargetByIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ReplaceUiIntegrationModel**](ReplaceUiIntegrationModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUiIntegrationsByTargetByIdtestGet

> UiIntegrationTestResultModel IntegrationUiIntegrationsByTargetByIdtestGet(ctx, target, id).Execute()

Tests a private source integration.



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
    target := "target_example" // string | Target for the integration
    id := "id_example" // string | Id of the integration to be tested

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdtestGet(context.Background(), target, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdtestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationUiIntegrationsByTargetByIdtestGet`: UiIntegrationTestResultModel
    fmt.Fprintf(os.Stdout, "Response from `UiIntegrationsApi.IntegrationUiIntegrationsByTargetByIdtestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** | Target for the integration | 
**id** | **string** | Id of the integration to be tested | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsByTargetByIdtestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UiIntegrationTestResultModel**](UiIntegrationTestResultModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUiIntegrationsByTargetGet

> UiIntegrationListModel IntegrationUiIntegrationsByTargetGet(ctx, target).Execute()





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
    target := "target_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsByTargetGet(context.Background(), target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsByTargetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationUiIntegrationsByTargetGet`: UiIntegrationListModel
    fmt.Fprintf(os.Stdout, "Response from `UiIntegrationsApi.IntegrationUiIntegrationsByTargetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsByTargetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UiIntegrationListModel**](UiIntegrationListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUiIntegrationsByTargetPost

> UiIntegrationCreatedModel IntegrationUiIntegrationsByTargetPost(ctx, target).Body(body).Execute()





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
    target := "target_example" // string | 
    body := *openapiclient.NewCreateUiIntegrationModel("Label_example", "SourceUrl_example", "SourceType_example") // CreateUiIntegrationModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsByTargetPost(context.Background(), target).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsByTargetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationUiIntegrationsByTargetPost`: UiIntegrationCreatedModel
    fmt.Fprintf(os.Stdout, "Response from `UiIntegrationsApi.IntegrationUiIntegrationsByTargetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsByTargetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateUiIntegrationModel**](CreateUiIntegrationModel.md) |  | 

### Return type

[**UiIntegrationCreatedModel**](UiIntegrationCreatedModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUiIntegrationsGet

> UiIntegrationListModel IntegrationUiIntegrationsGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UiIntegrationsApi.IntegrationUiIntegrationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiIntegrationsApi.IntegrationUiIntegrationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationUiIntegrationsGet`: UiIntegrationListModel
    fmt.Fprintf(os.Stdout, "Response from `UiIntegrationsApi.IntegrationUiIntegrationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUiIntegrationsGetRequest struct via the builder pattern


### Return type

[**UiIntegrationListModel**](UiIntegrationListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

