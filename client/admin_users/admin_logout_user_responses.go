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

// AdminLogoutUserReader is a Reader for the AdminLogoutUser structure.
type AdminLogoutUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminLogoutUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminLogoutUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminLogoutUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminLogoutUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminLogoutUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminLogoutUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminLogoutUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/users/{user_id}/logout] adminLogoutUser", response, response.Code())
	}
}

// NewAdminLogoutUserOK creates a AdminLogoutUserOK with default headers values
func NewAdminLogoutUserOK() *AdminLogoutUserOK {
	return &AdminLogoutUserOK{}
}

/*
AdminLogoutUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminLogoutUserOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this admin logout user Ok response has a 2xx status code
func (o *AdminLogoutUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin logout user Ok response has a 3xx status code
func (o *AdminLogoutUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin logout user Ok response has a 4xx status code
func (o *AdminLogoutUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin logout user Ok response has a 5xx status code
func (o *AdminLogoutUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin logout user Ok response a status code equal to that given
func (o *AdminLogoutUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin logout user Ok response
func (o *AdminLogoutUserOK) Code() int {
	return 200
}

func (o *AdminLogoutUserOK) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserOk  %+v", 200, o.Payload)
}

func (o *AdminLogoutUserOK) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserOk  %+v", 200, o.Payload)
}

func (o *AdminLogoutUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminLogoutUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminLogoutUserBadRequest creates a AdminLogoutUserBadRequest with default headers values
func NewAdminLogoutUserBadRequest() *AdminLogoutUserBadRequest {
	return &AdminLogoutUserBadRequest{}
}

/*
AdminLogoutUserBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminLogoutUserBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin logout user bad request response has a 2xx status code
func (o *AdminLogoutUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin logout user bad request response has a 3xx status code
func (o *AdminLogoutUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin logout user bad request response has a 4xx status code
func (o *AdminLogoutUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin logout user bad request response has a 5xx status code
func (o *AdminLogoutUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin logout user bad request response a status code equal to that given
func (o *AdminLogoutUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin logout user bad request response
func (o *AdminLogoutUserBadRequest) Code() int {
	return 400
}

func (o *AdminLogoutUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminLogoutUserBadRequest) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminLogoutUserBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminLogoutUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminLogoutUserUnauthorized creates a AdminLogoutUserUnauthorized with default headers values
func NewAdminLogoutUserUnauthorized() *AdminLogoutUserUnauthorized {
	return &AdminLogoutUserUnauthorized{}
}

/*
AdminLogoutUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminLogoutUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin logout user unauthorized response has a 2xx status code
func (o *AdminLogoutUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin logout user unauthorized response has a 3xx status code
func (o *AdminLogoutUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin logout user unauthorized response has a 4xx status code
func (o *AdminLogoutUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin logout user unauthorized response has a 5xx status code
func (o *AdminLogoutUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin logout user unauthorized response a status code equal to that given
func (o *AdminLogoutUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin logout user unauthorized response
func (o *AdminLogoutUserUnauthorized) Code() int {
	return 401
}

func (o *AdminLogoutUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminLogoutUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminLogoutUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminLogoutUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminLogoutUserForbidden creates a AdminLogoutUserForbidden with default headers values
func NewAdminLogoutUserForbidden() *AdminLogoutUserForbidden {
	return &AdminLogoutUserForbidden{}
}

/*
AdminLogoutUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminLogoutUserForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin logout user forbidden response has a 2xx status code
func (o *AdminLogoutUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin logout user forbidden response has a 3xx status code
func (o *AdminLogoutUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin logout user forbidden response has a 4xx status code
func (o *AdminLogoutUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin logout user forbidden response has a 5xx status code
func (o *AdminLogoutUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin logout user forbidden response a status code equal to that given
func (o *AdminLogoutUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin logout user forbidden response
func (o *AdminLogoutUserForbidden) Code() int {
	return 403
}

func (o *AdminLogoutUserForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminLogoutUserForbidden) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminLogoutUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminLogoutUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminLogoutUserNotFound creates a AdminLogoutUserNotFound with default headers values
func NewAdminLogoutUserNotFound() *AdminLogoutUserNotFound {
	return &AdminLogoutUserNotFound{}
}

/*
AdminLogoutUserNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AdminLogoutUserNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin logout user not found response has a 2xx status code
func (o *AdminLogoutUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin logout user not found response has a 3xx status code
func (o *AdminLogoutUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin logout user not found response has a 4xx status code
func (o *AdminLogoutUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin logout user not found response has a 5xx status code
func (o *AdminLogoutUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin logout user not found response a status code equal to that given
func (o *AdminLogoutUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin logout user not found response
func (o *AdminLogoutUserNotFound) Code() int {
	return 404
}

func (o *AdminLogoutUserNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminLogoutUserNotFound) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminLogoutUserNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminLogoutUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminLogoutUserInternalServerError creates a AdminLogoutUserInternalServerError with default headers values
func NewAdminLogoutUserInternalServerError() *AdminLogoutUserInternalServerError {
	return &AdminLogoutUserInternalServerError{}
}

/*
AdminLogoutUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminLogoutUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin logout user internal server error response has a 2xx status code
func (o *AdminLogoutUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin logout user internal server error response has a 3xx status code
func (o *AdminLogoutUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin logout user internal server error response has a 4xx status code
func (o *AdminLogoutUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin logout user internal server error response has a 5xx status code
func (o *AdminLogoutUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin logout user internal server error response a status code equal to that given
func (o *AdminLogoutUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin logout user internal server error response
func (o *AdminLogoutUserInternalServerError) Code() int {
	return 500
}

func (o *AdminLogoutUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminLogoutUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /admin/users/{user_id}/logout][%d] adminLogoutUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminLogoutUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminLogoutUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
