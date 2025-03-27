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

// ExpandRequestTupleKey struct for ExpandRequestTupleKey
type ExpandRequestTupleKey struct {
	Relation string `json:"relation"yaml:"relation"`
	Object   string `json:"object"yaml:"object"`
}

// NewExpandRequestTupleKey instantiates a new ExpandRequestTupleKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandRequestTupleKey(relation string, object string) *ExpandRequestTupleKey {
	this := ExpandRequestTupleKey{}
	this.Relation = relation
	this.Object = object
	return &this
}

// NewExpandRequestTupleKeyWithDefaults instantiates a new ExpandRequestTupleKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandRequestTupleKeyWithDefaults() *ExpandRequestTupleKey {
	this := ExpandRequestTupleKey{}
	return &this
}

// GetRelation returns the Relation field value
func (o *ExpandRequestTupleKey) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *ExpandRequestTupleKey) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *ExpandRequestTupleKey) SetRelation(v string) {
	o.Relation = v
}

// GetObject returns the Object field value
func (o *ExpandRequestTupleKey) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ExpandRequestTupleKey) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ExpandRequestTupleKey) SetObject(v string) {
	o.Object = v
}

func (o ExpandRequestTupleKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relation"] = o.Relation
	toSerialize["object"] = o.Object
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableExpandRequestTupleKey struct {
	value *ExpandRequestTupleKey
	isSet bool
}

func (v NullableExpandRequestTupleKey) Get() *ExpandRequestTupleKey {
	return v.value
}

func (v *NullableExpandRequestTupleKey) Set(val *ExpandRequestTupleKey) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandRequestTupleKey) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandRequestTupleKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandRequestTupleKey(val *ExpandRequestTupleKey) *NullableExpandRequestTupleKey {
	return &NullableExpandRequestTupleKey{value: val, isSet: true}
}

func (v NullableExpandRequestTupleKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandRequestTupleKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
