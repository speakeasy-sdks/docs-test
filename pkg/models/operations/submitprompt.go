// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test/pkg/models/shared"
)

type SubmitPromptResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApiextSubmitPromptResponse *shared.ApiextSubmitPromptResponse
}
