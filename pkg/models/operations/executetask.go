// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test/pkg/models/shared"
)

type ExecuteTaskRequest struct {
	// ExecuteTaskRequest
	ApiextExecuteTaskRequest shared.ApiextExecuteTaskRequest `request:"mediaType=application/json"`
	// Environment to execute the task in.
	EnvSlug *string `queryParam:"style=form,explode=true,name=envSlug"`
}

type ExecuteTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ApiextExecuteTaskResponse *shared.ApiextExecuteTaskResponse
}
