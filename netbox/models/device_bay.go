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
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceBay device bay
// swagger:model DeviceBay
type DeviceBay struct {

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// installed device
	InstalledDevice *NestedDevice `json:"installed_device,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Tags
	Tags []string `json:"tags"`
}

// Validate validates this device bay
func (m *DeviceBay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstalledDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceBay) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {

		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}

	}

	return nil
}

func (m *DeviceBay) validateInstalledDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.InstalledDevice) { // not required
		return nil
	}

	if m.InstalledDevice != nil {

		if err := m.InstalledDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installed_device")
			}
			return err
		}

	}

	return nil
}

func (m *DeviceBay) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceBay) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceBay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceBay) UnmarshalBinary(b []byte) error {
	var res DeviceBay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
