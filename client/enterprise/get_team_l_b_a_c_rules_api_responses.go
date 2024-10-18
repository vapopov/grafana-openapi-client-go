// Code generated by go-swagger; DO NOT EDIT.

package enterprise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetTeamLBACRulesAPIReader is a Reader for the GetTeamLBACRulesAPI structure.
type GetTeamLBACRulesAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamLBACRulesAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamLBACRulesAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTeamLBACRulesAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTeamLBACRulesAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTeamLBACRulesAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamLBACRulesAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamLBACRulesAPIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/uid/{uid}/lbac/teams] getTeamLBACRulesApi", response, response.Code())
	}
}

// NewGetTeamLBACRulesAPIOK creates a GetTeamLBACRulesAPIOK with default headers values
func NewGetTeamLBACRulesAPIOK() *GetTeamLBACRulesAPIOK {
	return &GetTeamLBACRulesAPIOK{}
}

/*
GetTeamLBACRulesAPIOK describes a response with status code 200, with default header values.

(empty)
*/
type GetTeamLBACRulesAPIOK struct {
	Payload *models.TeamLBACRules
}

// IsSuccess returns true when this get team l b a c rules Api Ok response has a 2xx status code
func (o *GetTeamLBACRulesAPIOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team l b a c rules Api Ok response has a 3xx status code
func (o *GetTeamLBACRulesAPIOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team l b a c rules Api Ok response has a 4xx status code
func (o *GetTeamLBACRulesAPIOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team l b a c rules Api Ok response has a 5xx status code
func (o *GetTeamLBACRulesAPIOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team l b a c rules Api Ok response a status code equal to that given
func (o *GetTeamLBACRulesAPIOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team l b a c rules Api Ok response
func (o *GetTeamLBACRulesAPIOK) Code() int {
	return 200
}

func (o *GetTeamLBACRulesAPIOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiOk %s", 200, payload)
}

func (o *GetTeamLBACRulesAPIOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiOk %s", 200, payload)
}

func (o *GetTeamLBACRulesAPIOK) GetPayload() *models.TeamLBACRules {
	return o.Payload
}

func (o *GetTeamLBACRulesAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamLBACRules)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamLBACRulesAPIBadRequest creates a GetTeamLBACRulesAPIBadRequest with default headers values
func NewGetTeamLBACRulesAPIBadRequest() *GetTeamLBACRulesAPIBadRequest {
	return &GetTeamLBACRulesAPIBadRequest{}
}

/*
GetTeamLBACRulesAPIBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetTeamLBACRulesAPIBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team l b a c rules Api bad request response has a 2xx status code
func (o *GetTeamLBACRulesAPIBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team l b a c rules Api bad request response has a 3xx status code
func (o *GetTeamLBACRulesAPIBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team l b a c rules Api bad request response has a 4xx status code
func (o *GetTeamLBACRulesAPIBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team l b a c rules Api bad request response has a 5xx status code
func (o *GetTeamLBACRulesAPIBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get team l b a c rules Api bad request response a status code equal to that given
func (o *GetTeamLBACRulesAPIBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get team l b a c rules Api bad request response
func (o *GetTeamLBACRulesAPIBadRequest) Code() int {
	return 400
}

func (o *GetTeamLBACRulesAPIBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiBadRequest %s", 400, payload)
}

func (o *GetTeamLBACRulesAPIBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiBadRequest %s", 400, payload)
}

func (o *GetTeamLBACRulesAPIBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamLBACRulesAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamLBACRulesAPIUnauthorized creates a GetTeamLBACRulesAPIUnauthorized with default headers values
func NewGetTeamLBACRulesAPIUnauthorized() *GetTeamLBACRulesAPIUnauthorized {
	return &GetTeamLBACRulesAPIUnauthorized{}
}

/*
GetTeamLBACRulesAPIUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetTeamLBACRulesAPIUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team l b a c rules Api unauthorized response has a 2xx status code
func (o *GetTeamLBACRulesAPIUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team l b a c rules Api unauthorized response has a 3xx status code
func (o *GetTeamLBACRulesAPIUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team l b a c rules Api unauthorized response has a 4xx status code
func (o *GetTeamLBACRulesAPIUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team l b a c rules Api unauthorized response has a 5xx status code
func (o *GetTeamLBACRulesAPIUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get team l b a c rules Api unauthorized response a status code equal to that given
func (o *GetTeamLBACRulesAPIUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get team l b a c rules Api unauthorized response
func (o *GetTeamLBACRulesAPIUnauthorized) Code() int {
	return 401
}

func (o *GetTeamLBACRulesAPIUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiUnauthorized %s", 401, payload)
}

func (o *GetTeamLBACRulesAPIUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiUnauthorized %s", 401, payload)
}

func (o *GetTeamLBACRulesAPIUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamLBACRulesAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamLBACRulesAPIForbidden creates a GetTeamLBACRulesAPIForbidden with default headers values
func NewGetTeamLBACRulesAPIForbidden() *GetTeamLBACRulesAPIForbidden {
	return &GetTeamLBACRulesAPIForbidden{}
}

/*
GetTeamLBACRulesAPIForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetTeamLBACRulesAPIForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team l b a c rules Api forbidden response has a 2xx status code
func (o *GetTeamLBACRulesAPIForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team l b a c rules Api forbidden response has a 3xx status code
func (o *GetTeamLBACRulesAPIForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team l b a c rules Api forbidden response has a 4xx status code
func (o *GetTeamLBACRulesAPIForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team l b a c rules Api forbidden response has a 5xx status code
func (o *GetTeamLBACRulesAPIForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team l b a c rules Api forbidden response a status code equal to that given
func (o *GetTeamLBACRulesAPIForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team l b a c rules Api forbidden response
func (o *GetTeamLBACRulesAPIForbidden) Code() int {
	return 403
}

func (o *GetTeamLBACRulesAPIForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiForbidden %s", 403, payload)
}

func (o *GetTeamLBACRulesAPIForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiForbidden %s", 403, payload)
}

func (o *GetTeamLBACRulesAPIForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamLBACRulesAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamLBACRulesAPINotFound creates a GetTeamLBACRulesAPINotFound with default headers values
func NewGetTeamLBACRulesAPINotFound() *GetTeamLBACRulesAPINotFound {
	return &GetTeamLBACRulesAPINotFound{}
}

/*
GetTeamLBACRulesAPINotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetTeamLBACRulesAPINotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team l b a c rules Api not found response has a 2xx status code
func (o *GetTeamLBACRulesAPINotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team l b a c rules Api not found response has a 3xx status code
func (o *GetTeamLBACRulesAPINotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team l b a c rules Api not found response has a 4xx status code
func (o *GetTeamLBACRulesAPINotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team l b a c rules Api not found response has a 5xx status code
func (o *GetTeamLBACRulesAPINotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team l b a c rules Api not found response a status code equal to that given
func (o *GetTeamLBACRulesAPINotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team l b a c rules Api not found response
func (o *GetTeamLBACRulesAPINotFound) Code() int {
	return 404
}

func (o *GetTeamLBACRulesAPINotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiNotFound %s", 404, payload)
}

func (o *GetTeamLBACRulesAPINotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiNotFound %s", 404, payload)
}

func (o *GetTeamLBACRulesAPINotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamLBACRulesAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamLBACRulesAPIInternalServerError creates a GetTeamLBACRulesAPIInternalServerError with default headers values
func NewGetTeamLBACRulesAPIInternalServerError() *GetTeamLBACRulesAPIInternalServerError {
	return &GetTeamLBACRulesAPIInternalServerError{}
}

/*
GetTeamLBACRulesAPIInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetTeamLBACRulesAPIInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team l b a c rules Api internal server error response has a 2xx status code
func (o *GetTeamLBACRulesAPIInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team l b a c rules Api internal server error response has a 3xx status code
func (o *GetTeamLBACRulesAPIInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team l b a c rules Api internal server error response has a 4xx status code
func (o *GetTeamLBACRulesAPIInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team l b a c rules Api internal server error response has a 5xx status code
func (o *GetTeamLBACRulesAPIInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team l b a c rules Api internal server error response a status code equal to that given
func (o *GetTeamLBACRulesAPIInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team l b a c rules Api internal server error response
func (o *GetTeamLBACRulesAPIInternalServerError) Code() int {
	return 500
}

func (o *GetTeamLBACRulesAPIInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiInternalServerError %s", 500, payload)
}

func (o *GetTeamLBACRulesAPIInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{uid}/lbac/teams][%d] getTeamLBACRulesApiInternalServerError %s", 500, payload)
}

func (o *GetTeamLBACRulesAPIInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamLBACRulesAPIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
