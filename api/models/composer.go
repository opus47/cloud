// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Composer composer
// swagger:model Composer
type Composer struct {

	// first
	First string `json:"first,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last
	Last string `json:"last,omitempty"`

	// middle
	Middle string `json:"middle,omitempty"`
}

// Validate validates this composer
func (m *Composer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Composer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Composer) UnmarshalBinary(b []byte) error {
	var res Composer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
