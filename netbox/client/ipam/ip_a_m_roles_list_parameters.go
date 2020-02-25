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
)

// NewIPAMRolesListParams creates a new IPAMRolesListParams object
// with the default values initialized.
func NewIPAMRolesListParams() *IPAMRolesListParams {
	var ()
	return &IPAMRolesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMRolesListParamsWithTimeout creates a new IPAMRolesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMRolesListParamsWithTimeout(timeout time.Duration) *IPAMRolesListParams {
	var ()
	return &IPAMRolesListParams{

		timeout: timeout,
	}
}

// NewIPAMRolesListParamsWithContext creates a new IPAMRolesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMRolesListParamsWithContext(ctx context.Context) *IPAMRolesListParams {
	var ()
	return &IPAMRolesListParams{

		Context: ctx,
	}
}

// NewIPAMRolesListParamsWithHTTPClient creates a new IPAMRolesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMRolesListParamsWithHTTPClient(client *http.Client) *IPAMRolesListParams {
	var ()
	return &IPAMRolesListParams{
		HTTPClient: client,
	}
}

/*IPAMRolesListParams contains all the parameters to send to the API endpoint
for the ipam roles list operation typically these are written to a http.Request
*/
type IPAMRolesListParams struct {

	/*ID*/
	ID *int64
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam roles list params
func (o *IPAMRolesListParams) WithTimeout(timeout time.Duration) *IPAMRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles list params
func (o *IPAMRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles list params
func (o *IPAMRolesListParams) WithContext(ctx context.Context) *IPAMRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles list params
func (o *IPAMRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles list params
func (o *IPAMRolesListParams) WithHTTPClient(client *http.Client) *IPAMRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles list params
func (o *IPAMRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam roles list params
func (o *IPAMRolesListParams) WithID(id *int64) *IPAMRolesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam roles list params
func (o *IPAMRolesListParams) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the ipam roles list params
func (o *IPAMRolesListParams) WithLimit(limit *int64) *IPAMRolesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam roles list params
func (o *IPAMRolesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam roles list params
func (o *IPAMRolesListParams) WithName(name *string) *IPAMRolesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam roles list params
func (o *IPAMRolesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam roles list params
func (o *IPAMRolesListParams) WithOffset(offset *int64) *IPAMRolesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam roles list params
func (o *IPAMRolesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam roles list params
func (o *IPAMRolesListParams) WithQ(q *string) *IPAMRolesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam roles list params
func (o *IPAMRolesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the ipam roles list params
func (o *IPAMRolesListParams) WithSlug(slug *string) *IPAMRolesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the ipam roles list params
func (o *IPAMRolesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
