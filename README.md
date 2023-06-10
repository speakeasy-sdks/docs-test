# test

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/docs-test
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
        RunID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListPromptsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Prompts](docs/sdks/prompts/README.md)

* [Cancel](docs/sdks/prompts/README.md#cancel) - Cancel Prompt
* [Get](docs/sdks/prompts/README.md#get) - Get Prompt
* [List](docs/sdks/prompts/README.md#list) - List Prompts
* [Submit](docs/sdks/prompts/README.md#submit) - Submit Prompt

### [Runbooks](docs/sdks/runbooks/README.md)

* [Execute](docs/sdks/runbooks/README.md#execute) - Execute Runbook

### [Runs](docs/sdks/runs/README.md)

* [Cancel](docs/sdks/runs/README.md#cancel) - Cancel Run
* [Get](docs/sdks/runs/README.md#get) - Cancel Run
* [GetOutputs](docs/sdks/runs/README.md#getoutputs) - Get Run Outputs
* [List](docs/sdks/runs/README.md#list) - List Runs

### [Sessions](docs/sdks/sessions/README.md)

* [Get](docs/sdks/sessions/README.md#get) - Get Session
* [List](docs/sdks/sessions/README.md#list) - List Sessions

### [Tasks](docs/sdks/tasks/README.md)

* [Execute](docs/sdks/tasks/README.md#execute) - Execute Task
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
