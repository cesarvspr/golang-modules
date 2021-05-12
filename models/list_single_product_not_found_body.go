// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListSingleProductNotFoundBody ListSingleProductNotFoundBody list single product not found body
//
// swagger:model ListSingleProductNotFoundBody
type ListSingleProductNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this list single product not found body
func (m *ListSingleProductNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list single product not found body based on context it is used
func (m *ListSingleProductNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListSingleProductNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListSingleProductNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListSingleProductNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
