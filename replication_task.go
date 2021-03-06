// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplicationTask The replication task
// swagger:model ReplicationTask
type ReplicationTask struct {

	// The destination resource that the task operates
	DstResource string `json:"dst_resource,omitempty"`

	// The end time of the task
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// The ID of the execution that the task belongs to
	ExecutionID int64 `json:"execution_id,omitempty"`

	// The ID of the task
	ID int64 `json:"id,omitempty"`

	// The ID of the underlying job that the task related to
	JobID string `json:"job_id,omitempty"`

	// The operation of the task
	Operation string `json:"operation,omitempty"`

	// The type of the resource that the task operates
	ResourceType string `json:"resource_type,omitempty"`

	// The source resource that the task operates
	SrcResource string `json:"src_resource,omitempty"`

	// The start time of the task
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// The status of the task
	Status string `json:"status,omitempty"`
}

// Validate validates this replication task
func (m *ReplicationTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationTask) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReplicationTask) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationTask) UnmarshalBinary(b []byte) error {
	var res ReplicationTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
