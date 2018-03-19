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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewExtrasExportTemplatesPartialUpdateParams creates a new ExtrasExportTemplatesPartialUpdateParams object
// with the default values initialized.
func NewExtrasExportTemplatesPartialUpdateParams() *ExtrasExportTemplatesPartialUpdateParams {
	var ()
	return &ExtrasExportTemplatesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesPartialUpdateParamsWithTimeout creates a new ExtrasExportTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasExportTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesPartialUpdateParams {
	var ()
	return &ExtrasExportTemplatesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasExportTemplatesPartialUpdateParamsWithContext creates a new ExtrasExportTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasExportTemplatesPartialUpdateParamsWithContext(ctx context.Context) *ExtrasExportTemplatesPartialUpdateParams {
	var ()
	return &ExtrasExportTemplatesPartialUpdateParams{

		Context: ctx,
	}
}

// NewExtrasExportTemplatesPartialUpdateParamsWithHTTPClient creates a new ExtrasExportTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasExportTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesPartialUpdateParams {
	var ()
	return &ExtrasExportTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasExportTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
for the extras export templates partial update operation typically these are written to a http.Request
*/
type ExtrasExportTemplatesPartialUpdateParams struct {

	/*Data*/
	Data *models.ExportTemplate
	/*ID
	  A unique integer value identifying this export template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) WithContext(ctx context.Context) *ExtrasExportTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) WithData(data *models.ExportTemplate) *ExtrasExportTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) SetData(data *models.ExportTemplate) {
	o.Data = data
}

// WithID adds the id to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) WithID(id int64) *ExtrasExportTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates partial update params
func (o *ExtrasExportTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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