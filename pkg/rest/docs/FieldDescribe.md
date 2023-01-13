# FieldDescribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregatable** | Pointer to **bool** |  | [optional] 
**AiPredictionField** | Pointer to **bool** |  | [optional] 
**AutoNumber** | Pointer to **bool** |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**Calculated** | Pointer to **bool** |  | [optional] 
**CalculatedFormula** | Pointer to **NullableString** |  | [optional] 
**CascadeDelete** | Pointer to **bool** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**CompoundFieldName** | Pointer to **NullableString** |  | [optional] 
**ControllerName** | Pointer to **NullableString** |  | [optional] 
**Createable** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**DefaultValueFormula** | Pointer to **NullableString** |  | [optional] 
**DefaultedOnCreate** | Pointer to **bool** |  | [optional] 
**DependentPicklist** | Pointer to **bool** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**Digits** | Pointer to **int32** |  | [optional] 
**DisplayLocationInDecimal** | Pointer to **bool** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **bool** |  | [optional] 
**ExtraTypeInfo** | Pointer to **NullableString** |  | [optional] 
**Filterable** | Pointer to **bool** |  | [optional] 
**FilteredLookupInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**FormulaTreatNullNumberAsZero** | Pointer to **bool** |  | [optional] 
**Groupable** | Pointer to **bool** |  | [optional] 
**HighScaleNumber** | Pointer to **bool** |  | [optional] 
**HtmlFormatted** | Pointer to **bool** |  | [optional] 
**IdLookup** | Pointer to **bool** |  | [optional] 
**InlineHelpText** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**Mask** | Pointer to **NullableString** |  | [optional] 
**MaskType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameField** | Pointer to **bool** |  | [optional] 
**NamePointing** | Pointer to **bool** |  | [optional] 
**Nillable** | Pointer to **bool** |  | [optional] 
**Permissionable** | Pointer to **bool** |  | [optional] 
**PicklistValues** | Pointer to [**[]FieldDescribePicklist**](FieldDescribePicklist.md) |  | [optional] 
**PolymorphicForeignKey** | Pointer to **bool** |  | [optional] 
**Precision** | Pointer to **int32** |  | [optional] 
**QueryByDistance** | Pointer to **bool** |  | [optional] 
**ReferenceTargetField** | Pointer to **NullableString** |  | [optional] 
**ReferenceTo** | Pointer to **[]string** |  | [optional] 
**RelationshipName** | Pointer to **NullableString** |  | [optional] 
**RelationshipOrder** | Pointer to **NullableInt32** |  | [optional] 
**RestrictedDelete** | Pointer to **bool** |  | [optional] 
**RestrictedPicklist** | Pointer to **bool** |  | [optional] 
**Scale** | Pointer to **int32** |  | [optional] 
**SearchPrefilterable** | Pointer to **bool** |  | [optional] 
**SoapType** | Pointer to **string** |  | [optional] 
**Sortable** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**FieldType**](FieldType.md) |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**Updateable** | Pointer to **bool** |  | [optional] 
**WriteRequiresMasterRead** | Pointer to **bool** |  | [optional] 

## Methods

### NewFieldDescribe

`func NewFieldDescribe() *FieldDescribe`

NewFieldDescribe instantiates a new FieldDescribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDescribeWithDefaults

`func NewFieldDescribeWithDefaults() *FieldDescribe`

NewFieldDescribeWithDefaults instantiates a new FieldDescribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatable

`func (o *FieldDescribe) GetAggregatable() bool`

GetAggregatable returns the Aggregatable field if non-nil, zero value otherwise.

### GetAggregatableOk

`func (o *FieldDescribe) GetAggregatableOk() (*bool, bool)`

GetAggregatableOk returns a tuple with the Aggregatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatable

`func (o *FieldDescribe) SetAggregatable(v bool)`

SetAggregatable sets Aggregatable field to given value.

### HasAggregatable

`func (o *FieldDescribe) HasAggregatable() bool`

HasAggregatable returns a boolean if a field has been set.

### GetAiPredictionField

`func (o *FieldDescribe) GetAiPredictionField() bool`

GetAiPredictionField returns the AiPredictionField field if non-nil, zero value otherwise.

### GetAiPredictionFieldOk

`func (o *FieldDescribe) GetAiPredictionFieldOk() (*bool, bool)`

GetAiPredictionFieldOk returns a tuple with the AiPredictionField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPredictionField

`func (o *FieldDescribe) SetAiPredictionField(v bool)`

SetAiPredictionField sets AiPredictionField field to given value.

### HasAiPredictionField

`func (o *FieldDescribe) HasAiPredictionField() bool`

HasAiPredictionField returns a boolean if a field has been set.

### GetAutoNumber

`func (o *FieldDescribe) GetAutoNumber() bool`

GetAutoNumber returns the AutoNumber field if non-nil, zero value otherwise.

### GetAutoNumberOk

`func (o *FieldDescribe) GetAutoNumberOk() (*bool, bool)`

GetAutoNumberOk returns a tuple with the AutoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNumber

`func (o *FieldDescribe) SetAutoNumber(v bool)`

SetAutoNumber sets AutoNumber field to given value.

### HasAutoNumber

`func (o *FieldDescribe) HasAutoNumber() bool`

HasAutoNumber returns a boolean if a field has been set.

### GetByteLength

`func (o *FieldDescribe) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *FieldDescribe) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *FieldDescribe) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *FieldDescribe) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetCalculated

`func (o *FieldDescribe) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *FieldDescribe) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *FieldDescribe) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *FieldDescribe) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetCalculatedFormula

`func (o *FieldDescribe) GetCalculatedFormula() string`

GetCalculatedFormula returns the CalculatedFormula field if non-nil, zero value otherwise.

### GetCalculatedFormulaOk

`func (o *FieldDescribe) GetCalculatedFormulaOk() (*string, bool)`

GetCalculatedFormulaOk returns a tuple with the CalculatedFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedFormula

`func (o *FieldDescribe) SetCalculatedFormula(v string)`

SetCalculatedFormula sets CalculatedFormula field to given value.

### HasCalculatedFormula

`func (o *FieldDescribe) HasCalculatedFormula() bool`

HasCalculatedFormula returns a boolean if a field has been set.

### SetCalculatedFormulaNil

`func (o *FieldDescribe) SetCalculatedFormulaNil(b bool)`

 SetCalculatedFormulaNil sets the value for CalculatedFormula to be an explicit nil

### UnsetCalculatedFormula
`func (o *FieldDescribe) UnsetCalculatedFormula()`

UnsetCalculatedFormula ensures that no value is present for CalculatedFormula, not even an explicit nil
### GetCascadeDelete

`func (o *FieldDescribe) GetCascadeDelete() bool`

GetCascadeDelete returns the CascadeDelete field if non-nil, zero value otherwise.

### GetCascadeDeleteOk

`func (o *FieldDescribe) GetCascadeDeleteOk() (*bool, bool)`

GetCascadeDeleteOk returns a tuple with the CascadeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeDelete

`func (o *FieldDescribe) SetCascadeDelete(v bool)`

SetCascadeDelete sets CascadeDelete field to given value.

### HasCascadeDelete

`func (o *FieldDescribe) HasCascadeDelete() bool`

HasCascadeDelete returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *FieldDescribe) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *FieldDescribe) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *FieldDescribe) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *FieldDescribe) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetCompoundFieldName

`func (o *FieldDescribe) GetCompoundFieldName() string`

GetCompoundFieldName returns the CompoundFieldName field if non-nil, zero value otherwise.

### GetCompoundFieldNameOk

`func (o *FieldDescribe) GetCompoundFieldNameOk() (*string, bool)`

GetCompoundFieldNameOk returns a tuple with the CompoundFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundFieldName

`func (o *FieldDescribe) SetCompoundFieldName(v string)`

SetCompoundFieldName sets CompoundFieldName field to given value.

### HasCompoundFieldName

`func (o *FieldDescribe) HasCompoundFieldName() bool`

HasCompoundFieldName returns a boolean if a field has been set.

### SetCompoundFieldNameNil

`func (o *FieldDescribe) SetCompoundFieldNameNil(b bool)`

 SetCompoundFieldNameNil sets the value for CompoundFieldName to be an explicit nil

### UnsetCompoundFieldName
`func (o *FieldDescribe) UnsetCompoundFieldName()`

UnsetCompoundFieldName ensures that no value is present for CompoundFieldName, not even an explicit nil
### GetControllerName

`func (o *FieldDescribe) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *FieldDescribe) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *FieldDescribe) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *FieldDescribe) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### SetControllerNameNil

`func (o *FieldDescribe) SetControllerNameNil(b bool)`

 SetControllerNameNil sets the value for ControllerName to be an explicit nil

### UnsetControllerName
`func (o *FieldDescribe) UnsetControllerName()`

UnsetControllerName ensures that no value is present for ControllerName, not even an explicit nil
### GetCreateable

`func (o *FieldDescribe) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *FieldDescribe) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *FieldDescribe) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *FieldDescribe) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *FieldDescribe) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *FieldDescribe) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *FieldDescribe) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *FieldDescribe) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDefaultValue

`func (o *FieldDescribe) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FieldDescribe) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FieldDescribe) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FieldDescribe) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *FieldDescribe) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *FieldDescribe) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDefaultValueFormula

`func (o *FieldDescribe) GetDefaultValueFormula() string`

GetDefaultValueFormula returns the DefaultValueFormula field if non-nil, zero value otherwise.

### GetDefaultValueFormulaOk

`func (o *FieldDescribe) GetDefaultValueFormulaOk() (*string, bool)`

GetDefaultValueFormulaOk returns a tuple with the DefaultValueFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueFormula

`func (o *FieldDescribe) SetDefaultValueFormula(v string)`

SetDefaultValueFormula sets DefaultValueFormula field to given value.

### HasDefaultValueFormula

`func (o *FieldDescribe) HasDefaultValueFormula() bool`

HasDefaultValueFormula returns a boolean if a field has been set.

### SetDefaultValueFormulaNil

`func (o *FieldDescribe) SetDefaultValueFormulaNil(b bool)`

 SetDefaultValueFormulaNil sets the value for DefaultValueFormula to be an explicit nil

### UnsetDefaultValueFormula
`func (o *FieldDescribe) UnsetDefaultValueFormula()`

UnsetDefaultValueFormula ensures that no value is present for DefaultValueFormula, not even an explicit nil
### GetDefaultedOnCreate

`func (o *FieldDescribe) GetDefaultedOnCreate() bool`

GetDefaultedOnCreate returns the DefaultedOnCreate field if non-nil, zero value otherwise.

### GetDefaultedOnCreateOk

`func (o *FieldDescribe) GetDefaultedOnCreateOk() (*bool, bool)`

GetDefaultedOnCreateOk returns a tuple with the DefaultedOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultedOnCreate

`func (o *FieldDescribe) SetDefaultedOnCreate(v bool)`

SetDefaultedOnCreate sets DefaultedOnCreate field to given value.

### HasDefaultedOnCreate

`func (o *FieldDescribe) HasDefaultedOnCreate() bool`

HasDefaultedOnCreate returns a boolean if a field has been set.

### GetDependentPicklist

`func (o *FieldDescribe) GetDependentPicklist() bool`

GetDependentPicklist returns the DependentPicklist field if non-nil, zero value otherwise.

### GetDependentPicklistOk

`func (o *FieldDescribe) GetDependentPicklistOk() (*bool, bool)`

GetDependentPicklistOk returns a tuple with the DependentPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentPicklist

`func (o *FieldDescribe) SetDependentPicklist(v bool)`

SetDependentPicklist sets DependentPicklist field to given value.

### HasDependentPicklist

`func (o *FieldDescribe) HasDependentPicklist() bool`

HasDependentPicklist returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *FieldDescribe) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *FieldDescribe) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *FieldDescribe) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *FieldDescribe) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetDigits

`func (o *FieldDescribe) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *FieldDescribe) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *FieldDescribe) SetDigits(v int32)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *FieldDescribe) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetDisplayLocationInDecimal

`func (o *FieldDescribe) GetDisplayLocationInDecimal() bool`

GetDisplayLocationInDecimal returns the DisplayLocationInDecimal field if non-nil, zero value otherwise.

### GetDisplayLocationInDecimalOk

`func (o *FieldDescribe) GetDisplayLocationInDecimalOk() (*bool, bool)`

GetDisplayLocationInDecimalOk returns a tuple with the DisplayLocationInDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocationInDecimal

`func (o *FieldDescribe) SetDisplayLocationInDecimal(v bool)`

SetDisplayLocationInDecimal sets DisplayLocationInDecimal field to given value.

### HasDisplayLocationInDecimal

`func (o *FieldDescribe) HasDisplayLocationInDecimal() bool`

HasDisplayLocationInDecimal returns a boolean if a field has been set.

### GetEncrypted

`func (o *FieldDescribe) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *FieldDescribe) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *FieldDescribe) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *FieldDescribe) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetExternalId

`func (o *FieldDescribe) GetExternalId() bool`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FieldDescribe) GetExternalIdOk() (*bool, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FieldDescribe) SetExternalId(v bool)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FieldDescribe) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExtraTypeInfo

`func (o *FieldDescribe) GetExtraTypeInfo() string`

GetExtraTypeInfo returns the ExtraTypeInfo field if non-nil, zero value otherwise.

### GetExtraTypeInfoOk

`func (o *FieldDescribe) GetExtraTypeInfoOk() (*string, bool)`

GetExtraTypeInfoOk returns a tuple with the ExtraTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraTypeInfo

`func (o *FieldDescribe) SetExtraTypeInfo(v string)`

SetExtraTypeInfo sets ExtraTypeInfo field to given value.

### HasExtraTypeInfo

`func (o *FieldDescribe) HasExtraTypeInfo() bool`

HasExtraTypeInfo returns a boolean if a field has been set.

### SetExtraTypeInfoNil

`func (o *FieldDescribe) SetExtraTypeInfoNil(b bool)`

 SetExtraTypeInfoNil sets the value for ExtraTypeInfo to be an explicit nil

### UnsetExtraTypeInfo
`func (o *FieldDescribe) UnsetExtraTypeInfo()`

UnsetExtraTypeInfo ensures that no value is present for ExtraTypeInfo, not even an explicit nil
### GetFilterable

`func (o *FieldDescribe) GetFilterable() bool`

GetFilterable returns the Filterable field if non-nil, zero value otherwise.

### GetFilterableOk

`func (o *FieldDescribe) GetFilterableOk() (*bool, bool)`

GetFilterableOk returns a tuple with the Filterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterable

`func (o *FieldDescribe) SetFilterable(v bool)`

SetFilterable sets Filterable field to given value.

### HasFilterable

`func (o *FieldDescribe) HasFilterable() bool`

HasFilterable returns a boolean if a field has been set.

### GetFilteredLookupInfo

`func (o *FieldDescribe) GetFilteredLookupInfo() map[string]interface{}`

GetFilteredLookupInfo returns the FilteredLookupInfo field if non-nil, zero value otherwise.

### GetFilteredLookupInfoOk

`func (o *FieldDescribe) GetFilteredLookupInfoOk() (*map[string]interface{}, bool)`

GetFilteredLookupInfoOk returns a tuple with the FilteredLookupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredLookupInfo

`func (o *FieldDescribe) SetFilteredLookupInfo(v map[string]interface{})`

SetFilteredLookupInfo sets FilteredLookupInfo field to given value.

### HasFilteredLookupInfo

`func (o *FieldDescribe) HasFilteredLookupInfo() bool`

HasFilteredLookupInfo returns a boolean if a field has been set.

### SetFilteredLookupInfoNil

`func (o *FieldDescribe) SetFilteredLookupInfoNil(b bool)`

 SetFilteredLookupInfoNil sets the value for FilteredLookupInfo to be an explicit nil

### UnsetFilteredLookupInfo
`func (o *FieldDescribe) UnsetFilteredLookupInfo()`

UnsetFilteredLookupInfo ensures that no value is present for FilteredLookupInfo, not even an explicit nil
### GetFormulaTreatNullNumberAsZero

`func (o *FieldDescribe) GetFormulaTreatNullNumberAsZero() bool`

GetFormulaTreatNullNumberAsZero returns the FormulaTreatNullNumberAsZero field if non-nil, zero value otherwise.

### GetFormulaTreatNullNumberAsZeroOk

`func (o *FieldDescribe) GetFormulaTreatNullNumberAsZeroOk() (*bool, bool)`

GetFormulaTreatNullNumberAsZeroOk returns a tuple with the FormulaTreatNullNumberAsZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaTreatNullNumberAsZero

`func (o *FieldDescribe) SetFormulaTreatNullNumberAsZero(v bool)`

SetFormulaTreatNullNumberAsZero sets FormulaTreatNullNumberAsZero field to given value.

### HasFormulaTreatNullNumberAsZero

`func (o *FieldDescribe) HasFormulaTreatNullNumberAsZero() bool`

HasFormulaTreatNullNumberAsZero returns a boolean if a field has been set.

### GetGroupable

`func (o *FieldDescribe) GetGroupable() bool`

GetGroupable returns the Groupable field if non-nil, zero value otherwise.

### GetGroupableOk

`func (o *FieldDescribe) GetGroupableOk() (*bool, bool)`

GetGroupableOk returns a tuple with the Groupable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupable

`func (o *FieldDescribe) SetGroupable(v bool)`

SetGroupable sets Groupable field to given value.

### HasGroupable

`func (o *FieldDescribe) HasGroupable() bool`

HasGroupable returns a boolean if a field has been set.

### GetHighScaleNumber

`func (o *FieldDescribe) GetHighScaleNumber() bool`

GetHighScaleNumber returns the HighScaleNumber field if non-nil, zero value otherwise.

### GetHighScaleNumberOk

`func (o *FieldDescribe) GetHighScaleNumberOk() (*bool, bool)`

GetHighScaleNumberOk returns a tuple with the HighScaleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighScaleNumber

`func (o *FieldDescribe) SetHighScaleNumber(v bool)`

SetHighScaleNumber sets HighScaleNumber field to given value.

### HasHighScaleNumber

`func (o *FieldDescribe) HasHighScaleNumber() bool`

HasHighScaleNumber returns a boolean if a field has been set.

### GetHtmlFormatted

`func (o *FieldDescribe) GetHtmlFormatted() bool`

GetHtmlFormatted returns the HtmlFormatted field if non-nil, zero value otherwise.

### GetHtmlFormattedOk

`func (o *FieldDescribe) GetHtmlFormattedOk() (*bool, bool)`

GetHtmlFormattedOk returns a tuple with the HtmlFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlFormatted

`func (o *FieldDescribe) SetHtmlFormatted(v bool)`

SetHtmlFormatted sets HtmlFormatted field to given value.

### HasHtmlFormatted

`func (o *FieldDescribe) HasHtmlFormatted() bool`

HasHtmlFormatted returns a boolean if a field has been set.

### GetIdLookup

`func (o *FieldDescribe) GetIdLookup() bool`

GetIdLookup returns the IdLookup field if non-nil, zero value otherwise.

### GetIdLookupOk

`func (o *FieldDescribe) GetIdLookupOk() (*bool, bool)`

GetIdLookupOk returns a tuple with the IdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdLookup

`func (o *FieldDescribe) SetIdLookup(v bool)`

SetIdLookup sets IdLookup field to given value.

### HasIdLookup

`func (o *FieldDescribe) HasIdLookup() bool`

HasIdLookup returns a boolean if a field has been set.

### GetInlineHelpText

`func (o *FieldDescribe) GetInlineHelpText() string`

GetInlineHelpText returns the InlineHelpText field if non-nil, zero value otherwise.

### GetInlineHelpTextOk

`func (o *FieldDescribe) GetInlineHelpTextOk() (*string, bool)`

GetInlineHelpTextOk returns a tuple with the InlineHelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHelpText

`func (o *FieldDescribe) SetInlineHelpText(v string)`

SetInlineHelpText sets InlineHelpText field to given value.

### HasInlineHelpText

`func (o *FieldDescribe) HasInlineHelpText() bool`

HasInlineHelpText returns a boolean if a field has been set.

### SetInlineHelpTextNil

`func (o *FieldDescribe) SetInlineHelpTextNil(b bool)`

 SetInlineHelpTextNil sets the value for InlineHelpText to be an explicit nil

### UnsetInlineHelpText
`func (o *FieldDescribe) UnsetInlineHelpText()`

UnsetInlineHelpText ensures that no value is present for InlineHelpText, not even an explicit nil
### GetLabel

`func (o *FieldDescribe) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldDescribe) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldDescribe) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FieldDescribe) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLength

`func (o *FieldDescribe) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *FieldDescribe) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *FieldDescribe) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *FieldDescribe) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMask

`func (o *FieldDescribe) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *FieldDescribe) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *FieldDescribe) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *FieldDescribe) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *FieldDescribe) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *FieldDescribe) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetMaskType

`func (o *FieldDescribe) GetMaskType() string`

GetMaskType returns the MaskType field if non-nil, zero value otherwise.

### GetMaskTypeOk

`func (o *FieldDescribe) GetMaskTypeOk() (*string, bool)`

GetMaskTypeOk returns a tuple with the MaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskType

`func (o *FieldDescribe) SetMaskType(v string)`

SetMaskType sets MaskType field to given value.

### HasMaskType

`func (o *FieldDescribe) HasMaskType() bool`

HasMaskType returns a boolean if a field has been set.

### SetMaskTypeNil

`func (o *FieldDescribe) SetMaskTypeNil(b bool)`

 SetMaskTypeNil sets the value for MaskType to be an explicit nil

### UnsetMaskType
`func (o *FieldDescribe) UnsetMaskType()`

UnsetMaskType ensures that no value is present for MaskType, not even an explicit nil
### GetName

`func (o *FieldDescribe) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDescribe) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDescribe) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldDescribe) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameField

`func (o *FieldDescribe) GetNameField() bool`

GetNameField returns the NameField field if non-nil, zero value otherwise.

### GetNameFieldOk

`func (o *FieldDescribe) GetNameFieldOk() (*bool, bool)`

GetNameFieldOk returns a tuple with the NameField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameField

`func (o *FieldDescribe) SetNameField(v bool)`

SetNameField sets NameField field to given value.

### HasNameField

`func (o *FieldDescribe) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetNamePointing

`func (o *FieldDescribe) GetNamePointing() bool`

GetNamePointing returns the NamePointing field if non-nil, zero value otherwise.

### GetNamePointingOk

`func (o *FieldDescribe) GetNamePointingOk() (*bool, bool)`

GetNamePointingOk returns a tuple with the NamePointing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePointing

`func (o *FieldDescribe) SetNamePointing(v bool)`

SetNamePointing sets NamePointing field to given value.

### HasNamePointing

`func (o *FieldDescribe) HasNamePointing() bool`

HasNamePointing returns a boolean if a field has been set.

### GetNillable

`func (o *FieldDescribe) GetNillable() bool`

GetNillable returns the Nillable field if non-nil, zero value otherwise.

### GetNillableOk

`func (o *FieldDescribe) GetNillableOk() (*bool, bool)`

GetNillableOk returns a tuple with the Nillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNillable

`func (o *FieldDescribe) SetNillable(v bool)`

SetNillable sets Nillable field to given value.

### HasNillable

`func (o *FieldDescribe) HasNillable() bool`

HasNillable returns a boolean if a field has been set.

### GetPermissionable

`func (o *FieldDescribe) GetPermissionable() bool`

GetPermissionable returns the Permissionable field if non-nil, zero value otherwise.

### GetPermissionableOk

`func (o *FieldDescribe) GetPermissionableOk() (*bool, bool)`

GetPermissionableOk returns a tuple with the Permissionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionable

`func (o *FieldDescribe) SetPermissionable(v bool)`

SetPermissionable sets Permissionable field to given value.

### HasPermissionable

`func (o *FieldDescribe) HasPermissionable() bool`

HasPermissionable returns a boolean if a field has been set.

### GetPicklistValues

`func (o *FieldDescribe) GetPicklistValues() []FieldDescribePicklist`

GetPicklistValues returns the PicklistValues field if non-nil, zero value otherwise.

### GetPicklistValuesOk

`func (o *FieldDescribe) GetPicklistValuesOk() (*[]FieldDescribePicklist, bool)`

GetPicklistValuesOk returns a tuple with the PicklistValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicklistValues

`func (o *FieldDescribe) SetPicklistValues(v []FieldDescribePicklist)`

SetPicklistValues sets PicklistValues field to given value.

### HasPicklistValues

`func (o *FieldDescribe) HasPicklistValues() bool`

HasPicklistValues returns a boolean if a field has been set.

### GetPolymorphicForeignKey

`func (o *FieldDescribe) GetPolymorphicForeignKey() bool`

GetPolymorphicForeignKey returns the PolymorphicForeignKey field if non-nil, zero value otherwise.

### GetPolymorphicForeignKeyOk

`func (o *FieldDescribe) GetPolymorphicForeignKeyOk() (*bool, bool)`

GetPolymorphicForeignKeyOk returns a tuple with the PolymorphicForeignKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolymorphicForeignKey

`func (o *FieldDescribe) SetPolymorphicForeignKey(v bool)`

SetPolymorphicForeignKey sets PolymorphicForeignKey field to given value.

### HasPolymorphicForeignKey

`func (o *FieldDescribe) HasPolymorphicForeignKey() bool`

HasPolymorphicForeignKey returns a boolean if a field has been set.

### GetPrecision

`func (o *FieldDescribe) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *FieldDescribe) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *FieldDescribe) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *FieldDescribe) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetQueryByDistance

`func (o *FieldDescribe) GetQueryByDistance() bool`

GetQueryByDistance returns the QueryByDistance field if non-nil, zero value otherwise.

### GetQueryByDistanceOk

`func (o *FieldDescribe) GetQueryByDistanceOk() (*bool, bool)`

GetQueryByDistanceOk returns a tuple with the QueryByDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryByDistance

`func (o *FieldDescribe) SetQueryByDistance(v bool)`

SetQueryByDistance sets QueryByDistance field to given value.

### HasQueryByDistance

`func (o *FieldDescribe) HasQueryByDistance() bool`

HasQueryByDistance returns a boolean if a field has been set.

### GetReferenceTargetField

`func (o *FieldDescribe) GetReferenceTargetField() string`

GetReferenceTargetField returns the ReferenceTargetField field if non-nil, zero value otherwise.

### GetReferenceTargetFieldOk

`func (o *FieldDescribe) GetReferenceTargetFieldOk() (*string, bool)`

GetReferenceTargetFieldOk returns a tuple with the ReferenceTargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTargetField

`func (o *FieldDescribe) SetReferenceTargetField(v string)`

SetReferenceTargetField sets ReferenceTargetField field to given value.

### HasReferenceTargetField

`func (o *FieldDescribe) HasReferenceTargetField() bool`

HasReferenceTargetField returns a boolean if a field has been set.

### SetReferenceTargetFieldNil

`func (o *FieldDescribe) SetReferenceTargetFieldNil(b bool)`

 SetReferenceTargetFieldNil sets the value for ReferenceTargetField to be an explicit nil

### UnsetReferenceTargetField
`func (o *FieldDescribe) UnsetReferenceTargetField()`

UnsetReferenceTargetField ensures that no value is present for ReferenceTargetField, not even an explicit nil
### GetReferenceTo

`func (o *FieldDescribe) GetReferenceTo() []string`

GetReferenceTo returns the ReferenceTo field if non-nil, zero value otherwise.

### GetReferenceToOk

`func (o *FieldDescribe) GetReferenceToOk() (*[]string, bool)`

GetReferenceToOk returns a tuple with the ReferenceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTo

`func (o *FieldDescribe) SetReferenceTo(v []string)`

SetReferenceTo sets ReferenceTo field to given value.

### HasReferenceTo

`func (o *FieldDescribe) HasReferenceTo() bool`

HasReferenceTo returns a boolean if a field has been set.

### GetRelationshipName

`func (o *FieldDescribe) GetRelationshipName() string`

GetRelationshipName returns the RelationshipName field if non-nil, zero value otherwise.

### GetRelationshipNameOk

`func (o *FieldDescribe) GetRelationshipNameOk() (*string, bool)`

GetRelationshipNameOk returns a tuple with the RelationshipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipName

`func (o *FieldDescribe) SetRelationshipName(v string)`

SetRelationshipName sets RelationshipName field to given value.

### HasRelationshipName

`func (o *FieldDescribe) HasRelationshipName() bool`

HasRelationshipName returns a boolean if a field has been set.

### SetRelationshipNameNil

`func (o *FieldDescribe) SetRelationshipNameNil(b bool)`

 SetRelationshipNameNil sets the value for RelationshipName to be an explicit nil

### UnsetRelationshipName
`func (o *FieldDescribe) UnsetRelationshipName()`

UnsetRelationshipName ensures that no value is present for RelationshipName, not even an explicit nil
### GetRelationshipOrder

`func (o *FieldDescribe) GetRelationshipOrder() int32`

GetRelationshipOrder returns the RelationshipOrder field if non-nil, zero value otherwise.

### GetRelationshipOrderOk

`func (o *FieldDescribe) GetRelationshipOrderOk() (*int32, bool)`

GetRelationshipOrderOk returns a tuple with the RelationshipOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipOrder

`func (o *FieldDescribe) SetRelationshipOrder(v int32)`

SetRelationshipOrder sets RelationshipOrder field to given value.

### HasRelationshipOrder

`func (o *FieldDescribe) HasRelationshipOrder() bool`

HasRelationshipOrder returns a boolean if a field has been set.

### SetRelationshipOrderNil

`func (o *FieldDescribe) SetRelationshipOrderNil(b bool)`

 SetRelationshipOrderNil sets the value for RelationshipOrder to be an explicit nil

### UnsetRelationshipOrder
`func (o *FieldDescribe) UnsetRelationshipOrder()`

UnsetRelationshipOrder ensures that no value is present for RelationshipOrder, not even an explicit nil
### GetRestrictedDelete

`func (o *FieldDescribe) GetRestrictedDelete() bool`

GetRestrictedDelete returns the RestrictedDelete field if non-nil, zero value otherwise.

### GetRestrictedDeleteOk

`func (o *FieldDescribe) GetRestrictedDeleteOk() (*bool, bool)`

GetRestrictedDeleteOk returns a tuple with the RestrictedDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDelete

`func (o *FieldDescribe) SetRestrictedDelete(v bool)`

SetRestrictedDelete sets RestrictedDelete field to given value.

### HasRestrictedDelete

`func (o *FieldDescribe) HasRestrictedDelete() bool`

HasRestrictedDelete returns a boolean if a field has been set.

### GetRestrictedPicklist

`func (o *FieldDescribe) GetRestrictedPicklist() bool`

GetRestrictedPicklist returns the RestrictedPicklist field if non-nil, zero value otherwise.

### GetRestrictedPicklistOk

`func (o *FieldDescribe) GetRestrictedPicklistOk() (*bool, bool)`

GetRestrictedPicklistOk returns a tuple with the RestrictedPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPicklist

`func (o *FieldDescribe) SetRestrictedPicklist(v bool)`

SetRestrictedPicklist sets RestrictedPicklist field to given value.

### HasRestrictedPicklist

`func (o *FieldDescribe) HasRestrictedPicklist() bool`

HasRestrictedPicklist returns a boolean if a field has been set.

### GetScale

`func (o *FieldDescribe) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *FieldDescribe) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *FieldDescribe) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *FieldDescribe) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetSearchPrefilterable

`func (o *FieldDescribe) GetSearchPrefilterable() bool`

GetSearchPrefilterable returns the SearchPrefilterable field if non-nil, zero value otherwise.

### GetSearchPrefilterableOk

`func (o *FieldDescribe) GetSearchPrefilterableOk() (*bool, bool)`

GetSearchPrefilterableOk returns a tuple with the SearchPrefilterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPrefilterable

`func (o *FieldDescribe) SetSearchPrefilterable(v bool)`

SetSearchPrefilterable sets SearchPrefilterable field to given value.

### HasSearchPrefilterable

`func (o *FieldDescribe) HasSearchPrefilterable() bool`

HasSearchPrefilterable returns a boolean if a field has been set.

### GetSoapType

`func (o *FieldDescribe) GetSoapType() string`

GetSoapType returns the SoapType field if non-nil, zero value otherwise.

### GetSoapTypeOk

`func (o *FieldDescribe) GetSoapTypeOk() (*string, bool)`

GetSoapTypeOk returns a tuple with the SoapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoapType

`func (o *FieldDescribe) SetSoapType(v string)`

SetSoapType sets SoapType field to given value.

### HasSoapType

`func (o *FieldDescribe) HasSoapType() bool`

HasSoapType returns a boolean if a field has been set.

### GetSortable

`func (o *FieldDescribe) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *FieldDescribe) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *FieldDescribe) SetSortable(v bool)`

SetSortable sets Sortable field to given value.

### HasSortable

`func (o *FieldDescribe) HasSortable() bool`

HasSortable returns a boolean if a field has been set.

### GetType

`func (o *FieldDescribe) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDescribe) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDescribe) SetType(v FieldType)`

SetType sets Type field to given value.

### HasType

`func (o *FieldDescribe) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnique

`func (o *FieldDescribe) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *FieldDescribe) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *FieldDescribe) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *FieldDescribe) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetUpdateable

`func (o *FieldDescribe) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *FieldDescribe) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *FieldDescribe) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *FieldDescribe) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetWriteRequiresMasterRead

`func (o *FieldDescribe) GetWriteRequiresMasterRead() bool`

GetWriteRequiresMasterRead returns the WriteRequiresMasterRead field if non-nil, zero value otherwise.

### GetWriteRequiresMasterReadOk

`func (o *FieldDescribe) GetWriteRequiresMasterReadOk() (*bool, bool)`

GetWriteRequiresMasterReadOk returns a tuple with the WriteRequiresMasterRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteRequiresMasterRead

`func (o *FieldDescribe) SetWriteRequiresMasterRead(v bool)`

SetWriteRequiresMasterRead sets WriteRequiresMasterRead field to given value.

### HasWriteRequiresMasterRead

`func (o *FieldDescribe) HasWriteRequiresMasterRead() bool`

HasWriteRequiresMasterRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


