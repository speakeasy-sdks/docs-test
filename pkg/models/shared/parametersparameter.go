// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ParametersParameterComponentEnum - Used to specify UI-only type modifiers
type ParametersParameterComponentEnum string

const (
	ParametersParameterComponentEnumUnknown   ParametersParameterComponentEnum = ""
	ParametersParameterComponentEnumEditorSQL ParametersParameterComponentEnum = "editor-sql"
	ParametersParameterComponentEnumTextarea  ParametersParameterComponentEnum = "textarea"
)

func (e *ParametersParameterComponentEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "":
		fallthrough
	case "editor-sql":
		fallthrough
	case "textarea":
		*e = ParametersParameterComponentEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ParametersParameterComponentEnum: %s", s)
	}
}

// ParametersParameterTypeEnum - Parameter data type.
type ParametersParameterTypeEnum string

const (
	ParametersParameterTypeEnumAny       ParametersParameterTypeEnum = "any"
	ParametersParameterTypeEnumString    ParametersParameterTypeEnum = "string"
	ParametersParameterTypeEnumBoolean   ParametersParameterTypeEnum = "boolean"
	ParametersParameterTypeEnumUpload    ParametersParameterTypeEnum = "upload"
	ParametersParameterTypeEnumInteger   ParametersParameterTypeEnum = "integer"
	ParametersParameterTypeEnumFloat     ParametersParameterTypeEnum = "float"
	ParametersParameterTypeEnumDate      ParametersParameterTypeEnum = "date"
	ParametersParameterTypeEnumDatetime  ParametersParameterTypeEnum = "datetime"
	ParametersParameterTypeEnumConfigvar ParametersParameterTypeEnum = "configvar"
	ParametersParameterTypeEnumList      ParametersParameterTypeEnum = "list"
	ParametersParameterTypeEnumMap       ParametersParameterTypeEnum = "map"
	ParametersParameterTypeEnumObject    ParametersParameterTypeEnum = "object"
)

func (e *ParametersParameterTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "any":
		fallthrough
	case "string":
		fallthrough
	case "boolean":
		fallthrough
	case "upload":
		fallthrough
	case "integer":
		fallthrough
	case "float":
		fallthrough
	case "date":
		fallthrough
	case "datetime":
		fallthrough
	case "configvar":
		fallthrough
	case "list":
		fallthrough
	case "map":
		fallthrough
	case "object":
		*e = ParametersParameterTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ParametersParameterTypeEnum: %s", s)
	}
}

type ParametersParameter struct {
	// Used to specify UI-only type modifiers
	Component   *ParametersParameterComponentEnum `json:"component,omitempty"`
	Constraints *ParametersConstraints            `json:"constraints,omitempty"`
	// Optional default value for this parameter, used if not set.
	Default interface{} `json:"default,omitempty"`
	// Description for this parameter.
	Desc *string `json:"desc,omitempty"`
	// Name for this parameter.
	Name *string `json:"name,omitempty"`
	// If this parameter has an object data type, represents an ordered list of key-value pairs that can be included in this object.
	Parameters []ParametersParameter `json:"parameters,omitempty"`
	// A human-friendly identifier for the parameter.
	Slug *string `json:"slug,omitempty"`
	// Parameter data type.
	Type   *ParametersParameterTypeEnum `json:"type,omitempty"`
	Values *ParametersParameter         `json:"values,omitempty"`
}