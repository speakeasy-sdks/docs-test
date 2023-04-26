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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetSessionRequest{
        ID: "7739251a-a52c-43f5-ad01-9da1ffe78f09",
    }

    res, err := s.Sessions.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextGetSessionResponse != nil {
        // handle response
    }
}
```

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListSessionsRequest{
        Limit: sdk.Int64(451159),
        Page: sdk.Int64(739264),
        RunbookID: sdk.String("perferendis"),
        UpdatedAfter: sdk.String("doloremque"),
        UpdatedBefore: sdk.String("reprehenderit"),
    }

    res, err := s.Sessions.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListSessionResponse != nil {
        // handle response
    }
}
```
