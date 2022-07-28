# \RolesApi

All URIs are relative to *https://identity.apaleo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRolesGet**](RolesApi.md#ApiRolesGet) | **Get** /api/v1/roles | Returns a static list of all roles.



## ApiRolesGet

> RoleListModel ApiRolesGet(ctx).Execute()

Returns a static list of all roles.

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
    resp, r, err := api_client.RolesApi.ApiRolesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ApiRolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRolesGet`: RoleListModel
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ApiRolesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRolesGetRequest struct via the builder pattern


### Return type

[**RoleListModel**](RoleListModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

