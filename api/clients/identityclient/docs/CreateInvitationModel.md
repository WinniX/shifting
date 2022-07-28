# CreateInvitationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address, used as the login | 
**Properties** | Pointer to **[]string** | Properties that user has access to. If the user is invited as account admin, this will be ignored. | [optional] 
**IsAccountAdmin** | Pointer to **bool** | If set to {true}, user has full access to all properties. Defaults to false. | [optional] 
**Roles** | Pointer to **[]string** | Roles that user has access to. If the user is invited as account admin, this will be ignored. | [optional] 
**Role** | Pointer to **string** | The role to be assigned to the user. If you specfiy this and &#39;Roles&#39;, the combination of both will be used.  If the user is invited as account admin, this will be ignored. | [optional] 

## Methods

### NewCreateInvitationModel

`func NewCreateInvitationModel(email string, ) *CreateInvitationModel`

NewCreateInvitationModel instantiates a new CreateInvitationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvitationModelWithDefaults

`func NewCreateInvitationModelWithDefaults() *CreateInvitationModel`

NewCreateInvitationModelWithDefaults instantiates a new CreateInvitationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateInvitationModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateInvitationModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateInvitationModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProperties

`func (o *CreateInvitationModel) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateInvitationModel) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateInvitationModel) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateInvitationModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetIsAccountAdmin

`func (o *CreateInvitationModel) GetIsAccountAdmin() bool`

GetIsAccountAdmin returns the IsAccountAdmin field if non-nil, zero value otherwise.

### GetIsAccountAdminOk

`func (o *CreateInvitationModel) GetIsAccountAdminOk() (*bool, bool)`

GetIsAccountAdminOk returns a tuple with the IsAccountAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountAdmin

`func (o *CreateInvitationModel) SetIsAccountAdmin(v bool)`

SetIsAccountAdmin sets IsAccountAdmin field to given value.

### HasIsAccountAdmin

`func (o *CreateInvitationModel) HasIsAccountAdmin() bool`

HasIsAccountAdmin returns a boolean if a field has been set.

### GetRoles

`func (o *CreateInvitationModel) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateInvitationModel) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateInvitationModel) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateInvitationModel) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRole

`func (o *CreateInvitationModel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateInvitationModel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateInvitationModel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateInvitationModel) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


