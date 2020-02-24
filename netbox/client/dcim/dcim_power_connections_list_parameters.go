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
)

// NewDcimPowerConnectionsListParams creates a new DcimPowerConnectionsListParams object
// with the default values initialized.
func NewDcimPowerConnectionsListParams() *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerConnectionsListParamsWithTimeout creates a new DcimPowerConnectionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerConnectionsListParamsWithTimeout(timeout time.Duration) *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{

		timeout: timeout,
	}
}

// NewDcimPowerConnectionsListParamsWithContext creates a new DcimPowerConnectionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerConnectionsListParamsWithContext(ctx context.Context) *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{

		Context: ctx,
	}
}

// NewDcimPowerConnectionsListParamsWithHTTPClient creates a new DcimPowerConnectionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerConnectionsListParamsWithHTTPClient(client *http.Client) *DcimPowerConnectionsListParams {
	var ()
	return &DcimPowerConnectionsListParams{
		HTTPClient: client,
	}
}

/*DcimPowerConnectionsListParams contains all the parameters to send to the API endpoint
for the dcim power connections list operation typically these are written to a http.Request
*/
type DcimPowerConnectionsListParams struct {

	/*ConnectionStatus*/
	ConnectionStatus *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *int64
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
	/*Site*/
	Site *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithTimeout(timeout time.Duration) *DcimPowerConnectionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithContext(ctx context.Context) *DcimPowerConnectionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithHTTPClient(client *http.Client) *DcimPowerConnectionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionStatus adds the connectionStatus to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithConnectionStatus(connectionStatus *string) *DcimPowerConnectionsListParams {
	o.SetConnectionStatus(connectionStatus)
	return o
}

// SetConnectionStatus adds the connectionStatus to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetConnectionStatus(connectionStatus *string) {
	o.ConnectionStatus = connectionStatus
}

// WithDevice adds the device to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithDevice(device *string) *DcimPowerConnectionsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithDeviceID(deviceID *int64) *DcimPowerConnectionsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetDeviceID(deviceID *int64) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithLimit(limit *int64) *DcimPowerConnectionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithName(name *string) *DcimPowerConnectionsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithOffset(offset *int64) *DcimPowerConnectionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSite adds the site to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) WithSite(site *string) *DcimPowerConnectionsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power connections list params
func (o *DcimPowerConnectionsListParams) SetSite(site *string) {
	o.Site = site
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerConnectionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectionStatus != nil {

		// query param connection_status
		var qrConnectionStatus string
		if o.ConnectionStatus != nil {
			qrConnectionStatus = *o.ConnectionStatus
		}
		qConnectionStatus := qrConnectionStatus
		if qConnectionStatus != "" {
			if err := r.SetQueryParam("connection_status", qConnectionStatus); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID int64
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := swag.FormatInt64(qrDeviceID)
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
