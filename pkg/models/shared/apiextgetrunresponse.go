// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ApiextGetRunResponseRuntime - Runtime that this run executed on.
type ApiextGetRunResponseRuntime string

const (
	ApiextGetRunResponseRuntimeStandard ApiextGetRunResponseRuntime = "standard"
	ApiextGetRunResponseRuntimeWorkflow ApiextGetRunResponseRuntime = "workflow"
)

func (e ApiextGetRunResponseRuntime) ToPointer() *ApiextGetRunResponseRuntime {
	return &e
}

func (e *ApiextGetRunResponseRuntime) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "workflow":
		*e = ApiextGetRunResponseRuntime(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApiextGetRunResponseRuntime: %v", v)
	}
}

// ApiextGetRunResponseStatus - Status of this run.
type ApiextGetRunResponseStatus string

const (
	ApiextGetRunResponseStatusNotStarted ApiextGetRunResponseStatus = "NotStarted"
	ApiextGetRunResponseStatusQueued     ApiextGetRunResponseStatus = "Queued"
	ApiextGetRunResponseStatusActive     ApiextGetRunResponseStatus = "Active"
	ApiextGetRunResponseStatusSucceeded  ApiextGetRunResponseStatus = "Succeeded"
	ApiextGetRunResponseStatusFailed     ApiextGetRunResponseStatus = "Failed"
	ApiextGetRunResponseStatusCancelled  ApiextGetRunResponseStatus = "Cancelled"
)

func (e ApiextGetRunResponseStatus) ToPointer() *ApiextGetRunResponseStatus {
	return &e
}

func (e *ApiextGetRunResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NotStarted":
		fallthrough
	case "Queued":
		fallthrough
	case "Active":
		fallthrough
	case "Succeeded":
		fallthrough
	case "Failed":
		fallthrough
	case "Cancelled":
		*e = ApiextGetRunResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApiextGetRunResponseStatus: %v", v)
	}
}

// ApiextGetRunResponse - OK
type ApiextGetRunResponse struct {
	// When the run became active. Empty if this run has not started yet.
	ActiveAt *string `json:"activeAt,omitempty"`
	// When the run was cancelled. Empty if this run was not cancelled.
	CancelledAt *string `json:"cancelledAt,omitempty"`
	// ID of the user who cancelled this run.
	CancelledBy *string               `json:"cancelledBy,omitempty"`
	Constraints *ApiextRunConstraints `json:"constraints,omitempty"`
	// When this run was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// ID of the user that created this run.
	CreatedBy *string `json:"createdBy,omitempty"`
	// When the run failed. Empty if this run did not fail.
	FailedAt *string `json:"failedAt,omitempty"`
	// Unique ID of this run.
	ID *string `json:"id,omitempty"`
	// Whether or not this run is private.
	IsPrivate *bool `json:"isPrivate,omitempty"`
	// Mapping of parameter slug to value used in this run's execution.
	ParamValues map[string]string `json:"paramValues,omitempty"`
	// Schema for the set of values users can provide when executing this run.
	Params []ApiextParameter `json:"params,omitempty"`
	// Explicit permissions of this run if it is private.
	Permissions []ApiextPermission `json:"permissions,omitempty"`
	Resources   map[string]string  `json:"resources,omitempty"`
	// Runtime that this run executed on.
	Runtime *ApiextGetRunResponseRuntime `json:"runtime,omitempty"`
	// ID of the session this run was spawned from if triggered by a session.
	SessionID *string `json:"sessionID,omitempty"`
	// Status of this run.
	Status *ApiextGetRunResponseStatus `json:"status,omitempty"`
	// When the run succeeded. Empty if this run did not succeed.
	SucceededAt *string `json:"succeededAt,omitempty"`
	// ID of the task this run was spawned from if triggered by a task.
	TaskID *string `json:"taskID,omitempty"`
	// ID of the team that owns this run.
	TeamID *string `json:"teamID,omitempty"`
	// Maximum amount of time in seconds the run could have executed for.
	Timeout *int64 `json:"timeout,omitempty"`
	// ID of the storage zone that the run used for its runs and outputs. Will be null if
	// there was no storage zone, in which case logs and outputs will be in the airplane API.
	ZoneID *string `json:"zoneID,omitempty"`
}
