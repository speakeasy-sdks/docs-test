// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test/pkg/models/shared"
)

type ListPromptsQueryParams struct {
	// ID of the run to retrieve prompts for.
	RunID string `queryParam:"style=form,explode=true,name=runID"`
}

type ListPromptsRequest struct {
	QueryParams ListPromptsQueryParams
}

type ListPromptsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApiextListPromptsResponse *shared.ApiextListPromptsResponse
}
