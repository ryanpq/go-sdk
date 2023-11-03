/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// UsersetTreeTupleToUserset struct for UsersetTreeTupleToUserset
type UsersetTreeTupleToUserset struct {
	Tupleset string     `json:"tupleset"yaml:"tupleset"`
	Computed []Computed `json:"computed"yaml:"computed"`
}

// NewUsersetTreeTupleToUserset instantiates a new UsersetTreeTupleToUserset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersetTreeTupleToUserset(tupleset string, computed []Computed) *UsersetTreeTupleToUserset {
	this := UsersetTreeTupleToUserset{}
	this.Tupleset = tupleset
	this.Computed = computed
	return &this
}

// NewUsersetTreeTupleToUsersetWithDefaults instantiates a new UsersetTreeTupleToUserset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersetTreeTupleToUsersetWithDefaults() *UsersetTreeTupleToUserset {
	this := UsersetTreeTupleToUserset{}
	return &this
}

// GetTupleset returns the Tupleset field value
func (o *UsersetTreeTupleToUserset) GetTupleset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tupleset
}

// GetTuplesetOk returns a tuple with the Tupleset field value
// and a boolean to check if the value has been set.
func (o *UsersetTreeTupleToUserset) GetTuplesetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tupleset, true
}

// SetTupleset sets field value
func (o *UsersetTreeTupleToUserset) SetTupleset(v string) {
	o.Tupleset = v
}

// GetComputed returns the Computed field value
func (o *UsersetTreeTupleToUserset) GetComputed() []Computed {
	if o == nil {
		var ret []Computed
		return ret
	}

	return o.Computed
}

// GetComputedOk returns a tuple with the Computed field value
// and a boolean to check if the value has been set.
func (o *UsersetTreeTupleToUserset) GetComputedOk() (*[]Computed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Computed, true
}

// SetComputed sets field value
func (o *UsersetTreeTupleToUserset) SetComputed(v []Computed) {
	o.Computed = v
}

func (o UsersetTreeTupleToUserset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tupleset"] = o.Tupleset
	toSerialize["computed"] = o.Computed
	return json.Marshal(toSerialize)
}

type NullableUsersetTreeTupleToUserset struct {
	value *UsersetTreeTupleToUserset
	isSet bool
}

func (v NullableUsersetTreeTupleToUserset) Get() *UsersetTreeTupleToUserset {
	return v.value
}

func (v *NullableUsersetTreeTupleToUserset) Set(val *UsersetTreeTupleToUserset) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersetTreeTupleToUserset) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersetTreeTupleToUserset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersetTreeTupleToUserset(val *UsersetTreeTupleToUserset) *NullableUsersetTreeTupleToUserset {
	return &NullableUsersetTreeTupleToUserset{value: val, isSet: true}
}

func (v NullableUsersetTreeTupleToUserset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersetTreeTupleToUserset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
