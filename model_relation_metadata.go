/**
 * Go SDK for OpenFGA
 *
 * API version: 1.x
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
)

// RelationMetadata struct for RelationMetadata
type RelationMetadata struct {
	DirectlyRelatedUserTypes *[]RelationReference `json:"directly_related_user_types,omitempty" yaml:"directly_related_user_types,omitempty"`
	Module                   *string              `json:"module,omitempty" yaml:"module,omitempty"`
	SourceInfo               *SourceInfo          `json:"source_info,omitempty" yaml:"source_info,omitempty"`
}

// NewRelationMetadata instantiates a new RelationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationMetadata() *RelationMetadata {
	this := RelationMetadata{}
	return &this
}

// NewRelationMetadataWithDefaults instantiates a new RelationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationMetadataWithDefaults() *RelationMetadata {
	this := RelationMetadata{}
	return &this
}

// GetDirectlyRelatedUserTypes returns the DirectlyRelatedUserTypes field value if set, zero value otherwise.
func (o *RelationMetadata) GetDirectlyRelatedUserTypes() []RelationReference {
	if o == nil || o.DirectlyRelatedUserTypes == nil {
		var ret []RelationReference
		return ret
	}
	return *o.DirectlyRelatedUserTypes
}

// GetDirectlyRelatedUserTypesOk returns a tuple with the DirectlyRelatedUserTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationMetadata) GetDirectlyRelatedUserTypesOk() (*[]RelationReference, bool) {
	if o == nil || o.DirectlyRelatedUserTypes == nil {
		return nil, false
	}
	return o.DirectlyRelatedUserTypes, true
}

// HasDirectlyRelatedUserTypes returns a boolean if a field has been set.
func (o *RelationMetadata) HasDirectlyRelatedUserTypes() bool {
	if o != nil && o.DirectlyRelatedUserTypes != nil {
		return true
	}

	return false
}

// SetDirectlyRelatedUserTypes gets a reference to the given []RelationReference and assigns it to the DirectlyRelatedUserTypes field.
func (o *RelationMetadata) SetDirectlyRelatedUserTypes(v []RelationReference) {
	o.DirectlyRelatedUserTypes = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *RelationMetadata) GetModule() string {
	if o == nil || o.Module == nil {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationMetadata) GetModuleOk() (*string, bool) {
	if o == nil || o.Module == nil {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *RelationMetadata) HasModule() bool {
	if o != nil && o.Module != nil {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *RelationMetadata) SetModule(v string) {
	o.Module = &v
}

// GetSourceInfo returns the SourceInfo field value if set, zero value otherwise.
func (o *RelationMetadata) GetSourceInfo() SourceInfo {
	if o == nil || o.SourceInfo == nil {
		var ret SourceInfo
		return ret
	}
	return *o.SourceInfo
}

// GetSourceInfoOk returns a tuple with the SourceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationMetadata) GetSourceInfoOk() (*SourceInfo, bool) {
	if o == nil || o.SourceInfo == nil {
		return nil, false
	}
	return o.SourceInfo, true
}

// HasSourceInfo returns a boolean if a field has been set.
func (o *RelationMetadata) HasSourceInfo() bool {
	if o != nil && o.SourceInfo != nil {
		return true
	}

	return false
}

// SetSourceInfo gets a reference to the given SourceInfo and assigns it to the SourceInfo field.
func (o *RelationMetadata) SetSourceInfo(v SourceInfo) {
	o.SourceInfo = &v
}

func (o RelationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DirectlyRelatedUserTypes != nil {
		toSerialize["directly_related_user_types"] = o.DirectlyRelatedUserTypes
	}
	if o.Module != nil {
		toSerialize["module"] = o.Module
	}
	if o.SourceInfo != nil {
		toSerialize["source_info"] = o.SourceInfo
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableRelationMetadata struct {
	value *RelationMetadata
	isSet bool
}

func (v NullableRelationMetadata) Get() *RelationMetadata {
	return v.value
}

func (v *NullableRelationMetadata) Set(val *RelationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationMetadata(val *RelationMetadata) *NullableRelationMetadata {
	return &NullableRelationMetadata{value: val, isSet: true}
}

func (v NullableRelationMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
