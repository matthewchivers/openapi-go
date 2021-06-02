# TestRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **bool** |  | [optional] 
**Runs** | Pointer to [**[]TestRun**](TestRun.md) |  | [optional] 

## Methods

### NewTestRuns

`func NewTestRuns() *TestRuns`

NewTestRuns instantiates a new TestRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunsWithDefaults

`func NewTestRunsWithDefaults() *TestRuns`

NewTestRunsWithDefaults instantiates a new TestRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *TestRuns) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *TestRuns) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *TestRuns) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *TestRuns) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetRuns

`func (o *TestRuns) GetRuns() []TestRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *TestRuns) GetRunsOk() (*[]TestRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *TestRuns) SetRuns(v []TestRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *TestRuns) HasRuns() bool`

HasRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


