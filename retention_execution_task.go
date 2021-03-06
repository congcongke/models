// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RetentionExecutionTask retention execution task
// swagger:model RetentionExecutionTask
type RetentionExecutionTask struct {

	// end time
	EndTime string `json:"end_time,omitempty"`

	// execution id
	ExecutionID int64 `json:"execution_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`

	// retained
	Retained int64 `json:"retained"`

	// start time
	StartTime string `json:"start_time,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status code
	StatusCode int64 `json:"status_code"`

	// status revision
	StatusRevision int64 `json:"status_revision,omitempty"`

	// total
	Total int64 `json:"total"`
}

// Validate validates this retention execution task
func (m *RetentionExecutionTask) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionExecutionTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionExecutionTask) UnmarshalBinary(b []byte) error {
	var res RetentionExecutionTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
