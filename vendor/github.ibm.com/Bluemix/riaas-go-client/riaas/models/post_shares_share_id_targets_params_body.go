// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostSharesShareIDTargetsParamsBody post shares share Id targets params body
// swagger:model postSharesShareIdTargetsParamsBody
type PostSharesShareIDTargetsParamsBody struct {

	// security groups
	SecurityGroups []*PostSharesShareIDTargetsParamsBodySecurityGroupsItems `json:"Security_groups"`

	// A single IPv4 or IPv6 address.
	Address string `json:"address,omitempty"`

	// The user-defined name for this file share mount target
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup *PostSharesShareIDTargetsParamsBodyResourceGroup `json:"resource_group,omitempty"`

	// subnet
	Subnet *PostSharesShareIDTargetsParamsBodySubnet `json:"subnet,omitempty"`

	// zone
	Zone *PostSharesShareIDTargetsParamsBodyZone `json:"zone,omitempty"`
}

// Validate validates this post shares share Id targets params body
func (m *PostSharesShareIDTargetsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSharesShareIDTargetsParamsBody) validateSecurityGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroups); i++ {
		if swag.IsZero(m.SecurityGroups[i]) { // not required
			continue
		}

		if m.SecurityGroups[i] != nil {
			if err := m.SecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Security_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostSharesShareIDTargetsParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostSharesShareIDTargetsParamsBody) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

func (m *PostSharesShareIDTargetsParamsBody) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

func (m *PostSharesShareIDTargetsParamsBody) validateZone(formats strfmt.Registry) error {

	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSharesShareIDTargetsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSharesShareIDTargetsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostSharesShareIDTargetsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}