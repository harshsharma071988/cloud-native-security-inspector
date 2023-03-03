/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: 0.1.0
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RegisterKubernetesClusterRequest Payload to request registering a Kubernetes cluster
type RegisterKubernetesClusterRequest struct {
	// Name of the cluster to be registered
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _RegisterKubernetesClusterRequest RegisterKubernetesClusterRequest

// NewRegisterKubernetesClusterRequest instantiates a new RegisterKubernetesClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterKubernetesClusterRequest(name string) *RegisterKubernetesClusterRequest {
	this := RegisterKubernetesClusterRequest{}
	this.Name = name
	return &this
}

// NewRegisterKubernetesClusterRequestWithDefaults instantiates a new RegisterKubernetesClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterKubernetesClusterRequestWithDefaults() *RegisterKubernetesClusterRequest {
	this := RegisterKubernetesClusterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RegisterKubernetesClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegisterKubernetesClusterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegisterKubernetesClusterRequest) SetName(v string) {
	o.Name = v
}

func (o RegisterKubernetesClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegisterKubernetesClusterRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRegisterKubernetesClusterRequest := _RegisterKubernetesClusterRequest{}

	if err = json.Unmarshal(bytes, &varRegisterKubernetesClusterRequest); err == nil {
		*o = RegisterKubernetesClusterRequest(varRegisterKubernetesClusterRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegisterKubernetesClusterRequest struct {
	value *RegisterKubernetesClusterRequest
	isSet bool
}

func (v NullableRegisterKubernetesClusterRequest) Get() *RegisterKubernetesClusterRequest {
	return v.value
}

func (v *NullableRegisterKubernetesClusterRequest) Set(val *RegisterKubernetesClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterKubernetesClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterKubernetesClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterKubernetesClusterRequest(val *RegisterKubernetesClusterRequest) *NullableRegisterKubernetesClusterRequest {
	return &NullableRegisterKubernetesClusterRequest{value: val, isSet: true}
}

func (v NullableRegisterKubernetesClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterKubernetesClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


