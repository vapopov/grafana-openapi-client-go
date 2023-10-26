// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetTeamByIDReader is a Reader for the GetTeamByID structure.
type GetTeamByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTeamByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTeamByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{team_id}] getTeamByID", response, response.Code())
	}
}

// NewGetTeamByIDOK creates a GetTeamByIDOK with default headers values
func NewGetTeamByIDOK() *GetTeamByIDOK {
	return &GetTeamByIDOK{}
}

/*
GetTeamByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetTeamByIDOK struct {
	Payload *models.TeamDTO
}

// IsSuccess returns true when this get team by Id Ok response has a 2xx status code
func (o *GetTeamByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team by Id Ok response has a 3xx status code
func (o *GetTeamByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team by Id Ok response has a 4xx status code
func (o *GetTeamByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team by Id Ok response has a 5xx status code
func (o *GetTeamByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team by Id Ok response a status code equal to that given
func (o *GetTeamByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team by Id Ok response
func (o *GetTeamByIDOK) Code() int {
	return 200
}

func (o *GetTeamByIDOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdOk  %+v", 200, o.Payload)
}

func (o *GetTeamByIDOK) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdOk  %+v", 200, o.Payload)
}

func (o *GetTeamByIDOK) GetPayload() *models.TeamDTO {
	return o.Payload
}

func (o *GetTeamByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamByIDUnauthorized creates a GetTeamByIDUnauthorized with default headers values
func NewGetTeamByIDUnauthorized() *GetTeamByIDUnauthorized {
	return &GetTeamByIDUnauthorized{}
}

/*
GetTeamByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetTeamByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team by Id unauthorized response has a 2xx status code
func (o *GetTeamByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team by Id unauthorized response has a 3xx status code
func (o *GetTeamByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team by Id unauthorized response has a 4xx status code
func (o *GetTeamByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team by Id unauthorized response has a 5xx status code
func (o *GetTeamByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get team by Id unauthorized response a status code equal to that given
func (o *GetTeamByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get team by Id unauthorized response
func (o *GetTeamByIDUnauthorized) Code() int {
	return 401
}

func (o *GetTeamByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamByIDForbidden creates a GetTeamByIDForbidden with default headers values
func NewGetTeamByIDForbidden() *GetTeamByIDForbidden {
	return &GetTeamByIDForbidden{}
}

/*
GetTeamByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetTeamByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team by Id forbidden response has a 2xx status code
func (o *GetTeamByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team by Id forbidden response has a 3xx status code
func (o *GetTeamByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team by Id forbidden response has a 4xx status code
func (o *GetTeamByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team by Id forbidden response has a 5xx status code
func (o *GetTeamByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team by Id forbidden response a status code equal to that given
func (o *GetTeamByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team by Id forbidden response
func (o *GetTeamByIDForbidden) Code() int {
	return 403
}

func (o *GetTeamByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamByIDForbidden) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamByIDNotFound creates a GetTeamByIDNotFound with default headers values
func NewGetTeamByIDNotFound() *GetTeamByIDNotFound {
	return &GetTeamByIDNotFound{}
}

/*
GetTeamByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetTeamByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team by Id not found response has a 2xx status code
func (o *GetTeamByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team by Id not found response has a 3xx status code
func (o *GetTeamByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team by Id not found response has a 4xx status code
func (o *GetTeamByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team by Id not found response has a 5xx status code
func (o *GetTeamByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team by Id not found response a status code equal to that given
func (o *GetTeamByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team by Id not found response
func (o *GetTeamByIDNotFound) Code() int {
	return 404
}

func (o *GetTeamByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamByIDNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamByIDInternalServerError creates a GetTeamByIDInternalServerError with default headers values
func NewGetTeamByIDInternalServerError() *GetTeamByIDInternalServerError {
	return &GetTeamByIDInternalServerError{}
}

/*
GetTeamByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetTeamByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team by Id internal server error response has a 2xx status code
func (o *GetTeamByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team by Id internal server error response has a 3xx status code
func (o *GetTeamByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team by Id internal server error response has a 4xx status code
func (o *GetTeamByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team by Id internal server error response has a 5xx status code
func (o *GetTeamByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team by Id internal server error response a status code equal to that given
func (o *GetTeamByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team by Id internal server error response
func (o *GetTeamByIDInternalServerError) Code() int {
	return 500
}

func (o *GetTeamByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] getTeamByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
