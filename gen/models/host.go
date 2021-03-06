// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Host host
//
// swagger:model Host
type Host struct {

	// ID
	ID string `json:"ID,omitempty"`

	// IP
	IP string `json:"IP,omitempty"`

	// a (key, value) map
	MetaData map[string]interface{} `json:"MetaData,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// status
	Status bool `json:"Status,omitempty"`
}

// Validate validates this host
func (m *Host) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Host) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Host) UnmarshalBinary(b []byte) error {
	var res Host
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
