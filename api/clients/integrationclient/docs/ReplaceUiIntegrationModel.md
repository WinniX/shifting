# ReplaceUiIntegrationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label to be displayed as the name of your integration.  For example, this is the text that is displayed on your user&#39;s apps list. | 
**LabelTranslations** | Pointer to **map[string]string** | Different translations for the integration label.  You can provide as many translations as you want.  You should only provide the iso code of the language for example &#39;en&#39; for english and &#39;de&#39; for German. No need for the country.  If there is not a translation for a specific language, we will display the value of the Label as default. | [optional] 
**IconSource** | Pointer to **string** |  | [optional] 
**SourceUrl** | **string** |  | 
**SourceType** | **string** | Public source type means, that the browser can see the URL directly. This is true for instance, if you have Hybrid OAuth flow active.                Private source type means, that the server will request from that source URL a public URL, which is pre-authenticated and forward that to the client. | 
**PropertyIds** | Pointer to **[]string** | List of Ids for the properties for which the integration is configured  If the list is empty then the integration is configured for all properties  Remark: Only applicable to non account level integrations | [optional] 
**Roles** | Pointer to **[]string** | List of Roles for which the integration is configured to be used  If the list is empty then the integration is configured for all roles | [optional] 

## Methods

### NewReplaceUiIntegrationModel

`func NewReplaceUiIntegrationModel(label string, sourceUrl string, sourceType string, ) *ReplaceUiIntegrationModel`

NewReplaceUiIntegrationModel instantiates a new ReplaceUiIntegrationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceUiIntegrationModelWithDefaults

`func NewReplaceUiIntegrationModelWithDefaults() *ReplaceUiIntegrationModel`

NewReplaceUiIntegrationModelWithDefaults instantiates a new ReplaceUiIntegrationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ReplaceUiIntegrationModel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ReplaceUiIntegrationModel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ReplaceUiIntegrationModel) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelTranslations

`func (o *ReplaceUiIntegrationModel) GetLabelTranslations() map[string]string`

GetLabelTranslations returns the LabelTranslations field if non-nil, zero value otherwise.

### GetLabelTranslationsOk

`func (o *ReplaceUiIntegrationModel) GetLabelTranslationsOk() (*map[string]string, bool)`

GetLabelTranslationsOk returns a tuple with the LabelTranslations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelTranslations

`func (o *ReplaceUiIntegrationModel) SetLabelTranslations(v map[string]string)`

SetLabelTranslations sets LabelTranslations field to given value.

### HasLabelTranslations

`func (o *ReplaceUiIntegrationModel) HasLabelTranslations() bool`

HasLabelTranslations returns a boolean if a field has been set.

### GetIconSource

`func (o *ReplaceUiIntegrationModel) GetIconSource() string`

GetIconSource returns the IconSource field if non-nil, zero value otherwise.

### GetIconSourceOk

`func (o *ReplaceUiIntegrationModel) GetIconSourceOk() (*string, bool)`

GetIconSourceOk returns a tuple with the IconSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSource

`func (o *ReplaceUiIntegrationModel) SetIconSource(v string)`

SetIconSource sets IconSource field to given value.

### HasIconSource

`func (o *ReplaceUiIntegrationModel) HasIconSource() bool`

HasIconSource returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ReplaceUiIntegrationModel) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ReplaceUiIntegrationModel) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ReplaceUiIntegrationModel) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetSourceType

`func (o *ReplaceUiIntegrationModel) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ReplaceUiIntegrationModel) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ReplaceUiIntegrationModel) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetPropertyIds

`func (o *ReplaceUiIntegrationModel) GetPropertyIds() []string`

GetPropertyIds returns the PropertyIds field if non-nil, zero value otherwise.

### GetPropertyIdsOk

`func (o *ReplaceUiIntegrationModel) GetPropertyIdsOk() (*[]string, bool)`

GetPropertyIdsOk returns a tuple with the PropertyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIds

`func (o *ReplaceUiIntegrationModel) SetPropertyIds(v []string)`

SetPropertyIds sets PropertyIds field to given value.

### HasPropertyIds

`func (o *ReplaceUiIntegrationModel) HasPropertyIds() bool`

HasPropertyIds returns a boolean if a field has been set.

### GetRoles

`func (o *ReplaceUiIntegrationModel) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ReplaceUiIntegrationModel) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ReplaceUiIntegrationModel) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ReplaceUiIntegrationModel) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


