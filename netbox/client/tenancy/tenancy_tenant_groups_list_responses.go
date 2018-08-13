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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// TenancyTenantGroupsListReader is a Reader for the TenancyTenantGroupsList structure.
type TenancyTenantGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTenancyTenantGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTenancyTenantGroupsListOK creates a TenancyTenantGroupsListOK with default headers values
func NewTenancyTenantGroupsListOK() *TenancyTenantGroupsListOK {
	return &TenancyTenantGroupsListOK{}
}

/*TenancyTenantGroupsListOK handles this case with default header values.

TenancyTenantGroupsListOK tenancy tenant groups list o k
*/
type TenancyTenantGroupsListOK struct {
	Payload *models.TenancyTenantGroupsListOKBody
}

func (o *TenancyTenantGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/tenant-groups/][%d] tenancyTenantGroupsListOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenancyTenantGroupsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
