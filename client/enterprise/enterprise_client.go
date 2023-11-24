// Code generated by go-swagger; DO NOT EDIT.

package enterprise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new enterprise API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for enterprise API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SearchResult(params *SearchResultParams, opts ...ClientOption) (*SearchResultOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	SearchResult debugs permissions

	Returns the result of the search through access-control role assignments.

You need to have a permission with action `teams.roles:read` on scope `teams:*`
and a permission with action `users.roles:read` on scope `users:*`.
*/
func (a *Client) SearchResult(params *SearchResultParams, opts ...ClientOption) (*SearchResultOK, error) {
	if params == nil {
		params = NewSearchResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchResult",
		Method:             "POST",
		PathPattern:        "/access-control/assignments/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchResultReader{formats: a.formats},
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
	success, ok := result.(*SearchResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
