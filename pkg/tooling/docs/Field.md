# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autonumber** | Pointer to **bool** | Indicates whether this field is an autonumber field (true) or not (false). Analogous to a SQL IDENTITY type, autonumber fields are read only, non-createable text fields with a maximum length of 30 characters. Autonumber fields are read-only fields used to provide a unique ID that is independent of the internal object ID (such as a purchase order number or invoice number). Autonumber fields are configured entirely in the Salesforce user interface. The API provides access to this attribute so that client applications can determine whether a given field is an autonumber field. | [optional] 
**ByteLength** | Pointer to **int32** | For variable-length fields (including binary fields), the maximum size of the field, in bytes. | [optional] 
**Calculated** | Pointer to **bool** | Indicates whether the field is a custom formula field (true) or not (false). Note that custom formula fields are always read-only. | [optional] 
**CaseSensitive** | Pointer to **bool** | Indicates whether the field is case sensitive (true) or not (false). | [optional] 
**ControllerName** | Pointer to **string** | The name of the field that controls the values of this picklist. It only applies if type is picklist or multipicklist and dependentPicklist is true. The mapping of controlling field to dependent field is stored in the validFor attribute of each PicklistEntry for this picklist. | [optional] 
**Createable** | Pointer to **bool** | Indicates whether the field can be created (true) or not (false). If true, then this field value can be set in a create() call. | [optional] 
**Custom** | Pointer to **bool** | Indicates whether the field is a custom field (true) or not (false). | [optional] 
**DataTranslationEnabled** | Pointer to **bool** | Indicates whether data translation is enabled for the field (true) or not (false). Available in API version 49.0 and later. | [optional] 
**DefaultedOnCreate** | Pointer to **bool** | Indicates whether this field is defaulted when created (true) or not (false). If true, then Salesforce implicitly assigns a value for this field when the object is created, even if a value for this field is not passed in on the create() call. For example, in the Opportunity object, the Probability field has this attribute because its value is derived from the Stage field. Similarly, the Owner has this attribute on most objects because its value is derived from the current user (if the Owner field is not specified). | [optional] 
**DefaultValueFormula** | Pointer to **string** | The default value specified for this field if the formula is not used. If no value has been specified, this field is not returned. | [optional] 
**DependentPicklist** | Pointer to **bool** | Indicates whether a picklist is a dependent picklist (true) where available values depend on the chosen values from a controlling field, or not (false). See About Dependent Picklists. | [optional] 
**DeprecatedAndHidden** | Pointer to **bool** | Reserved for future use. | [optional] 
**Digits** | Pointer to **int32** | For fields of type integer. Maximum number of digits. The API returns an error if an integer value exceeds the number of digits. | [optional] 
**DisplayLocationInDecimal** | Pointer to **bool** | Indicates how the geolocation values of a Location custom field appears in the user interface. If true, the geolocation values appear in decimal notation. If false, the geolocation values appear as degrees, minutes, and seconds. | [optional] 
**Encrypted** | Pointer to **bool** | Indicates whether this field is encrypted. This value only appears in the results of a describeSObjects() call when it is true; otherwise, it is omitted from the results. This field is available in API version 31.0 and later. | [optional] 
**ExtraTypeInfo** | Pointer to **string** | If the field is a textarea field type, indicates if the text area is plain text (plaintextarea) or rich text (richtextarea). If the field is a url field type, if this value is imageurl, the URL references an image file. Available on standard fields on standard objects only, for example, Account.photoUrl, Contact.photoUrl, and so on. If the field is a reference field type, indicates the type of external object relationship. Available on external objects only. null—lookup relationship externallookup—external lookup relationship indirectlookup—indirect lookup relationship filterable boolean Indicates whether the field is filterable (true) or not (false). If true, then this field can be specified in the WHERE clause of a query string in a query() call. filteredLookupInfo FilteredLookupInfo If the field is a reference field type with a lookup filter, filteredLookupInfo contains the lookup filter information for the field. If there is no lookup filter, or the filter is inactive, this field is null. This field is available in API version 31.0 and later. formula string The formula specified for this field. If no formula is specified for this field, it is not returned. | [optional] 
**Groupable** | Pointer to **bool** | Indicates whether the field can be included in the GROUP BY clause of a SOQL query (true) or not (false). See GROUP BY in the Salesforce SOQL and SOSL Reference Guide. Available in API version 18.0 and later. | [optional] 
**HighScaleNumber** | Pointer to **bool** | Indicates whether the field stores numbers to 8 decimal places regardless of what’s specified in the field details (true) or not (false). Used to handle currencies for products that cost fractions of a cent, in large quantities. If high-scale unit pricing isn’t enabled in your organization, this field isn’t returned. Available in API version 33.0 and later. | [optional] 
**HtmlFormatted** | Pointer to **bool** | Indicates whether a field such as a hyperlink custom formula field has been formatted for HTML and should be encoded for display in HTML (true) or not (false). Also indicates whether a field is a custom formula field that has an IMAGE text function. | [optional] 
**IdLookup** | Pointer to **bool** | Indicates whether the field can be used to specify a record in an upsert() call (true) or not (false). | [optional] 
**InlineHelpText** | Pointer to **string** | The text that displays in the field-level help hover text for this field. | [optional] 
**Label** | Pointer to **string** | Text label that is displayed next to the field in the Salesforce user interface. This label can be localized. | [optional] 
**Length** | Pointer to **int32** | Returns the maximum size of the field in Unicode characters (not bytes) or 255, whichever is less. The maximum value returned by the getLength() property is 255. Available in API version 49.0 and later. | [optional] 
**Mask** | Pointer to **string** | Reserved for future use. | [optional] 
**MaskType** | Pointer to **string** | Reserved for future use. | [optional] 
**Name** | Pointer to **string** | Field name used in API calls, such as create(), delete(), and query(). | [optional] 
**NameField** | Pointer to **bool** | Indicates whether this field is a name field (true) or not (false). Used to identify the name field for standard objects (such as AccountName for an Account object) and custom objects. Limited to one per object, except where FirstName and LastName fields are used (such as in the Contact object). If a compound name is present, for example the Name field on a person account, nameField is set to true for that record. If no compound name is present, FirstName and LastName have this field set to true. | [optional] 
**NamePointing** | Pointer to **bool** | Indicates whether the field&#39;s value is the Name of the parent of this object (true) or not (false). Used for objects whose parents may be more than one type of object, for example a task may have an account or a contact as a parent. | [optional] 
**Nillable** | Pointer to **bool** | Indicates whether the field is nillable (true) or not (false). A nillable field can have empty content. A non-nillable field must have a value in order for the object to be created or saved. | [optional] 
**Permissionable** | Pointer to **bool** | Indicates whether FieldPermissions can be specified for the field (true) or not (false). | [optional] 
**PicklistValues** | Pointer to [**[]PicklistEntry**](PicklistEntry.md) | Provides the list of valid values for the picklist. Specified only if restrictedPicklist is true. | [optional] 
**PolymorphicForeignKey** | Pointer to **bool** | Indicates whether the foreign key includes multiple entity types (true) or not (false). | [optional] 
**Precision** | Pointer to **int32** | For fields of type double. Maximum number of digits that can be stored, including all numbers to the left and to the right of the decimal point (but excluding the decimal point character). | [optional] 
**RelationshipName** | Pointer to **string** | The name of the relationship, if this is a master-detail relationship field. | [optional] 
**RelationshipOrder** | Pointer to **int32** | The type of relationship for a master-detail relationship field. Valid values are: &#x60;0&#x60; if the field is the primary relationship &#x60;1&#x60; if the field is the secondary relationship | [optional] 
**ReferenceTargetField** | Pointer to **string** | Applies only to indirect lookup relationships on external objects. Name of the custom field on the parent standard or custom object whose values are matched against the values of the child external object&#39;s indirect lookup relationship field. This matching is done to determine which records are related to each other. This field is available in API version 32.0 and later. | [optional] 
**ReferenceTo** | Pointer to **[]string** | For fields that refer to other objects, this array indicates the object types of the referenced objects. | [optional] 
**RestrictedPicklist** | Pointer to **bool** | Indicates whether the field is a restricted picklist (true) or not (false). | [optional] 
**Scale** | Pointer to **int32** | For fields of type double. Number of digits to the right of the decimal point. The API silently truncates any extra digits to the right of the decimal point, but it returns a fault response if the number has too many digits to the left of the decimal point. | [optional] 
**SearchPrefilterable** | Pointer to **bool** | Indicates whether a foreign key can be included in prefiltering (true) or not (false) when used in a SOSL WHERE clause. Prefiltering means to filter by a specific field value before executing the full search query. Available in API version 40.0 and later. | [optional] 
**SoapType** | Pointer to **string** |  | [optional] 
**Sortable** | Pointer to **bool** | Indicates whether a query can sort on this field (true) or not (false). | [optional] 
**Type** | Pointer to [**FieldType**](FieldType.md) |  | [optional] 
**Unique** | Pointer to **bool** | Indicates whether the value must be unique true) or not false). | [optional] 
**Updateable** | Pointer to **bool** | Indicates one of the following: Whether the field is updateable, (true) or not (false). If true, then this field value can be set in an update() call. If the field is in a master-detail relationship on a custom object, indicates whether the child records can be reparented to different parent records (true), false otherwise. | [optional] 
**WriteRequiresMasterRead** | Pointer to **bool** | This field only applies to master-detail relationships. Indicates whether a user requires read sharing access (true) or write sharing access (false) to the parent record to insert, update, and delete a child record. In both cases, a user also needs Create, Edit, and Delete object permissions for the child object. | [optional] 

## Methods

### NewField

`func NewField() *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonumber

`func (o *Field) GetAutonumber() bool`

GetAutonumber returns the Autonumber field if non-nil, zero value otherwise.

### GetAutonumberOk

`func (o *Field) GetAutonumberOk() (*bool, bool)`

GetAutonumberOk returns a tuple with the Autonumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonumber

`func (o *Field) SetAutonumber(v bool)`

SetAutonumber sets Autonumber field to given value.

### HasAutonumber

`func (o *Field) HasAutonumber() bool`

HasAutonumber returns a boolean if a field has been set.

### GetByteLength

`func (o *Field) GetByteLength() int32`

GetByteLength returns the ByteLength field if non-nil, zero value otherwise.

### GetByteLengthOk

`func (o *Field) GetByteLengthOk() (*int32, bool)`

GetByteLengthOk returns a tuple with the ByteLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteLength

`func (o *Field) SetByteLength(v int32)`

SetByteLength sets ByteLength field to given value.

### HasByteLength

`func (o *Field) HasByteLength() bool`

HasByteLength returns a boolean if a field has been set.

### GetCalculated

`func (o *Field) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *Field) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *Field) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *Field) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *Field) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *Field) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *Field) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *Field) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetControllerName

`func (o *Field) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *Field) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *Field) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *Field) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetCreateable

`func (o *Field) GetCreateable() bool`

GetCreateable returns the Createable field if non-nil, zero value otherwise.

### GetCreateableOk

`func (o *Field) GetCreateableOk() (*bool, bool)`

GetCreateableOk returns a tuple with the Createable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateable

`func (o *Field) SetCreateable(v bool)`

SetCreateable sets Createable field to given value.

### HasCreateable

`func (o *Field) HasCreateable() bool`

HasCreateable returns a boolean if a field has been set.

### GetCustom

`func (o *Field) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *Field) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *Field) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *Field) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDataTranslationEnabled

`func (o *Field) GetDataTranslationEnabled() bool`

GetDataTranslationEnabled returns the DataTranslationEnabled field if non-nil, zero value otherwise.

### GetDataTranslationEnabledOk

`func (o *Field) GetDataTranslationEnabledOk() (*bool, bool)`

GetDataTranslationEnabledOk returns a tuple with the DataTranslationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTranslationEnabled

`func (o *Field) SetDataTranslationEnabled(v bool)`

SetDataTranslationEnabled sets DataTranslationEnabled field to given value.

### HasDataTranslationEnabled

`func (o *Field) HasDataTranslationEnabled() bool`

HasDataTranslationEnabled returns a boolean if a field has been set.

### GetDefaultedOnCreate

`func (o *Field) GetDefaultedOnCreate() bool`

GetDefaultedOnCreate returns the DefaultedOnCreate field if non-nil, zero value otherwise.

### GetDefaultedOnCreateOk

`func (o *Field) GetDefaultedOnCreateOk() (*bool, bool)`

GetDefaultedOnCreateOk returns a tuple with the DefaultedOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultedOnCreate

`func (o *Field) SetDefaultedOnCreate(v bool)`

SetDefaultedOnCreate sets DefaultedOnCreate field to given value.

### HasDefaultedOnCreate

`func (o *Field) HasDefaultedOnCreate() bool`

HasDefaultedOnCreate returns a boolean if a field has been set.

### GetDefaultValueFormula

`func (o *Field) GetDefaultValueFormula() string`

GetDefaultValueFormula returns the DefaultValueFormula field if non-nil, zero value otherwise.

### GetDefaultValueFormulaOk

`func (o *Field) GetDefaultValueFormulaOk() (*string, bool)`

GetDefaultValueFormulaOk returns a tuple with the DefaultValueFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueFormula

`func (o *Field) SetDefaultValueFormula(v string)`

SetDefaultValueFormula sets DefaultValueFormula field to given value.

### HasDefaultValueFormula

`func (o *Field) HasDefaultValueFormula() bool`

HasDefaultValueFormula returns a boolean if a field has been set.

### GetDependentPicklist

`func (o *Field) GetDependentPicklist() bool`

GetDependentPicklist returns the DependentPicklist field if non-nil, zero value otherwise.

### GetDependentPicklistOk

`func (o *Field) GetDependentPicklistOk() (*bool, bool)`

GetDependentPicklistOk returns a tuple with the DependentPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentPicklist

`func (o *Field) SetDependentPicklist(v bool)`

SetDependentPicklist sets DependentPicklist field to given value.

### HasDependentPicklist

`func (o *Field) HasDependentPicklist() bool`

HasDependentPicklist returns a boolean if a field has been set.

### GetDeprecatedAndHidden

`func (o *Field) GetDeprecatedAndHidden() bool`

GetDeprecatedAndHidden returns the DeprecatedAndHidden field if non-nil, zero value otherwise.

### GetDeprecatedAndHiddenOk

`func (o *Field) GetDeprecatedAndHiddenOk() (*bool, bool)`

GetDeprecatedAndHiddenOk returns a tuple with the DeprecatedAndHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedAndHidden

`func (o *Field) SetDeprecatedAndHidden(v bool)`

SetDeprecatedAndHidden sets DeprecatedAndHidden field to given value.

### HasDeprecatedAndHidden

`func (o *Field) HasDeprecatedAndHidden() bool`

HasDeprecatedAndHidden returns a boolean if a field has been set.

### GetDigits

`func (o *Field) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *Field) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *Field) SetDigits(v int32)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *Field) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetDisplayLocationInDecimal

`func (o *Field) GetDisplayLocationInDecimal() bool`

GetDisplayLocationInDecimal returns the DisplayLocationInDecimal field if non-nil, zero value otherwise.

### GetDisplayLocationInDecimalOk

`func (o *Field) GetDisplayLocationInDecimalOk() (*bool, bool)`

GetDisplayLocationInDecimalOk returns a tuple with the DisplayLocationInDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocationInDecimal

`func (o *Field) SetDisplayLocationInDecimal(v bool)`

SetDisplayLocationInDecimal sets DisplayLocationInDecimal field to given value.

### HasDisplayLocationInDecimal

`func (o *Field) HasDisplayLocationInDecimal() bool`

HasDisplayLocationInDecimal returns a boolean if a field has been set.

### GetEncrypted

`func (o *Field) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *Field) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *Field) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *Field) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetExtraTypeInfo

`func (o *Field) GetExtraTypeInfo() string`

GetExtraTypeInfo returns the ExtraTypeInfo field if non-nil, zero value otherwise.

### GetExtraTypeInfoOk

`func (o *Field) GetExtraTypeInfoOk() (*string, bool)`

GetExtraTypeInfoOk returns a tuple with the ExtraTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraTypeInfo

`func (o *Field) SetExtraTypeInfo(v string)`

SetExtraTypeInfo sets ExtraTypeInfo field to given value.

### HasExtraTypeInfo

`func (o *Field) HasExtraTypeInfo() bool`

HasExtraTypeInfo returns a boolean if a field has been set.

### GetGroupable

`func (o *Field) GetGroupable() bool`

GetGroupable returns the Groupable field if non-nil, zero value otherwise.

### GetGroupableOk

`func (o *Field) GetGroupableOk() (*bool, bool)`

GetGroupableOk returns a tuple with the Groupable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupable

`func (o *Field) SetGroupable(v bool)`

SetGroupable sets Groupable field to given value.

### HasGroupable

`func (o *Field) HasGroupable() bool`

HasGroupable returns a boolean if a field has been set.

### GetHighScaleNumber

`func (o *Field) GetHighScaleNumber() bool`

GetHighScaleNumber returns the HighScaleNumber field if non-nil, zero value otherwise.

### GetHighScaleNumberOk

`func (o *Field) GetHighScaleNumberOk() (*bool, bool)`

GetHighScaleNumberOk returns a tuple with the HighScaleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighScaleNumber

`func (o *Field) SetHighScaleNumber(v bool)`

SetHighScaleNumber sets HighScaleNumber field to given value.

### HasHighScaleNumber

`func (o *Field) HasHighScaleNumber() bool`

HasHighScaleNumber returns a boolean if a field has been set.

### GetHtmlFormatted

`func (o *Field) GetHtmlFormatted() bool`

GetHtmlFormatted returns the HtmlFormatted field if non-nil, zero value otherwise.

### GetHtmlFormattedOk

`func (o *Field) GetHtmlFormattedOk() (*bool, bool)`

GetHtmlFormattedOk returns a tuple with the HtmlFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlFormatted

`func (o *Field) SetHtmlFormatted(v bool)`

SetHtmlFormatted sets HtmlFormatted field to given value.

### HasHtmlFormatted

`func (o *Field) HasHtmlFormatted() bool`

HasHtmlFormatted returns a boolean if a field has been set.

### GetIdLookup

`func (o *Field) GetIdLookup() bool`

GetIdLookup returns the IdLookup field if non-nil, zero value otherwise.

### GetIdLookupOk

`func (o *Field) GetIdLookupOk() (*bool, bool)`

GetIdLookupOk returns a tuple with the IdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdLookup

`func (o *Field) SetIdLookup(v bool)`

SetIdLookup sets IdLookup field to given value.

### HasIdLookup

`func (o *Field) HasIdLookup() bool`

HasIdLookup returns a boolean if a field has been set.

### GetInlineHelpText

`func (o *Field) GetInlineHelpText() string`

GetInlineHelpText returns the InlineHelpText field if non-nil, zero value otherwise.

### GetInlineHelpTextOk

`func (o *Field) GetInlineHelpTextOk() (*string, bool)`

GetInlineHelpTextOk returns a tuple with the InlineHelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHelpText

`func (o *Field) SetInlineHelpText(v string)`

SetInlineHelpText sets InlineHelpText field to given value.

### HasInlineHelpText

`func (o *Field) HasInlineHelpText() bool`

HasInlineHelpText returns a boolean if a field has been set.

### GetLabel

`func (o *Field) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Field) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Field) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Field) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLength

`func (o *Field) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Field) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Field) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Field) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMask

`func (o *Field) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *Field) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *Field) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *Field) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetMaskType

`func (o *Field) GetMaskType() string`

GetMaskType returns the MaskType field if non-nil, zero value otherwise.

### GetMaskTypeOk

`func (o *Field) GetMaskTypeOk() (*string, bool)`

GetMaskTypeOk returns a tuple with the MaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskType

`func (o *Field) SetMaskType(v string)`

SetMaskType sets MaskType field to given value.

### HasMaskType

`func (o *Field) HasMaskType() bool`

HasMaskType returns a boolean if a field has been set.

### GetName

`func (o *Field) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Field) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Field) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Field) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameField

`func (o *Field) GetNameField() bool`

GetNameField returns the NameField field if non-nil, zero value otherwise.

### GetNameFieldOk

`func (o *Field) GetNameFieldOk() (*bool, bool)`

GetNameFieldOk returns a tuple with the NameField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameField

`func (o *Field) SetNameField(v bool)`

SetNameField sets NameField field to given value.

### HasNameField

`func (o *Field) HasNameField() bool`

HasNameField returns a boolean if a field has been set.

### GetNamePointing

`func (o *Field) GetNamePointing() bool`

GetNamePointing returns the NamePointing field if non-nil, zero value otherwise.

### GetNamePointingOk

`func (o *Field) GetNamePointingOk() (*bool, bool)`

GetNamePointingOk returns a tuple with the NamePointing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePointing

`func (o *Field) SetNamePointing(v bool)`

SetNamePointing sets NamePointing field to given value.

### HasNamePointing

`func (o *Field) HasNamePointing() bool`

HasNamePointing returns a boolean if a field has been set.

### GetNillable

`func (o *Field) GetNillable() bool`

GetNillable returns the Nillable field if non-nil, zero value otherwise.

### GetNillableOk

`func (o *Field) GetNillableOk() (*bool, bool)`

GetNillableOk returns a tuple with the Nillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNillable

`func (o *Field) SetNillable(v bool)`

SetNillable sets Nillable field to given value.

### HasNillable

`func (o *Field) HasNillable() bool`

HasNillable returns a boolean if a field has been set.

### GetPermissionable

`func (o *Field) GetPermissionable() bool`

GetPermissionable returns the Permissionable field if non-nil, zero value otherwise.

### GetPermissionableOk

`func (o *Field) GetPermissionableOk() (*bool, bool)`

GetPermissionableOk returns a tuple with the Permissionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionable

`func (o *Field) SetPermissionable(v bool)`

SetPermissionable sets Permissionable field to given value.

### HasPermissionable

`func (o *Field) HasPermissionable() bool`

HasPermissionable returns a boolean if a field has been set.

### GetPicklistValues

`func (o *Field) GetPicklistValues() []PicklistEntry`

GetPicklistValues returns the PicklistValues field if non-nil, zero value otherwise.

### GetPicklistValuesOk

`func (o *Field) GetPicklistValuesOk() (*[]PicklistEntry, bool)`

GetPicklistValuesOk returns a tuple with the PicklistValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicklistValues

`func (o *Field) SetPicklistValues(v []PicklistEntry)`

SetPicklistValues sets PicklistValues field to given value.

### HasPicklistValues

`func (o *Field) HasPicklistValues() bool`

HasPicklistValues returns a boolean if a field has been set.

### GetPolymorphicForeignKey

`func (o *Field) GetPolymorphicForeignKey() bool`

GetPolymorphicForeignKey returns the PolymorphicForeignKey field if non-nil, zero value otherwise.

### GetPolymorphicForeignKeyOk

`func (o *Field) GetPolymorphicForeignKeyOk() (*bool, bool)`

GetPolymorphicForeignKeyOk returns a tuple with the PolymorphicForeignKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolymorphicForeignKey

`func (o *Field) SetPolymorphicForeignKey(v bool)`

SetPolymorphicForeignKey sets PolymorphicForeignKey field to given value.

### HasPolymorphicForeignKey

`func (o *Field) HasPolymorphicForeignKey() bool`

HasPolymorphicForeignKey returns a boolean if a field has been set.

### GetPrecision

`func (o *Field) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *Field) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *Field) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *Field) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetRelationshipName

`func (o *Field) GetRelationshipName() string`

GetRelationshipName returns the RelationshipName field if non-nil, zero value otherwise.

### GetRelationshipNameOk

`func (o *Field) GetRelationshipNameOk() (*string, bool)`

GetRelationshipNameOk returns a tuple with the RelationshipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipName

`func (o *Field) SetRelationshipName(v string)`

SetRelationshipName sets RelationshipName field to given value.

### HasRelationshipName

`func (o *Field) HasRelationshipName() bool`

HasRelationshipName returns a boolean if a field has been set.

### GetRelationshipOrder

`func (o *Field) GetRelationshipOrder() int32`

GetRelationshipOrder returns the RelationshipOrder field if non-nil, zero value otherwise.

### GetRelationshipOrderOk

`func (o *Field) GetRelationshipOrderOk() (*int32, bool)`

GetRelationshipOrderOk returns a tuple with the RelationshipOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipOrder

`func (o *Field) SetRelationshipOrder(v int32)`

SetRelationshipOrder sets RelationshipOrder field to given value.

### HasRelationshipOrder

`func (o *Field) HasRelationshipOrder() bool`

HasRelationshipOrder returns a boolean if a field has been set.

### GetReferenceTargetField

`func (o *Field) GetReferenceTargetField() string`

GetReferenceTargetField returns the ReferenceTargetField field if non-nil, zero value otherwise.

### GetReferenceTargetFieldOk

`func (o *Field) GetReferenceTargetFieldOk() (*string, bool)`

GetReferenceTargetFieldOk returns a tuple with the ReferenceTargetField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTargetField

`func (o *Field) SetReferenceTargetField(v string)`

SetReferenceTargetField sets ReferenceTargetField field to given value.

### HasReferenceTargetField

`func (o *Field) HasReferenceTargetField() bool`

HasReferenceTargetField returns a boolean if a field has been set.

### GetReferenceTo

`func (o *Field) GetReferenceTo() []string`

GetReferenceTo returns the ReferenceTo field if non-nil, zero value otherwise.

### GetReferenceToOk

`func (o *Field) GetReferenceToOk() (*[]string, bool)`

GetReferenceToOk returns a tuple with the ReferenceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTo

`func (o *Field) SetReferenceTo(v []string)`

SetReferenceTo sets ReferenceTo field to given value.

### HasReferenceTo

`func (o *Field) HasReferenceTo() bool`

HasReferenceTo returns a boolean if a field has been set.

### GetRestrictedPicklist

`func (o *Field) GetRestrictedPicklist() bool`

GetRestrictedPicklist returns the RestrictedPicklist field if non-nil, zero value otherwise.

### GetRestrictedPicklistOk

`func (o *Field) GetRestrictedPicklistOk() (*bool, bool)`

GetRestrictedPicklistOk returns a tuple with the RestrictedPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPicklist

`func (o *Field) SetRestrictedPicklist(v bool)`

SetRestrictedPicklist sets RestrictedPicklist field to given value.

### HasRestrictedPicklist

`func (o *Field) HasRestrictedPicklist() bool`

HasRestrictedPicklist returns a boolean if a field has been set.

### GetScale

`func (o *Field) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *Field) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *Field) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *Field) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetSearchPrefilterable

`func (o *Field) GetSearchPrefilterable() bool`

GetSearchPrefilterable returns the SearchPrefilterable field if non-nil, zero value otherwise.

### GetSearchPrefilterableOk

`func (o *Field) GetSearchPrefilterableOk() (*bool, bool)`

GetSearchPrefilterableOk returns a tuple with the SearchPrefilterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPrefilterable

`func (o *Field) SetSearchPrefilterable(v bool)`

SetSearchPrefilterable sets SearchPrefilterable field to given value.

### HasSearchPrefilterable

`func (o *Field) HasSearchPrefilterable() bool`

HasSearchPrefilterable returns a boolean if a field has been set.

### GetSoapType

`func (o *Field) GetSoapType() string`

GetSoapType returns the SoapType field if non-nil, zero value otherwise.

### GetSoapTypeOk

`func (o *Field) GetSoapTypeOk() (*string, bool)`

GetSoapTypeOk returns a tuple with the SoapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoapType

`func (o *Field) SetSoapType(v string)`

SetSoapType sets SoapType field to given value.

### HasSoapType

`func (o *Field) HasSoapType() bool`

HasSoapType returns a boolean if a field has been set.

### GetSortable

`func (o *Field) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *Field) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *Field) SetSortable(v bool)`

SetSortable sets Sortable field to given value.

### HasSortable

`func (o *Field) HasSortable() bool`

HasSortable returns a boolean if a field has been set.

### GetType

`func (o *Field) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Field) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Field) SetType(v FieldType)`

SetType sets Type field to given value.

### HasType

`func (o *Field) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnique

`func (o *Field) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *Field) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *Field) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *Field) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetUpdateable

`func (o *Field) GetUpdateable() bool`

GetUpdateable returns the Updateable field if non-nil, zero value otherwise.

### GetUpdateableOk

`func (o *Field) GetUpdateableOk() (*bool, bool)`

GetUpdateableOk returns a tuple with the Updateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateable

`func (o *Field) SetUpdateable(v bool)`

SetUpdateable sets Updateable field to given value.

### HasUpdateable

`func (o *Field) HasUpdateable() bool`

HasUpdateable returns a boolean if a field has been set.

### GetWriteRequiresMasterRead

`func (o *Field) GetWriteRequiresMasterRead() bool`

GetWriteRequiresMasterRead returns the WriteRequiresMasterRead field if non-nil, zero value otherwise.

### GetWriteRequiresMasterReadOk

`func (o *Field) GetWriteRequiresMasterReadOk() (*bool, bool)`

GetWriteRequiresMasterReadOk returns a tuple with the WriteRequiresMasterRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteRequiresMasterRead

`func (o *Field) SetWriteRequiresMasterRead(v bool)`

SetWriteRequiresMasterRead sets WriteRequiresMasterRead field to given value.

### HasWriteRequiresMasterRead

`func (o *Field) HasWriteRequiresMasterRead() bool`

HasWriteRequiresMasterRead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


