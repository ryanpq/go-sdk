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

// ListObjectsResponse struct for ListObjectsResponse
type ListObjectsResponse struct {
	Objects []string `json:"objects"yaml:"objects"`
}

// NewListObjectsResponse instantiates a new ListObjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectsResponse(objects []string) *ListObjectsResponse {
	this := ListObjectsResponse{}
	this.Objects = objects
	return &this
}

// NewListObjectsResponseWithDefaults instantiates a new ListObjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectsResponseWithDefaults() *ListObjectsResponse {
	this := ListObjectsResponse{}
	return &this
}

// GetObjects returns the Objects field value
func (o *ListObjectsResponse) GetObjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *ListObjectsResponse) GetObjectsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *ListObjectsResponse) SetObjects(v []string) {
	o.Objects = v
}

func (o ListObjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objects"] = o.Objects
	return json.Marshal(toSerialize)
}

type NullableListObjectsResponse struct {
	value *ListObjectsResponse
	isSet bool
}

func (v NullableListObjectsResponse) Get() *ListObjectsResponse {
	return v.value
}

func (v *NullableListObjectsResponse) Set(val *ListObjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectsResponse(val *ListObjectsResponse) *NullableListObjectsResponse {
	return &NullableListObjectsResponse{value: val, isSet: true}
}

func (v NullableListObjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
