// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// DeleteAlertNotificationChannelReader is a Reader for the DeleteAlertNotificationChannel structure.
type DeleteAlertNotificationChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertNotificationChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertNotificationChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAlertNotificationChannelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAlertNotificationChannelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAlertNotificationChannelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAlertNotificationChannelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /alert-notifications/{notification_channel_id}] deleteAlertNotificationChannel", response, response.Code())
	}
}

// NewDeleteAlertNotificationChannelOK creates a DeleteAlertNotificationChannelOK with default headers values
func NewDeleteAlertNotificationChannelOK() *DeleteAlertNotificationChannelOK {
	return &DeleteAlertNotificationChannelOK{}
}

/*
DeleteAlertNotificationChannelOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteAlertNotificationChannelOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete alert notification channel Ok response has a 2xx status code
func (o *DeleteAlertNotificationChannelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete alert notification channel Ok response has a 3xx status code
func (o *DeleteAlertNotificationChannelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel Ok response has a 4xx status code
func (o *DeleteAlertNotificationChannelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alert notification channel Ok response has a 5xx status code
func (o *DeleteAlertNotificationChannelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel Ok response a status code equal to that given
func (o *DeleteAlertNotificationChannelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete alert notification channel Ok response
func (o *DeleteAlertNotificationChannelOK) Code() int {
	return 200
}

func (o *DeleteAlertNotificationChannelOK) Error() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelOk  %+v", 200, o.Payload)
}

func (o *DeleteAlertNotificationChannelOK) String() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelOk  %+v", 200, o.Payload)
}

func (o *DeleteAlertNotificationChannelOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelUnauthorized creates a DeleteAlertNotificationChannelUnauthorized with default headers values
func NewDeleteAlertNotificationChannelUnauthorized() *DeleteAlertNotificationChannelUnauthorized {
	return &DeleteAlertNotificationChannelUnauthorized{}
}

/*
DeleteAlertNotificationChannelUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteAlertNotificationChannelUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel unauthorized response has a 2xx status code
func (o *DeleteAlertNotificationChannelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel unauthorized response has a 3xx status code
func (o *DeleteAlertNotificationChannelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel unauthorized response has a 4xx status code
func (o *DeleteAlertNotificationChannelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alert notification channel unauthorized response has a 5xx status code
func (o *DeleteAlertNotificationChannelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel unauthorized response a status code equal to that given
func (o *DeleteAlertNotificationChannelUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete alert notification channel unauthorized response
func (o *DeleteAlertNotificationChannelUnauthorized) Code() int {
	return 401
}

func (o *DeleteAlertNotificationChannelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAlertNotificationChannelUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAlertNotificationChannelUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelForbidden creates a DeleteAlertNotificationChannelForbidden with default headers values
func NewDeleteAlertNotificationChannelForbidden() *DeleteAlertNotificationChannelForbidden {
	return &DeleteAlertNotificationChannelForbidden{}
}

/*
DeleteAlertNotificationChannelForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteAlertNotificationChannelForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel forbidden response has a 2xx status code
func (o *DeleteAlertNotificationChannelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel forbidden response has a 3xx status code
func (o *DeleteAlertNotificationChannelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel forbidden response has a 4xx status code
func (o *DeleteAlertNotificationChannelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alert notification channel forbidden response has a 5xx status code
func (o *DeleteAlertNotificationChannelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel forbidden response a status code equal to that given
func (o *DeleteAlertNotificationChannelForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete alert notification channel forbidden response
func (o *DeleteAlertNotificationChannelForbidden) Code() int {
	return 403
}

func (o *DeleteAlertNotificationChannelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAlertNotificationChannelForbidden) String() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAlertNotificationChannelForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelNotFound creates a DeleteAlertNotificationChannelNotFound with default headers values
func NewDeleteAlertNotificationChannelNotFound() *DeleteAlertNotificationChannelNotFound {
	return &DeleteAlertNotificationChannelNotFound{}
}

/*
DeleteAlertNotificationChannelNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteAlertNotificationChannelNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel not found response has a 2xx status code
func (o *DeleteAlertNotificationChannelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel not found response has a 3xx status code
func (o *DeleteAlertNotificationChannelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel not found response has a 4xx status code
func (o *DeleteAlertNotificationChannelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete alert notification channel not found response has a 5xx status code
func (o *DeleteAlertNotificationChannelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete alert notification channel not found response a status code equal to that given
func (o *DeleteAlertNotificationChannelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete alert notification channel not found response
func (o *DeleteAlertNotificationChannelNotFound) Code() int {
	return 404
}

func (o *DeleteAlertNotificationChannelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAlertNotificationChannelNotFound) String() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAlertNotificationChannelNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertNotificationChannelInternalServerError creates a DeleteAlertNotificationChannelInternalServerError with default headers values
func NewDeleteAlertNotificationChannelInternalServerError() *DeleteAlertNotificationChannelInternalServerError {
	return &DeleteAlertNotificationChannelInternalServerError{}
}

/*
DeleteAlertNotificationChannelInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteAlertNotificationChannelInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete alert notification channel internal server error response has a 2xx status code
func (o *DeleteAlertNotificationChannelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete alert notification channel internal server error response has a 3xx status code
func (o *DeleteAlertNotificationChannelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete alert notification channel internal server error response has a 4xx status code
func (o *DeleteAlertNotificationChannelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete alert notification channel internal server error response has a 5xx status code
func (o *DeleteAlertNotificationChannelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete alert notification channel internal server error response a status code equal to that given
func (o *DeleteAlertNotificationChannelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete alert notification channel internal server error response
func (o *DeleteAlertNotificationChannelInternalServerError) Code() int {
	return 500
}

func (o *DeleteAlertNotificationChannelInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAlertNotificationChannelInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /alert-notifications/{notification_channel_id}][%d] deleteAlertNotificationChannelInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAlertNotificationChannelInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAlertNotificationChannelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
