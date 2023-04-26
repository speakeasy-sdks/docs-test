# Prompts

## Overview

A prompt is used to gather user input during a task's execution. See Prompts to see how prompts are used.

### Available Operations

* [Cancel](#cancel) - Cancel Prompt
* [Get](#get) - Get Prompt
* [List](#list) - List Prompts
* [Submit](#submit) - Submit Prompt

## Cancel

Cancel a prompt.

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.ApiextCancelPromptRequest{
        ID: sdk.String("pmt20221122zyydx3rho2t"),
    }

    res, err := s.Prompts.Cancel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextCancelPromptResponse != nil {
        // handle response
    }
}
```

## Get

Get information about an existing prompt.

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
    req := operations.GetPromptRequest{
        ID: "9bd9d8d6-9a67-44e0-b467-cc8796ed151a",
    }

    res, err := s.Prompts.Get(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextGetPromptResponse != nil {
        // handle response
    }
}
```

## List

List prompts from an existing run.

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
    req := operations.ListPromptsRequest{
        RunID: "perferendis",
    }

    res, err := s.Prompts.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListPromptsResponse != nil {
        // handle response
    }
}
```

## Submit

Submit a prompt with a set of parameter values.

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.ApiextSubmitPromptRequest{
        ID: sdk.String("pmt20221122zyydx3rho2t"),
        Values: map[string]string{
            "repellendus": "sapiente",
            "quo": "odit",
        },
    }

    res, err := s.Prompts.Submit(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextSubmitPromptResponse != nil {
        // handle response
    }
}
```
