# Runs

## Overview

A run represents an instance of a task's execution. See Tasks API for how to execute tasks.

### Available Operations

* [Cancel](#cancel) - Cancel Run
* [Get](#get) - Cancel Run
* [GetOutputs](#getoutputs) - Get Run Outputs
* [List](#list) - List Runs

## Cancel

Cancel a run.
Check on the status of your run with [/runs/get](/api/runs#runs-get).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test"
	"test/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.Cancel(ctx, shared.ApiextCancelRunRequest{
        RunID: sdk.String("run20220111zlq2ig4"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Get

Get information about an existing run.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test"
	"test/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.Get(ctx, operations.GetRunRequest{
        ID: "1ba928fc-8167-442c-b739-205929396fea",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextGetRunResponse != nil {
        // handle response
    }
}
```

## GetOutputs

Get outputs from an existing run.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test"
	"test/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.GetOutputs(ctx, operations.GetOutputsRequest{
        ID: "7596eb10-faaa-4235-ac59-55907aff1a3a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextGetOutputsResponse != nil {
        // handle response
    }
}
```

## List

List Runs

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test"
	"test/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Runs.List(ctx, operations.ListRunsRequest{
        Limit: sdk.Int64(161309),
        Page: sdk.Int64(995300),
        Since: sdk.String("mollitia"),
        TaskID: sdk.String("occaecati"),
        TaskSlug: sdk.String("numquam"),
        Until: sdk.String("commodi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListRunsResponse != nil {
        // handle response
    }
}
```
