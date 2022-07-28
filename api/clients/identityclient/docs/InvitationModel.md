# InvitationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address, used as the login | 
**Properties** | Pointer to **[]string** | Properties that user has access to | [optional] 
**IsAccountAdmin** | **bool** | If set to {true}, user has full access to all properties. | 
**Role** | Pointer to **string** | Role this user is invited to. If more than one, returns the first. | [optional] 
**Roles** | Pointer to **[]string** | Roles that user has access to | [optional] 
**InvitedBy** | **string** | Email of the user making the invitation | 
**Created** | **time.Time** | Date the invitation was made | 

## Methods

### NewInvitationModel

`func NewInvitationModel(email string, isAccountAdmin bool, invitedBy string, created time.Time, ) *InvitationModel`

NewInvitationModel instantiates a new InvitationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationModelWithDefaults

`func NewInvitationModelWithDefaults() *InvitationModel`

NewInvitationModelWithDefaults instantiates a new InvitationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InvitationModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProperties

`func (o *InvitationModel) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InvitationModel) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InvitationModel) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InvitationModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetIsAccountAdmin

`func (o *InvitationModel) GetIsAccountAdmin() bool`

GetIsAccountAdmin returns the IsAccountAdmin field if non-nil, zero value otherwise.

### GetIsAccountAdminOk

`func (o *InvitationModel) GetIsAccountAdminOk() (*bool, bool)`

GetIsAccountAdminOk returns a tuple with the IsAccountAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountAdmin

`func (o *InvitationModel) SetIsAccountAdmin(v bool)`

SetIsAccountAdmin sets IsAccountAdmin field to given value.


### GetRole

`func (o *InvitationModel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InvitationModel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InvitationModel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InvitationModel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *InvitationModel) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InvitationModel) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InvitationModel) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InvitationModel) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetInvitedBy

`func (o *InvitationModel) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *InvitationModel) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *InvitationModel) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.


### GetCreated

`func (o *InvitationModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InvitationModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InvitationModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


