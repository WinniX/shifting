# \InvitationApi

All URIs are relative to *https://identity.apaleo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountInvitationsByEmailDelete**](InvitationApi.md#ApiAccountInvitationsByEmailDelete) | **Delete** /api/v1/account/invitations/{email} | Deletes an invitation for the current account if it exists
[**ApiAccountInvitationsGet**](InvitationApi.md#ApiAccountInvitationsGet) | **Get** /api/v1/account/invitations | Returns a list of invitations for the current account
[**ApiAccountInvitationsPost**](InvitationApi.md#ApiAccountInvitationsPost) | **Post** /api/v1/account/invitations | Invites a user to the current account



## ApiAccountInvitationsByEmailDelete

> ApiAccountInvitationsByEmailDelete(ctx, email).Execute()

Deletes an invitation for the current account if it exists



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
    email := "email_example" // string | email of the invited person

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationApi.ApiAccountInvitationsByEmailDelete(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationApi.ApiAccountInvitationsByEmailDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | email of the invited person | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountInvitationsByEmailDeleteRequest struct via the builder pattern


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


## ApiAccountInvitationsGet

> InvitationListModel ApiAccountInvitationsGet(ctx).PropertyId(propertyId).Execute()

Returns a list of invitations for the current account



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
    propertyId := "propertyId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationApi.ApiAccountInvitationsGet(context.Background()).PropertyId(propertyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationApi.ApiAccountInvitationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountInvitationsGet`: InvitationListModel
    fmt.Fprintf(os.Stdout, "Response from `InvitationApi.ApiAccountInvitationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountInvitationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **propertyId** | **string** |  | 

### Return type

[**InvitationListModel**](InvitationListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountInvitationsPost

> ApiAccountInvitationsPost(ctx).Body(body).Execute()

Invites a user to the current account



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
    body := *openapiclient.NewCreateInvitationModel("Email_example") // CreateInvitationModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationApi.ApiAccountInvitationsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationApi.ApiAccountInvitationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountInvitationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateInvitationModel**](CreateInvitationModel.md) |  | 

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

