# \DefaultApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRunsByGroup**](DefaultApi.md#GetRunsByGroup) | **Get** /runs/{groupId} | Retrieve all active runs for a group id
[**SubmitRunsByGroup**](DefaultApi.md#SubmitRunsByGroup) | **Post** /runs/{groupId} | Submit a set of test runs for a group ID can be run multiple times against the same group ID, the new test runs will be appended



## GetRunsByGroup

> TestRuns GetRunsByGroup(ctx, groupId).Execute()

Retrieve all active runs for a group id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "groupId_example" // string | Run group ID can be any type of string value that is URL safe

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetRunsByGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRunsByGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunsByGroup`: TestRuns
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRunsByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Run group ID can be any type of string value that is URL safe | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunsByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestRuns**](TestRuns.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRunsByGroup

> SubmitRunsByGroup(ctx, groupId).TestRunRequest(testRunRequest).Execute()

Submit a set of test runs for a group ID can be run multiple times against the same group ID, the new test runs will be appended

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "groupId_example" // string | Run group ID can be any type of string value that is URL safe
    testRunRequest := *openapiclient.NewTestRunRequest() // TestRunRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SubmitRunsByGroup(context.Background(), groupId).TestRunRequest(testRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubmitRunsByGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Run group ID can be any type of string value that is URL safe | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRunsByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRunRequest** | [**TestRunRequest**](TestRunRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

