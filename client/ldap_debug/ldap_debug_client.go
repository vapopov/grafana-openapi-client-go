// Code generated by go-swagger; DO NOT EDIT.

package ldap_debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ldap debug API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ldap debug API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSyncStatus(opts ...ClientOption) (*GetSyncStatusOK, error)
	GetSyncStatusWithParams(params *GetSyncStatusParams, opts ...ClientOption) (*GetSyncStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetSyncStatus returns the current state of the LDAP background sync integration

You need to have a permission with action `ldap.status:read`.
*/
func (a *Client) GetSyncStatus(opts ...ClientOption) (*GetSyncStatusOK, error) {
	params := NewGetSyncStatusParams()
	return a.GetSyncStatusWithParams(params, opts...)
}

func (a *Client) GetSyncStatusWithParams(params *GetSyncStatusParams, opts ...ClientOption) (*GetSyncStatusOK, error) {
	if params == nil {
		params = NewGetSyncStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSyncStatus",
		Method:             "GET",
		PathPattern:        "/admin/ldap-sync-status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSyncStatusReader{formats: a.formats},
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
	success, ok := result.(*GetSyncStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSyncStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
