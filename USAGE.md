<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "test"
    "test/pkg/models/shared"
    "test/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: shared.SchemeAPIKeyAuth{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.ListPromptsRequest{
        QueryParams: operations.ListPromptsQueryParams{
            RunID: "unde",
        },
    }

    ctx := context.Background()
    res, err := s.Prompts.List(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApiextListPromptsResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->