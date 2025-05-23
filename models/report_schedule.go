// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportSchedule report schedule
//
// swagger:model ReportSchedule
type ReportSchedule struct {

	// day of month
	DayOfMonth string `json:"dayOfMonth,omitempty"`

	// end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// frequency
	Frequency string `json:"frequency,omitempty"`

	// interval amount
	IntervalAmount int64 `json:"intervalAmount,omitempty"`

	// interval frequency
	IntervalFrequency string `json:"intervalFrequency,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`

	// workdays only
	WorkdaysOnly bool `json:"workdaysOnly,omitempty"`
}

// Validate validates this report schedule
func (m *ReportSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportSchedule) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report schedule based on context it is used
func (m *ReportSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportSchedule) UnmarshalBinary(b []byte) error {
	var res ReportSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
