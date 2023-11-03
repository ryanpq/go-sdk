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

// WriteRequest struct for WriteRequest
type WriteRequest struct {
	Writes               *WriteRequestTupleKeys `json:"writes,omitempty"yaml:"writes,omitempty"`
	Deletes              *WriteRequestTupleKeys `json:"deletes,omitempty"yaml:"deletes,omitempty"`
	AuthorizationModelId *string                `json:"authorization_model_id,omitempty"yaml:"authorization_model_id,omitempty"`
}

// NewWriteRequest instantiates a new WriteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteRequest() *WriteRequest {
	this := WriteRequest{}
	return &this
}

// NewWriteRequestWithDefaults instantiates a new WriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRequestWithDefaults() *WriteRequest {
	this := WriteRequest{}
	return &this
}

// GetWrites returns the Writes field value if set, zero value otherwise.
func (o *WriteRequest) GetWrites() WriteRequestTupleKeys {
	if o == nil || o.Writes == nil {
		var ret WriteRequestTupleKeys
		return ret
	}
	return *o.Writes
}

// GetWritesOk returns a tuple with the Writes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRequest) GetWritesOk() (*WriteRequestTupleKeys, bool) {
	if o == nil || o.Writes == nil {
		return nil, false
	}
	return o.Writes, true
}

// HasWrites returns a boolean if a field has been set.
func (o *WriteRequest) HasWrites() bool {
	if o != nil && o.Writes != nil {
		return true
	}

	return false
}

// SetWrites gets a reference to the given WriteRequestTupleKeys and assigns it to the Writes field.
func (o *WriteRequest) SetWrites(v WriteRequestTupleKeys) {
	o.Writes = &v
}

// GetDeletes returns the Deletes field value if set, zero value otherwise.
func (o *WriteRequest) GetDeletes() WriteRequestTupleKeys {
	if o == nil || o.Deletes == nil {
		var ret WriteRequestTupleKeys
		return ret
	}
	return *o.Deletes
}

// GetDeletesOk returns a tuple with the Deletes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRequest) GetDeletesOk() (*WriteRequestTupleKeys, bool) {
	if o == nil || o.Deletes == nil {
		return nil, false
	}
	return o.Deletes, true
}

// HasDeletes returns a boolean if a field has been set.
func (o *WriteRequest) HasDeletes() bool {
	if o != nil && o.Deletes != nil {
		return true
	}

	return false
}

// SetDeletes gets a reference to the given WriteRequestTupleKeys and assigns it to the Deletes field.
func (o *WriteRequest) SetDeletes(v WriteRequestTupleKeys) {
	o.Deletes = &v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *WriteRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *WriteRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *WriteRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

func (o WriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Writes != nil {
		toSerialize["writes"] = o.Writes
	}
	if o.Deletes != nil {
		toSerialize["deletes"] = o.Deletes
	}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	return json.Marshal(toSerialize)
}

type NullableWriteRequest struct {
	value *WriteRequest
	isSet bool
}

func (v NullableWriteRequest) Get() *WriteRequest {
	return v.value
}

func (v *NullableWriteRequest) Set(val *WriteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteRequest(val *WriteRequest) *NullableWriteRequest {
	return &NullableWriteRequest{value: val, isSet: true}
}

func (v NullableWriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
