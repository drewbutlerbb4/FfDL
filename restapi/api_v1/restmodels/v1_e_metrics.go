// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1EMetrics v1 e metrics
// swagger:model v1EMetrics

type V1EMetrics struct {

	// Repeated, order-dependent list of temporal keys
	// Example: {"iteration": 209}
	Etimes map[string]V1Any `json:"etimes,omitempty"`

	// Group label, such as test, train, or validate
	Grouplabel string `json:"grouplabel,omitempty"`

	// meta
	Meta *V1MetaInfo `json:"meta,omitempty"`

	// / {"cross_entropy": 0.4430539906024933,	"accuracy": 0.8999999761581421}
	Values map[string]V1Any `json:"values,omitempty"`
}

/* polymorph v1EMetrics etimes false */

/* polymorph v1EMetrics grouplabel false */

/* polymorph v1EMetrics meta false */

/* polymorph v1EMetrics values false */

// Validate validates this v1 e metrics
func (m *V1EMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEtimes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EMetrics) validateEtimes(formats strfmt.Registry) error {

	if swag.IsZero(m.Etimes) { // not required
		return nil
	}

	if err := validate.Required("etimes", "body", m.Etimes); err != nil {
		return err
	}

	return nil
}

func (m *V1EMetrics) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {

		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *V1EMetrics) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EMetrics) UnmarshalBinary(b []byte) error {
	var res V1EMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}