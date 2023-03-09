/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: ${project.version}
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ConstraintsViolationsError Error object extension for returning a constraints violation error
type ConstraintsViolationsError struct {
	Violations []ConstraintsViolation `json:"violations"`
	// A URI reference that identifies the problem type  When dereferenced, it provide human-readable documentation for the problem type using HTML
	Type string `json:"type"`
	// A short, human-readable summary of the problem type
	Title string `json:"title"`
	// The HTTP status code generated by the origin server for this occurrence of the problem
	Status *int32 `json:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem
	Detail *string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced
	Instance             *string `json:"instance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstraintsViolationsError ConstraintsViolationsError

// NewConstraintsViolationsError instantiates a new ConstraintsViolationsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraintsViolationsError(violations []ConstraintsViolation, type_ string, title string) *ConstraintsViolationsError {
	this := ConstraintsViolationsError{}
	this.Type = type_
	this.Title = title
	return &this
}

// NewConstraintsViolationsErrorWithDefaults instantiates a new ConstraintsViolationsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintsViolationsErrorWithDefaults() *ConstraintsViolationsError {
	this := ConstraintsViolationsError{}
	return &this
}

// GetViolations returns the Violations field value
func (o *ConstraintsViolationsError) GetViolations() []ConstraintsViolation {
	if o == nil {
		var ret []ConstraintsViolation
		return ret
	}

	return o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value
// and a boolean to check if the value has been set.
func (o *ConstraintsViolationsError) GetViolationsOk() ([]ConstraintsViolation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Violations, true
}

// SetViolations sets field value
func (o *ConstraintsViolationsError) SetViolations(v []ConstraintsViolation) {
	o.Violations = v
}

// GetType returns the Type field value
func (o *ConstraintsViolationsError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConstraintsViolationsError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConstraintsViolationsError) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ConstraintsViolationsError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ConstraintsViolationsError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ConstraintsViolationsError) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConstraintsViolationsError) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintsViolationsError) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConstraintsViolationsError) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ConstraintsViolationsError) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ConstraintsViolationsError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintsViolationsError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ConstraintsViolationsError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ConstraintsViolationsError) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ConstraintsViolationsError) GetInstance() string {
	if o == nil || o.Instance == nil {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintsViolationsError) GetInstanceOk() (*string, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ConstraintsViolationsError) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ConstraintsViolationsError) SetInstance(v string) {
	o.Instance = &v
}

func (o ConstraintsViolationsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["violations"] = o.Violations
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConstraintsViolationsError) UnmarshalJSON(bytes []byte) (err error) {
	varConstraintsViolationsError := _ConstraintsViolationsError{}

	if err = json.Unmarshal(bytes, &varConstraintsViolationsError); err == nil {
		*o = ConstraintsViolationsError(varConstraintsViolationsError)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "violations")
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "status")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "instance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstraintsViolationsError struct {
	value *ConstraintsViolationsError
	isSet bool
}

func (v NullableConstraintsViolationsError) Get() *ConstraintsViolationsError {
	return v.value
}

func (v *NullableConstraintsViolationsError) Set(val *ConstraintsViolationsError) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintsViolationsError) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintsViolationsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintsViolationsError(val *ConstraintsViolationsError) *NullableConstraintsViolationsError {
	return &NullableConstraintsViolationsError{value: val, isSet: true}
}

func (v NullableConstraintsViolationsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraintsViolationsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
