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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ExecuteTaskRequest{
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
    }

    res, err := s.Tasks.Execute(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextExecuteTaskResponse != nil {
        // handle response
    }
}
```