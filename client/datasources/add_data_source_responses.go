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

// AddDataSourceReader is a Reader for the AddDataSource structure.
type AddDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddDataSourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddDataSourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddDataSourceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDataSourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /datasources] addDataSource", response, response.Code())
	}
}

// NewAddDataSourceOK creates a AddDataSourceOK with default headers values
func NewAddDataSourceOK() *AddDataSourceOK {
	return &AddDataSourceOK{}
}

/*
AddDataSourceOK describes a response with status code 200, with default header values.

(empty)
*/
type AddDataSourceOK struct {
	Payload *models.AddDataSourceOKBody
}

// IsSuccess returns true when this add data source Ok response has a 2xx status code
func (o *AddDataSourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add data source Ok response has a 3xx status code
func (o *AddDataSourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add data source Ok response has a 4xx status code
func (o *AddDataSourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add data source Ok response has a 5xx status code
func (o *AddDataSourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add data source Ok response a status code equal to that given
func (o *AddDataSourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add data source Ok response
func (o *AddDataSourceOK) Code() int {
	return 200
}

func (o *AddDataSourceOK) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceOk  %+v", 200, o.Payload)
}

func (o *AddDataSourceOK) String() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceOk  %+v", 200, o.Payload)
}

func (o *AddDataSourceOK) GetPayload() *models.AddDataSourceOKBody {
	return o.Payload
}

func (o *AddDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddDataSourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceUnauthorized creates a AddDataSourceUnauthorized with default headers values
func NewAddDataSourceUnauthorized() *AddDataSourceUnauthorized {
	return &AddDataSourceUnauthorized{}
}

/*
AddDataSourceUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddDataSourceUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add data source unauthorized response has a 2xx status code
func (o *AddDataSourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add data source unauthorized response has a 3xx status code
func (o *AddDataSourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add data source unauthorized response has a 4xx status code
func (o *AddDataSourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add data source unauthorized response has a 5xx status code
func (o *AddDataSourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add data source unauthorized response a status code equal to that given
func (o *AddDataSourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add data source unauthorized response
func (o *AddDataSourceUnauthorized) Code() int {
	return 401
}

func (o *AddDataSourceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceUnauthorized  %+v", 401, o.Payload)
}

func (o *AddDataSourceUnauthorized) String() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceUnauthorized  %+v", 401, o.Payload)
}

func (o *AddDataSourceUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceForbidden creates a AddDataSourceForbidden with default headers values
func NewAddDataSourceForbidden() *AddDataSourceForbidden {
	return &AddDataSourceForbidden{}
}

/*
AddDataSourceForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddDataSourceForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add data source forbidden response has a 2xx status code
func (o *AddDataSourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add data source forbidden response has a 3xx status code
func (o *AddDataSourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add data source forbidden response has a 4xx status code
func (o *AddDataSourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add data source forbidden response has a 5xx status code
func (o *AddDataSourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add data source forbidden response a status code equal to that given
func (o *AddDataSourceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add data source forbidden response
func (o *AddDataSourceForbidden) Code() int {
	return 403
}

func (o *AddDataSourceForbidden) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceForbidden  %+v", 403, o.Payload)
}

func (o *AddDataSourceForbidden) String() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceForbidden  %+v", 403, o.Payload)
}

func (o *AddDataSourceForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceConflict creates a AddDataSourceConflict with default headers values
func NewAddDataSourceConflict() *AddDataSourceConflict {
	return &AddDataSourceConflict{}
}

/*
AddDataSourceConflict describes a response with status code 409, with default header values.

ConflictError
*/
type AddDataSourceConflict struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add data source conflict response has a 2xx status code
func (o *AddDataSourceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add data source conflict response has a 3xx status code
func (o *AddDataSourceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add data source conflict response has a 4xx status code
func (o *AddDataSourceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add data source conflict response has a 5xx status code
func (o *AddDataSourceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add data source conflict response a status code equal to that given
func (o *AddDataSourceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add data source conflict response
func (o *AddDataSourceConflict) Code() int {
	return 409
}

func (o *AddDataSourceConflict) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceConflict  %+v", 409, o.Payload)
}

func (o *AddDataSourceConflict) String() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceConflict  %+v", 409, o.Payload)
}

func (o *AddDataSourceConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceInternalServerError creates a AddDataSourceInternalServerError with default headers values
func NewAddDataSourceInternalServerError() *AddDataSourceInternalServerError {
	return &AddDataSourceInternalServerError{}
}

/*
AddDataSourceInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddDataSourceInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this add data source internal server error response has a 2xx status code
func (o *AddDataSourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add data source internal server error response has a 3xx status code
func (o *AddDataSourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add data source internal server error response has a 4xx status code
func (o *AddDataSourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add data source internal server error response has a 5xx status code
func (o *AddDataSourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add data source internal server error response a status code equal to that given
func (o *AddDataSourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add data source internal server error response
func (o *AddDataSourceInternalServerError) Code() int {
	return 500
}

func (o *AddDataSourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDataSourceInternalServerError) String() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDataSourceInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
