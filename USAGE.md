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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListPromptsRequest{
        RunID: "corrupti",
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
<!-- End SDK Example Usage -->