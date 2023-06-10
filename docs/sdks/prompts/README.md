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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Cancel(ctx, shared.ApiextCancelPromptRequest{
        ID: sdk.String("pmt20221122zyydx3rho2t"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextCancelPromptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.ApiextCancelPromptRequest](../../models/shared/apiextcancelpromptrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CancelPromptResponse](../../models/operations/cancelpromptresponse.md), error**


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Get(ctx, operations.GetPromptRequest{
        ID: "9bd9d8d6-9a67-44e0-b467-cc8796ed151a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextGetPromptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetPromptRequest](../../models/operations/getpromptrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetPromptResponse](../../models/operations/getpromptresponse.md), error**


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.List(ctx, operations.ListPromptsRequest{
        RunID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListPromptsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListPromptsRequest](../../models/operations/listpromptsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListPromptsResponse](../../models/operations/listpromptsresponse.md), error**


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Prompts.Submit(ctx, shared.ApiextSubmitPromptRequest{
        ID: sdk.String("pmt20221122zyydx3rho2t"),
        Values: map[string]string{
            "repellendus": "sapiente",
            "quo": "odit",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextSubmitPromptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.ApiextSubmitPromptRequest](../../models/shared/apiextsubmitpromptrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SubmitPromptResponse](../../models/operations/submitpromptresponse.md), error**

