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
	"fmt"
)

// InternalErrorCode the model 'InternalErrorCode'
type InternalErrorCode string

// List of InternalErrorCode
const (
	NO_INTERNAL_ERROR   InternalErrorCode = "no_internal_error"
	INTERNAL_ERROR      InternalErrorCode = "internal_error"
	CANCELLED           InternalErrorCode = "cancelled"
	DEADLINE_EXCEEDED   InternalErrorCode = "deadline_exceeded"
	ALREADY_EXISTS      InternalErrorCode = "already_exists"
	RESOURCE_EXHAUSTED  InternalErrorCode = "resource_exhausted"
	FAILED_PRECONDITION InternalErrorCode = "failed_precondition"
	ABORTED             InternalErrorCode = "aborted"
	OUT_OF_RANGE        InternalErrorCode = "out_of_range"
	UNAVAILABLE         InternalErrorCode = "unavailable"
	DATA_LOSS           InternalErrorCode = "data_loss"
)

var allowedInternalErrorCodeEnumValues = []InternalErrorCode{
	"no_internal_error",
	"internal_error",
	"cancelled",
	"deadline_exceeded",
	"already_exists",
	"resource_exhausted",
	"failed_precondition",
	"aborted",
	"out_of_range",
	"unavailable",
	"data_loss",
}

func (v *InternalErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InternalErrorCode(value)
	for _, existing := range allowedInternalErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InternalErrorCode", value)
}

// NewInternalErrorCodeFromValue returns a pointer to a valid InternalErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInternalErrorCodeFromValue(v string) (*InternalErrorCode, error) {
	ev := InternalErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InternalErrorCode: valid values are %v", v, allowedInternalErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InternalErrorCode) IsValid() bool {
	for _, existing := range allowedInternalErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InternalErrorCode value
func (v InternalErrorCode) Ptr() *InternalErrorCode {
	return &v
}

type NullableInternalErrorCode struct {
	value *InternalErrorCode
	isSet bool
}

func (v NullableInternalErrorCode) Get() *InternalErrorCode {
	return v.value
}

func (v *NullableInternalErrorCode) Set(val *InternalErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalErrorCode(val *InternalErrorCode) *NullableInternalErrorCode {
	return &NullableInternalErrorCode{value: val, isSet: true}
}

func (v NullableInternalErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}