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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableVirtualMachineInterface writable virtual machine interface
// swagger:model WritableVirtualMachineInterface
type WritableVirtualMachineInterface struct {

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// Mode
	// Enum: [access tagged tagged-all]
	Mode string `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []int64 `json:"tagged_vlans"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 128gfc-sfp28 inifiband-sdr inifiband-ddr inifiband-qdr inifiband-fdr10 inifiband-fdr inifiband-edr inifiband-hdr inifiband-ndr inifiband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Type *string `json:"type"`

	// Untagged VLAN
	UntaggedVlan *int64 `json:"untagged_vlan,omitempty"`

	// Virtual machine
	VirtualMachine *int64 `json:"virtual_machine,omitempty"`
}

// Validate validates this writable virtual machine interface
func (m *WritableVirtualMachineInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVirtualMachineInterface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

var writableVirtualMachineInterfaceTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualMachineInterfaceTypeModePropEnum = append(writableVirtualMachineInterfaceTypeModePropEnum, v)
	}
}

const (

	// WritableVirtualMachineInterfaceModeAccess captures enum value "access"
	WritableVirtualMachineInterfaceModeAccess string = "access"

	// WritableVirtualMachineInterfaceModeTagged captures enum value "tagged"
	WritableVirtualMachineInterfaceModeTagged string = "tagged"

	// WritableVirtualMachineInterfaceModeTaggedAll captures enum value "tagged-all"
	WritableVirtualMachineInterfaceModeTaggedAll string = "tagged-all"
)

// prop value enum
func (m *WritableVirtualMachineInterface) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableVirtualMachineInterfaceTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualMachineInterface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualMachineInterface) validateTags(formats strfmt.Registry) error {

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

var writableVirtualMachineInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","128gfc-sfp28","inifiband-sdr","inifiband-ddr","inifiband-qdr","inifiband-fdr10","inifiband-fdr","inifiband-edr","inifiband-hdr","inifiband-ndr","inifiband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualMachineInterfaceTypeTypePropEnum = append(writableVirtualMachineInterfaceTypeTypePropEnum, v)
	}
}

const (

	// WritableVirtualMachineInterfaceTypeVirtual captures enum value "virtual"
	WritableVirtualMachineInterfaceTypeVirtual string = "virtual"

	// WritableVirtualMachineInterfaceTypeLag captures enum value "lag"
	WritableVirtualMachineInterfaceTypeLag string = "lag"

	// WritableVirtualMachineInterfaceTypeNr100baseTx captures enum value "100base-tx"
	WritableVirtualMachineInterfaceTypeNr100baseTx string = "100base-tx"

	// WritableVirtualMachineInterfaceTypeNr1000baseT captures enum value "1000base-t"
	WritableVirtualMachineInterfaceTypeNr1000baseT string = "1000base-t"

	// WritableVirtualMachineInterfaceTypeNr25gbaseT captures enum value "2.5gbase-t"
	WritableVirtualMachineInterfaceTypeNr25gbaseT string = "2.5gbase-t"

	// WritableVirtualMachineInterfaceTypeNr5gbaseT captures enum value "5gbase-t"
	WritableVirtualMachineInterfaceTypeNr5gbaseT string = "5gbase-t"

	// WritableVirtualMachineInterfaceTypeNr10gbaseT captures enum value "10gbase-t"
	WritableVirtualMachineInterfaceTypeNr10gbaseT string = "10gbase-t"

	// WritableVirtualMachineInterfaceTypeNr10gbaseCx4 captures enum value "10gbase-cx4"
	WritableVirtualMachineInterfaceTypeNr10gbaseCx4 string = "10gbase-cx4"

	// WritableVirtualMachineInterfaceTypeNr1000baseXGbic captures enum value "1000base-x-gbic"
	WritableVirtualMachineInterfaceTypeNr1000baseXGbic string = "1000base-x-gbic"

	// WritableVirtualMachineInterfaceTypeNr1000baseXSfp captures enum value "1000base-x-sfp"
	WritableVirtualMachineInterfaceTypeNr1000baseXSfp string = "1000base-x-sfp"

	// WritableVirtualMachineInterfaceTypeNr10gbaseXSfpp captures enum value "10gbase-x-sfpp"
	WritableVirtualMachineInterfaceTypeNr10gbaseXSfpp string = "10gbase-x-sfpp"

	// WritableVirtualMachineInterfaceTypeNr10gbaseXXfp captures enum value "10gbase-x-xfp"
	WritableVirtualMachineInterfaceTypeNr10gbaseXXfp string = "10gbase-x-xfp"

	// WritableVirtualMachineInterfaceTypeNr10gbaseXXenpak captures enum value "10gbase-x-xenpak"
	WritableVirtualMachineInterfaceTypeNr10gbaseXXenpak string = "10gbase-x-xenpak"

	// WritableVirtualMachineInterfaceTypeNr10gbaseXX2 captures enum value "10gbase-x-x2"
	WritableVirtualMachineInterfaceTypeNr10gbaseXX2 string = "10gbase-x-x2"

	// WritableVirtualMachineInterfaceTypeNr25gbaseXSfp28 captures enum value "25gbase-x-sfp28"
	WritableVirtualMachineInterfaceTypeNr25gbaseXSfp28 string = "25gbase-x-sfp28"

	// WritableVirtualMachineInterfaceTypeNr40gbaseXQsfpp captures enum value "40gbase-x-qsfpp"
	WritableVirtualMachineInterfaceTypeNr40gbaseXQsfpp string = "40gbase-x-qsfpp"

	// WritableVirtualMachineInterfaceTypeNr50gbaseXSfp28 captures enum value "50gbase-x-sfp28"
	WritableVirtualMachineInterfaceTypeNr50gbaseXSfp28 string = "50gbase-x-sfp28"

	// WritableVirtualMachineInterfaceTypeNr100gbaseXCfp captures enum value "100gbase-x-cfp"
	WritableVirtualMachineInterfaceTypeNr100gbaseXCfp string = "100gbase-x-cfp"

	// WritableVirtualMachineInterfaceTypeNr100gbaseXCfp2 captures enum value "100gbase-x-cfp2"
	WritableVirtualMachineInterfaceTypeNr100gbaseXCfp2 string = "100gbase-x-cfp2"

	// WritableVirtualMachineInterfaceTypeNr200gbaseXCfp2 captures enum value "200gbase-x-cfp2"
	WritableVirtualMachineInterfaceTypeNr200gbaseXCfp2 string = "200gbase-x-cfp2"

	// WritableVirtualMachineInterfaceTypeNr100gbaseXCfp4 captures enum value "100gbase-x-cfp4"
	WritableVirtualMachineInterfaceTypeNr100gbaseXCfp4 string = "100gbase-x-cfp4"

	// WritableVirtualMachineInterfaceTypeNr100gbaseXCpak captures enum value "100gbase-x-cpak"
	WritableVirtualMachineInterfaceTypeNr100gbaseXCpak string = "100gbase-x-cpak"

	// WritableVirtualMachineInterfaceTypeNr100gbaseXQsfp28 captures enum value "100gbase-x-qsfp28"
	WritableVirtualMachineInterfaceTypeNr100gbaseXQsfp28 string = "100gbase-x-qsfp28"

	// WritableVirtualMachineInterfaceTypeNr200gbaseXQsfp56 captures enum value "200gbase-x-qsfp56"
	WritableVirtualMachineInterfaceTypeNr200gbaseXQsfp56 string = "200gbase-x-qsfp56"

	// WritableVirtualMachineInterfaceTypeNr400gbaseXQsfpdd captures enum value "400gbase-x-qsfpdd"
	WritableVirtualMachineInterfaceTypeNr400gbaseXQsfpdd string = "400gbase-x-qsfpdd"

	// WritableVirtualMachineInterfaceTypeNr400gbaseXOsfp captures enum value "400gbase-x-osfp"
	WritableVirtualMachineInterfaceTypeNr400gbaseXOsfp string = "400gbase-x-osfp"

	// WritableVirtualMachineInterfaceTypeIeee80211a captures enum value "ieee802.11a"
	WritableVirtualMachineInterfaceTypeIeee80211a string = "ieee802.11a"

	// WritableVirtualMachineInterfaceTypeIeee80211g captures enum value "ieee802.11g"
	WritableVirtualMachineInterfaceTypeIeee80211g string = "ieee802.11g"

	// WritableVirtualMachineInterfaceTypeIeee80211n captures enum value "ieee802.11n"
	WritableVirtualMachineInterfaceTypeIeee80211n string = "ieee802.11n"

	// WritableVirtualMachineInterfaceTypeIeee80211ac captures enum value "ieee802.11ac"
	WritableVirtualMachineInterfaceTypeIeee80211ac string = "ieee802.11ac"

	// WritableVirtualMachineInterfaceTypeIeee80211ad captures enum value "ieee802.11ad"
	WritableVirtualMachineInterfaceTypeIeee80211ad string = "ieee802.11ad"

	// WritableVirtualMachineInterfaceTypeIeee80211ax captures enum value "ieee802.11ax"
	WritableVirtualMachineInterfaceTypeIeee80211ax string = "ieee802.11ax"

	// WritableVirtualMachineInterfaceTypeGsm captures enum value "gsm"
	WritableVirtualMachineInterfaceTypeGsm string = "gsm"

	// WritableVirtualMachineInterfaceTypeCdma captures enum value "cdma"
	WritableVirtualMachineInterfaceTypeCdma string = "cdma"

	// WritableVirtualMachineInterfaceTypeLte captures enum value "lte"
	WritableVirtualMachineInterfaceTypeLte string = "lte"

	// WritableVirtualMachineInterfaceTypeSonetOc3 captures enum value "sonet-oc3"
	WritableVirtualMachineInterfaceTypeSonetOc3 string = "sonet-oc3"

	// WritableVirtualMachineInterfaceTypeSonetOc12 captures enum value "sonet-oc12"
	WritableVirtualMachineInterfaceTypeSonetOc12 string = "sonet-oc12"

	// WritableVirtualMachineInterfaceTypeSonetOc48 captures enum value "sonet-oc48"
	WritableVirtualMachineInterfaceTypeSonetOc48 string = "sonet-oc48"

	// WritableVirtualMachineInterfaceTypeSonetOc192 captures enum value "sonet-oc192"
	WritableVirtualMachineInterfaceTypeSonetOc192 string = "sonet-oc192"

	// WritableVirtualMachineInterfaceTypeSonetOc768 captures enum value "sonet-oc768"
	WritableVirtualMachineInterfaceTypeSonetOc768 string = "sonet-oc768"

	// WritableVirtualMachineInterfaceTypeSonetOc1920 captures enum value "sonet-oc1920"
	WritableVirtualMachineInterfaceTypeSonetOc1920 string = "sonet-oc1920"

	// WritableVirtualMachineInterfaceTypeSonetOc3840 captures enum value "sonet-oc3840"
	WritableVirtualMachineInterfaceTypeSonetOc3840 string = "sonet-oc3840"

	// WritableVirtualMachineInterfaceTypeNr1gfcSfp captures enum value "1gfc-sfp"
	WritableVirtualMachineInterfaceTypeNr1gfcSfp string = "1gfc-sfp"

	// WritableVirtualMachineInterfaceTypeNr2gfcSfp captures enum value "2gfc-sfp"
	WritableVirtualMachineInterfaceTypeNr2gfcSfp string = "2gfc-sfp"

	// WritableVirtualMachineInterfaceTypeNr4gfcSfp captures enum value "4gfc-sfp"
	WritableVirtualMachineInterfaceTypeNr4gfcSfp string = "4gfc-sfp"

	// WritableVirtualMachineInterfaceTypeNr8gfcSfpp captures enum value "8gfc-sfpp"
	WritableVirtualMachineInterfaceTypeNr8gfcSfpp string = "8gfc-sfpp"

	// WritableVirtualMachineInterfaceTypeNr16gfcSfpp captures enum value "16gfc-sfpp"
	WritableVirtualMachineInterfaceTypeNr16gfcSfpp string = "16gfc-sfpp"

	// WritableVirtualMachineInterfaceTypeNr32gfcSfp28 captures enum value "32gfc-sfp28"
	WritableVirtualMachineInterfaceTypeNr32gfcSfp28 string = "32gfc-sfp28"

	// WritableVirtualMachineInterfaceTypeNr128gfcSfp28 captures enum value "128gfc-sfp28"
	WritableVirtualMachineInterfaceTypeNr128gfcSfp28 string = "128gfc-sfp28"

	// WritableVirtualMachineInterfaceTypeInifibandSdr captures enum value "inifiband-sdr"
	WritableVirtualMachineInterfaceTypeInifibandSdr string = "inifiband-sdr"

	// WritableVirtualMachineInterfaceTypeInifibandDdr captures enum value "inifiband-ddr"
	WritableVirtualMachineInterfaceTypeInifibandDdr string = "inifiband-ddr"

	// WritableVirtualMachineInterfaceTypeInifibandQdr captures enum value "inifiband-qdr"
	WritableVirtualMachineInterfaceTypeInifibandQdr string = "inifiband-qdr"

	// WritableVirtualMachineInterfaceTypeInifibandFdr10 captures enum value "inifiband-fdr10"
	WritableVirtualMachineInterfaceTypeInifibandFdr10 string = "inifiband-fdr10"

	// WritableVirtualMachineInterfaceTypeInifibandFdr captures enum value "inifiband-fdr"
	WritableVirtualMachineInterfaceTypeInifibandFdr string = "inifiband-fdr"

	// WritableVirtualMachineInterfaceTypeInifibandEdr captures enum value "inifiband-edr"
	WritableVirtualMachineInterfaceTypeInifibandEdr string = "inifiband-edr"

	// WritableVirtualMachineInterfaceTypeInifibandHdr captures enum value "inifiband-hdr"
	WritableVirtualMachineInterfaceTypeInifibandHdr string = "inifiband-hdr"

	// WritableVirtualMachineInterfaceTypeInifibandNdr captures enum value "inifiband-ndr"
	WritableVirtualMachineInterfaceTypeInifibandNdr string = "inifiband-ndr"

	// WritableVirtualMachineInterfaceTypeInifibandXdr captures enum value "inifiband-xdr"
	WritableVirtualMachineInterfaceTypeInifibandXdr string = "inifiband-xdr"

	// WritableVirtualMachineInterfaceTypeT1 captures enum value "t1"
	WritableVirtualMachineInterfaceTypeT1 string = "t1"

	// WritableVirtualMachineInterfaceTypeE1 captures enum value "e1"
	WritableVirtualMachineInterfaceTypeE1 string = "e1"

	// WritableVirtualMachineInterfaceTypeT3 captures enum value "t3"
	WritableVirtualMachineInterfaceTypeT3 string = "t3"

	// WritableVirtualMachineInterfaceTypeE3 captures enum value "e3"
	WritableVirtualMachineInterfaceTypeE3 string = "e3"

	// WritableVirtualMachineInterfaceTypeCiscoStackwise captures enum value "cisco-stackwise"
	WritableVirtualMachineInterfaceTypeCiscoStackwise string = "cisco-stackwise"

	// WritableVirtualMachineInterfaceTypeCiscoStackwisePlus captures enum value "cisco-stackwise-plus"
	WritableVirtualMachineInterfaceTypeCiscoStackwisePlus string = "cisco-stackwise-plus"

	// WritableVirtualMachineInterfaceTypeCiscoFlexstack captures enum value "cisco-flexstack"
	WritableVirtualMachineInterfaceTypeCiscoFlexstack string = "cisco-flexstack"

	// WritableVirtualMachineInterfaceTypeCiscoFlexstackPlus captures enum value "cisco-flexstack-plus"
	WritableVirtualMachineInterfaceTypeCiscoFlexstackPlus string = "cisco-flexstack-plus"

	// WritableVirtualMachineInterfaceTypeJuniperVcp captures enum value "juniper-vcp"
	WritableVirtualMachineInterfaceTypeJuniperVcp string = "juniper-vcp"

	// WritableVirtualMachineInterfaceTypeExtremeSummitstack captures enum value "extreme-summitstack"
	WritableVirtualMachineInterfaceTypeExtremeSummitstack string = "extreme-summitstack"

	// WritableVirtualMachineInterfaceTypeExtremeSummitstack128 captures enum value "extreme-summitstack-128"
	WritableVirtualMachineInterfaceTypeExtremeSummitstack128 string = "extreme-summitstack-128"

	// WritableVirtualMachineInterfaceTypeExtremeSummitstack256 captures enum value "extreme-summitstack-256"
	WritableVirtualMachineInterfaceTypeExtremeSummitstack256 string = "extreme-summitstack-256"

	// WritableVirtualMachineInterfaceTypeExtremeSummitstack512 captures enum value "extreme-summitstack-512"
	WritableVirtualMachineInterfaceTypeExtremeSummitstack512 string = "extreme-summitstack-512"

	// WritableVirtualMachineInterfaceTypeOther captures enum value "other"
	WritableVirtualMachineInterfaceTypeOther string = "other"
)

// prop value enum
func (m *WritableVirtualMachineInterface) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableVirtualMachineInterfaceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualMachineInterface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVirtualMachineInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVirtualMachineInterface) UnmarshalBinary(b []byte) error {
	var res WritableVirtualMachineInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
