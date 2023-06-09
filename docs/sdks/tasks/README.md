# Tasks

## Overview

A Task is a lightweight app that represents a single business operation for people at your company to execute.

### Available Operations

* [Execute](#execute) - Execute Task

## Execute

Execute a task with a set of parameter values and receive a run ID to track the task's execution.
Check on the status of your newly created run with [/runs/get](/api/runs#runs-get).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test"
	"test/pkg/models/operations"
	"test/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Tasks.Execute(ctx, operations.ExecuteTaskRequest{
        ApiextExecuteTaskRequest: shared.ApiextExecuteTaskRequest{
            ID: sdk.String("tsk20210728zxb2vxn"),
            ParamValues: map[string]string{
                "maiores": "dicta",
                "corporis": "dolore",
            },
            Resources: map[string]string{
                "dicta": "harum",
                "enim": "accusamus",
            },
            Slug: sdk.String("hello_world"),
        },
        EnvSlug: sdk.String("commodi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextExecuteTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ExecuteTaskRequest](../../models/operations/executetaskrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ExecuteTaskResponse](../../models/operations/executetaskresponse.md), error**

