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
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewIPAMServicesCreateParams creates a new IPAMServicesCreateParams object
// with the default values initialized.
func NewIPAMServicesCreateParams() *IPAMServicesCreateParams {
	var ()
	return &IPAMServicesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMServicesCreateParamsWithTimeout creates a new IPAMServicesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMServicesCreateParamsWithTimeout(timeout time.Duration) *IPAMServicesCreateParams {
	var ()
	return &IPAMServicesCreateParams{

		timeout: timeout,
	}
}

// NewIPAMServicesCreateParamsWithContext creates a new IPAMServicesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMServicesCreateParamsWithContext(ctx context.Context) *IPAMServicesCreateParams {
	var ()
	return &IPAMServicesCreateParams{

		Context: ctx,
	}
}

// NewIPAMServicesCreateParamsWithHTTPClient creates a new IPAMServicesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMServicesCreateParamsWithHTTPClient(client *http.Client) *IPAMServicesCreateParams {
	var ()
	return &IPAMServicesCreateParams{
		HTTPClient: client,
	}
}

/*IPAMServicesCreateParams contains all the parameters to send to the API endpoint
for the ipam services create operation typically these are written to a http.Request
*/
type IPAMServicesCreateParams struct {

	/*Data*/
	Data *models.Service

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam services create params
func (o *IPAMServicesCreateParams) WithTimeout(timeout time.Duration) *IPAMServicesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services create params
func (o *IPAMServicesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services create params
func (o *IPAMServicesCreateParams) WithContext(ctx context.Context) *IPAMServicesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services create params
func (o *IPAMServicesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services create params
func (o *IPAMServicesCreateParams) WithHTTPClient(client *http.Client) *IPAMServicesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services create params
func (o *IPAMServicesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services create params
func (o *IPAMServicesCreateParams) WithData(data *models.Service) *IPAMServicesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services create params
func (o *IPAMServicesCreateParams) SetData(data *models.Service) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMServicesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
