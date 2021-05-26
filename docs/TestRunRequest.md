# TestRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassNames** | Pointer to **[]string** |  | [optional] 
**RequestorType** | Pointer to **string** | The request type, normally REQUEST, JENKINS, can be others | [optional] 
**Requestor** | Pointer to **string** | The requestor, usually email address but doesnt need to be | [optional] 
**TestStream** | Pointer to **string** | The test stream to be used | [optional] 
**Obr** | Pointer to **string** | an overriding OBR | [optional] 
**MavenRepository** | Pointer to **string** | an overriding maven repo | [optional] 
**SharedEnvironmentPhase** | Pointer to **string** | The shared env phase, build or discard, for normal runs, null | [optional] 
**SharedEnvironmentRunName** | Pointer to **string** | The shared env name,  for normal runs, null | [optional] 
**Overrides** | Pointer to **map[string]interface{}** | Override properties | [optional] 
**Trace** | Pointer to **bool** | enable trace | [optional] 

## Methods

### NewTestRunRequest

`func NewTestRunRequest() *TestRunRequest`

NewTestRunRequest instantiates a new TestRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunRequestWithDefaults

`func NewTestRunRequestWithDefaults() *TestRunRequest`

NewTestRunRequestWithDefaults instantiates a new TestRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassNames

`func (o *TestRunRequest) GetClassNames() []string`

GetClassNames returns the ClassNames field if non-nil, zero value otherwise.

### GetClassNamesOk

`func (o *TestRunRequest) GetClassNamesOk() (*[]string, bool)`

GetClassNamesOk returns a tuple with the ClassNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassNames

`func (o *TestRunRequest) SetClassNames(v []string)`

SetClassNames sets ClassNames field to given value.

### HasClassNames

`func (o *TestRunRequest) HasClassNames() bool`

HasClassNames returns a boolean if a field has been set.

### GetRequestorType

`func (o *TestRunRequest) GetRequestorType() string`

GetRequestorType returns the RequestorType field if non-nil, zero value otherwise.

### GetRequestorTypeOk

`func (o *TestRunRequest) GetRequestorTypeOk() (*string, bool)`

GetRequestorTypeOk returns a tuple with the RequestorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorType

`func (o *TestRunRequest) SetRequestorType(v string)`

SetRequestorType sets RequestorType field to given value.

### HasRequestorType

`func (o *TestRunRequest) HasRequestorType() bool`

HasRequestorType returns a boolean if a field has been set.

### GetRequestor

`func (o *TestRunRequest) GetRequestor() string`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *TestRunRequest) GetRequestorOk() (*string, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *TestRunRequest) SetRequestor(v string)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *TestRunRequest) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetTestStream

`func (o *TestRunRequest) GetTestStream() string`

GetTestStream returns the TestStream field if non-nil, zero value otherwise.

### GetTestStreamOk

`func (o *TestRunRequest) GetTestStreamOk() (*string, bool)`

GetTestStreamOk returns a tuple with the TestStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStream

`func (o *TestRunRequest) SetTestStream(v string)`

SetTestStream sets TestStream field to given value.

### HasTestStream

`func (o *TestRunRequest) HasTestStream() bool`

HasTestStream returns a boolean if a field has been set.

### GetObr

`func (o *TestRunRequest) GetObr() string`

GetObr returns the Obr field if non-nil, zero value otherwise.

### GetObrOk

`func (o *TestRunRequest) GetObrOk() (*string, bool)`

GetObrOk returns a tuple with the Obr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObr

`func (o *TestRunRequest) SetObr(v string)`

SetObr sets Obr field to given value.

### HasObr

`func (o *TestRunRequest) HasObr() bool`

HasObr returns a boolean if a field has been set.

### GetMavenRepository

`func (o *TestRunRequest) GetMavenRepository() string`

GetMavenRepository returns the MavenRepository field if non-nil, zero value otherwise.

### GetMavenRepositoryOk

`func (o *TestRunRequest) GetMavenRepositoryOk() (*string, bool)`

GetMavenRepositoryOk returns a tuple with the MavenRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMavenRepository

`func (o *TestRunRequest) SetMavenRepository(v string)`

SetMavenRepository sets MavenRepository field to given value.

### HasMavenRepository

`func (o *TestRunRequest) HasMavenRepository() bool`

HasMavenRepository returns a boolean if a field has been set.

### GetSharedEnvironmentPhase

`func (o *TestRunRequest) GetSharedEnvironmentPhase() string`

GetSharedEnvironmentPhase returns the SharedEnvironmentPhase field if non-nil, zero value otherwise.

### GetSharedEnvironmentPhaseOk

`func (o *TestRunRequest) GetSharedEnvironmentPhaseOk() (*string, bool)`

GetSharedEnvironmentPhaseOk returns a tuple with the SharedEnvironmentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedEnvironmentPhase

`func (o *TestRunRequest) SetSharedEnvironmentPhase(v string)`

SetSharedEnvironmentPhase sets SharedEnvironmentPhase field to given value.

### HasSharedEnvironmentPhase

`func (o *TestRunRequest) HasSharedEnvironmentPhase() bool`

HasSharedEnvironmentPhase returns a boolean if a field has been set.

### GetSharedEnvironmentRunName

`func (o *TestRunRequest) GetSharedEnvironmentRunName() string`

GetSharedEnvironmentRunName returns the SharedEnvironmentRunName field if non-nil, zero value otherwise.

### GetSharedEnvironmentRunNameOk

`func (o *TestRunRequest) GetSharedEnvironmentRunNameOk() (*string, bool)`

GetSharedEnvironmentRunNameOk returns a tuple with the SharedEnvironmentRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedEnvironmentRunName

`func (o *TestRunRequest) SetSharedEnvironmentRunName(v string)`

SetSharedEnvironmentRunName sets SharedEnvironmentRunName field to given value.

### HasSharedEnvironmentRunName

`func (o *TestRunRequest) HasSharedEnvironmentRunName() bool`

HasSharedEnvironmentRunName returns a boolean if a field has been set.

### GetOverrides

`func (o *TestRunRequest) GetOverrides() map[string]interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *TestRunRequest) GetOverridesOk() (*map[string]interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *TestRunRequest) SetOverrides(v map[string]interface{})`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *TestRunRequest) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetTrace

`func (o *TestRunRequest) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *TestRunRequest) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *TestRunRequest) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *TestRunRequest) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


