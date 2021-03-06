// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListenerPolicyTargetReference listener policy target reference
// swagger:model ListenerPolicyTargetReference
type ListenerPolicyTargetReference struct {

	// The pool's canonical URL.
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The http status code in the redirect response.
	// Enum: [301 302 303 307 308]
	HTTPStatusCode int64 `json:"http_status_code,omitempty"`

	// The unique identifier for this load balancer pool
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this load balancer pool
	// Pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$
	Name string `json:"name,omitempty"`

	// The redirect target URL.
	URL string `json:"url,omitempty"`
}

// Validate validates this listener policy target reference
func (m *ListenerPolicyTargetReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListenerPolicyTargetReference) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

var listenerPolicyTargetReferenceTypeHTTPStatusCodePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[301,302,303,307,308]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerPolicyTargetReferenceTypeHTTPStatusCodePropEnum = append(listenerPolicyTargetReferenceTypeHTTPStatusCodePropEnum, v)
	}
}

// prop value enum
func (m *ListenerPolicyTargetReference) validateHTTPStatusCodeEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, listenerPolicyTargetReferenceTypeHTTPStatusCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ListenerPolicyTargetReference) validateHTTPStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPStatusCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPStatusCodeEnum("http_status_code", "body", m.HTTPStatusCode); err != nil {
		return err
	}

	return nil
}

func (m *ListenerPolicyTargetReference) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ListenerPolicyTargetReference) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerPolicyTargetReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerPolicyTargetReference) UnmarshalBinary(b []byte) error {
	var res ListenerPolicyTargetReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
