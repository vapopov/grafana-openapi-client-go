// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewRoutePutMuteTimingParams creates a new RoutePutMuteTimingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRoutePutMuteTimingParams() *RoutePutMuteTimingParams {
	return &RoutePutMuteTimingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRoutePutMuteTimingParamsWithTimeout creates a new RoutePutMuteTimingParams object
// with the ability to set a timeout on a request.
func NewRoutePutMuteTimingParamsWithTimeout(timeout time.Duration) *RoutePutMuteTimingParams {
	return &RoutePutMuteTimingParams{
		timeout: timeout,
	}
}

// NewRoutePutMuteTimingParamsWithContext creates a new RoutePutMuteTimingParams object
// with the ability to set a context for a request.
func NewRoutePutMuteTimingParamsWithContext(ctx context.Context) *RoutePutMuteTimingParams {
	return &RoutePutMuteTimingParams{
		Context: ctx,
	}
}

// NewRoutePutMuteTimingParamsWithHTTPClient creates a new RoutePutMuteTimingParams object
// with the ability to set a custom HTTPClient for a request.
func NewRoutePutMuteTimingParamsWithHTTPClient(client *http.Client) *RoutePutMuteTimingParams {
	return &RoutePutMuteTimingParams{
		HTTPClient: client,
	}
}

/*
RoutePutMuteTimingParams contains all the parameters to send to the API endpoint

	for the route put mute timing operation.

	Typically these are written to a http.Request.
*/
type RoutePutMuteTimingParams struct {

	// Body.
	Body *models.MuteTimeInterval

	// XDisableProvenance.
	XDisableProvenance *string

	/* Name.

	   Mute timing name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route put mute timing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoutePutMuteTimingParams) WithDefaults() *RoutePutMuteTimingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route put mute timing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoutePutMuteTimingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route put mute timing params
func (o *RoutePutMuteTimingParams) WithTimeout(timeout time.Duration) *RoutePutMuteTimingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route put mute timing params
func (o *RoutePutMuteTimingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route put mute timing params
func (o *RoutePutMuteTimingParams) WithContext(ctx context.Context) *RoutePutMuteTimingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route put mute timing params
func (o *RoutePutMuteTimingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route put mute timing params
func (o *RoutePutMuteTimingParams) WithHTTPClient(client *http.Client) *RoutePutMuteTimingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route put mute timing params
func (o *RoutePutMuteTimingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the route put mute timing params
func (o *RoutePutMuteTimingParams) WithBody(body *models.MuteTimeInterval) *RoutePutMuteTimingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the route put mute timing params
func (o *RoutePutMuteTimingParams) SetBody(body *models.MuteTimeInterval) {
	o.Body = body
}

// WithXDisableProvenance adds the xDisableProvenance to the route put mute timing params
func (o *RoutePutMuteTimingParams) WithXDisableProvenance(xDisableProvenance *string) *RoutePutMuteTimingParams {
	o.SetXDisableProvenance(xDisableProvenance)
	return o
}

// SetXDisableProvenance adds the xDisableProvenance to the route put mute timing params
func (o *RoutePutMuteTimingParams) SetXDisableProvenance(xDisableProvenance *string) {
	o.XDisableProvenance = xDisableProvenance
}

// WithName adds the name to the route put mute timing params
func (o *RoutePutMuteTimingParams) WithName(name string) *RoutePutMuteTimingParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the route put mute timing params
func (o *RoutePutMuteTimingParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RoutePutMuteTimingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.XDisableProvenance != nil {

		// header param X-Disable-Provenance
		if err := r.SetHeaderParam("X-Disable-Provenance", *o.XDisableProvenance); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
