# Go API client for identityclient

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./identityclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://identity.apaleo.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InvitationApi* | [**ApiAccountInvitationsByEmailDelete**](docs/InvitationApi.md#apiaccountinvitationsbyemaildelete) | **Delete** /api/v1/account/invitations/{email} | Deletes an invitation for the current account if it exists
*InvitationApi* | [**ApiAccountInvitationsGet**](docs/InvitationApi.md#apiaccountinvitationsget) | **Get** /api/v1/account/invitations | Returns a list of invitations for the current account
*InvitationApi* | [**ApiAccountInvitationsPost**](docs/InvitationApi.md#apiaccountinvitationspost) | **Post** /api/v1/account/invitations | Invites a user to the current account
*RolesApi* | [**ApiRolesGet**](docs/RolesApi.md#apirolesget) | **Get** /api/v1/roles | Returns a static list of all roles.
*UsersApi* | [**ApiUsersByUserIdGet**](docs/UsersApi.md#apiusersbyuseridget) | **Get** /api/v1/users/{userId} | Returns a user for the current account.
*UsersApi* | [**ApiUsersByUserIdPatch**](docs/UsersApi.md#apiusersbyuseridpatch) | **Patch** /api/v1/users/{userId} | Modify user in an account.
*UsersApi* | [**ApiUsersGet**](docs/UsersApi.md#apiusersget) | **Get** /api/v1/users | Returns a list of users for the current account.


## Documentation For Models

 - [CreateInvitationModel](docs/CreateInvitationModel.md)
 - [InvitationListModel](docs/InvitationListModel.md)
 - [InvitationModel](docs/InvitationModel.md)
 - [InvitedUserToAccountResponseModel](docs/InvitedUserToAccountResponseModel.md)
 - [MessageItemCollection](docs/MessageItemCollection.md)
 - [Operation](docs/Operation.md)
 - [PropertyRolesItemModel](docs/PropertyRolesItemModel.md)
 - [RoleListModel](docs/RoleListModel.md)
 - [UserItemModel](docs/UserItemModel.md)
 - [UserModel](docs/UserModel.md)
 - [UsersListModel](docs/UsersListModel.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://identity.apaleo.com/connect/authorize
- **Scopes**: 
 - **identity:account-users.manage**: Manage account users
 - **identity:clients.manage**: Manage clients

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



