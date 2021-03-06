// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RobotCreateV1 robot create v1
// swagger:model RobotCreateV1
type RobotCreateV1 struct {

	// The permission of robot account
	Access []*Access `json:"access"`

	// The description of robot account
	Description string `json:"description,omitempty"`

	// The expiration time on or after which the JWT MUST NOT be accepted for processing.
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// The name of robot account
	Name string `json:"name,omitempty"`
}

// Validate validates this robot create v1
func (m *RobotCreateV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RobotCreateV1) validateAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Access) { // not required
		return nil
	}

	for i := 0; i < len(m.Access); i++ {
		if swag.IsZero(m.Access[i]) { // not required
			continue
		}

		if m.Access[i] != nil {
			if err := m.Access[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RobotCreateV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotCreateV1) UnmarshalBinary(b []byte) error {
	var res RobotCreateV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
