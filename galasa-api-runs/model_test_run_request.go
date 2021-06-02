/*
 * Galasa Runs Requests
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.15.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runs.api.galasa.dev

import (
	"encoding/json"
)

// TestRunRequest struct for TestRunRequest
type TestRunRequest struct {
	ClassNames *[]string `json:"classNames,omitempty"`
	// The request type, normally REQUEST, JENKINS, can be others
	RequestorType *string `json:"requestorType,omitempty"`
	// The requestor, usually email address but doesnt need to be
	Requestor *string `json:"requestor,omitempty"`
	// The test stream to be used
	TestStream *string `json:"testStream,omitempty"`
	// an overriding OBR
	Obr *string `json:"obr,omitempty"`
	// an overriding maven repo
	MavenRepository *string `json:"mavenRepository,omitempty"`
	// The shared env phase, build or discard, for normal runs, null
	SharedEnvironmentPhase *string `json:"sharedEnvironmentPhase,omitempty"`
	// The shared env name,  for normal runs, null
	SharedEnvironmentRunName *string `json:"sharedEnvironmentRunName,omitempty"`
	// Override properties
	Overrides *map[string]interface{} `json:"overrides,omitempty"`
	// enable trace
	Trace *bool `json:"trace,omitempty"`
}

// NewTestRunRequest instantiates a new TestRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunRequest() *TestRunRequest {
	this := TestRunRequest{}
	return &this
}

// NewTestRunRequestWithDefaults instantiates a new TestRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunRequestWithDefaults() *TestRunRequest {
	this := TestRunRequest{}
	return &this
}

// GetClassNames returns the ClassNames field value if set, zero value otherwise.
func (o *TestRunRequest) GetClassNames() []string {
	if o == nil || o.ClassNames == nil {
		var ret []string
		return ret
	}
	return *o.ClassNames
}

// GetClassNamesOk returns a tuple with the ClassNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetClassNamesOk() (*[]string, bool) {
	if o == nil || o.ClassNames == nil {
		return nil, false
	}
	return o.ClassNames, true
}

// HasClassNames returns a boolean if a field has been set.
func (o *TestRunRequest) HasClassNames() bool {
	if o != nil && o.ClassNames != nil {
		return true
	}

	return false
}

// SetClassNames gets a reference to the given []string and assigns it to the ClassNames field.
func (o *TestRunRequest) SetClassNames(v []string) {
	o.ClassNames = &v
}

// GetRequestorType returns the RequestorType field value if set, zero value otherwise.
func (o *TestRunRequest) GetRequestorType() string {
	if o == nil || o.RequestorType == nil {
		var ret string
		return ret
	}
	return *o.RequestorType
}

// GetRequestorTypeOk returns a tuple with the RequestorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetRequestorTypeOk() (*string, bool) {
	if o == nil || o.RequestorType == nil {
		return nil, false
	}
	return o.RequestorType, true
}

// HasRequestorType returns a boolean if a field has been set.
func (o *TestRunRequest) HasRequestorType() bool {
	if o != nil && o.RequestorType != nil {
		return true
	}

	return false
}

// SetRequestorType gets a reference to the given string and assigns it to the RequestorType field.
func (o *TestRunRequest) SetRequestorType(v string) {
	o.RequestorType = &v
}

// GetRequestor returns the Requestor field value if set, zero value otherwise.
func (o *TestRunRequest) GetRequestor() string {
	if o == nil || o.Requestor == nil {
		var ret string
		return ret
	}
	return *o.Requestor
}

// GetRequestorOk returns a tuple with the Requestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetRequestorOk() (*string, bool) {
	if o == nil || o.Requestor == nil {
		return nil, false
	}
	return o.Requestor, true
}

// HasRequestor returns a boolean if a field has been set.
func (o *TestRunRequest) HasRequestor() bool {
	if o != nil && o.Requestor != nil {
		return true
	}

	return false
}

// SetRequestor gets a reference to the given string and assigns it to the Requestor field.
func (o *TestRunRequest) SetRequestor(v string) {
	o.Requestor = &v
}

// GetTestStream returns the TestStream field value if set, zero value otherwise.
func (o *TestRunRequest) GetTestStream() string {
	if o == nil || o.TestStream == nil {
		var ret string
		return ret
	}
	return *o.TestStream
}

// GetTestStreamOk returns a tuple with the TestStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetTestStreamOk() (*string, bool) {
	if o == nil || o.TestStream == nil {
		return nil, false
	}
	return o.TestStream, true
}

// HasTestStream returns a boolean if a field has been set.
func (o *TestRunRequest) HasTestStream() bool {
	if o != nil && o.TestStream != nil {
		return true
	}

	return false
}

// SetTestStream gets a reference to the given string and assigns it to the TestStream field.
func (o *TestRunRequest) SetTestStream(v string) {
	o.TestStream = &v
}

// GetObr returns the Obr field value if set, zero value otherwise.
func (o *TestRunRequest) GetObr() string {
	if o == nil || o.Obr == nil {
		var ret string
		return ret
	}
	return *o.Obr
}

// GetObrOk returns a tuple with the Obr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetObrOk() (*string, bool) {
	if o == nil || o.Obr == nil {
		return nil, false
	}
	return o.Obr, true
}

// HasObr returns a boolean if a field has been set.
func (o *TestRunRequest) HasObr() bool {
	if o != nil && o.Obr != nil {
		return true
	}

	return false
}

// SetObr gets a reference to the given string and assigns it to the Obr field.
func (o *TestRunRequest) SetObr(v string) {
	o.Obr = &v
}

// GetMavenRepository returns the MavenRepository field value if set, zero value otherwise.
func (o *TestRunRequest) GetMavenRepository() string {
	if o == nil || o.MavenRepository == nil {
		var ret string
		return ret
	}
	return *o.MavenRepository
}

// GetMavenRepositoryOk returns a tuple with the MavenRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetMavenRepositoryOk() (*string, bool) {
	if o == nil || o.MavenRepository == nil {
		return nil, false
	}
	return o.MavenRepository, true
}

// HasMavenRepository returns a boolean if a field has been set.
func (o *TestRunRequest) HasMavenRepository() bool {
	if o != nil && o.MavenRepository != nil {
		return true
	}

	return false
}

// SetMavenRepository gets a reference to the given string and assigns it to the MavenRepository field.
func (o *TestRunRequest) SetMavenRepository(v string) {
	o.MavenRepository = &v
}

// GetSharedEnvironmentPhase returns the SharedEnvironmentPhase field value if set, zero value otherwise.
func (o *TestRunRequest) GetSharedEnvironmentPhase() string {
	if o == nil || o.SharedEnvironmentPhase == nil {
		var ret string
		return ret
	}
	return *o.SharedEnvironmentPhase
}

// GetSharedEnvironmentPhaseOk returns a tuple with the SharedEnvironmentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetSharedEnvironmentPhaseOk() (*string, bool) {
	if o == nil || o.SharedEnvironmentPhase == nil {
		return nil, false
	}
	return o.SharedEnvironmentPhase, true
}

// HasSharedEnvironmentPhase returns a boolean if a field has been set.
func (o *TestRunRequest) HasSharedEnvironmentPhase() bool {
	if o != nil && o.SharedEnvironmentPhase != nil {
		return true
	}

	return false
}

// SetSharedEnvironmentPhase gets a reference to the given string and assigns it to the SharedEnvironmentPhase field.
func (o *TestRunRequest) SetSharedEnvironmentPhase(v string) {
	o.SharedEnvironmentPhase = &v
}

// GetSharedEnvironmentRunName returns the SharedEnvironmentRunName field value if set, zero value otherwise.
func (o *TestRunRequest) GetSharedEnvironmentRunName() string {
	if o == nil || o.SharedEnvironmentRunName == nil {
		var ret string
		return ret
	}
	return *o.SharedEnvironmentRunName
}

// GetSharedEnvironmentRunNameOk returns a tuple with the SharedEnvironmentRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetSharedEnvironmentRunNameOk() (*string, bool) {
	if o == nil || o.SharedEnvironmentRunName == nil {
		return nil, false
	}
	return o.SharedEnvironmentRunName, true
}

// HasSharedEnvironmentRunName returns a boolean if a field has been set.
func (o *TestRunRequest) HasSharedEnvironmentRunName() bool {
	if o != nil && o.SharedEnvironmentRunName != nil {
		return true
	}

	return false
}

// SetSharedEnvironmentRunName gets a reference to the given string and assigns it to the SharedEnvironmentRunName field.
func (o *TestRunRequest) SetSharedEnvironmentRunName(v string) {
	o.SharedEnvironmentRunName = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *TestRunRequest) GetOverrides() map[string]interface{} {
	if o == nil || o.Overrides == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetOverridesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *TestRunRequest) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given map[string]interface{} and assigns it to the Overrides field.
func (o *TestRunRequest) SetOverrides(v map[string]interface{}) {
	o.Overrides = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *TestRunRequest) GetTrace() bool {
	if o == nil || o.Trace == nil {
		var ret bool
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunRequest) GetTraceOk() (*bool, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *TestRunRequest) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given bool and assigns it to the Trace field.
func (o *TestRunRequest) SetTrace(v bool) {
	o.Trace = &v
}

func (o TestRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClassNames != nil {
		toSerialize["classNames"] = o.ClassNames
	}
	if o.RequestorType != nil {
		toSerialize["requestorType"] = o.RequestorType
	}
	if o.Requestor != nil {
		toSerialize["requestor"] = o.Requestor
	}
	if o.TestStream != nil {
		toSerialize["testStream"] = o.TestStream
	}
	if o.Obr != nil {
		toSerialize["obr"] = o.Obr
	}
	if o.MavenRepository != nil {
		toSerialize["mavenRepository"] = o.MavenRepository
	}
	if o.SharedEnvironmentPhase != nil {
		toSerialize["sharedEnvironmentPhase"] = o.SharedEnvironmentPhase
	}
	if o.SharedEnvironmentRunName != nil {
		toSerialize["sharedEnvironmentRunName"] = o.SharedEnvironmentRunName
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	return json.Marshal(toSerialize)
}

type NullableTestRunRequest struct {
	value *TestRunRequest
	isSet bool
}

func (v NullableTestRunRequest) Get() *TestRunRequest {
	return v.value
}

func (v *NullableTestRunRequest) Set(val *TestRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunRequest(val *TestRunRequest) *NullableTestRunRequest {
	return &NullableTestRunRequest{value: val, isSet: true}
}

func (v NullableTestRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

