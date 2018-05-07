// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// WritablePlatform writable platform
// swagger:model WritablePlatform
type WritablePlatform struct {

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Manufacturer
	//
	// Optionally limit this platform to devices of a certain manufacturer
	Manufacturer int64 `json:"manufacturer,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	Name *string `json:"name"`

	// NAPALM driver
	//
	// The name of the NAPALM driver to use when interacting with devices
	// Max Length: 50
	NapalmDriver string `json:"napalm_driver,omitempty"`

	// Legacy RPC client
	RPCClient string `json:"rpc_client,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`
}

// Validate validates this writable platform
func (m *WritablePlatform) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNapalmDriver(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRPCClient(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePlatform) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritablePlatform) validateNapalmDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.NapalmDriver) { // not required
		return nil
	}

	if err := validate.MaxLength("napalm_driver", "body", string(m.NapalmDriver), 50); err != nil {
		return err
	}

	return nil
}

var writablePlatformTypeRPCClientPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["juniper-junos","cisco-ios","opengear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePlatformTypeRPCClientPropEnum = append(writablePlatformTypeRPCClientPropEnum, v)
	}
}

const (

	// WritablePlatformRPCClientJuniperJunos captures enum value "juniper-junos"
	WritablePlatformRPCClientJuniperJunos string = "juniper-junos"

	// WritablePlatformRPCClientCiscoIos captures enum value "cisco-ios"
	WritablePlatformRPCClientCiscoIos string = "cisco-ios"

	// WritablePlatformRPCClientOpengear captures enum value "opengear"
	WritablePlatformRPCClientOpengear string = "opengear"
)

// prop value enum
func (m *WritablePlatform) validateRPCClientEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePlatformTypeRPCClientPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePlatform) validateRPCClient(formats strfmt.Registry) error {

	if swag.IsZero(m.RPCClient) { // not required
		return nil
	}

	// value enum
	if err := m.validateRPCClientEnum("rpc_client", "body", m.RPCClient); err != nil {
		return err
	}

	return nil
}

func (m *WritablePlatform) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 50); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePlatform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePlatform) UnmarshalBinary(b []byte) error {
	var res WritablePlatform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
