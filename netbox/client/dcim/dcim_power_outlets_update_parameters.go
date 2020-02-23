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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewDcimPowerOutletsUpdateParams creates a new DcimPowerOutletsUpdateParams object
// with the default values initialized.
func NewDcimPowerOutletsUpdateParams() *DcimPowerOutletsUpdateParams {
	var ()
	return &DcimPowerOutletsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsUpdateParamsWithTimeout creates a new DcimPowerOutletsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletsUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsUpdateParams {
	var ()
	return &DcimPowerOutletsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletsUpdateParamsWithContext creates a new DcimPowerOutletsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletsUpdateParamsWithContext(ctx context.Context) *DcimPowerOutletsUpdateParams {
	var ()
	return &DcimPowerOutletsUpdateParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletsUpdateParamsWithHTTPClient creates a new DcimPowerOutletsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletsUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsUpdateParams {
	var ()
	return &DcimPowerOutletsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletsUpdateParams contains all the parameters to send to the API endpoint
for the dcim power outlets update operation typically these are written to a http.Request
*/
type DcimPowerOutletsUpdateParams struct {

	/*Data*/
	Data *models.WritablePowerOutlet
	/*ID
	  A unique integer value identifying this power outlet.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithContext(ctx context.Context) *DcimPowerOutletsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithData(data *models.WritablePowerOutlet) *DcimPowerOutletsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetData(data *models.WritablePowerOutlet) {
	o.Data = data
}

// WithID adds the id to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) WithID(id int64) *DcimPowerOutletsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets update params
func (o *DcimPowerOutletsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
