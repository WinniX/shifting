# UserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** | Unique user identifier. | 
**FirstName** | **string** | First name | 
**LastName** | **string** | Last name | 
**Email** | **string** | Email address, used as the login | 
**Enabled** | Pointer to **bool** | If set to {false}, the user is disabled for this account and cannot log in | [optional] 
**IsAccountAdmin** | **bool** | If set to {true}, user has full access to all properties. | 
**Properties** | Pointer to **[]string** | List of properties to which user has access. | [optional] 
**PropertyRoles** | Pointer to [**[]PropertyRolesItemModel**](PropertyRolesItemModel.md) | List of properties to which user has access. | [optional] 

## Methods

### NewUserModel

`func NewUserModel(subjectId string, firstName string, lastName string, email string, isAccountAdmin bool, ) *UserModel`

NewUserModel instantiates a new UserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelWithDefaults

`func NewUserModelWithDefaults() *UserModel`

NewUserModelWithDefaults instantiates a new UserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *UserModel) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *UserModel) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *UserModel) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetFirstName

`func (o *UserModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *UserModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsAccountAdmin

`func (o *UserModel) GetIsAccountAdmin() bool`

GetIsAccountAdmin returns the IsAccountAdmin field if non-nil, zero value otherwise.

### GetIsAccountAdminOk

`func (o *UserModel) GetIsAccountAdminOk() (*bool, bool)`

GetIsAccountAdminOk returns a tuple with the IsAccountAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountAdmin

`func (o *UserModel) SetIsAccountAdmin(v bool)`

SetIsAccountAdmin sets IsAccountAdmin field to given value.


### GetProperties

`func (o *UserModel) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserModel) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserModel) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UserModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertyRoles

`func (o *UserModel) GetPropertyRoles() []PropertyRolesItemModel`

GetPropertyRoles returns the PropertyRoles field if non-nil, zero value otherwise.

### GetPropertyRolesOk

`func (o *UserModel) GetPropertyRolesOk() (*[]PropertyRolesItemModel, bool)`

GetPropertyRolesOk returns a tuple with the PropertyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRoles

`func (o *UserModel) SetPropertyRoles(v []PropertyRolesItemModel)`

SetPropertyRoles sets PropertyRoles field to given value.

### HasPropertyRoles

`func (o *UserModel) HasPropertyRoles() bool`

HasPropertyRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


