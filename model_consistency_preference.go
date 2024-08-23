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
	"encoding/json"
	"fmt"
)

// ConsistencyPreference - UNSPECIFIED: Default if not set. Behavior will be the same as MINIMIZE_LATENCY  - MINIMIZE_LATENCY: Minimize latency at the potential expense of lower consistency.  - HIGHER_CONSISTENCY: Prefer higher consistency, at the potential expense of increased latency.
type ConsistencyPreference string

// List of ConsistencyPreference
const (
	CONSISTENCYPREFERENCE_UNSPECIFIED        ConsistencyPreference = "UNSPECIFIED"
	CONSISTENCYPREFERENCE_MINIMIZE_LATENCY   ConsistencyPreference = "MINIMIZE_LATENCY"
	CONSISTENCYPREFERENCE_HIGHER_CONSISTENCY ConsistencyPreference = "HIGHER_CONSISTENCY"
)

var allowedConsistencyPreferenceEnumValues = []ConsistencyPreference{
	"UNSPECIFIED",
	"MINIMIZE_LATENCY",
	"HIGHER_CONSISTENCY",
}

func (v *ConsistencyPreference) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsistencyPreference(value)
	for _, existing := range allowedConsistencyPreferenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsistencyPreference", value)
}

// NewConsistencyPreferenceFromValue returns a pointer to a valid ConsistencyPreference
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsistencyPreferenceFromValue(v string) (*ConsistencyPreference, error) {
	ev := ConsistencyPreference(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsistencyPreference: valid values are %v", v, allowedConsistencyPreferenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsistencyPreference) IsValid() bool {
	for _, existing := range allowedConsistencyPreferenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsistencyPreference value
func (v ConsistencyPreference) Ptr() *ConsistencyPreference {
	return &v
}

type NullableConsistencyPreference struct {
	value *ConsistencyPreference
	isSet bool
}

func (v NullableConsistencyPreference) Get() *ConsistencyPreference {
	return v.value
}

func (v *NullableConsistencyPreference) Set(val *ConsistencyPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistencyPreference) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistencyPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistencyPreference(val *ConsistencyPreference) *NullableConsistencyPreference {
	return &NullableConsistencyPreference{value: val, isSet: true}
}

func (v NullableConsistencyPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistencyPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}