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

// RoleListModel struct for RoleListModel
type RoleListModel struct {
	// A list of all existing roles
	Roles []string `json:"roles"`
}

// NewRoleListModel instantiates a new RoleListModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleListModel(roles []string) *RoleListModel {
	this := RoleListModel{}
	this.Roles = roles
	return &this
}

// NewRoleListModelWithDefaults instantiates a new RoleListModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleListModelWithDefaults() *RoleListModel {
	this := RoleListModel{}
	return &this
}

// GetRoles returns the Roles field value
func (o *RoleListModel) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RoleListModel) GetRolesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *RoleListModel) SetRoles(v []string) {
	o.Roles = v
}

func (o RoleListModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableRoleListModel struct {
	value *RoleListModel
	isSet bool
}

func (v NullableRoleListModel) Get() *RoleListModel {
	return v.value
}

func (v *NullableRoleListModel) Set(val *RoleListModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleListModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleListModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleListModel(val *RoleListModel) *NullableRoleListModel {
	return &NullableRoleListModel{value: val, isSet: true}
}

func (v NullableRoleListModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleListModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


