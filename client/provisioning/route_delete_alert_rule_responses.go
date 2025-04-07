// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RouteDeleteAlertRuleReader is a Reader for the RouteDeleteAlertRule structure.
type RouteDeleteAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteDeleteAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRouteDeleteAlertRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/provisioning/alert-rules/{UID}] RouteDeleteAlertRule", response, response.Code())
	}
}

// NewRouteDeleteAlertRuleNoContent creates a RouteDeleteAlertRuleNoContent with default headers values
func NewRouteDeleteAlertRuleNoContent() *RouteDeleteAlertRuleNoContent {
	return &RouteDeleteAlertRuleNoContent{}
}

/*
RouteDeleteAlertRuleNoContent describes a response with status code 204, with default header values.

	The alert rule was deleted successfully.
*/
type RouteDeleteAlertRuleNoContent struct {
}

// IsSuccess returns true when this route delete alert rule no content response has a 2xx status code
func (o *RouteDeleteAlertRuleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route delete alert rule no content response has a 3xx status code
func (o *RouteDeleteAlertRuleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route delete alert rule no content response has a 4xx status code
func (o *RouteDeleteAlertRuleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this route delete alert rule no content response has a 5xx status code
func (o *RouteDeleteAlertRuleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this route delete alert rule no content response a status code equal to that given
func (o *RouteDeleteAlertRuleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the route delete alert rule no content response
func (o *RouteDeleteAlertRuleNoContent) Code() int {
	return 204
}

func (o *RouteDeleteAlertRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/provisioning/alert-rules/{UID}][%d] routeDeleteAlertRuleNoContent", 204)
}

func (o *RouteDeleteAlertRuleNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/provisioning/alert-rules/{UID}][%d] routeDeleteAlertRuleNoContent", 204)
}

func (o *RouteDeleteAlertRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
