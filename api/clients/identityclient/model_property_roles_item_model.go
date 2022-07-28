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

// PropertyRolesItemModel struct for PropertyRolesItemModel
type PropertyRolesItemModel struct {
	Id string `json:"id"`
	Roles []string `json:"roles"`
}

// NewPropertyRolesItemModel instantiates a new PropertyRolesItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyRolesItemModel(id string, roles []string) *PropertyRolesItemModel {
	this := PropertyRolesItemModel{}
	this.Id = id
	this.Roles = roles
	return &this
}

// NewPropertyRolesItemModelWithDefaults instantiates a new PropertyRolesItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyRolesItemModelWithDefaults() *PropertyRolesItemModel {
	this := PropertyRolesItemModel{}
	return &this
}

// GetId returns the Id field value
func (o *PropertyRolesItemModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PropertyRolesItemModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PropertyRolesItemModel) SetId(v string) {
	o.Id = v
}

// GetRoles returns the Roles field value
func (o *PropertyRolesItemModel) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *PropertyRolesItemModel) GetRolesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *PropertyRolesItemModel) SetRoles(v []string) {
	o.Roles = v
}

func (o PropertyRolesItemModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyRolesItemModel struct {
	value *PropertyRolesItemModel
	isSet bool
}

func (v NullablePropertyRolesItemModel) Get() *PropertyRolesItemModel {
	return v.value
}

func (v *NullablePropertyRolesItemModel) Set(val *PropertyRolesItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyRolesItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyRolesItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyRolesItemModel(val *PropertyRolesItemModel) *NullablePropertyRolesItemModel {
	return &NullablePropertyRolesItemModel{value: val, isSet: true}
}

func (v NullablePropertyRolesItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyRolesItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


