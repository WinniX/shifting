# UiIntegrationItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Code** | **string** |  | 
**Label** | **string** |  | 
**LabelTranslations** | Pointer to **map[string]string** |  | [optional] 
**IconSource** | Pointer to **string** |  | [optional] 
**SourceUrl** | **string** |  | 
**PropertyIds** | Pointer to **[]string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**SourceType** | **string** | Public source type means, that the browser can see the URL directly. This is true for instance, if you have Hybrid OAuth flow active.                Private source type means, that the server will request from that source URL a public URL, which is pre-authenticated and forward that to the client. | 
**Target** | **string** |  | 

## Methods

### NewUiIntegrationItemModel

`func NewUiIntegrationItemModel(id string, code string, label string, sourceUrl string, sourceType string, target string, ) *UiIntegrationItemModel`

NewUiIntegrationItemModel instantiates a new UiIntegrationItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiIntegrationItemModelWithDefaults

`func NewUiIntegrationItemModelWithDefaults() *UiIntegrationItemModel`

NewUiIntegrationItemModelWithDefaults instantiates a new UiIntegrationItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UiIntegrationItemModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiIntegrationItemModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiIntegrationItemModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *UiIntegrationItemModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UiIntegrationItemModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UiIntegrationItemModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetLabel

`func (o *UiIntegrationItemModel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UiIntegrationItemModel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UiIntegrationItemModel) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelTranslations

`func (o *UiIntegrationItemModel) GetLabelTranslations() map[string]string`

GetLabelTranslations returns the LabelTranslations field if non-nil, zero value otherwise.

### GetLabelTranslationsOk

`func (o *UiIntegrationItemModel) GetLabelTranslationsOk() (*map[string]string, bool)`

GetLabelTranslationsOk returns a tuple with the LabelTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelTranslations

`func (o *UiIntegrationItemModel) SetLabelTranslations(v map[string]string)`

SetLabelTranslations sets LabelTranslations field to given value.

### HasLabelTranslations

`func (o *UiIntegrationItemModel) HasLabelTranslations() bool`

HasLabelTranslations returns a boolean if a field has been set.

### GetIconSource

`func (o *UiIntegrationItemModel) GetIconSource() string`

GetIconSource returns the IconSource field if non-nil, zero value otherwise.

### GetIconSourceOk

`func (o *UiIntegrationItemModel) GetIconSourceOk() (*string, bool)`

GetIconSourceOk returns a tuple with the IconSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSource

`func (o *UiIntegrationItemModel) SetIconSource(v string)`

SetIconSource sets IconSource field to given value.

### HasIconSource

`func (o *UiIntegrationItemModel) HasIconSource() bool`

HasIconSource returns a boolean if a field has been set.

### GetSourceUrl

`func (o *UiIntegrationItemModel) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *UiIntegrationItemModel) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *UiIntegrationItemModel) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetPropertyIds

`func (o *UiIntegrationItemModel) GetPropertyIds() []string`

GetPropertyIds returns the PropertyIds field if non-nil, zero value otherwise.

### GetPropertyIdsOk

`func (o *UiIntegrationItemModel) GetPropertyIdsOk() (*[]string, bool)`

GetPropertyIdsOk returns a tuple with the PropertyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIds

`func (o *UiIntegrationItemModel) SetPropertyIds(v []string)`

SetPropertyIds sets PropertyIds field to given value.

### HasPropertyIds

`func (o *UiIntegrationItemModel) HasPropertyIds() bool`

HasPropertyIds returns a boolean if a field has been set.

### GetRoles

`func (o *UiIntegrationItemModel) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UiIntegrationItemModel) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UiIntegrationItemModel) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UiIntegrationItemModel) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSourceType

`func (o *UiIntegrationItemModel) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UiIntegrationItemModel) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UiIntegrationItemModel) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetTarget

`func (o *UiIntegrationItemModel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UiIntegrationItemModel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UiIntegrationItemModel) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


