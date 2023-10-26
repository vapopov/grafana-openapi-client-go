// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetDataSourcesReader is a Reader for the GetDataSources structure.
type GetDataSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDataSourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDataSourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataSourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources] getDataSources", response, response.Code())
	}
}

// NewGetDataSourcesOK creates a GetDataSourcesOK with default headers values
func NewGetDataSourcesOK() *GetDataSourcesOK {
	return &GetDataSourcesOK{}
}

/*
GetDataSourcesOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDataSourcesOK struct {
	Payload models.DataSourceList
}

// IsSuccess returns true when this get data sources Ok response has a 2xx status code
func (o *GetDataSourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data sources Ok response has a 3xx status code
func (o *GetDataSourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data sources Ok response has a 4xx status code
func (o *GetDataSourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data sources Ok response has a 5xx status code
func (o *GetDataSourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data sources Ok response a status code equal to that given
func (o *GetDataSourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get data sources Ok response
func (o *GetDataSourcesOK) Code() int {
	return 200
}

func (o *GetDataSourcesOK) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesOk  %+v", 200, o.Payload)
}

func (o *GetDataSourcesOK) String() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesOk  %+v", 200, o.Payload)
}

func (o *GetDataSourcesOK) GetPayload() models.DataSourceList {
	return o.Payload
}

func (o *GetDataSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourcesUnauthorized creates a GetDataSourcesUnauthorized with default headers values
func NewGetDataSourcesUnauthorized() *GetDataSourcesUnauthorized {
	return &GetDataSourcesUnauthorized{}
}

/*
GetDataSourcesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDataSourcesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data sources unauthorized response has a 2xx status code
func (o *GetDataSourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data sources unauthorized response has a 3xx status code
func (o *GetDataSourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data sources unauthorized response has a 4xx status code
func (o *GetDataSourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data sources unauthorized response has a 5xx status code
func (o *GetDataSourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get data sources unauthorized response a status code equal to that given
func (o *GetDataSourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get data sources unauthorized response
func (o *GetDataSourcesUnauthorized) Code() int {
	return 401
}

func (o *GetDataSourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDataSourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDataSourcesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourcesForbidden creates a GetDataSourcesForbidden with default headers values
func NewGetDataSourcesForbidden() *GetDataSourcesForbidden {
	return &GetDataSourcesForbidden{}
}

/*
GetDataSourcesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDataSourcesForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data sources forbidden response has a 2xx status code
func (o *GetDataSourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data sources forbidden response has a 3xx status code
func (o *GetDataSourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data sources forbidden response has a 4xx status code
func (o *GetDataSourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data sources forbidden response has a 5xx status code
func (o *GetDataSourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get data sources forbidden response a status code equal to that given
func (o *GetDataSourcesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get data sources forbidden response
func (o *GetDataSourcesForbidden) Code() int {
	return 403
}

func (o *GetDataSourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetDataSourcesForbidden) String() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetDataSourcesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourcesInternalServerError creates a GetDataSourcesInternalServerError with default headers values
func NewGetDataSourcesInternalServerError() *GetDataSourcesInternalServerError {
	return &GetDataSourcesInternalServerError{}
}

/*
GetDataSourcesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDataSourcesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get data sources internal server error response has a 2xx status code
func (o *GetDataSourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data sources internal server error response has a 3xx status code
func (o *GetDataSourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data sources internal server error response has a 4xx status code
func (o *GetDataSourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data sources internal server error response has a 5xx status code
func (o *GetDataSourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get data sources internal server error response a status code equal to that given
func (o *GetDataSourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get data sources internal server error response
func (o *GetDataSourcesInternalServerError) Code() int {
	return 500
}

func (o *GetDataSourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataSourcesInternalServerError) String() string {
	return fmt.Sprintf("[GET /datasources][%d] getDataSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDataSourcesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDataSourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
