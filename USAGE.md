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