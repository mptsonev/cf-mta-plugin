package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ValidationStatus Result of parameter validation
// swagger:model ValidationStatus
type ValidationStatus struct {
	XMLName xml.Name `xml:"http://www.sap.com/lmsl/slp ValidationStatus"`

	// message
	// Required: true
	Message *string `xml:"message"`

	// status
	// Required: true
	Status SlpParameterValidationStatus `xml:"status"`
}

// Validate validates this validation status
func (m *ValidationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidationStatus) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *ValidationStatus) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		return err
	}

	return nil
}