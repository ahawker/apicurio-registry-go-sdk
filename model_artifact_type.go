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
	"fmt"
)

// ArtifactType 
type ArtifactType string

// List of ArtifactType
const (
	AVRO ArtifactType = "AVRO"
	PROTOBUF ArtifactType = "PROTOBUF"
	JSON ArtifactType = "JSON"
	OPENAPI ArtifactType = "OPENAPI"
	ASYNCAPI ArtifactType = "ASYNCAPI"
	GRAPHQL ArtifactType = "GRAPHQL"
	KCONNECT ArtifactType = "KCONNECT"
	WSDL ArtifactType = "WSDL"
	XSD ArtifactType = "XSD"
	XML ArtifactType = "XML"
)

// All allowed values of ArtifactType enum
var AllowedArtifactTypeEnumValues = []ArtifactType{
	"AVRO",
	"PROTOBUF",
	"JSON",
	"OPENAPI",
	"ASYNCAPI",
	"GRAPHQL",
	"KCONNECT",
	"WSDL",
	"XSD",
	"XML",
}

func (v *ArtifactType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ArtifactType(value)
	for _, existing := range AllowedArtifactTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ArtifactType", value)
}

// NewArtifactTypeFromValue returns a pointer to a valid ArtifactType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewArtifactTypeFromValue(v string) (*ArtifactType, error) {
	ev := ArtifactType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ArtifactType: valid values are %v", v, AllowedArtifactTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ArtifactType) IsValid() bool {
	for _, existing := range AllowedArtifactTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArtifactType value
func (v ArtifactType) Ptr() *ArtifactType {
	return &v
}

type NullableArtifactType struct {
	value *ArtifactType
	isSet bool
}

func (v NullableArtifactType) Get() *ArtifactType {
	return v.value
}

func (v *NullableArtifactType) Set(val *ArtifactType) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactType) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactType(val *ArtifactType) *NullableArtifactType {
	return &NullableArtifactType{value: val, isSet: true}
}

func (v NullableArtifactType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

