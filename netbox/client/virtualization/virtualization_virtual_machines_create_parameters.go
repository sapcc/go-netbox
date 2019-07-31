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

package virtualization

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

// NewVirtualizationVirtualMachinesCreateParams creates a new VirtualizationVirtualMachinesCreateParams object
// with the default values initialized.
func NewVirtualizationVirtualMachinesCreateParams() *VirtualizationVirtualMachinesCreateParams {
	var ()
	return &VirtualizationVirtualMachinesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesCreateParamsWithTimeout creates a new VirtualizationVirtualMachinesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationVirtualMachinesCreateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesCreateParams {
	var ()
	return &VirtualizationVirtualMachinesCreateParams{

		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesCreateParamsWithContext creates a new VirtualizationVirtualMachinesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationVirtualMachinesCreateParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesCreateParams {
	var ()
	return &VirtualizationVirtualMachinesCreateParams{

		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesCreateParamsWithHTTPClient creates a new VirtualizationVirtualMachinesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationVirtualMachinesCreateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesCreateParams {
	var ()
	return &VirtualizationVirtualMachinesCreateParams{
		HTTPClient: client,
	}
}

/*VirtualizationVirtualMachinesCreateParams contains all the parameters to send to the API endpoint
for the virtualization virtual machines create operation typically these are written to a http.Request
*/
type VirtualizationVirtualMachinesCreateParams struct {

	/*Data*/
	Data *models.WritableVirtualMachine

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) WithData(data *models.WritableVirtualMachine) *VirtualizationVirtualMachinesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual machines create params
func (o *VirtualizationVirtualMachinesCreateParams) SetData(data *models.WritableVirtualMachine) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
