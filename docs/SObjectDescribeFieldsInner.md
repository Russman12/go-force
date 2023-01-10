# SObjectDescribeFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregatable** | Pointer to **bool** |  | [optional] 
**AiPredictionField** | Pointer to **bool** |  | [optional] 
**AutoNumber** | Pointer to **bool** |  | [optional] 
**ByteLength** | Pointer to **int32** |  | [optional] 
**Calculated** | Pointer to **bool** |  | [optional] 
**CalculatedFormula** | Pointer to **string** |  | [optional] 
**CascadeDelete** | Pointer to **bool** |  | [optional] 
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**CompoundFieldName** | Pointer to **string** |  | [optional] 
**ControllerName** | Pointer to **string** |  | [optional] 
**Createable** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**DefaultValueFormula** | Pointer to **string** |  | [optional] 
**DefaultedOnCreate** | Pointer to **bool** |  | [optional] 
**DependentPicklist** | Pointer to **bool** |  | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** |  | [optional] 
**Digits** | Pointer to **int32** |  | [optional] 
**DisplayLocationInDecimal** | Pointer to **bool** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **bool** |  | [optional] 
**ExtraTypeInfo** | Pointer to **string** |  | [optional] 
**Filterable** | Pointer to **bool** |  | [optional] 
**FilteredLookupInfo** | Pointer to **string** |  | [optional] 
**FormulaTreatNullNumberAsZero** | Pointer to **bool** |  | [optional] 
**Groupable** | Pointer to **bool** |  | [optional] 
**HighScaleNumber** | Pointer to **bool** |  | [optional] 
**HtmlFormatted** | Pointer to **bool** |  | [optional] 
**IdLookup** | Pointer to **bool** |  | [optional] 
**InlineHelpText** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**Mask** | Pointer to **string** |  | [optional] 
**MaskType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameField** | Pointer to **bool** |  | [optional] 
**NamePointing** | Pointer to **bool** |  | [optional] 
**Nillable** | Pointer to **bool** |  | [optional] 
**Permissionable** | Pointer to **bool** |  | [optional] 
**PicklistValues** | Pointer to [**[]SObjectDescribeFieldsInnerPicklistValuesInner**](SObjectDescribeFieldsInnerPicklistValuesInner.md) |  | [optional] 
**PolymorphicForeignKey** | Pointer to **bool** |  | [optional] 
**Precision** | Pointer to **int32** |  | [optional] 
**QueryByDistance** | Pointer to **bool** |  | [optional] 
**ReferenceTargetField** | Pointer to **string** |  | [optional] 
**ReferenceTo** | Pointer to **[]string** |  | [optional] 
**RelationshipName** | Pointer to **string** |  | [optional] 
**RelationshipOrder** | Pointer to **string** |  | [optional] 
**RestrictedDelete** | Pointer to **bool** |  | [optional] 
**RestrictedPicklist** | Pointer to **bool** |  | [optional] 
**Scale** | Pointer to **int32** |  | [optional] 
**SearchPrefilterable** | Pointer to **bool** |  | [optional] 
**SoapType** | Pointer to **string** |  | [optional] 
**Sortable** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**Updateable** | Pointer to **bool** |  | [optional] 
**WriteRequiresMasterRead** | Pointer to **bool** |  | [optional] 

## Methods

### NewSObjectDescribeFieldsInner

`func NewSObjectDescribeFieldsInner() *SObjectDescribeFieldsInner`

NewSObjectDescribeFieldsInner instantiates a new SObjectDescribeFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSObjectDescribeFieldsInnerWithDefaults

`func NewSObjectDescribeFieldsInnerWithDefaults() *SObjectDescribeFieldsInner`

NewSObjectDescribeFieldsInnerWithDefaults instantiates a new SObjectDescribeFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatable

`func (o *SObjectDescribeFieldsInner) GetAggregatable() bool`

GetAggregatable returns the Aggregatable field if non-nil, zero value otherwise.

### GetAggregatableOk

`func (o *SObjectDescribeFieldsInner) GetAggregatableOk() (*bool, bool)`

GetAggregatableOk returns a tuple with the Aggregatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatable

`func (o *SObjectDescribeFieldsInner) SetAggregatable(v bool)`

SetAggregatable sets Aggregatable field to given value.

### HasAggregatable

`func (o *SObjectDescribeFieldsInner) HasAggregatable() bool`

HasAggregatable returns a boolean if a field has been set.

### GetAiPredictionField

`func (o *SObjectDescribeFieldsInner) GetAiPredictionField() bool`

GetAiPredictionField returns the AiPredictionField field if non-nil, zero value otherwise.

### GetAiPredictionFieldOk

`func (o *SObjectDescribeFieldsInner) GetAiPredictionFieldOk() (*bool, bool)`

GetAiPredictionFieldOk returns a tuple with the AiPredictionField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPredictionField

`func (o *SObjectDescribeFieldsInner) SetAiPredictionField(v bool)`

SetAiPredictionField sets AiPredictionField field to given value.

### HasAiPredictionField

`func (o *SObjectDescribeFieldsInner) HasAiPredictionField() bool`

HasAiPredictionField returns a boolean if a field has been set.

### GetAutoNumber

`func (o *SObjectDescribeFieldsInner) GetAutoNumber() bool`

GetAutoNumber returns the AutoNumber field if non-nil, zero value otherwise.

### GetAutoNumberOk

`func (o *SObjectDescribeFieldsInner) GetAutoNumberOk() (*bool, bool)`

GetAutoNumberOk returns a tuple with the AutoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNumber

`func (o *SObjectDescribeFieldsInner) SetAutoNumber(v bool)`

SetAutoNumber sets AutoNumber field to given value.

### HasAutoNumber

`func (o *SObjectDescribeFieldsInner) HasAutoNumber() bool`

HasAutoNumber returns a boolean if a field has been set.

### GetByteLength

`func (o *SObjectDescribeFieldsInner) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *SObjectDescribeFieldsInner) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *SObjectDescribeFieldsInner) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *SObjectDescribeFieldsInner) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetCalculated

`func (o *SObjectDescribeFieldsInner) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *SObjectDescribeFieldsInner) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *SObjectDescribeFieldsInner) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *SObjectDescribeFieldsInner) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetCalculatedFormula

`func (o *SObjectDescribeFieldsInner) GetCalculatedFormula() string`

GetCalculatedFormula returns the CalculatedFormula field if non-nil, zero value otherwise.

### GetCalculatedFormulaOk

`func (o *SObjectDescribeFieldsInner) GetCalculatedFormulaOk() (*string, bool)`

GetCalculatedFormulaOk returns a tuple with the CalculatedFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedFormula

`func (o *SObjectDescribeFieldsInner) SetCalculatedFormula(v string)`

SetCalculatedFormula sets CalculatedFormula field to given value.

### HasCalculatedFormula

`func (o *SObjectDescribeFieldsInner) HasCalculatedFormula() bool`

HasCalculatedFormula returns a boolean if a field has been set.

### GetCascadeDelete

`func (o *SObjectDescribeFieldsInner) GetCascadeDelete() bool`

GetCascadeDelete returns the CascadeDelete field if non-nil, zero value otherwise.

### GetCascadeDeleteOk

`func (o *SObjectDescribeFieldsInner) GetCascadeDeleteOk() (*bool, bool)`

GetCascadeDeleteOk returns a tuple with the CascadeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeDelete

`func (o *SObjectDescribeFieldsInner) SetCascadeDelete(v bool)`

SetCascadeDelete sets CascadeDelete field to given value.

### HasCascadeDelete

`func (o *SObjectDescribeFieldsInner) HasCascadeDelete() bool`

HasCascadeDelete returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *SObjectDescribeFieldsInner) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *SObjectDescribeFieldsInner) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *SObjectDescribeFieldsInner) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *SObjectDescribeFieldsInner) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetCompoundFieldName

`func (o *SObjectDescribeFieldsInner) GetCompoundFieldName() string`

GetCompoundFieldName returns the CompoundFieldName field if non-nil, zero value otherwise.

### GetCompoundFieldNameOk

`func (o *SObjectDescribeFieldsInner) GetCompoundFieldNameOk() (*string, bool)`

GetCompoundFieldNameOk returns a tuple with the CompoundFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompoundFieldName

`func (o *SObjectDescribeFieldsInner) SetCompoundFieldName(v string)`

SetCompoundFieldName sets CompoundFieldName field to given value.

### HasCompoundFieldName

`func (o *SObjectDescribeFieldsInner) HasCompoundFieldName() bool`

HasCompoundFieldName returns a boolean if a field has been set.

### GetControllerName

`func (o *SObjectDescribeFieldsInner) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *SObjectDescribeFieldsInner) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *SObjectDescribeFieldsInner) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *SObjectDescribeFieldsInner) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetCreateable

`func (o *SObjectDescribeFieldsInner) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *SObjectDescribeFieldsInner) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *SObjectDescribeFieldsInner) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *SObjectDescribeFieldsInner) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *SObjectDescribeFieldsInner) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *SObjectDescribeFieldsInner) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *SObjectDescribeFieldsInner) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *SObjectDescribeFieldsInner) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SObjectDescribeFieldsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SObjectDescribeFieldsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SObjectDescribeFieldsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SObjectDescribeFieldsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDefaultValueFormula

`func (o *SObjectDescribeFieldsInner) GetDefaultValueFormula() string`

GetDefaultValueFormula returns the DefaultValueFormula field if non-nil, zero value otherwise.

### GetDefaultValueFormulaOk

`func (o *SObjectDescribeFieldsInner) GetDefaultValueFormulaOk() (*string, bool)`

GetDefaultValueFormulaOk returns a tuple with the DefaultValueFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueFormula

`func (o *SObjectDescribeFieldsInner) SetDefaultValueFormula(v string)`

SetDefaultValueFormula sets DefaultValueFormula field to given value.

### HasDefaultValueFormula

`func (o *SObjectDescribeFieldsInner) HasDefaultValueFormula() bool`

HasDefaultValueFormula returns a boolean if a field has been set.

### GetDefaultedOnCreate

`func (o *SObjectDescribeFieldsInner) GetDefaultedOnCreate() bool`

GetDefaultedOnCreate returns the DefaultedOnCreate field if non-nil, zero value otherwise.

### GetDefaultedOnCreateOk

`func (o *SObjectDescribeFieldsInner) GetDefaultedOnCreateOk() (*bool, bool)`

GetDefaultedOnCreateOk returns a tuple with the DefaultedOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultedOnCreate

`func (o *SObjectDescribeFieldsInner) SetDefaultedOnCreate(v bool)`

SetDefaultedOnCreate sets DefaultedOnCreate field to given value.

### HasDefaultedOnCreate

`func (o *SObjectDescribeFieldsInner) HasDefaultedOnCreate() bool`

HasDefaultedOnCreate returns a boolean if a field has been set.

### GetDependentPicklist

`func (o *SObjectDescribeFieldsInner) GetDependentPicklist() bool`

GetDependentPicklist returns the DependentPicklist field if non-nil, zero value otherwise.

### GetDependentPicklistOk

`func (o *SObjectDescribeFieldsInner) GetDependentPicklistOk() (*bool, bool)`

GetDependentPicklistOk returns a tuple with the DependentPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentPicklist

`func (o *SObjectDescribeFieldsInner) SetDependentPicklist(v bool)`

SetDependentPicklist sets DependentPicklist field to given value.

### HasDependentPicklist

`func (o *SObjectDescribeFieldsInner) HasDependentPicklist() bool`

HasDependentPicklist returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *SObjectDescribeFieldsInner) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *SObjectDescribeFieldsInner) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *SObjectDescribeFieldsInner) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *SObjectDescribeFieldsInner) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetDigits

`func (o *SObjectDescribeFieldsInner) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *SObjectDescribeFieldsInner) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *SObjectDescribeFieldsInner) SetDigits(v int32)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *SObjectDescribeFieldsInner) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetDisplayLocationInDecimal

`func (o *SObjectDescribeFieldsInner) GetDisplayLocationInDecimal() bool`

GetDisplayLocationInDecimal returns the DisplayLocationInDecimal field if non-nil, zero value otherwise.

### GetDisplayLocationInDecimalOk

`func (o *SObjectDescribeFieldsInner) GetDisplayLocationInDecimalOk() (*bool, bool)`

GetDisplayLocationInDecimalOk returns a tuple with the DisplayLocationInDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocationInDecimal

`func (o *SObjectDescribeFieldsInner) SetDisplayLocationInDecimal(v bool)`

SetDisplayLocationInDecimal sets DisplayLocationInDecimal field to given value.

### HasDisplayLocationInDecimal

`func (o *SObjectDescribeFieldsInner) HasDisplayLocationInDecimal() bool`

HasDisplayLocationInDecimal returns a boolean if a field has been set.

### GetEncrypted

`func (o *SObjectDescribeFieldsInner) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *SObjectDescribeFieldsInner) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *SObjectDescribeFieldsInner) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *SObjectDescribeFieldsInner) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetExternalId

`func (o *SObjectDescribeFieldsInner) GetExternalId() bool`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SObjectDescribeFieldsInner) GetExternalIdOk() (*bool, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SObjectDescribeFieldsInner) SetExternalId(v bool)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SObjectDescribeFieldsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExtraTypeInfo

`func (o *SObjectDescribeFieldsInner) GetExtraTypeInfo() string`

GetExtraTypeInfo returns the ExtraTypeInfo field if non-nil, zero value otherwise.

### GetExtraTypeInfoOk

`func (o *SObjectDescribeFieldsInner) GetExtraTypeInfoOk() (*string, bool)`

GetExtraTypeInfoOk returns a tuple with the ExtraTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraTypeInfo

`func (o *SObjectDescribeFieldsInner) SetExtraTypeInfo(v string)`

SetExtraTypeInfo sets ExtraTypeInfo field to given value.

### HasExtraTypeInfo

`func (o *SObjectDescribeFieldsInner) HasExtraTypeInfo() bool`

HasExtraTypeInfo returns a boolean if a field has been set.

### GetFilterable

`func (o *SObjectDescribeFieldsInner) GetFilterable() bool`

GetFilterable returns the Filterable field if non-nil, zero value otherwise.

### GetFilterableOk

`func (o *SObjectDescribeFieldsInner) GetFilterableOk() (*bool, bool)`

GetFilterableOk returns a tuple with the Filterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterable

`func (o *SObjectDescribeFieldsInner) SetFilterable(v bool)`

SetFilterable sets Filterable field to given value.

### HasFilterable

`func (o *SObjectDescribeFieldsInner) HasFilterable() bool`

HasFilterable returns a boolean if a field has been set.

### GetFilteredLookupInfo

`func (o *SObjectDescribeFieldsInner) GetFilteredLookupInfo() string`

GetFilteredLookupInfo returns the FilteredLookupInfo field if non-nil, zero value otherwise.

### GetFilteredLookupInfoOk

`func (o *SObjectDescribeFieldsInner) GetFilteredLookupInfoOk() (*string, bool)`

GetFilteredLookupInfoOk returns a tuple with the FilteredLookupInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredLookupInfo

`func (o *SObjectDescribeFieldsInner) SetFilteredLookupInfo(v string)`

SetFilteredLookupInfo sets FilteredLookupInfo field to given value.

### HasFilteredLookupInfo

`func (o *SObjectDescribeFieldsInner) HasFilteredLookupInfo() bool`

HasFilteredLookupInfo returns a boolean if a field has been set.

### GetFormulaTreatNullNumberAsZero

`func (o *SObjectDescribeFieldsInner) GetFormulaTreatNullNumberAsZero() bool`

GetFormulaTreatNullNumberAsZero returns the FormulaTreatNullNumberAsZero field if non-nil, zero value otherwise.

### GetFormulaTreatNullNumberAsZeroOk

`func (o *SObjectDescribeFieldsInner) GetFormulaTreatNullNumberAsZeroOk() (*bool, bool)`

GetFormulaTreatNullNumberAsZeroOk returns a tuple with the FormulaTreatNullNumberAsZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaTreatNullNumberAsZero

`func (o *SObjectDescribeFieldsInner) SetFormulaTreatNullNumberAsZero(v bool)`

SetFormulaTreatNullNumberAsZero sets FormulaTreatNullNumberAsZero field to given value.

### HasFormulaTreatNullNumberAsZero

`func (o *SObjectDescribeFieldsInner) HasFormulaTreatNullNumberAsZero() bool`

HasFormulaTreatNullNumberAsZero returns a boolean if a field has been set.

### GetGroupable

`func (o *SObjectDescribeFieldsInner) GetGroupable() bool`

GetGroupable returns the Groupable field if non-nil, zero value otherwise.

### GetGroupableOk

`func (o *SObjectDescribeFieldsInner) GetGroupableOk() (*bool, bool)`

GetGroupableOk returns a tuple with the Groupable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupable

`func (o *SObjectDescribeFieldsInner) SetGroupable(v bool)`

SetGroupable sets Groupable field to given value.

### HasGroupable

`func (o *SObjectDescribeFieldsInner) HasGroupable() bool`

HasGroupable returns a boolean if a field has been set.

### GetHighScaleNumber

`func (o *SObjectDescribeFieldsInner) GetHighScaleNumber() bool`

GetHighScaleNumber returns the HighScaleNumber field if non-nil, zero value otherwise.

### GetHighScaleNumberOk

`func (o *SObjectDescribeFieldsInner) GetHighScaleNumberOk() (*bool, bool)`

GetHighScaleNumberOk returns a tuple with the HighScaleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighScaleNumber

`func (o *SObjectDescribeFieldsInner) SetHighScaleNumber(v bool)`

SetHighScaleNumber sets HighScaleNumber field to given value.

### HasHighScaleNumber

`func (o *SObjectDescribeFieldsInner) HasHighScaleNumber() bool`

HasHighScaleNumber returns a boolean if a field has been set.

### GetHtmlFormatted

`func (o *SObjectDescribeFieldsInner) GetHtmlFormatted() bool`

GetHtmlFormatted returns the HtmlFormatted field if non-nil, zero value otherwise.

### GetHtmlFormattedOk

`func (o *SObjectDescribeFieldsInner) GetHtmlFormattedOk() (*bool, bool)`

GetHtmlFormattedOk returns a tuple with the HtmlFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlFormatted

`func (o *SObjectDescribeFieldsInner) SetHtmlFormatted(v bool)`

SetHtmlFormatted sets HtmlFormatted field to given value.

### HasHtmlFormatted

`func (o *SObjectDescribeFieldsInner) HasHtmlFormatted() bool`

HasHtmlFormatted returns a boolean if a field has been set.

### GetIdLookup

`func (o *SObjectDescribeFieldsInner) GetIdLookup() bool`

GetIdLookup returns the IdLookup field if non-nil, zero value otherwise.

### GetIdLookupOk

`func (o *SObjectDescribeFieldsInner) GetIdLookupOk() (*bool, bool)`

GetIdLookupOk returns a tuple with the IdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdLookup

`func (o *SObjectDescribeFieldsInner) SetIdLookup(v bool)`

SetIdLookup sets IdLookup field to given value.

### HasIdLookup

`func (o *SObjectDescribeFieldsInner) HasIdLookup() bool`

HasIdLookup returns a boolean if a field has been set.

### GetInlineHelpText

`func (o *SObjectDescribeFieldsInner) GetInlineHelpText() string`

GetInlineHelpText returns the InlineHelpText field if non-nil, zero value otherwise.

### GetInlineHelpTextOk

`func (o *SObjectDescribeFieldsInner) GetInlineHelpTextOk() (*string, bool)`

GetInlineHelpTextOk returns a tuple with the InlineHelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHelpText

`func (o *SObjectDescribeFieldsInner) SetInlineHelpText(v string)`

SetInlineHelpText sets InlineHelpText field to given value.

### HasInlineHelpText

`func (o *SObjectDescribeFieldsInner) HasInlineHelpText() bool`

HasInlineHelpText returns a boolean if a field has been set.

### GetLabel

`func (o *SObjectDescribeFieldsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SObjectDescribeFieldsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SObjectDescribeFieldsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SObjectDescribeFieldsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLength

`func (o *SObjectDescribeFieldsInner) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *SObjectDescribeFieldsInner) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *SObjectDescribeFieldsInner) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *SObjectDescribeFieldsInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMask

`func (o *SObjectDescribeFieldsInner) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *SObjectDescribeFieldsInner) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *SObjectDescribeFieldsInner) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *SObjectDescribeFieldsInner) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetMaskType

`func (o *SObjectDescribeFieldsInner) GetMaskType() string`

GetMaskType returns the MaskType field if non-nil, zero value otherwise.

### GetMaskTypeOk

`func (o *SObjectDescribeFieldsInner) GetMaskTypeOk() (*string, bool)`

GetMaskTypeOk returns a tuple with the MaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskType

`func (o *SObjectDescribeFieldsInner) SetMaskType(v string)`

SetMaskType sets MaskType field to given value.

### HasMaskType

`func (o *SObjectDescribeFieldsInner) HasMaskType() bool`

HasMaskType returns a boolean if a field has been set.

### GetName

`func (o *SObjectDescribeFieldsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SObjectDescribeFieldsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SObjectDescribeFieldsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SObjectDescribeFieldsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameField

`func (o *SObjectDescribeFieldsInner) GetNameField() bool`

GetNameField returns the NameField field if non-nil, zero value otherwise.

### GetNameFieldOk

`func (o *SObjectDescribeFieldsInner) GetNameFieldOk() (*bool, bool)`

GetNameFieldOk returns a tuple with the NameField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameField

`func (o *SObjectDescribeFieldsInner) SetNameField(v bool)`

SetNameField sets NameField field to given value.

### HasNameField

`func (o *SObjectDescribeFieldsInner) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetNamePointing

`func (o *SObjectDescribeFieldsInner) GetNamePointing() bool`

GetNamePointing returns the NamePointing field if non-nil, zero value otherwise.

### GetNamePointingOk

`func (o *SObjectDescribeFieldsInner) GetNamePointingOk() (*bool, bool)`

GetNamePointingOk returns a tuple with the NamePointing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePointing

`func (o *SObjectDescribeFieldsInner) SetNamePointing(v bool)`

SetNamePointing sets NamePointing field to given value.

### HasNamePointing

`func (o *SObjectDescribeFieldsInner) HasNamePointing() bool`

HasNamePointing returns a boolean if a field has been set.

### GetNillable

`func (o *SObjectDescribeFieldsInner) GetNillable() bool`

GetNillable returns the Nillable field if non-nil, zero value otherwise.

### GetNillableOk

`func (o *SObjectDescribeFieldsInner) GetNillableOk() (*bool, bool)`

GetNillableOk returns a tuple with the Nillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNillable

`func (o *SObjectDescribeFieldsInner) SetNillable(v bool)`

SetNillable sets Nillable field to given value.

### HasNillable

`func (o *SObjectDescribeFieldsInner) HasNillable() bool`

HasNillable returns a boolean if a field has been set.

### GetPermissionable

`func (o *SObjectDescribeFieldsInner) GetPermissionable() bool`

GetPermissionable returns the Permissionable field if non-nil, zero value otherwise.

### GetPermissionableOk

`func (o *SObjectDescribeFieldsInner) GetPermissionableOk() (*bool, bool)`

GetPermissionableOk returns a tuple with the Permissionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionable

`func (o *SObjectDescribeFieldsInner) SetPermissionable(v bool)`

SetPermissionable sets Permissionable field to given value.

### HasPermissionable

`func (o *SObjectDescribeFieldsInner) HasPermissionable() bool`

HasPermissionable returns a boolean if a field has been set.

### GetPicklistValues

`func (o *SObjectDescribeFieldsInner) GetPicklistValues() []SObjectDescribeFieldsInnerPicklistValuesInner`

GetPicklistValues returns the PicklistValues field if non-nil, zero value otherwise.

### GetPicklistValuesOk

`func (o *SObjectDescribeFieldsInner) GetPicklistValuesOk() (*[]SObjectDescribeFieldsInnerPicklistValuesInner, bool)`

GetPicklistValuesOk returns a tuple with the PicklistValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicklistValues

`func (o *SObjectDescribeFieldsInner) SetPicklistValues(v []SObjectDescribeFieldsInnerPicklistValuesInner)`

SetPicklistValues sets PicklistValues field to given value.

### HasPicklistValues

`func (o *SObjectDescribeFieldsInner) HasPicklistValues() bool`

HasPicklistValues returns a boolean if a field has been set.

### GetPolymorphicForeignKey

`func (o *SObjectDescribeFieldsInner) GetPolymorphicForeignKey() bool`

GetPolymorphicForeignKey returns the PolymorphicForeignKey field if non-nil, zero value otherwise.

### GetPolymorphicForeignKeyOk

`func (o *SObjectDescribeFieldsInner) GetPolymorphicForeignKeyOk() (*bool, bool)`

GetPolymorphicForeignKeyOk returns a tuple with the PolymorphicForeignKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolymorphicForeignKey

`func (o *SObjectDescribeFieldsInner) SetPolymorphicForeignKey(v bool)`

SetPolymorphicForeignKey sets PolymorphicForeignKey field to given value.

### HasPolymorphicForeignKey

`func (o *SObjectDescribeFieldsInner) HasPolymorphicForeignKey() bool`

HasPolymorphicForeignKey returns a boolean if a field has been set.

### GetPrecision

`func (o *SObjectDescribeFieldsInner) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *SObjectDescribeFieldsInner) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *SObjectDescribeFieldsInner) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *SObjectDescribeFieldsInner) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetQueryByDistance

`func (o *SObjectDescribeFieldsInner) GetQueryByDistance() bool`

GetQueryByDistance returns the QueryByDistance field if non-nil, zero value otherwise.

### GetQueryByDistanceOk

`func (o *SObjectDescribeFieldsInner) GetQueryByDistanceOk() (*bool, bool)`

GetQueryByDistanceOk returns a tuple with the QueryByDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryByDistance

`func (o *SObjectDescribeFieldsInner) SetQueryByDistance(v bool)`

SetQueryByDistance sets QueryByDistance field to given value.

### HasQueryByDistance

`func (o *SObjectDescribeFieldsInner) HasQueryByDistance() bool`

HasQueryByDistance returns a boolean if a field has been set.

### GetReferenceTargetField

`func (o *SObjectDescribeFieldsInner) GetReferenceTargetField() string`

GetReferenceTargetField returns the ReferenceTargetField field if non-nil, zero value otherwise.

### GetReferenceTargetFieldOk

`func (o *SObjectDescribeFieldsInner) GetReferenceTargetFieldOk() (*string, bool)`

GetReferenceTargetFieldOk returns a tuple with the ReferenceTargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTargetField

`func (o *SObjectDescribeFieldsInner) SetReferenceTargetField(v string)`

SetReferenceTargetField sets ReferenceTargetField field to given value.

### HasReferenceTargetField

`func (o *SObjectDescribeFieldsInner) HasReferenceTargetField() bool`

HasReferenceTargetField returns a boolean if a field has been set.

### GetReferenceTo

`func (o *SObjectDescribeFieldsInner) GetReferenceTo() []string`

GetReferenceTo returns the ReferenceTo field if non-nil, zero value otherwise.

### GetReferenceToOk

`func (o *SObjectDescribeFieldsInner) GetReferenceToOk() (*[]string, bool)`

GetReferenceToOk returns a tuple with the ReferenceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTo

`func (o *SObjectDescribeFieldsInner) SetReferenceTo(v []string)`

SetReferenceTo sets ReferenceTo field to given value.

### HasReferenceTo

`func (o *SObjectDescribeFieldsInner) HasReferenceTo() bool`

HasReferenceTo returns a boolean if a field has been set.

### GetRelationshipName

`func (o *SObjectDescribeFieldsInner) GetRelationshipName() string`

GetRelationshipName returns the RelationshipName field if non-nil, zero value otherwise.

### GetRelationshipNameOk

`func (o *SObjectDescribeFieldsInner) GetRelationshipNameOk() (*string, bool)`

GetRelationshipNameOk returns a tuple with the RelationshipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipName

`func (o *SObjectDescribeFieldsInner) SetRelationshipName(v string)`

SetRelationshipName sets RelationshipName field to given value.

### HasRelationshipName

`func (o *SObjectDescribeFieldsInner) HasRelationshipName() bool`

HasRelationshipName returns a boolean if a field has been set.

### GetRelationshipOrder

`func (o *SObjectDescribeFieldsInner) GetRelationshipOrder() string`

GetRelationshipOrder returns the RelationshipOrder field if non-nil, zero value otherwise.

### GetRelationshipOrderOk

`func (o *SObjectDescribeFieldsInner) GetRelationshipOrderOk() (*string, bool)`

GetRelationshipOrderOk returns a tuple with the RelationshipOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipOrder

`func (o *SObjectDescribeFieldsInner) SetRelationshipOrder(v string)`

SetRelationshipOrder sets RelationshipOrder field to given value.

### HasRelationshipOrder

`func (o *SObjectDescribeFieldsInner) HasRelationshipOrder() bool`

HasRelationshipOrder returns a boolean if a field has been set.

### GetRestrictedDelete

`func (o *SObjectDescribeFieldsInner) GetRestrictedDelete() bool`

GetRestrictedDelete returns the RestrictedDelete field if non-nil, zero value otherwise.

### GetRestrictedDeleteOk

`func (o *SObjectDescribeFieldsInner) GetRestrictedDeleteOk() (*bool, bool)`

GetRestrictedDeleteOk returns a tuple with the RestrictedDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDelete

`func (o *SObjectDescribeFieldsInner) SetRestrictedDelete(v bool)`

SetRestrictedDelete sets RestrictedDelete field to given value.

### HasRestrictedDelete

`func (o *SObjectDescribeFieldsInner) HasRestrictedDelete() bool`

HasRestrictedDelete returns a boolean if a field has been set.

### GetRestrictedPicklist

`func (o *SObjectDescribeFieldsInner) GetRestrictedPicklist() bool`

GetRestrictedPicklist returns the RestrictedPicklist field if non-nil, zero value otherwise.

### GetRestrictedPicklistOk

`func (o *SObjectDescribeFieldsInner) GetRestrictedPicklistOk() (*bool, bool)`

GetRestrictedPicklistOk returns a tuple with the RestrictedPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPicklist

`func (o *SObjectDescribeFieldsInner) SetRestrictedPicklist(v bool)`

SetRestrictedPicklist sets RestrictedPicklist field to given value.

### HasRestrictedPicklist

`func (o *SObjectDescribeFieldsInner) HasRestrictedPicklist() bool`

HasRestrictedPicklist returns a boolean if a field has been set.

### GetScale

`func (o *SObjectDescribeFieldsInner) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *SObjectDescribeFieldsInner) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *SObjectDescribeFieldsInner) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *SObjectDescribeFieldsInner) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetSearchPrefilterable

`func (o *SObjectDescribeFieldsInner) GetSearchPrefilterable() bool`

GetSearchPrefilterable returns the SearchPrefilterable field if non-nil, zero value otherwise.

### GetSearchPrefilterableOk

`func (o *SObjectDescribeFieldsInner) GetSearchPrefilterableOk() (*bool, bool)`

GetSearchPrefilterableOk returns a tuple with the SearchPrefilterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPrefilterable

`func (o *SObjectDescribeFieldsInner) SetSearchPrefilterable(v bool)`

SetSearchPrefilterable sets SearchPrefilterable field to given value.

### HasSearchPrefilterable

`func (o *SObjectDescribeFieldsInner) HasSearchPrefilterable() bool`

HasSearchPrefilterable returns a boolean if a field has been set.

### GetSoapType

`func (o *SObjectDescribeFieldsInner) GetSoapType() string`

GetSoapType returns the SoapType field if non-nil, zero value otherwise.

### GetSoapTypeOk

`func (o *SObjectDescribeFieldsInner) GetSoapTypeOk() (*string, bool)`

GetSoapTypeOk returns a tuple with the SoapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoapType

`func (o *SObjectDescribeFieldsInner) SetSoapType(v string)`

SetSoapType sets SoapType field to given value.

### HasSoapType

`func (o *SObjectDescribeFieldsInner) HasSoapType() bool`

HasSoapType returns a boolean if a field has been set.

### GetSortable

`func (o *SObjectDescribeFieldsInner) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *SObjectDescribeFieldsInner) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *SObjectDescribeFieldsInner) SetSortable(v bool)`

SetSortable sets Sortable field to given value.

### HasSortable

`func (o *SObjectDescribeFieldsInner) HasSortable() bool`

HasSortable returns a boolean if a field has been set.

### GetType

`func (o *SObjectDescribeFieldsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SObjectDescribeFieldsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SObjectDescribeFieldsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SObjectDescribeFieldsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnique

`func (o *SObjectDescribeFieldsInner) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *SObjectDescribeFieldsInner) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *SObjectDescribeFieldsInner) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *SObjectDescribeFieldsInner) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetUpdateable

`func (o *SObjectDescribeFieldsInner) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *SObjectDescribeFieldsInner) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *SObjectDescribeFieldsInner) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *SObjectDescribeFieldsInner) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetWriteRequiresMasterRead

`func (o *SObjectDescribeFieldsInner) GetWriteRequiresMasterRead() bool`

GetWriteRequiresMasterRead returns the WriteRequiresMasterRead field if non-nil, zero value otherwise.

### GetWriteRequiresMasterReadOk

`func (o *SObjectDescribeFieldsInner) GetWriteRequiresMasterReadOk() (*bool, bool)`

GetWriteRequiresMasterReadOk returns a tuple with the WriteRequiresMasterRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteRequiresMasterRead

`func (o *SObjectDescribeFieldsInner) SetWriteRequiresMasterRead(v bool)`

SetWriteRequiresMasterRead sets WriteRequiresMasterRead field to given value.

### HasWriteRequiresMasterRead

`func (o *SObjectDescribeFieldsInner) HasWriteRequiresMasterRead() bool`

HasWriteRequiresMasterRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


