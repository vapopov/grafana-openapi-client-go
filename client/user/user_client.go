// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UpdateUserEmail(opts ...ClientOption) error
	UpdateUserEmailWithParams(params *UpdateUserEmailParams, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
UpdateUserEmail updates user email

Update the email of user given a verification code.
*/
func (a *Client) UpdateUserEmail(opts ...ClientOption) error {
	params := NewUpdateUserEmailParams()
	return a.UpdateUserEmailWithParams(params, opts...)
}

func (a *Client) UpdateUserEmailWithParams(params *UpdateUserEmailParams, opts ...ClientOption) error {
	if params == nil {
		params = NewUpdateUserEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserEmail",
		Method:             "GET",
		PathPattern:        "/user/email/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserEmailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
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
