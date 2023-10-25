// Code generated by go-swagger; DO NOT EDIT.

package access_control_provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access control provisioning API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access control provisioning API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AdminProvisioningReloadAccessControl(params *AdminProvisioningReloadAccessControlParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadAccessControlAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AdminProvisioningReloadAccessControl yous need to have a permission with action provisioning reload with scope provisioners accesscontrol
*/
func (a *Client) AdminProvisioningReloadAccessControl(params *AdminProvisioningReloadAccessControlParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminProvisioningReloadAccessControlAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminProvisioningReloadAccessControlParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminProvisioningReloadAccessControl",
		Method:             "POST",
		PathPattern:        "/admin/provisioning/access-control/reload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminProvisioningReloadAccessControlReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminProvisioningReloadAccessControlAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminProvisioningReloadAccessControl: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
