# \UsersApi

All URIs are relative to *https://identity.apaleo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUsersByUserIdGet**](UsersApi.md#ApiUsersByUserIdGet) | **Get** /api/v1/users/{userId} | Returns a user for the current account.
[**ApiUsersByUserIdPatch**](UsersApi.md#ApiUsersByUserIdPatch) | **Patch** /api/v1/users/{userId} | Modify user in an account.
[**ApiUsersGet**](UsersApi.md#ApiUsersGet) | **Get** /api/v1/users | Returns a list of users for the current account.



## ApiUsersByUserIdGet

> UserModel ApiUsersByUserIdGet(ctx, userId).Execute()

Returns a user for the current account.



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
    userId := "userId_example" // string | User's subjectId.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ApiUsersByUserIdGet(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUsersByUserIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersByUserIdGet`: UserModel
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUsersByUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User&#39;s subjectId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersByUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserModel**](UserModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersByUserIdPatch

> ApiUsersByUserIdPatch(ctx, userId).Body(body).Execute()

Modify user in an account.



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
    userId := "userId_example" // string | User's subjectId
    body := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ApiUsersByUserIdPatch(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUsersByUserIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User&#39;s subjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersByUserIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]Operation**](Operation.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersGet

> UsersListModel ApiUsersGet(ctx).PropertyIds(propertyIds).Email(email).Execute()

Returns a list of users for the current account.



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
    propertyIds := []string{"Inner_example"} // []string | List of property ids to filter users by (optional)
    email := "email_example" // string | Filter users by email (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ApiUsersGet(context.Background()).PropertyIds(propertyIds).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ApiUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersGet`: UsersListModel
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ApiUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **propertyIds** | **[]string** | List of property ids to filter users by | 
 **email** | **string** | Filter users by email | 

### Return type

[**UsersListModel**](UsersListModel.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

