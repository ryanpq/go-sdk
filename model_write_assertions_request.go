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

// WriteAssertionsRequest struct for WriteAssertionsRequest
type WriteAssertionsRequest struct {
	Assertions []Assertion `json:"assertions"yaml:"assertions"`
}

// NewWriteAssertionsRequest instantiates a new WriteAssertionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteAssertionsRequest(assertions []Assertion) *WriteAssertionsRequest {
	this := WriteAssertionsRequest{}
	this.Assertions = assertions
	return &this
}

// NewWriteAssertionsRequestWithDefaults instantiates a new WriteAssertionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteAssertionsRequestWithDefaults() *WriteAssertionsRequest {
	this := WriteAssertionsRequest{}
	return &this
}

// GetAssertions returns the Assertions field value
func (o *WriteAssertionsRequest) GetAssertions() []Assertion {
	if o == nil {
		var ret []Assertion
		return ret
	}

	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value
// and a boolean to check if the value has been set.
func (o *WriteAssertionsRequest) GetAssertionsOk() (*[]Assertion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// SetAssertions sets field value
func (o *WriteAssertionsRequest) SetAssertions(v []Assertion) {
	o.Assertions = v
}

func (o WriteAssertionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assertions"] = o.Assertions
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableWriteAssertionsRequest struct {
	value *WriteAssertionsRequest
	isSet bool
}

func (v NullableWriteAssertionsRequest) Get() *WriteAssertionsRequest {
	return v.value
}

func (v *NullableWriteAssertionsRequest) Set(val *WriteAssertionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteAssertionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteAssertionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteAssertionsRequest(val *WriteAssertionsRequest) *NullableWriteAssertionsRequest {
	return &NullableWriteAssertionsRequest{value: val, isSet: true}
}

func (v NullableWriteAssertionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteAssertionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
