# UserItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** | Unique user identifier. | 
**FirstName** | **string** | First name | 
**LastName** | **string** | Last name | 
**Email** | **string** | Email address, used as the login | 
**Properties** | Pointer to [**[]PropertyRolesItemModel**](PropertyRolesItemModel.md) | Properties that user has access to | [optional] 
**Enabled** | **bool** | If set to {false}, the user is disabled for this account and cannot log in | 
**IsAccountAdmin** | **bool** | If set to {true}, user has full access to all properties. | 

## Methods

### NewUserItemModel

`func NewUserItemModel(subjectId string, firstName string, lastName string, email string, enabled bool, isAccountAdmin bool, ) *UserItemModel`

NewUserItemModel instantiates a new UserItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserItemModelWithDefaults

`func NewUserItemModelWithDefaults() *UserItemModel`

NewUserItemModelWithDefaults instantiates a new UserItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *UserItemModel) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *UserItemModel) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *UserItemModel) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetFirstName

`func (o *UserItemModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserItemModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserItemModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserItemModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserItemModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserItemModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UserItemModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserItemModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserItemModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProperties

`func (o *UserItemModel) GetProperties() []PropertyRolesItemModel`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserItemModel) GetPropertiesOk() (*[]PropertyRolesItemModel, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserItemModel) SetProperties(v []PropertyRolesItemModel)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UserItemModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetEnabled

`func (o *UserItemModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserItemModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserItemModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIsAccountAdmin

`func (o *UserItemModel) GetIsAccountAdmin() bool`

GetIsAccountAdmin returns the IsAccountAdmin field if non-nil, zero value otherwise.

### GetIsAccountAdminOk

`func (o *UserItemModel) GetIsAccountAdminOk() (*bool, bool)`

GetIsAccountAdminOk returns a tuple with the IsAccountAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountAdmin

`func (o *UserItemModel) SetIsAccountAdmin(v bool)`

SetIsAccountAdmin sets IsAccountAdmin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


