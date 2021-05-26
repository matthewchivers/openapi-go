# TestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The run name | [optional] 
**Type** | Pointer to **string** | The type of request, ie REQUEST, JENKINS, LOCAL | [optional] 
**Group** | Pointer to **string** | the group ID assigned to this run | [optional] 
**Test** | Pointer to **string** | the full test name in bundle/test format | [optional] 
**BundleName** | Pointer to **string** | the bundle name | [optional] 
**TestName** | Pointer to **string** | the test class name | [optional] 
**Status** | Pointer to **string** | the current status of the run, ie allocated, starting, building etc | [optional] 
**Queued** | Pointer to **string** | when the test was queued | [optional] 
**Requestor** | Pointer to **string** | who requested the run | [optional] 
**Stream** | Pointer to **string** | the test stream the test should be loaded from | [optional] 
**Repo** | Pointer to **string** | the maven repositories to be used | [optional] 
**Obr** | Pointer to **string** | the obrs to be used | [optional] 
**Local** | Pointer to **bool** | is this a local run | [optional] 
**Trace** | Pointer to **bool** | has trace been requested | [optional] 

## Methods

### NewTestRun

`func NewTestRun() *TestRun`

NewTestRun instantiates a new TestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunWithDefaults

`func NewTestRunWithDefaults() *TestRun`

NewTestRunWithDefaults instantiates a new TestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TestRun) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestRun) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestRun) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TestRun) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *TestRun) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TestRun) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TestRun) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TestRun) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetTest

`func (o *TestRun) GetTest() string`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *TestRun) GetTestOk() (*string, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *TestRun) SetTest(v string)`

SetTest sets Test field to given value.

### HasTest

`func (o *TestRun) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetBundleName

`func (o *TestRun) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *TestRun) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *TestRun) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *TestRun) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetTestName

`func (o *TestRun) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestRun) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestRun) SetTestName(v string)`

SetTestName sets TestName field to given value.

### HasTestName

`func (o *TestRun) HasTestName() bool`

HasTestName returns a boolean if a field has been set.

### GetStatus

`func (o *TestRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQueued

`func (o *TestRun) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *TestRun) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *TestRun) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *TestRun) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRequestor

`func (o *TestRun) GetRequestor() string`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *TestRun) GetRequestorOk() (*string, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *TestRun) SetRequestor(v string)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *TestRun) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetStream

`func (o *TestRun) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *TestRun) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *TestRun) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *TestRun) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetRepo

`func (o *TestRun) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *TestRun) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *TestRun) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *TestRun) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetObr

`func (o *TestRun) GetObr() string`

GetObr returns the Obr field if non-nil, zero value otherwise.

### GetObrOk

`func (o *TestRun) GetObrOk() (*string, bool)`

GetObrOk returns a tuple with the Obr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObr

`func (o *TestRun) SetObr(v string)`

SetObr sets Obr field to given value.

### HasObr

`func (o *TestRun) HasObr() bool`

HasObr returns a boolean if a field has been set.

### GetLocal

`func (o *TestRun) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *TestRun) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *TestRun) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *TestRun) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetTrace

`func (o *TestRun) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *TestRun) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *TestRun) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *TestRun) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


