// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// RouteGetPolicyTreeExportReader is a Reader for the RouteGetPolicyTreeExport structure.
type RouteGetPolicyTreeExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetPolicyTreeExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetPolicyTreeExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRouteGetPolicyTreeExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/policies/export] RouteGetPolicyTreeExport", response, response.Code())
	}
}

// NewRouteGetPolicyTreeExportOK creates a RouteGetPolicyTreeExportOK with default headers values
func NewRouteGetPolicyTreeExportOK() *RouteGetPolicyTreeExportOK {
	return &RouteGetPolicyTreeExportOK{}
}

/*
RouteGetPolicyTreeExportOK describes a response with status code 200, with default header values.

AlertingFileExport
*/
type RouteGetPolicyTreeExportOK struct {
	Payload *models.AlertingFileExport
}

// IsSuccess returns true when this route get policy tree export Ok response has a 2xx status code
func (o *RouteGetPolicyTreeExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route get policy tree export Ok response has a 3xx status code
func (o *RouteGetPolicyTreeExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get policy tree export Ok response has a 4xx status code
func (o *RouteGetPolicyTreeExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this route get policy tree export Ok response has a 5xx status code
func (o *RouteGetPolicyTreeExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this route get policy tree export Ok response a status code equal to that given
func (o *RouteGetPolicyTreeExportOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the route get policy tree export Ok response
func (o *RouteGetPolicyTreeExportOK) Code() int {
	return 200
}

func (o *RouteGetPolicyTreeExportOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] routeGetPolicyTreeExportOk %s", 200, payload)
}

func (o *RouteGetPolicyTreeExportOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] routeGetPolicyTreeExportOk %s", 200, payload)
}

func (o *RouteGetPolicyTreeExportOK) GetPayload() *models.AlertingFileExport {
	return o.Payload
}

func (o *RouteGetPolicyTreeExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingFileExport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetPolicyTreeExportNotFound creates a RouteGetPolicyTreeExportNotFound with default headers values
func NewRouteGetPolicyTreeExportNotFound() *RouteGetPolicyTreeExportNotFound {
	return &RouteGetPolicyTreeExportNotFound{}
}

/*
RouteGetPolicyTreeExportNotFound describes a response with status code 404, with default header values.

NotFound
*/
type RouteGetPolicyTreeExportNotFound struct {
	Payload models.NotFound
}

// IsSuccess returns true when this route get policy tree export not found response has a 2xx status code
func (o *RouteGetPolicyTreeExportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this route get policy tree export not found response has a 3xx status code
func (o *RouteGetPolicyTreeExportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get policy tree export not found response has a 4xx status code
func (o *RouteGetPolicyTreeExportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this route get policy tree export not found response has a 5xx status code
func (o *RouteGetPolicyTreeExportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this route get policy tree export not found response a status code equal to that given
func (o *RouteGetPolicyTreeExportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the route get policy tree export not found response
func (o *RouteGetPolicyTreeExportNotFound) Code() int {
	return 404
}

func (o *RouteGetPolicyTreeExportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] routeGetPolicyTreeExportNotFound %s", 404, payload)
}

func (o *RouteGetPolicyTreeExportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/policies/export][%d] routeGetPolicyTreeExportNotFound %s", 404, payload)
}

func (o *RouteGetPolicyTreeExportNotFound) GetPayload() models.NotFound {
	return o.Payload
}

func (o *RouteGetPolicyTreeExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
