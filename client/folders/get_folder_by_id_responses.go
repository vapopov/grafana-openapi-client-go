// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetFolderByIDReader is a Reader for the GetFolderByID structure.
type GetFolderByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFolderByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFolderByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFolderByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFolderByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFolderByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFolderByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /folders/id/{folder_id}] getFolderByID", response, response.Code())
	}
}

// NewGetFolderByIDOK creates a GetFolderByIDOK with default headers values
func NewGetFolderByIDOK() *GetFolderByIDOK {
	return &GetFolderByIDOK{}
}

/*
GetFolderByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetFolderByIDOK struct {
	Payload *models.Folder
}

// IsSuccess returns true when this get folder by Id Ok response has a 2xx status code
func (o *GetFolderByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get folder by Id Ok response has a 3xx status code
func (o *GetFolderByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Id Ok response has a 4xx status code
func (o *GetFolderByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder by Id Ok response has a 5xx status code
func (o *GetFolderByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Id Ok response a status code equal to that given
func (o *GetFolderByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get folder by Id Ok response
func (o *GetFolderByIDOK) Code() int {
	return 200
}

func (o *GetFolderByIDOK) Error() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdOk  %+v", 200, o.Payload)
}

func (o *GetFolderByIDOK) String() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdOk  %+v", 200, o.Payload)
}

func (o *GetFolderByIDOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *GetFolderByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByIDUnauthorized creates a GetFolderByIDUnauthorized with default headers values
func NewGetFolderByIDUnauthorized() *GetFolderByIDUnauthorized {
	return &GetFolderByIDUnauthorized{}
}

/*
GetFolderByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetFolderByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Id unauthorized response has a 2xx status code
func (o *GetFolderByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Id unauthorized response has a 3xx status code
func (o *GetFolderByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Id unauthorized response has a 4xx status code
func (o *GetFolderByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder by Id unauthorized response has a 5xx status code
func (o *GetFolderByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Id unauthorized response a status code equal to that given
func (o *GetFolderByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get folder by Id unauthorized response
func (o *GetFolderByIDUnauthorized) Code() int {
	return 401
}

func (o *GetFolderByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFolderByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFolderByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByIDForbidden creates a GetFolderByIDForbidden with default headers values
func NewGetFolderByIDForbidden() *GetFolderByIDForbidden {
	return &GetFolderByIDForbidden{}
}

/*
GetFolderByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetFolderByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Id forbidden response has a 2xx status code
func (o *GetFolderByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Id forbidden response has a 3xx status code
func (o *GetFolderByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Id forbidden response has a 4xx status code
func (o *GetFolderByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder by Id forbidden response has a 5xx status code
func (o *GetFolderByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Id forbidden response a status code equal to that given
func (o *GetFolderByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get folder by Id forbidden response
func (o *GetFolderByIDForbidden) Code() int {
	return 403
}

func (o *GetFolderByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetFolderByIDForbidden) String() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetFolderByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByIDNotFound creates a GetFolderByIDNotFound with default headers values
func NewGetFolderByIDNotFound() *GetFolderByIDNotFound {
	return &GetFolderByIDNotFound{}
}

/*
GetFolderByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetFolderByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Id not found response has a 2xx status code
func (o *GetFolderByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Id not found response has a 3xx status code
func (o *GetFolderByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Id not found response has a 4xx status code
func (o *GetFolderByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get folder by Id not found response has a 5xx status code
func (o *GetFolderByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get folder by Id not found response a status code equal to that given
func (o *GetFolderByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get folder by Id not found response
func (o *GetFolderByIDNotFound) Code() int {
	return 404
}

func (o *GetFolderByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderByIDNotFound) String() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetFolderByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFolderByIDInternalServerError creates a GetFolderByIDInternalServerError with default headers values
func NewGetFolderByIDInternalServerError() *GetFolderByIDInternalServerError {
	return &GetFolderByIDInternalServerError{}
}

/*
GetFolderByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetFolderByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get folder by Id internal server error response has a 2xx status code
func (o *GetFolderByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get folder by Id internal server error response has a 3xx status code
func (o *GetFolderByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get folder by Id internal server error response has a 4xx status code
func (o *GetFolderByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get folder by Id internal server error response has a 5xx status code
func (o *GetFolderByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get folder by Id internal server error response a status code equal to that given
func (o *GetFolderByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get folder by Id internal server error response
func (o *GetFolderByIDInternalServerError) Code() int {
	return 500
}

func (o *GetFolderByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFolderByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /folders/id/{folder_id}][%d] getFolderByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFolderByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFolderByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
