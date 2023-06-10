# Runbooks

## Overview

A Runbook is a multi-step, human-in-the-loop workflow. Runbooks are able to take a set of top-level parameters, run one or more functions, and generate output at each step of the way.

### Available Operations

* [Execute](#execute) - Execute Runbook

## Execute

Execute a runbook and receive a session ID to track the runbook's execution.
Check on the status of your newly created session with [/sessions/get](/api/sessions#sessions-get).

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
    res, err := s.Runbooks.Execute(ctx, operations.ExecuteRunbookRequest{
        ApiextExecuteRunbookRequest: shared.ApiextExecuteRunbookRequest{
            ID: "rbk20220120z15kl79",
            ParamValues: map[string]string{
                "at": "maiores",
                "molestiae": "quod",
                "quod": "esse",
                "totam": "porro",
            },
            Slug: sdk.String("hello_world"),
        },
        EnvSlug: sdk.String("dolorum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextExecuteRunbookResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ExecuteRunbookRequest](../../models/operations/executerunbookrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ExecuteRunbookResponse](../../models/operations/executerunbookresponse.md), error**

