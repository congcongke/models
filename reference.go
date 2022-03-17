// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Reference reference
// swagger:model Reference
type Reference struct {

	// annotations
	Annotations Annotations `json:"annotations,omitempty"`

	// The digest of the child artifact
	ChildDigest string `json:"child_digest,omitempty"`

	// The child ID of the reference
	ChildID int64 `json:"child_id,omitempty"`

	// The parent ID of the reference
	ParentID int64 `json:"parent_id,omitempty"`

	// platform
	Platform *Platform `json:"platform,omitempty"`

	// The download URLs
	Urls []string `json:"urls"`
}

// Validate validates this reference
func (m *Reference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Reference) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if err := m.Annotations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *Reference) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Reference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reference) UnmarshalBinary(b []byte) error {
	var res Reference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
