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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewDcimInterfacesCreateParams creates a new DcimInterfacesCreateParams object
// with the default values initialized.
func NewDcimInterfacesCreateParams() *DcimInterfacesCreateParams {
	var ()
	return &DcimInterfacesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesCreateParamsWithTimeout creates a new DcimInterfacesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfacesCreateParamsWithTimeout(timeout time.Duration) *DcimInterfacesCreateParams {
	var ()
	return &DcimInterfacesCreateParams{

		timeout: timeout,
	}
}

// NewDcimInterfacesCreateParamsWithContext creates a new DcimInterfacesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfacesCreateParamsWithContext(ctx context.Context) *DcimInterfacesCreateParams {
	var ()
	return &DcimInterfacesCreateParams{

		Context: ctx,
	}
}

// NewDcimInterfacesCreateParamsWithHTTPClient creates a new DcimInterfacesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfacesCreateParamsWithHTTPClient(client *http.Client) *DcimInterfacesCreateParams {
	var ()
	return &DcimInterfacesCreateParams{
		HTTPClient: client,
	}
}

/*DcimInterfacesCreateParams contains all the parameters to send to the API endpoint
for the dcim interfaces create operation typically these are written to a http.Request
*/
type DcimInterfacesCreateParams struct {

	/*Data*/
	Data *models.WritableDeviceInterface

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) WithTimeout(timeout time.Duration) *DcimInterfacesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) WithContext(ctx context.Context) *DcimInterfacesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) WithHTTPClient(client *http.Client) *DcimInterfacesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) WithData(data *models.WritableDeviceInterface) *DcimInterfacesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interfaces create params
func (o *DcimInterfacesCreateParams) SetData(data *models.WritableDeviceInterface) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
