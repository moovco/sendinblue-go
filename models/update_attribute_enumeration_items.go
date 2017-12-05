// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateAttributeEnumerationItems update attribute enumeration items
// swagger:model updateAttributeEnumerationItems
type UpdateAttributeEnumerationItems struct {

	// Id of the value
	ID int64 `json:"id,omitempty"`

	// Label of the value
	Label string `json:"label,omitempty"`
}

// Validate validates this update attribute enumeration items
func (m *UpdateAttributeEnumerationItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAttributeEnumerationItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAttributeEnumerationItems) UnmarshalBinary(b []byte) error {
	var res UpdateAttributeEnumerationItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}