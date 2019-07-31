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

// NewIPAMPrefixesAvailablePrefixesCreateParams creates a new IPAMPrefixesAvailablePrefixesCreateParams object
// with the default values initialized.
func NewIPAMPrefixesAvailablePrefixesCreateParams() *IPAMPrefixesAvailablePrefixesCreateParams {
	var ()
	return &IPAMPrefixesAvailablePrefixesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMPrefixesAvailablePrefixesCreateParamsWithTimeout creates a new IPAMPrefixesAvailablePrefixesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMPrefixesAvailablePrefixesCreateParamsWithTimeout(timeout time.Duration) *IPAMPrefixesAvailablePrefixesCreateParams {
	var ()
	return &IPAMPrefixesAvailablePrefixesCreateParams{

		timeout: timeout,
	}
}

// NewIPAMPrefixesAvailablePrefixesCreateParamsWithContext creates a new IPAMPrefixesAvailablePrefixesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMPrefixesAvailablePrefixesCreateParamsWithContext(ctx context.Context) *IPAMPrefixesAvailablePrefixesCreateParams {
	var ()
	return &IPAMPrefixesAvailablePrefixesCreateParams{

		Context: ctx,
	}
}

// NewIPAMPrefixesAvailablePrefixesCreateParamsWithHTTPClient creates a new IPAMPrefixesAvailablePrefixesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMPrefixesAvailablePrefixesCreateParamsWithHTTPClient(client *http.Client) *IPAMPrefixesAvailablePrefixesCreateParams {
	var ()
	return &IPAMPrefixesAvailablePrefixesCreateParams{
		HTTPClient: client,
	}
}

/*IPAMPrefixesAvailablePrefixesCreateParams contains all the parameters to send to the API endpoint
for the ipam prefixes available prefixes create operation typically these are written to a http.Request
*/
type IPAMPrefixesAvailablePrefixesCreateParams struct {

	/*Data*/
	Data *models.WritablePrefix
	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) WithTimeout(timeout time.Duration) *IPAMPrefixesAvailablePrefixesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) WithContext(ctx context.Context) *IPAMPrefixesAvailablePrefixesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) WithHTTPClient(client *http.Client) *IPAMPrefixesAvailablePrefixesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) WithData(data *models.WritablePrefix) *IPAMPrefixesAvailablePrefixesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) WithID(id int64) *IPAMPrefixesAvailablePrefixesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available prefixes create params
func (o *IPAMPrefixesAvailablePrefixesCreateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMPrefixesAvailablePrefixesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
