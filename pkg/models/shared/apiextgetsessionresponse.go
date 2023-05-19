// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ApiextGetSessionResponseStatus - Status of this session.
type ApiextGetSessionResponseStatus string

const (
	ApiextGetSessionResponseStatusPending    ApiextGetSessionResponseStatus = "Pending"
	ApiextGetSessionResponseStatusActive     ApiextGetSessionResponseStatus = "Active"
	ApiextGetSessionResponseStatusWaiting    ApiextGetSessionResponseStatus = "Waiting"
	ApiextGetSessionResponseStatusSucceeded  ApiextGetSessionResponseStatus = "Succeeded"
	ApiextGetSessionResponseStatusFailed     ApiextGetSessionResponseStatus = "Failed"
	ApiextGetSessionResponseStatusCancelling ApiextGetSessionResponseStatus = "Cancelling"
	ApiextGetSessionResponseStatusCancelled  ApiextGetSessionResponseStatus = "Cancelled"
)

func (e ApiextGetSessionResponseStatus) ToPointer() *ApiextGetSessionResponseStatus {
	return &e
}

func (e *ApiextGetSessionResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Active":
		fallthrough
	case "Waiting":
		fallthrough
	case "Succeeded":
		fallthrough
	case "Failed":
		fallthrough
	case "Cancelling":
		fallthrough
	case "Cancelled":
		*e = ApiextGetSessionResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApiextGetSessionResponseStatus: %v", v)
	}
}

// ApiextGetSessionResponse - OK
type ApiextGetSessionResponse struct {
	// When this session was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// ID of the user that created this session.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Unique ID of this session.
	ID *string `json:"id,omitempty"`
	// Whether or not the session is private.
	IsPrivate *bool `json:"isPrivate,omitempty"`
	// Name of this session.
	Name *string `json:"name,omitempty"`
	// Mapping of parameter slug to value used in this session's execution.
	ParamValues map[string]string `json:"paramValues,omitempty"`
	// Schema for the set of values users can provide when executing this session.
	Params []ApiextParameter `json:"params,omitempty"`
	// Explicit permissions of this session if it is private.
	Permissions []ApiextPermission `json:"permissions,omitempty"`
	// ID of the runbook this session was spawned from if triggered from a runbook.
	RunbookID *string `json:"runbookID,omitempty"`
	// Status of this session.
	Status *ApiextGetSessionResponseStatus `json:"status,omitempty"`
	// ID of the team that owns this session.
	TeamID *string `json:"teamID,omitempty"`
	// When this session was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// ID of the user who updated this session.
	UpdatedBy *string `json:"updatedBy,omitempty"`
}
