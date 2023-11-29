// Code generated by go-swagger; DO NOT EDIT.

package folder_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new folder permissions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for folder permissions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFolderPermissionList(folderUID string, opts ...ClientOption) (*GetFolderPermissionListOK, error)
	GetFolderPermissionListWithParams(params *GetFolderPermissionListParams, opts ...ClientOption) (*GetFolderPermissionListOK, error)

	UpdateFolderPermissions(folderUID string, body *models.UpdateDashboardACLCommand, opts ...ClientOption) (*UpdateFolderPermissionsOK, error)
	UpdateFolderPermissionsWithParams(params *UpdateFolderPermissionsParams, opts ...ClientOption) (*UpdateFolderPermissionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetFolderPermissionList gets all existing permissions for the folder with the given uid
*/
func (a *Client) GetFolderPermissionList(folderUID string, opts ...ClientOption) (*GetFolderPermissionListOK, error) {
	params := NewGetFolderPermissionListParams().WithFolderUID(folderUID)
	return a.GetFolderPermissionListWithParams(params, opts...)
}

func (a *Client) GetFolderPermissionListWithParams(params *GetFolderPermissionListParams, opts ...ClientOption) (*GetFolderPermissionListOK, error) {
	if params == nil {
		params = NewGetFolderPermissionListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolderPermissionList",
		Method:             "GET",
		PathPattern:        "/folders/{folder_uid}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFolderPermissionListReader{formats: a.formats},
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
	success, ok := result.(*GetFolderPermissionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolderPermissionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFolderPermissions updates permissions for a folder this operation will remove existing permissions if they re not included in the request
*/
func (a *Client) UpdateFolderPermissions(folderUID string, body *models.UpdateDashboardACLCommand, opts ...ClientOption) (*UpdateFolderPermissionsOK, error) {
	params := NewUpdateFolderPermissionsParams().WithBody(body).WithFolderUID(folderUID)
	return a.UpdateFolderPermissionsWithParams(params, opts...)
}

func (a *Client) UpdateFolderPermissionsWithParams(params *UpdateFolderPermissionsParams, opts ...ClientOption) (*UpdateFolderPermissionsOK, error) {
	if params == nil {
		params = NewUpdateFolderPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFolderPermissions",
		Method:             "POST",
		PathPattern:        "/folders/{folder_uid}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateFolderPermissionsReader{formats: a.formats},
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
	success, ok := result.(*UpdateFolderPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFolderPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
