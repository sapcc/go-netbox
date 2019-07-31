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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewIPAMVlanGroupsUpdateParams creates a new IPAMVlanGroupsUpdateParams object
// with the default values initialized.
func NewIPAMVlanGroupsUpdateParams() *IPAMVlanGroupsUpdateParams {
	var ()
	return &IPAMVlanGroupsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVlanGroupsUpdateParamsWithTimeout creates a new IPAMVlanGroupsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVlanGroupsUpdateParamsWithTimeout(timeout time.Duration) *IPAMVlanGroupsUpdateParams {
	var ()
	return &IPAMVlanGroupsUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMVlanGroupsUpdateParamsWithContext creates a new IPAMVlanGroupsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVlanGroupsUpdateParamsWithContext(ctx context.Context) *IPAMVlanGroupsUpdateParams {
	var ()
	return &IPAMVlanGroupsUpdateParams{

		Context: ctx,
	}
}

// NewIPAMVlanGroupsUpdateParamsWithHTTPClient creates a new IPAMVlanGroupsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVlanGroupsUpdateParamsWithHTTPClient(client *http.Client) *IPAMVlanGroupsUpdateParams {
	var ()
	return &IPAMVlanGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMVlanGroupsUpdateParams contains all the parameters to send to the API endpoint
for the ipam vlan groups update operation typically these are written to a http.Request
*/
type IPAMVlanGroupsUpdateParams struct {

	/*Data*/
	Data *models.WritableVLANGroup
	/*ID
	  A unique integer value identifying this VLAN group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) WithTimeout(timeout time.Duration) *IPAMVlanGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) WithContext(ctx context.Context) *IPAMVlanGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) WithHTTPClient(client *http.Client) *IPAMVlanGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) WithData(data *models.WritableVLANGroup) *IPAMVlanGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) SetData(data *models.WritableVLANGroup) {
	o.Data = data
}

// WithID adds the id to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) WithID(id int64) *IPAMVlanGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups update params
func (o *IPAMVlanGroupsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVlanGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
