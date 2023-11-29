// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new dashboard versions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboard versions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDashboardVersionByID(dashboardVersionID int64, dashboardID int64, opts ...ClientOption) (*GetDashboardVersionByIDOK, error)
	GetDashboardVersionByIDWithParams(params *GetDashboardVersionByIDParams, opts ...ClientOption) (*GetDashboardVersionByIDOK, error)

	GetDashboardVersionByUID(uid string, dashboardVersionID int64, opts ...ClientOption) (*GetDashboardVersionByUIDOK, error)
	GetDashboardVersionByUIDWithParams(params *GetDashboardVersionByUIDParams, opts ...ClientOption) (*GetDashboardVersionByUIDOK, error)

	GetDashboardVersionsByID(dashboardID int64, opts ...ClientOption) (*GetDashboardVersionsByIDOK, error)
	GetDashboardVersionsByIDWithParams(params *GetDashboardVersionsByIDParams, opts ...ClientOption) (*GetDashboardVersionsByIDOK, error)

	GetDashboardVersionsByUID(params *GetDashboardVersionsByUIDParams, opts ...ClientOption) (*GetDashboardVersionsByUIDOK, error)

	RestoreDashboardVersionByID(dashboardID int64, body *models.RestoreDashboardVersionCommand, opts ...ClientOption) (*RestoreDashboardVersionByIDOK, error)
	RestoreDashboardVersionByIDWithParams(params *RestoreDashboardVersionByIDParams, opts ...ClientOption) (*RestoreDashboardVersionByIDOK, error)

	RestoreDashboardVersionByUID(uid string, body *models.RestoreDashboardVersionCommand, opts ...ClientOption) (*RestoreDashboardVersionByUIDOK, error)
	RestoreDashboardVersionByUIDWithParams(params *RestoreDashboardVersionByUIDParams, opts ...ClientOption) (*RestoreDashboardVersionByUIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDashboardVersionByID gets a specific dashboard version

Please refer to [updated API](#/dashboard_versions/getDashboardVersionByUID) instead
*/
func (a *Client) GetDashboardVersionByID(dashboardVersionID int64, dashboardID int64, opts ...ClientOption) (*GetDashboardVersionByIDOK, error) {
	params := NewGetDashboardVersionByIDParams().WithDashboardID(dashboardID).WithDashboardVersionID(dashboardVersionID)
	return a.GetDashboardVersionByIDWithParams(params, opts...)
}

func (a *Client) GetDashboardVersionByIDWithParams(params *GetDashboardVersionByIDParams, opts ...ClientOption) (*GetDashboardVersionByIDOK, error) {
	if params == nil {
		params = NewGetDashboardVersionByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardVersionByID",
		Method:             "GET",
		PathPattern:        "/dashboards/id/{DashboardID}/versions/{DashboardVersionID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardVersionByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardVersionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardVersionByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardVersionByUID gets a specific dashboard version using UID
*/
func (a *Client) GetDashboardVersionByUID(uid string, dashboardVersionID int64, opts ...ClientOption) (*GetDashboardVersionByUIDOK, error) {
	params := NewGetDashboardVersionByUIDParams().WithDashboardVersionID(dashboardVersionID).WithUID(uid)
	return a.GetDashboardVersionByUIDWithParams(params, opts...)
}

func (a *Client) GetDashboardVersionByUIDWithParams(params *GetDashboardVersionByUIDParams, opts ...ClientOption) (*GetDashboardVersionByUIDOK, error) {
	if params == nil {
		params = NewGetDashboardVersionByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardVersionByUID",
		Method:             "GET",
		PathPattern:        "/dashboards/uid/{uid}/versions/{DashboardVersionID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardVersionByUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardVersionByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardVersionByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardVersionsByID gets all existing versions for the dashboard

Please refer to [updated API](#/dashboard_versions/getDashboardVersionsByUID) instead
*/
func (a *Client) GetDashboardVersionsByID(dashboardID int64, opts ...ClientOption) (*GetDashboardVersionsByIDOK, error) {
	params := NewGetDashboardVersionsByIDParams().WithDashboardID(dashboardID)
	return a.GetDashboardVersionsByIDWithParams(params, opts...)
}

func (a *Client) GetDashboardVersionsByIDWithParams(params *GetDashboardVersionsByIDParams, opts ...ClientOption) (*GetDashboardVersionsByIDOK, error) {
	if params == nil {
		params = NewGetDashboardVersionsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardVersionsByID",
		Method:             "GET",
		PathPattern:        "/dashboards/id/{DashboardID}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardVersionsByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardVersionsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardVersionsByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboardVersionsByUID gets all existing versions for the dashboard using UID
*/

func (a *Client) GetDashboardVersionsByUID(params *GetDashboardVersionsByUIDParams, opts ...ClientOption) (*GetDashboardVersionsByUIDOK, error) {
	if params == nil {
		params = NewGetDashboardVersionsByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDashboardVersionsByUID",
		Method:             "GET",
		PathPattern:        "/dashboards/uid/{uid}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDashboardVersionsByUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardVersionsByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDashboardVersionsByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestoreDashboardVersionByID restores a dashboard to a given dashboard version

Please refer to [updated API](#/dashboard_versions/restoreDashboardVersionByUID) instead
*/
func (a *Client) RestoreDashboardVersionByID(dashboardID int64, body *models.RestoreDashboardVersionCommand, opts ...ClientOption) (*RestoreDashboardVersionByIDOK, error) {
	params := NewRestoreDashboardVersionByIDParams().WithBody(body).WithDashboardID(dashboardID)
	return a.RestoreDashboardVersionByIDWithParams(params, opts...)
}

func (a *Client) RestoreDashboardVersionByIDWithParams(params *RestoreDashboardVersionByIDParams, opts ...ClientOption) (*RestoreDashboardVersionByIDOK, error) {
	if params == nil {
		params = NewRestoreDashboardVersionByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restoreDashboardVersionByID",
		Method:             "POST",
		PathPattern:        "/dashboards/id/{DashboardID}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreDashboardVersionByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreDashboardVersionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreDashboardVersionByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestoreDashboardVersionByUID restores a dashboard to a given dashboard version using UID
*/
func (a *Client) RestoreDashboardVersionByUID(uid string, body *models.RestoreDashboardVersionCommand, opts ...ClientOption) (*RestoreDashboardVersionByUIDOK, error) {
	params := NewRestoreDashboardVersionByUIDParams().WithBody(body).WithUID(uid)
	return a.RestoreDashboardVersionByUIDWithParams(params, opts...)
}

func (a *Client) RestoreDashboardVersionByUIDWithParams(params *RestoreDashboardVersionByUIDParams, opts ...ClientOption) (*RestoreDashboardVersionByUIDOK, error) {
	if params == nil {
		params = NewRestoreDashboardVersionByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restoreDashboardVersionByUID",
		Method:             "POST",
		PathPattern:        "/dashboards/uid/{uid}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreDashboardVersionByUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreDashboardVersionByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreDashboardVersionByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
