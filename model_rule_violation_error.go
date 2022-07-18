/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.5.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RuleViolationError All error responses, whether `4xx` or `5xx` will include one of these as the response body.
type RuleViolationError struct {
	// List of rule violation causes.
	Causes []RuleViolationCause `json:"causes"`
	// The short error message.
	Message *string `json:"message,omitempty"`
	// The server-side error code.
	ErrorCode *int32 `json:"error_code,omitempty"`
	// Full details about the error.  This might contain a server stack trace, for example.
	Detail *string `json:"detail,omitempty"`
	// The error name - typically the classname of the exception thrown by the server.
	Name *string `json:"name,omitempty"`
}

// NewRuleViolationError instantiates a new RuleViolationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleViolationError(causes []RuleViolationCause) *RuleViolationError {
	this := RuleViolationError{}
	this.Causes = causes
	return &this
}

// NewRuleViolationErrorWithDefaults instantiates a new RuleViolationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleViolationErrorWithDefaults() *RuleViolationError {
	this := RuleViolationError{}
	return &this
}

// GetCauses returns the Causes field value
func (o *RuleViolationError) GetCauses() []RuleViolationCause {
	if o == nil {
		var ret []RuleViolationCause
		return ret
	}

	return o.Causes
}

// GetCausesOk returns a tuple with the Causes field value
// and a boolean to check if the value has been set.
func (o *RuleViolationError) GetCausesOk() ([]RuleViolationCause, bool) {
	if o == nil {
		return nil, false
	}
	return o.Causes, true
}

// SetCauses sets field value
func (o *RuleViolationError) SetCauses(v []RuleViolationCause) {
	o.Causes = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RuleViolationError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleViolationError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RuleViolationError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RuleViolationError) SetMessage(v string) {
	o.Message = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RuleViolationError) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleViolationError) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RuleViolationError) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *RuleViolationError) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *RuleViolationError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleViolationError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *RuleViolationError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *RuleViolationError) SetDetail(v string) {
	o.Detail = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleViolationError) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleViolationError) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleViolationError) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleViolationError) SetName(v string) {
	o.Name = &v
}

func (o RuleViolationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["causes"] = o.Causes
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRuleViolationError struct {
	value *RuleViolationError
	isSet bool
}

func (v NullableRuleViolationError) Get() *RuleViolationError {
	return v.value
}

func (v *NullableRuleViolationError) Set(val *RuleViolationError) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleViolationError) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleViolationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleViolationError(val *RuleViolationError) *NullableRuleViolationError {
	return &NullableRuleViolationError{value: val, isSet: true}
}

func (v NullableRuleViolationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleViolationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


