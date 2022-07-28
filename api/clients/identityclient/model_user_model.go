/*
 * apaleo Identity API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package identityclient

import (
	"encoding/json"
)

// UserModel struct for UserModel
type UserModel struct {
	// Unique user identifier.
	SubjectId string `json:"subjectId"`
	// First name
	FirstName string `json:"firstName"`
	// Last name
	LastName string `json:"lastName"`
	// Email address, used as the login
	Email string `json:"email"`
	// If set to {false}, the user is disabled for this account and cannot log in
	Enabled *bool `json:"enabled,omitempty"`
	// If set to {true}, user has full access to all properties.
	IsAccountAdmin bool `json:"isAccountAdmin"`
	// List of properties to which user has access.
	Properties *[]string `json:"properties,omitempty"`
	// List of properties to which user has access.
	PropertyRoles *[]PropertyRolesItemModel `json:"propertyRoles,omitempty"`
}

// NewUserModel instantiates a new UserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModel(subjectId string, firstName string, lastName string, email string, isAccountAdmin bool) *UserModel {
	this := UserModel{}
	this.SubjectId = subjectId
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.IsAccountAdmin = isAccountAdmin
	return &this
}

// NewUserModelWithDefaults instantiates a new UserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelWithDefaults() *UserModel {
	this := UserModel{}
	return &this
}

// GetSubjectId returns the SubjectId field value
func (o *UserModel) GetSubjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetSubjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *UserModel) SetSubjectId(v string) {
	o.SubjectId = v
}

// GetFirstName returns the FirstName field value
func (o *UserModel) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserModel) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserModel) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserModel) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UserModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserModel) SetEmail(v string) {
	o.Email = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserModel) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserModel) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserModel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsAccountAdmin returns the IsAccountAdmin field value
func (o *UserModel) GetIsAccountAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAccountAdmin
}

// GetIsAccountAdminOk returns a tuple with the IsAccountAdmin field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetIsAccountAdminOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAccountAdmin, true
}

// SetIsAccountAdmin sets field value
func (o *UserModel) SetIsAccountAdmin(v bool) {
	o.IsAccountAdmin = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserModel) GetProperties() []string {
	if o == nil || o.Properties == nil {
		var ret []string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetPropertiesOk() (*[]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserModel) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []string and assigns it to the Properties field.
func (o *UserModel) SetProperties(v []string) {
	o.Properties = &v
}

// GetPropertyRoles returns the PropertyRoles field value if set, zero value otherwise.
func (o *UserModel) GetPropertyRoles() []PropertyRolesItemModel {
	if o == nil || o.PropertyRoles == nil {
		var ret []PropertyRolesItemModel
		return ret
	}
	return *o.PropertyRoles
}

// GetPropertyRolesOk returns a tuple with the PropertyRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetPropertyRolesOk() (*[]PropertyRolesItemModel, bool) {
	if o == nil || o.PropertyRoles == nil {
		return nil, false
	}
	return o.PropertyRoles, true
}

// HasPropertyRoles returns a boolean if a field has been set.
func (o *UserModel) HasPropertyRoles() bool {
	if o != nil && o.PropertyRoles != nil {
		return true
	}

	return false
}

// SetPropertyRoles gets a reference to the given []PropertyRolesItemModel and assigns it to the PropertyRoles field.
func (o *UserModel) SetPropertyRoles(v []PropertyRolesItemModel) {
	o.PropertyRoles = &v
}

func (o UserModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subjectId"] = o.SubjectId
	}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["isAccountAdmin"] = o.IsAccountAdmin
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.PropertyRoles != nil {
		toSerialize["propertyRoles"] = o.PropertyRoles
	}
	return json.Marshal(toSerialize)
}

type NullableUserModel struct {
	value *UserModel
	isSet bool
}

func (v NullableUserModel) Get() *UserModel {
	return v.value
}

func (v *NullableUserModel) Set(val *UserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModel(val *UserModel) *NullableUserModel {
	return &NullableUserModel{value: val, isSet: true}
}

func (v NullableUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


