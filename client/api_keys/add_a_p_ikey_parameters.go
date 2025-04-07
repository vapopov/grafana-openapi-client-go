// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// NewAddAPIkeyParams creates a new AddAPIkeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAPIkeyParams() *AddAPIkeyParams {
	return &AddAPIkeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAPIkeyParamsWithTimeout creates a new AddAPIkeyParams object
// with the ability to set a timeout on a request.
func NewAddAPIkeyParamsWithTimeout(timeout time.Duration) *AddAPIkeyParams {
	return &AddAPIkeyParams{
		timeout: timeout,
	}
}

// NewAddAPIkeyParamsWithContext creates a new AddAPIkeyParams object
// with the ability to set a context for a request.
func NewAddAPIkeyParamsWithContext(ctx context.Context) *AddAPIkeyParams {
	return &AddAPIkeyParams{
		Context: ctx,
	}
}

// NewAddAPIkeyParamsWithHTTPClient creates a new AddAPIkeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAPIkeyParamsWithHTTPClient(client *http.Client) *AddAPIkeyParams {
	return &AddAPIkeyParams{
		HTTPClient: client,
	}
}

/*
AddAPIkeyParams contains all the parameters to send to the API endpoint

	for the add a p ikey operation.

	Typically these are written to a http.Request.
*/
type AddAPIkeyParams struct {

	// Body.
	Body *models.AddAPIKeyCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add a p ikey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAPIkeyParams) WithDefaults() *AddAPIkeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add a p ikey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAPIkeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add a p ikey params
func (o *AddAPIkeyParams) WithTimeout(timeout time.Duration) *AddAPIkeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add a p ikey params
func (o *AddAPIkeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add a p ikey params
func (o *AddAPIkeyParams) WithContext(ctx context.Context) *AddAPIkeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add a p ikey params
func (o *AddAPIkeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add a p ikey params
func (o *AddAPIkeyParams) WithHTTPClient(client *http.Client) *AddAPIkeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add a p ikey params
func (o *AddAPIkeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add a p ikey params
func (o *AddAPIkeyParams) WithBody(body *models.AddAPIKeyCommand) *AddAPIkeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add a p ikey params
func (o *AddAPIkeyParams) SetBody(body *models.AddAPIKeyCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAPIkeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
