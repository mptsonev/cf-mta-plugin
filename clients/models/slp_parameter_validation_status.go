package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SlpParameterValidationStatus Parameter validation status
// swagger:model slp_parameter_validation_status
type SlpParameterValidationStatus string

const (
	SlpParameterValidationStatusSlpParameterValidationStatusERROR   SlpParameterValidationStatus = "slp.parameter.validation.status.ERROR"
	SlpParameterValidationStatusSlpParameterValidationStatusOK      SlpParameterValidationStatus = "slp.parameter.validation.status.OK"
	SlpParameterValidationStatusSlpParameterValidationStatusWARNING SlpParameterValidationStatus = "slp.parameter.validation.status.WARNING"
	SlpParameterValidationStatusSlpParameterValidationStatusINFO    SlpParameterValidationStatus = "slp.parameter.validation.status.INFO"
)

// for schema
var slpParameterValidationStatusEnum []interface{}

func (m SlpParameterValidationStatus) validateSlpParameterValidationStatusEnum(path, location string, value SlpParameterValidationStatus) error {
	if slpParameterValidationStatusEnum == nil {
		var res []SlpParameterValidationStatus
		if err := json.Unmarshal([]byte(`["slp.parameter.validation.status.ERROR","slp.parameter.validation.status.OK","slp.parameter.validation.status.WARNING","slp.parameter.validation.status.INFO"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			slpParameterValidationStatusEnum = append(slpParameterValidationStatusEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, slpParameterValidationStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this slp parameter validation status
func (m SlpParameterValidationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSlpParameterValidationStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}