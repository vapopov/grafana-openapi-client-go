// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminCreateUserForm admin create user form
//
// swagger:model AdminCreateUserForm
type AdminCreateUserForm struct {

	// email
	Email string `json:"email,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// password
	Password Password `json:"password,omitempty"`
}

// Validate validates this admin create user form
func (m *AdminCreateUserForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminCreateUserForm) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := m.Password.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("password")
		}
		return err
	}

	return nil
}

// ContextValidate validate this admin create user form based on the context it is used
func (m *AdminCreateUserForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePassword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminCreateUserForm) contextValidatePassword(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := m.Password.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("password")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminCreateUserForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminCreateUserForm) UnmarshalBinary(b []byte) error {
	var res AdminCreateUserForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
