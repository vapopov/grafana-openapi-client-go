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

// AdminDisableUserReader is a Reader for the AdminDisableUser structure.
type AdminDisableUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDisableUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminDisableUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDisableUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminDisableUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminDisableUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminDisableUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/users/{user_id}/disable] adminDisableUser", response, response.Code())
	}
}

// NewAdminDisableUserOK creates a AdminDisableUserOK with default headers values
func NewAdminDisableUserOK() *AdminDisableUserOK {
	return &AdminDisableUserOK{}
}

/*
AdminDisableUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminDisableUserOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this admin disable user Ok response has a 2xx status code
func (o *AdminDisableUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin disable user Ok response has a 3xx status code
func (o *AdminDisableUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin disable user Ok response has a 4xx status code
func (o *AdminDisableUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin disable user Ok response has a 5xx status code
func (o *AdminDisableUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin disable user Ok response a status code equal to that given
func (o *AdminDisableUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin disable user Ok response
func (o *AdminDisableUserOK) Code() int {
	return 200
}

func (o *AdminDisableUserOK) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserOk  %+v", 200, o.Payload)
}

func (o *AdminDisableUserOK) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserOk  %+v", 200, o.Payload)
}

func (o *AdminDisableUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminDisableUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDisableUserUnauthorized creates a AdminDisableUserUnauthorized with default headers values
func NewAdminDisableUserUnauthorized() *AdminDisableUserUnauthorized {
	return &AdminDisableUserUnauthorized{}
}

/*
AdminDisableUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminDisableUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin disable user unauthorized response has a 2xx status code
func (o *AdminDisableUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin disable user unauthorized response has a 3xx status code
func (o *AdminDisableUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin disable user unauthorized response has a 4xx status code
func (o *AdminDisableUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin disable user unauthorized response has a 5xx status code
func (o *AdminDisableUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin disable user unauthorized response a status code equal to that given
func (o *AdminDisableUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin disable user unauthorized response
func (o *AdminDisableUserUnauthorized) Code() int {
	return 401
}

func (o *AdminDisableUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminDisableUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminDisableUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDisableUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDisableUserForbidden creates a AdminDisableUserForbidden with default headers values
func NewAdminDisableUserForbidden() *AdminDisableUserForbidden {
	return &AdminDisableUserForbidden{}
}

/*
AdminDisableUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminDisableUserForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin disable user forbidden response has a 2xx status code
func (o *AdminDisableUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin disable user forbidden response has a 3xx status code
func (o *AdminDisableUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin disable user forbidden response has a 4xx status code
func (o *AdminDisableUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin disable user forbidden response has a 5xx status code
func (o *AdminDisableUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin disable user forbidden response a status code equal to that given
func (o *AdminDisableUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin disable user forbidden response
func (o *AdminDisableUserForbidden) Code() int {
	return 403
}

func (o *AdminDisableUserForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminDisableUserForbidden) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminDisableUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDisableUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDisableUserNotFound creates a AdminDisableUserNotFound with default headers values
func NewAdminDisableUserNotFound() *AdminDisableUserNotFound {
	return &AdminDisableUserNotFound{}
}

/*
AdminDisableUserNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AdminDisableUserNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin disable user not found response has a 2xx status code
func (o *AdminDisableUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin disable user not found response has a 3xx status code
func (o *AdminDisableUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin disable user not found response has a 4xx status code
func (o *AdminDisableUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin disable user not found response has a 5xx status code
func (o *AdminDisableUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin disable user not found response a status code equal to that given
func (o *AdminDisableUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin disable user not found response
func (o *AdminDisableUserNotFound) Code() int {
	return 404
}

func (o *AdminDisableUserNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminDisableUserNotFound) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminDisableUserNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDisableUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDisableUserInternalServerError creates a AdminDisableUserInternalServerError with default headers values
func NewAdminDisableUserInternalServerError() *AdminDisableUserInternalServerError {
	return &AdminDisableUserInternalServerError{}
}

/*
AdminDisableUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminDisableUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin disable user internal server error response has a 2xx status code
func (o *AdminDisableUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin disable user internal server error response has a 3xx status code
func (o *AdminDisableUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin disable user internal server error response has a 4xx status code
func (o *AdminDisableUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin disable user internal server error response has a 5xx status code
func (o *AdminDisableUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin disable user internal server error response a status code equal to that given
func (o *AdminDisableUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin disable user internal server error response
func (o *AdminDisableUserInternalServerError) Code() int {
	return 500
}

func (o *AdminDisableUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminDisableUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/disable][%d] adminDisableUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminDisableUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDisableUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
