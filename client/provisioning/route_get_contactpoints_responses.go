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

// RouteGetContactpointsReader is a Reader for the RouteGetContactpoints structure.
type RouteGetContactpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetContactpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetContactpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/contact-points] RouteGetContactpoints", response, response.Code())
	}
}

// NewRouteGetContactpointsOK creates a RouteGetContactpointsOK with default headers values
func NewRouteGetContactpointsOK() *RouteGetContactpointsOK {
	return &RouteGetContactpointsOK{}
}

/*
RouteGetContactpointsOK describes a response with status code 200, with default header values.

ContactPoints
*/
type RouteGetContactpointsOK struct {
	Payload models.ContactPoints
}

// IsSuccess returns true when this route get contactpoints Ok response has a 2xx status code
func (o *RouteGetContactpointsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route get contactpoints Ok response has a 3xx status code
func (o *RouteGetContactpointsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get contactpoints Ok response has a 4xx status code
func (o *RouteGetContactpointsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this route get contactpoints Ok response has a 5xx status code
func (o *RouteGetContactpointsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this route get contactpoints Ok response a status code equal to that given
func (o *RouteGetContactpointsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the route get contactpoints Ok response
func (o *RouteGetContactpointsOK) Code() int {
	return 200
}

func (o *RouteGetContactpointsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/contact-points][%d] routeGetContactpointsOk %s", 200, payload)
}

func (o *RouteGetContactpointsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/contact-points][%d] routeGetContactpointsOk %s", 200, payload)
}

func (o *RouteGetContactpointsOK) GetPayload() models.ContactPoints {
	return o.Payload
}

func (o *RouteGetContactpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
