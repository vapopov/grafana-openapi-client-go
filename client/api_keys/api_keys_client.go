// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new api keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddAPIkey(body *models.AddCommand, opts ...ClientOption) (*AddAPIkeyOK, error)
	AddAPIkeyWithParams(params *AddAPIkeyParams, opts ...ClientOption) (*AddAPIkeyOK, error)

	DeleteAPIkey(id int64, opts ...ClientOption) (*DeleteAPIkeyOK, error)
	DeleteAPIkeyWithParams(params *DeleteAPIkeyParams, opts ...ClientOption) (*DeleteAPIkeyOK, error)

	GetAPIkeys(params *GetAPIkeysParams, opts ...ClientOption) (*GetAPIkeysOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddAPIkey creates an API key

Will return details of the created API key.
*/
func (a *Client) AddAPIkey(body *models.AddCommand, opts ...ClientOption) (*AddAPIkeyOK, error) {
	params := NewAddAPIkeyParams().WithBody(body)
	return a.AddAPIkeyWithParams(params, opts...)
}

func (a *Client) AddAPIkeyWithParams(params *AddAPIkeyParams, opts ...ClientOption) (*AddAPIkeyOK, error) {
	if params == nil {
		params = NewAddAPIkeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addAPIkey",
		Method:             "POST",
		PathPattern:        "/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddAPIkeyReader{formats: a.formats},
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
	success, ok := result.(*AddAPIkeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addAPIkey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAPIkey deletes API key

Deletes an API key.
Deprecated. See: https://grafana.com/docs/grafana/next/administration/api-keys/#migrate-api-keys-to-grafana-service-accounts-using-the-api.
*/
func (a *Client) DeleteAPIkey(id int64, opts ...ClientOption) (*DeleteAPIkeyOK, error) {
	params := NewDeleteAPIkeyParams().WithID(id)
	return a.DeleteAPIkeyWithParams(params, opts...)
}

func (a *Client) DeleteAPIkeyWithParams(params *DeleteAPIkeyParams, opts ...ClientOption) (*DeleteAPIkeyOK, error) {
	if params == nil {
		params = NewDeleteAPIkeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAPIkey",
		Method:             "DELETE",
		PathPattern:        "/auth/keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAPIkeyReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPIkeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAPIkey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIkeys gets auth keys

Will return auth keys.

Deprecated: true.

Deprecated. Please use GET /api/serviceaccounts and GET /api/serviceaccounts/{id}/tokens instead
see https://grafana.com/docs/grafana/next/administration/api-keys/#migrate-api-keys-to-grafana-service-accounts-using-the-api.
*/

func (a *Client) GetAPIkeys(params *GetAPIkeysParams, opts ...ClientOption) (*GetAPIkeysOK, error) {
	if params == nil {
		params = NewGetAPIkeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAPIkeys",
		Method:             "GET",
		PathPattern:        "/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAPIkeysReader{formats: a.formats},
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
	success, ok := result.(*GetAPIkeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAPIkeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
