// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// AdminRevokeUserAuthTokenReader is a Reader for the AdminRevokeUserAuthToken structure.
type AdminRevokeUserAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminRevokeUserAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminRevokeUserAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminRevokeUserAuthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminRevokeUserAuthTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminRevokeUserAuthTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminRevokeUserAuthTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminRevokeUserAuthTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/users/{user_id}/revoke-auth-token] adminRevokeUserAuthToken", response, response.Code())
	}
}

// NewAdminRevokeUserAuthTokenOK creates a AdminRevokeUserAuthTokenOK with default headers values
func NewAdminRevokeUserAuthTokenOK() *AdminRevokeUserAuthTokenOK {
	return &AdminRevokeUserAuthTokenOK{}
}

/*
AdminRevokeUserAuthTokenOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminRevokeUserAuthTokenOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this admin revoke user auth token Ok response has a 2xx status code
func (o *AdminRevokeUserAuthTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin revoke user auth token Ok response has a 3xx status code
func (o *AdminRevokeUserAuthTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin revoke user auth token Ok response has a 4xx status code
func (o *AdminRevokeUserAuthTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin revoke user auth token Ok response has a 5xx status code
func (o *AdminRevokeUserAuthTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin revoke user auth token Ok response a status code equal to that given
func (o *AdminRevokeUserAuthTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin revoke user auth token Ok response
func (o *AdminRevokeUserAuthTokenOK) Code() int {
	return 200
}

func (o *AdminRevokeUserAuthTokenOK) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenOk  %+v", 200, o.Payload)
}

func (o *AdminRevokeUserAuthTokenOK) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenOk  %+v", 200, o.Payload)
}

func (o *AdminRevokeUserAuthTokenOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenBadRequest creates a AdminRevokeUserAuthTokenBadRequest with default headers values
func NewAdminRevokeUserAuthTokenBadRequest() *AdminRevokeUserAuthTokenBadRequest {
	return &AdminRevokeUserAuthTokenBadRequest{}
}

/*
AdminRevokeUserAuthTokenBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminRevokeUserAuthTokenBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin revoke user auth token bad request response has a 2xx status code
func (o *AdminRevokeUserAuthTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin revoke user auth token bad request response has a 3xx status code
func (o *AdminRevokeUserAuthTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin revoke user auth token bad request response has a 4xx status code
func (o *AdminRevokeUserAuthTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin revoke user auth token bad request response has a 5xx status code
func (o *AdminRevokeUserAuthTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin revoke user auth token bad request response a status code equal to that given
func (o *AdminRevokeUserAuthTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin revoke user auth token bad request response
func (o *AdminRevokeUserAuthTokenBadRequest) Code() int {
	return 400
}

func (o *AdminRevokeUserAuthTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenBadRequest  %+v", 400, o.Payload)
}

func (o *AdminRevokeUserAuthTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenBadRequest  %+v", 400, o.Payload)
}

func (o *AdminRevokeUserAuthTokenBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenUnauthorized creates a AdminRevokeUserAuthTokenUnauthorized with default headers values
func NewAdminRevokeUserAuthTokenUnauthorized() *AdminRevokeUserAuthTokenUnauthorized {
	return &AdminRevokeUserAuthTokenUnauthorized{}
}

/*
AdminRevokeUserAuthTokenUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminRevokeUserAuthTokenUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin revoke user auth token unauthorized response has a 2xx status code
func (o *AdminRevokeUserAuthTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin revoke user auth token unauthorized response has a 3xx status code
func (o *AdminRevokeUserAuthTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin revoke user auth token unauthorized response has a 4xx status code
func (o *AdminRevokeUserAuthTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin revoke user auth token unauthorized response has a 5xx status code
func (o *AdminRevokeUserAuthTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin revoke user auth token unauthorized response a status code equal to that given
func (o *AdminRevokeUserAuthTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin revoke user auth token unauthorized response
func (o *AdminRevokeUserAuthTokenUnauthorized) Code() int {
	return 401
}

func (o *AdminRevokeUserAuthTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminRevokeUserAuthTokenUnauthorized) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminRevokeUserAuthTokenUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenForbidden creates a AdminRevokeUserAuthTokenForbidden with default headers values
func NewAdminRevokeUserAuthTokenForbidden() *AdminRevokeUserAuthTokenForbidden {
	return &AdminRevokeUserAuthTokenForbidden{}
}

/*
AdminRevokeUserAuthTokenForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminRevokeUserAuthTokenForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin revoke user auth token forbidden response has a 2xx status code
func (o *AdminRevokeUserAuthTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin revoke user auth token forbidden response has a 3xx status code
func (o *AdminRevokeUserAuthTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin revoke user auth token forbidden response has a 4xx status code
func (o *AdminRevokeUserAuthTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin revoke user auth token forbidden response has a 5xx status code
func (o *AdminRevokeUserAuthTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin revoke user auth token forbidden response a status code equal to that given
func (o *AdminRevokeUserAuthTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin revoke user auth token forbidden response
func (o *AdminRevokeUserAuthTokenForbidden) Code() int {
	return 403
}

func (o *AdminRevokeUserAuthTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenForbidden  %+v", 403, o.Payload)
}

func (o *AdminRevokeUserAuthTokenForbidden) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenForbidden  %+v", 403, o.Payload)
}

func (o *AdminRevokeUserAuthTokenForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenNotFound creates a AdminRevokeUserAuthTokenNotFound with default headers values
func NewAdminRevokeUserAuthTokenNotFound() *AdminRevokeUserAuthTokenNotFound {
	return &AdminRevokeUserAuthTokenNotFound{}
}

/*
AdminRevokeUserAuthTokenNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AdminRevokeUserAuthTokenNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin revoke user auth token not found response has a 2xx status code
func (o *AdminRevokeUserAuthTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin revoke user auth token not found response has a 3xx status code
func (o *AdminRevokeUserAuthTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin revoke user auth token not found response has a 4xx status code
func (o *AdminRevokeUserAuthTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin revoke user auth token not found response has a 5xx status code
func (o *AdminRevokeUserAuthTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin revoke user auth token not found response a status code equal to that given
func (o *AdminRevokeUserAuthTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin revoke user auth token not found response
func (o *AdminRevokeUserAuthTokenNotFound) Code() int {
	return 404
}

func (o *AdminRevokeUserAuthTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenNotFound  %+v", 404, o.Payload)
}

func (o *AdminRevokeUserAuthTokenNotFound) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenNotFound  %+v", 404, o.Payload)
}

func (o *AdminRevokeUserAuthTokenNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminRevokeUserAuthTokenInternalServerError creates a AdminRevokeUserAuthTokenInternalServerError with default headers values
func NewAdminRevokeUserAuthTokenInternalServerError() *AdminRevokeUserAuthTokenInternalServerError {
	return &AdminRevokeUserAuthTokenInternalServerError{}
}

/*
AdminRevokeUserAuthTokenInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminRevokeUserAuthTokenInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin revoke user auth token internal server error response has a 2xx status code
func (o *AdminRevokeUserAuthTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin revoke user auth token internal server error response has a 3xx status code
func (o *AdminRevokeUserAuthTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin revoke user auth token internal server error response has a 4xx status code
func (o *AdminRevokeUserAuthTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin revoke user auth token internal server error response has a 5xx status code
func (o *AdminRevokeUserAuthTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin revoke user auth token internal server error response a status code equal to that given
func (o *AdminRevokeUserAuthTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin revoke user auth token internal server error response
func (o *AdminRevokeUserAuthTokenInternalServerError) Code() int {
	return 500
}

func (o *AdminRevokeUserAuthTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminRevokeUserAuthTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/revoke-auth-token][%d] adminRevokeUserAuthTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminRevokeUserAuthTokenInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminRevokeUserAuthTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
