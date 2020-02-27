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

package circuits

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

	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewCircuitsProvidersCreateParams creates a new CircuitsProvidersCreateParams object
// with the default values initialized.
func NewCircuitsProvidersCreateParams() *CircuitsProvidersCreateParams {
	var ()
	return &CircuitsProvidersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersCreateParamsWithTimeout creates a new CircuitsProvidersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsProvidersCreateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersCreateParams {
	var ()
	return &CircuitsProvidersCreateParams{

		timeout: timeout,
	}
}

// NewCircuitsProvidersCreateParamsWithContext creates a new CircuitsProvidersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsProvidersCreateParamsWithContext(ctx context.Context) *CircuitsProvidersCreateParams {
	var ()
	return &CircuitsProvidersCreateParams{

		Context: ctx,
	}
}

// NewCircuitsProvidersCreateParamsWithHTTPClient creates a new CircuitsProvidersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsProvidersCreateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersCreateParams {
	var ()
	return &CircuitsProvidersCreateParams{
		HTTPClient: client,
	}
}

/*CircuitsProvidersCreateParams contains all the parameters to send to the API endpoint
for the circuits providers create operation typically these are written to a http.Request
*/
type CircuitsProvidersCreateParams struct {

	/*Data*/
	Data *models.Provider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithContext(ctx context.Context) *CircuitsProvidersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithData(data *models.Provider) *CircuitsProvidersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
