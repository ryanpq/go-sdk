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

// Computed struct for Computed
type Computed struct {
	Userset string `json:"userset" yaml:"userset"`
}

// NewComputed instantiates a new Computed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputed(userset string) *Computed {
	this := Computed{}
	this.Userset = userset
	return &this
}

// NewComputedWithDefaults instantiates a new Computed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputedWithDefaults() *Computed {
	this := Computed{}
	return &this
}

// GetUserset returns the Userset field value
func (o *Computed) GetUserset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Userset
}

// GetUsersetOk returns a tuple with the Userset field value
// and a boolean to check if the value has been set.
func (o *Computed) GetUsersetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userset, true
}

// SetUserset sets field value
func (o *Computed) SetUserset(v string) {
	o.Userset = v
}

func (o Computed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userset"] = o.Userset
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableComputed struct {
	value *Computed
	isSet bool
}

func (v NullableComputed) Get() *Computed {
	return v.value
}

func (v *NullableComputed) Set(val *Computed) {
	v.value = val
	v.isSet = true
}

func (v NullableComputed) IsSet() bool {
	return v.isSet
}

func (v *NullableComputed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputed(val *Computed) *NullableComputed {
	return &NullableComputed{value: val, isSet: true}
}

func (v NullableComputed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
