# Sessions

## Overview

A session represents an instance of a runbook's execution. See Runbooks API for how to execute runbooks.

### Available Operations

* [Get](#get) - Get Session
* [List](#list) - List Sessions

## Get

Get information about an existing session.

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
    res, err := s.Sessions.Get(ctx, operations.GetSessionRequest{
        ID: "7739251a-a52c-43f5-ad01-9da1ffe78f09",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextGetSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetSessionRequest](../../models/operations/getsessionrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetSessionResponse](../../models/operations/getsessionresponse.md), error**


## List

List Sessions

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
    res, err := s.Sessions.List(ctx, operations.ListSessionsRequest{
        Limit: sdk.Int64(451159),
        Page: sdk.Int64(739264),
        RunbookID: sdk.String("perferendis"),
        UpdatedAfter: sdk.String("doloremque"),
        UpdatedBefore: sdk.String("reprehenderit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListSessionsRequest](../../models/operations/listsessionsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListSessionsResponse](../../models/operations/listsessionsresponse.md), error**

