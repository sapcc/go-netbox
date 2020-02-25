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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rack rack
// swagger:model Rack
type Rack struct {

	// Asset tag
	//
	// A unique tag used to identify this rack
	// Max Length: 50
	AssetTag *string `json:"asset_tag,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Descending units
	//
	// Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty"`

	// Device count
	// Read Only: true
	DeviceCount int64 `json:"device_count,omitempty"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// Facility ID
	// Max Length: 50
	FacilityID *string `json:"facility_id,omitempty"`

	// group
	Group *NestedRackGroup `json:"group,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Outer depth
	// Maximum: 32767
	// Minimum: 0
	OuterDepth *int64 `json:"outer_depth,omitempty"`

	// outer unit
	OuterUnit *RackOuterUnit `json:"outer_unit,omitempty"`

	// Outer width
	// Maximum: 32767
	// Minimum: 0
	OuterWidth *int64 `json:"outer_width,omitempty"`

	// Powerfeed count
	// Read Only: true
	PowerfeedCount int64 `json:"powerfeed_count,omitempty"`

	// role
	Role *NestedRackRole `json:"role,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// site
	// Required: true
	Site *NestedSite `json:"site"`

	// status
	Status *RackStatus `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// type
	Type *RackType `json:"type,omitempty"`

	// Height (U)
	// Maximum: 100
	// Minimum: 1
	UHeight int64 `json:"u_height,omitempty"`

	// width
	Width *RackWidth `json:"width,omitempty"`
}

// Validate validates this rack
func (m *Rack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOuterWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rack) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(*m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateFacilityID(formats strfmt.Registry) error {

	if swag.IsZero(m.FacilityID) { // not required
		return nil
	}

	if err := validate.MaxLength("facility_id", "body", string(*m.FacilityID), 50); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateName(formats strfmt.Registry) error {

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

func (m *Rack) validateOuterDepth(formats strfmt.Registry) error {

	if swag.IsZero(m.OuterDepth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_depth", "body", int64(*m.OuterDepth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_depth", "body", int64(*m.OuterDepth), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateOuterUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.OuterUnit) { // not required
		return nil
	}

	if m.OuterUnit != nil {
		if err := m.OuterUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outer_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateOuterWidth(formats strfmt.Registry) error {

	if swag.IsZero(m.OuterWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("outer_width", "body", int64(*m.OuterWidth), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("outer_width", "body", int64(*m.OuterWidth), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *Rack) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *Rack) validateUHeight(formats strfmt.Registry) error {

	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", int64(m.UHeight), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", int64(m.UHeight), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Rack) validateWidth(formats strfmt.Registry) error {

	if swag.IsZero(m.Width) { // not required
		return nil
	}

	if m.Width != nil {
		if err := m.Width.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("width")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rack) UnmarshalBinary(b []byte) error {
	var res Rack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackOuterUnit Outer unit
// swagger:model RackOuterUnit
type RackOuterUnit struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this rack outer unit
func (m *RackOuterUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackOuterUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("outer_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *RackOuterUnit) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("outer_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackOuterUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackOuterUnit) UnmarshalBinary(b []byte) error {
	var res RackOuterUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackStatus Status
// swagger:model RackStatus
type RackStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this rack status
func (m *RackStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *RackStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackStatus) UnmarshalBinary(b []byte) error {
	var res RackStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackType Type
// swagger:model RackType
type RackType struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this rack type
func (m *RackType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *RackType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackType) UnmarshalBinary(b []byte) error {
	var res RackType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackWidth Width
// swagger:model RackWidth
type RackWidth struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this rack width
func (m *RackWidth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackWidth) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("width"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *RackWidth) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("width"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackWidth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackWidth) UnmarshalBinary(b []byte) error {
	var res RackWidth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
