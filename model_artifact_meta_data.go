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
	"time"
)

// ArtifactMetaData 
type ArtifactMetaData struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedBy string `json:"createdBy"`
	CreatedOn time.Time `json:"createdOn"`
	ModifiedBy string `json:"modifiedBy"`
	ModifiedOn time.Time `json:"modifiedOn"`
	// The ID of a single artifact.
	Id string `json:"id"`
	// 
	Version string `json:"version"`
	Type ArtifactType `json:"type"`
	// 
	GlobalId int64 `json:"globalId"`
	State ArtifactState `json:"state"`
	// 
	Labels []string `json:"labels,omitempty"`
	// User-defined name-value pairs. Name and value must be strings.
	Properties *map[string]string `json:"properties,omitempty"`
	// An ID of a single artifact group.
	GroupId *string `json:"groupId,omitempty"`
	// 
	ContentId int64 `json:"contentId"`
	// 
	References []ArtifactReference `json:"references,omitempty"`
}

// NewArtifactMetaData instantiates a new ArtifactMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactMetaData(createdBy string, createdOn time.Time, modifiedBy string, modifiedOn time.Time, id string, version string, type_ ArtifactType, globalId int64, state ArtifactState, contentId int64) *ArtifactMetaData {
	this := ArtifactMetaData{}
	this.CreatedBy = createdBy
	this.CreatedOn = createdOn
	this.ModifiedBy = modifiedBy
	this.ModifiedOn = modifiedOn
	this.Id = id
	this.Version = version
	this.Type = type_
	this.GlobalId = globalId
	this.State = state
	this.ContentId = contentId
	return &this
}

// NewArtifactMetaDataWithDefaults instantiates a new ArtifactMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactMetaDataWithDefaults() *ArtifactMetaData {
	this := ArtifactMetaData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArtifactMetaData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArtifactMetaData) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArtifactMetaData) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArtifactMetaData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArtifactMetaData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArtifactMetaData) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ArtifactMetaData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ArtifactMetaData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *ArtifactMetaData) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *ArtifactMetaData) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *ArtifactMetaData) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *ArtifactMetaData) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *ArtifactMetaData) GetModifiedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *ArtifactMetaData) SetModifiedOn(v time.Time) {
	o.ModifiedOn = v
}

// GetId returns the Id field value
func (o *ArtifactMetaData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArtifactMetaData) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *ArtifactMetaData) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ArtifactMetaData) SetVersion(v string) {
	o.Version = v
}

// GetType returns the Type field value
func (o *ArtifactMetaData) GetType() ArtifactType {
	if o == nil {
		var ret ArtifactType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetTypeOk() (*ArtifactType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArtifactMetaData) SetType(v ArtifactType) {
	o.Type = v
}

// GetGlobalId returns the GlobalId field value
func (o *ArtifactMetaData) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *ArtifactMetaData) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetState returns the State field value
func (o *ArtifactMetaData) GetState() ArtifactState {
	if o == nil {
		var ret ArtifactState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetStateOk() (*ArtifactState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ArtifactMetaData) SetState(v ArtifactState) {
	o.State = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ArtifactMetaData) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ArtifactMetaData) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ArtifactMetaData) SetLabels(v []string) {
	o.Labels = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ArtifactMetaData) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ArtifactMetaData) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *ArtifactMetaData) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ArtifactMetaData) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ArtifactMetaData) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ArtifactMetaData) SetGroupId(v string) {
	o.GroupId = &v
}

// GetContentId returns the ContentId field value
func (o *ArtifactMetaData) GetContentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetContentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *ArtifactMetaData) SetContentId(v int64) {
	o.ContentId = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ArtifactMetaData) GetReferences() []ArtifactReference {
	if o == nil || o.References == nil {
		var ret []ArtifactReference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetaData) GetReferencesOk() ([]ArtifactReference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ArtifactMetaData) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []ArtifactReference and assigns it to the References field.
func (o *ArtifactMetaData) SetReferences(v []ArtifactReference) {
	o.References = v
}

func (o ArtifactMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if true {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["globalId"] = o.GlobalId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["contentId"] = o.ContentId
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactMetaData struct {
	value *ArtifactMetaData
	isSet bool
}

func (v NullableArtifactMetaData) Get() *ArtifactMetaData {
	return v.value
}

func (v *NullableArtifactMetaData) Set(val *ArtifactMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactMetaData(val *ArtifactMetaData) *NullableArtifactMetaData {
	return &NullableArtifactMetaData{value: val, isSet: true}
}

func (v NullableArtifactMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


