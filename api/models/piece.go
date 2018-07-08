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

// Piece piece
// swagger:model Piece
type Piece struct {

	// catalog
	Catalog string `json:"catalog,omitempty"`

	// composer
	Composer *Composer `json:"composer,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// key
	Key *Key `json:"key,omitempty"`

	// movements
	Movements []*Movement `json:"movements"`

	// number
	Number int64 `json:"number,omitempty"`

	// parts
	Parts []*Part `json:"parts"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this piece
func (m *Piece) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComposer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMovements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Piece) validateComposer(formats strfmt.Registry) error {

	if swag.IsZero(m.Composer) { // not required
		return nil
	}

	if m.Composer != nil {
		if err := m.Composer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("composer")
			}
			return err
		}
	}

	return nil
}

func (m *Piece) validateKey(formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {
		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("key")
			}
			return err
		}
	}

	return nil
}

func (m *Piece) validateMovements(formats strfmt.Registry) error {

	if swag.IsZero(m.Movements) { // not required
		return nil
	}

	for i := 0; i < len(m.Movements); i++ {
		if swag.IsZero(m.Movements[i]) { // not required
			continue
		}

		if m.Movements[i] != nil {
			if err := m.Movements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("movements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Piece) validateParts(formats strfmt.Registry) error {

	if swag.IsZero(m.Parts) { // not required
		return nil
	}

	for i := 0; i < len(m.Parts); i++ {
		if swag.IsZero(m.Parts[i]) { // not required
			continue
		}

		if m.Parts[i] != nil {
			if err := m.Parts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Piece) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Piece) UnmarshalBinary(b []byte) error {
	var res Piece
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
