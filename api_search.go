/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.5.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"os"
)


// SearchApiService SearchApi service
type SearchApiService service

type SearchApiSearchArtifactsRequest struct {
	ctx context.Context
	ApiService *SearchApiService
	name *string
	offset *int32
	limit *int32
	order *SortOrder
	orderby *SortBy
	labels *[]string
	properties *[]string
	description *string
	group *string
	globalId *int64
	contentId *int64
}

// Filter by artifact name.
func (r SearchApiSearchArtifactsRequest) Name(name string) SearchApiSearchArtifactsRequest {
	r.name = &name
	return r
}

// The number of artifacts to skip before starting to collect the result set.  Defaults to 0.
func (r SearchApiSearchArtifactsRequest) Offset(offset int32) SearchApiSearchArtifactsRequest {
	r.offset = &offset
	return r
}

// The number of artifacts to return.  Defaults to 20.
func (r SearchApiSearchArtifactsRequest) Limit(limit int32) SearchApiSearchArtifactsRequest {
	r.limit = &limit
	return r
}

// Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;).
func (r SearchApiSearchArtifactsRequest) Order(order SortOrder) SearchApiSearchArtifactsRequest {
	r.order = &order
	return r
}

// The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60; 
func (r SearchApiSearchArtifactsRequest) Orderby(orderby SortBy) SearchApiSearchArtifactsRequest {
	r.orderby = &orderby
	return r
}

// Filter by label.  Include one or more label to only return artifacts containing all of the specified labels.
func (r SearchApiSearchArtifactsRequest) Labels(labels []string) SearchApiSearchArtifactsRequest {
	r.labels = &labels
	return r
}

// Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example &#x60;properties&#x3D;foo:bar&#x60; will return only artifacts with a custom property named &#x60;foo&#x60; and value &#x60;bar&#x60;.
func (r SearchApiSearchArtifactsRequest) Properties(properties []string) SearchApiSearchArtifactsRequest {
	r.properties = &properties
	return r
}

// Filter by description.
func (r SearchApiSearchArtifactsRequest) Description(description string) SearchApiSearchArtifactsRequest {
	r.description = &description
	return r
}

// Filter by artifact group.
func (r SearchApiSearchArtifactsRequest) Group(group string) SearchApiSearchArtifactsRequest {
	r.group = &group
	return r
}

// Filter by globalId.
func (r SearchApiSearchArtifactsRequest) GlobalId(globalId int64) SearchApiSearchArtifactsRequest {
	r.globalId = &globalId
	return r
}

// Filter by contentId.
func (r SearchApiSearchArtifactsRequest) ContentId(contentId int64) SearchApiSearchArtifactsRequest {
	r.contentId = &contentId
	return r
}

func (r SearchApiSearchArtifactsRequest) Execute() (*ArtifactSearchResults, *http.Response, error) {
	return r.ApiService.SearchArtifactsExecute(r)
}

/*
SearchArtifacts Search for artifacts

Returns a paginated list of all artifacts that match the provided filter criteria.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SearchApiSearchArtifactsRequest
*/
func (a *SearchApiService) SearchArtifacts(ctx context.Context) SearchApiSearchArtifactsRequest {
	return SearchApiSearchArtifactsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactSearchResults
func (a *SearchApiService) SearchArtifactsExecute(r SearchApiSearchArtifactsRequest) (*ArtifactSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.SearchArtifacts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/artifacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.orderby != nil {
		localVarQueryParams.Add("orderby", parameterToString(*r.orderby, ""))
	}
	if r.labels != nil {
		t := *r.labels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("labels", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("labels", parameterToString(t, "multi"))
		}
	}
	if r.properties != nil {
		t := *r.properties
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("properties", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("properties", parameterToString(t, "multi"))
		}
	}
	if r.description != nil {
		localVarQueryParams.Add("description", parameterToString(*r.description, ""))
	}
	if r.group != nil {
		localVarQueryParams.Add("group", parameterToString(*r.group, ""))
	}
	if r.globalId != nil {
		localVarQueryParams.Add("globalId", parameterToString(*r.globalId, ""))
	}
	if r.contentId != nil {
		localVarQueryParams.Add("contentId", parameterToString(*r.contentId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SearchApiSearchArtifactsByContentRequest struct {
	ctx context.Context
	ApiService *SearchApiService
	body **os.File
	canonical *bool
	artifactType *ArtifactType
	offset *int32
	limit *int32
	order *string
	orderby *string
}

// The content to search for.
func (r SearchApiSearchArtifactsByContentRequest) Body(body *os.File) SearchApiSearchArtifactsByContentRequest {
	r.body = &body
	return r
}

// Parameter that can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the &#x60;artifactType&#x60; query parameter.
func (r SearchApiSearchArtifactsByContentRequest) Canonical(canonical bool) SearchApiSearchArtifactsByContentRequest {
	r.canonical = &canonical
	return r
}

// Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the &#x60;canonical&#x60; query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts.
func (r SearchApiSearchArtifactsByContentRequest) ArtifactType(artifactType ArtifactType) SearchApiSearchArtifactsByContentRequest {
	r.artifactType = &artifactType
	return r
}

// The number of artifacts to skip before starting to collect the result set.  Defaults to 0.
func (r SearchApiSearchArtifactsByContentRequest) Offset(offset int32) SearchApiSearchArtifactsByContentRequest {
	r.offset = &offset
	return r
}

// The number of artifacts to return.  Defaults to 20.
func (r SearchApiSearchArtifactsByContentRequest) Limit(limit int32) SearchApiSearchArtifactsByContentRequest {
	r.limit = &limit
	return r
}

// Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;).
func (r SearchApiSearchArtifactsByContentRequest) Order(order string) SearchApiSearchArtifactsByContentRequest {
	r.order = &order
	return r
}

// The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60; 
func (r SearchApiSearchArtifactsByContentRequest) Orderby(orderby string) SearchApiSearchArtifactsByContentRequest {
	r.orderby = &orderby
	return r
}

func (r SearchApiSearchArtifactsByContentRequest) Execute() (*ArtifactSearchResults, *http.Response, error) {
	return r.ApiService.SearchArtifactsByContentExecute(r)
}

/*
SearchArtifactsByContent Search for artifacts by content

Returns a paginated list of all artifacts with at least one version that matches the
posted content.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SearchApiSearchArtifactsByContentRequest
*/
func (a *SearchApiService) SearchArtifactsByContent(ctx context.Context) SearchApiSearchArtifactsByContentRequest {
	return SearchApiSearchArtifactsByContentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ArtifactSearchResults
func (a *SearchApiService) SearchArtifactsByContentExecute(r SearchApiSearchArtifactsByContentRequest) (*ArtifactSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.SearchArtifactsByContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search/artifacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.canonical != nil {
		localVarQueryParams.Add("canonical", parameterToString(*r.canonical, ""))
	}
	if r.artifactType != nil {
		localVarQueryParams.Add("artifactType", parameterToString(*r.artifactType, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.orderby != nil {
		localVarQueryParams.Add("orderby", parameterToString(*r.orderby, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
