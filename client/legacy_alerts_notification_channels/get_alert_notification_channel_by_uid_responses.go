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

// GetAlertNotificationChannelByUIDReader is a Reader for the GetAlertNotificationChannelByUID structure.
type GetAlertNotificationChannelByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertNotificationChannelByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertNotificationChannelByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAlertNotificationChannelByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertNotificationChannelByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertNotificationChannelByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertNotificationChannelByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /alert-notifications/uid/{notification_channel_uid}] getAlertNotificationChannelByUID", response, response.Code())
	}
}

// NewGetAlertNotificationChannelByUIDOK creates a GetAlertNotificationChannelByUIDOK with default headers values
func NewGetAlertNotificationChannelByUIDOK() *GetAlertNotificationChannelByUIDOK {
	return &GetAlertNotificationChannelByUIDOK{}
}

/*
GetAlertNotificationChannelByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAlertNotificationChannelByUIDOK struct {
	Payload *models.AlertNotification
}

// IsSuccess returns true when this get alert notification channel by Uid Ok response has a 2xx status code
func (o *GetAlertNotificationChannelByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert notification channel by Uid Ok response has a 3xx status code
func (o *GetAlertNotificationChannelByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification channel by Uid Ok response has a 4xx status code
func (o *GetAlertNotificationChannelByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert notification channel by Uid Ok response has a 5xx status code
func (o *GetAlertNotificationChannelByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification channel by Uid Ok response a status code equal to that given
func (o *GetAlertNotificationChannelByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert notification channel by Uid Ok response
func (o *GetAlertNotificationChannelByUIDOK) Code() int {
	return 200
}

func (o *GetAlertNotificationChannelByUIDOK) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidOk  %+v", 200, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDOK) String() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidOk  %+v", 200, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDOK) GetPayload() *models.AlertNotification {
	return o.Payload
}

func (o *GetAlertNotificationChannelByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByUIDUnauthorized creates a GetAlertNotificationChannelByUIDUnauthorized with default headers values
func NewGetAlertNotificationChannelByUIDUnauthorized() *GetAlertNotificationChannelByUIDUnauthorized {
	return &GetAlertNotificationChannelByUIDUnauthorized{}
}

/*
GetAlertNotificationChannelByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAlertNotificationChannelByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification channel by Uid unauthorized response has a 2xx status code
func (o *GetAlertNotificationChannelByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification channel by Uid unauthorized response has a 3xx status code
func (o *GetAlertNotificationChannelByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification channel by Uid unauthorized response has a 4xx status code
func (o *GetAlertNotificationChannelByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert notification channel by Uid unauthorized response has a 5xx status code
func (o *GetAlertNotificationChannelByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification channel by Uid unauthorized response a status code equal to that given
func (o *GetAlertNotificationChannelByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get alert notification channel by Uid unauthorized response
func (o *GetAlertNotificationChannelByUIDUnauthorized) Code() int {
	return 401
}

func (o *GetAlertNotificationChannelByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByUIDForbidden creates a GetAlertNotificationChannelByUIDForbidden with default headers values
func NewGetAlertNotificationChannelByUIDForbidden() *GetAlertNotificationChannelByUIDForbidden {
	return &GetAlertNotificationChannelByUIDForbidden{}
}

/*
GetAlertNotificationChannelByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetAlertNotificationChannelByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification channel by Uid forbidden response has a 2xx status code
func (o *GetAlertNotificationChannelByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification channel by Uid forbidden response has a 3xx status code
func (o *GetAlertNotificationChannelByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification channel by Uid forbidden response has a 4xx status code
func (o *GetAlertNotificationChannelByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert notification channel by Uid forbidden response has a 5xx status code
func (o *GetAlertNotificationChannelByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification channel by Uid forbidden response a status code equal to that given
func (o *GetAlertNotificationChannelByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get alert notification channel by Uid forbidden response
func (o *GetAlertNotificationChannelByUIDForbidden) Code() int {
	return 403
}

func (o *GetAlertNotificationChannelByUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDForbidden) String() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByUIDNotFound creates a GetAlertNotificationChannelByUIDNotFound with default headers values
func NewGetAlertNotificationChannelByUIDNotFound() *GetAlertNotificationChannelByUIDNotFound {
	return &GetAlertNotificationChannelByUIDNotFound{}
}

/*
GetAlertNotificationChannelByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetAlertNotificationChannelByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification channel by Uid not found response has a 2xx status code
func (o *GetAlertNotificationChannelByUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification channel by Uid not found response has a 3xx status code
func (o *GetAlertNotificationChannelByUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification channel by Uid not found response has a 4xx status code
func (o *GetAlertNotificationChannelByUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert notification channel by Uid not found response has a 5xx status code
func (o *GetAlertNotificationChannelByUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert notification channel by Uid not found response a status code equal to that given
func (o *GetAlertNotificationChannelByUIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get alert notification channel by Uid not found response
func (o *GetAlertNotificationChannelByUIDNotFound) Code() int {
	return 404
}

func (o *GetAlertNotificationChannelByUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDNotFound) String() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotificationChannelByUIDInternalServerError creates a GetAlertNotificationChannelByUIDInternalServerError with default headers values
func NewGetAlertNotificationChannelByUIDInternalServerError() *GetAlertNotificationChannelByUIDInternalServerError {
	return &GetAlertNotificationChannelByUIDInternalServerError{}
}

/*
GetAlertNotificationChannelByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAlertNotificationChannelByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get alert notification channel by Uid internal server error response has a 2xx status code
func (o *GetAlertNotificationChannelByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert notification channel by Uid internal server error response has a 3xx status code
func (o *GetAlertNotificationChannelByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert notification channel by Uid internal server error response has a 4xx status code
func (o *GetAlertNotificationChannelByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert notification channel by Uid internal server error response has a 5xx status code
func (o *GetAlertNotificationChannelByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alert notification channel by Uid internal server error response a status code equal to that given
func (o *GetAlertNotificationChannelByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get alert notification channel by Uid internal server error response
func (o *GetAlertNotificationChannelByUIDInternalServerError) Code() int {
	return 500
}

func (o *GetAlertNotificationChannelByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /alert-notifications/uid/{notification_channel_uid}][%d] getAlertNotificationChannelByUidInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertNotificationChannelByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAlertNotificationChannelByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
