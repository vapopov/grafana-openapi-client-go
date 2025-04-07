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

// RouteDeleteAlertRuleGroupReader is a Reader for the RouteDeleteAlertRuleGroup structure.
type RouteDeleteAlertRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteDeleteAlertRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRouteDeleteAlertRuleGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRouteDeleteAlertRuleGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRouteDeleteAlertRuleGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}] RouteDeleteAlertRuleGroup", response, response.Code())
	}
}

// NewRouteDeleteAlertRuleGroupNoContent creates a RouteDeleteAlertRuleGroupNoContent with default headers values
func NewRouteDeleteAlertRuleGroupNoContent() *RouteDeleteAlertRuleGroupNoContent {
	return &RouteDeleteAlertRuleGroupNoContent{}
}

/*
RouteDeleteAlertRuleGroupNoContent describes a response with status code 204, with default header values.

	The alert rule group was deleted successfully.
*/
type RouteDeleteAlertRuleGroupNoContent struct {
}

// IsSuccess returns true when this route delete alert rule group no content response has a 2xx status code
func (o *RouteDeleteAlertRuleGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route delete alert rule group no content response has a 3xx status code
func (o *RouteDeleteAlertRuleGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route delete alert rule group no content response has a 4xx status code
func (o *RouteDeleteAlertRuleGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this route delete alert rule group no content response has a 5xx status code
func (o *RouteDeleteAlertRuleGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this route delete alert rule group no content response a status code equal to that given
func (o *RouteDeleteAlertRuleGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the route delete alert rule group no content response
func (o *RouteDeleteAlertRuleGroupNoContent) Code() int {
	return 204
}

func (o *RouteDeleteAlertRuleGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routeDeleteAlertRuleGroupNoContent", 204)
}

func (o *RouteDeleteAlertRuleGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routeDeleteAlertRuleGroupNoContent", 204)
}

func (o *RouteDeleteAlertRuleGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRouteDeleteAlertRuleGroupForbidden creates a RouteDeleteAlertRuleGroupForbidden with default headers values
func NewRouteDeleteAlertRuleGroupForbidden() *RouteDeleteAlertRuleGroupForbidden {
	return &RouteDeleteAlertRuleGroupForbidden{}
}

/*
RouteDeleteAlertRuleGroupForbidden describes a response with status code 403, with default header values.

ForbiddenError
*/
type RouteDeleteAlertRuleGroupForbidden struct {
	Payload *models.ForbiddenError
}

// IsSuccess returns true when this route delete alert rule group forbidden response has a 2xx status code
func (o *RouteDeleteAlertRuleGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this route delete alert rule group forbidden response has a 3xx status code
func (o *RouteDeleteAlertRuleGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route delete alert rule group forbidden response has a 4xx status code
func (o *RouteDeleteAlertRuleGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this route delete alert rule group forbidden response has a 5xx status code
func (o *RouteDeleteAlertRuleGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this route delete alert rule group forbidden response a status code equal to that given
func (o *RouteDeleteAlertRuleGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the route delete alert rule group forbidden response
func (o *RouteDeleteAlertRuleGroupForbidden) Code() int {
	return 403
}

func (o *RouteDeleteAlertRuleGroupForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routeDeleteAlertRuleGroupForbidden %s", 403, payload)
}

func (o *RouteDeleteAlertRuleGroupForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routeDeleteAlertRuleGroupForbidden %s", 403, payload)
}

func (o *RouteDeleteAlertRuleGroupForbidden) GetPayload() *models.ForbiddenError {
	return o.Payload
}

func (o *RouteDeleteAlertRuleGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteDeleteAlertRuleGroupNotFound creates a RouteDeleteAlertRuleGroupNotFound with default headers values
func NewRouteDeleteAlertRuleGroupNotFound() *RouteDeleteAlertRuleGroupNotFound {
	return &RouteDeleteAlertRuleGroupNotFound{}
}

/*
RouteDeleteAlertRuleGroupNotFound describes a response with status code 404, with default header values.

NotFound
*/
type RouteDeleteAlertRuleGroupNotFound struct {
	Payload models.NotFound
}

// IsSuccess returns true when this route delete alert rule group not found response has a 2xx status code
func (o *RouteDeleteAlertRuleGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this route delete alert rule group not found response has a 3xx status code
func (o *RouteDeleteAlertRuleGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route delete alert rule group not found response has a 4xx status code
func (o *RouteDeleteAlertRuleGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this route delete alert rule group not found response has a 5xx status code
func (o *RouteDeleteAlertRuleGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this route delete alert rule group not found response a status code equal to that given
func (o *RouteDeleteAlertRuleGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the route delete alert rule group not found response
func (o *RouteDeleteAlertRuleGroupNotFound) Code() int {
	return 404
}

func (o *RouteDeleteAlertRuleGroupNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routeDeleteAlertRuleGroupNotFound %s", 404, payload)
}

func (o *RouteDeleteAlertRuleGroupNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routeDeleteAlertRuleGroupNotFound %s", 404, payload)
}

func (o *RouteDeleteAlertRuleGroupNotFound) GetPayload() models.NotFound {
	return o.Payload
}

func (o *RouteDeleteAlertRuleGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
