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
)

// NewRouteDeleteAlertRuleGroupParams creates a new RouteDeleteAlertRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteDeleteAlertRuleGroupParams() *RouteDeleteAlertRuleGroupParams {
	return &RouteDeleteAlertRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteDeleteAlertRuleGroupParamsWithTimeout creates a new RouteDeleteAlertRuleGroupParams object
// with the ability to set a timeout on a request.
func NewRouteDeleteAlertRuleGroupParamsWithTimeout(timeout time.Duration) *RouteDeleteAlertRuleGroupParams {
	return &RouteDeleteAlertRuleGroupParams{
		timeout: timeout,
	}
}

// NewRouteDeleteAlertRuleGroupParamsWithContext creates a new RouteDeleteAlertRuleGroupParams object
// with the ability to set a context for a request.
func NewRouteDeleteAlertRuleGroupParamsWithContext(ctx context.Context) *RouteDeleteAlertRuleGroupParams {
	return &RouteDeleteAlertRuleGroupParams{
		Context: ctx,
	}
}

// NewRouteDeleteAlertRuleGroupParamsWithHTTPClient creates a new RouteDeleteAlertRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteDeleteAlertRuleGroupParamsWithHTTPClient(client *http.Client) *RouteDeleteAlertRuleGroupParams {
	return &RouteDeleteAlertRuleGroupParams{
		HTTPClient: client,
	}
}

/*
RouteDeleteAlertRuleGroupParams contains all the parameters to send to the API endpoint

	for the route delete alert rule group operation.

	Typically these are written to a http.Request.
*/
type RouteDeleteAlertRuleGroupParams struct {

	// FolderUID.
	FolderUID string

	// Group.
	Group string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route delete alert rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteDeleteAlertRuleGroupParams) WithDefaults() *RouteDeleteAlertRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route delete alert rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteDeleteAlertRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) WithTimeout(timeout time.Duration) *RouteDeleteAlertRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) WithContext(ctx context.Context) *RouteDeleteAlertRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) WithHTTPClient(client *http.Client) *RouteDeleteAlertRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderUID adds the folderUID to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) WithFolderUID(folderUID string) *RouteDeleteAlertRuleGroupParams {
	o.SetFolderUID(folderUID)
	return o
}

// SetFolderUID adds the folderUid to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) SetFolderUID(folderUID string) {
	o.FolderUID = folderUID
}

// WithGroup adds the group to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) WithGroup(group string) *RouteDeleteAlertRuleGroupParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the route delete alert rule group params
func (o *RouteDeleteAlertRuleGroupParams) SetGroup(group string) {
	o.Group = group
}

// WriteToRequest writes these params to a swagger request
func (o *RouteDeleteAlertRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FolderUID
	if err := r.SetPathParam("FolderUID", o.FolderUID); err != nil {
		return err
	}

	// path param Group
	if err := r.SetPathParam("Group", o.Group); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
