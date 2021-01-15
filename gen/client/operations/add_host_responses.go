// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kashifkhan0771/HostService/gen/models"
)

// AddHostReader is a Reader for the AddHost structure.
type AddHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddHostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddHostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddHostCreated creates a AddHostCreated with default headers values
func NewAddHostCreated() *AddHostCreated {
	return &AddHostCreated{}
}

/*AddHostCreated handles this case with default header values.

host added
*/
type AddHostCreated struct {
	Payload *models.Host
}

func (o *AddHostCreated) Error() string {
	return fmt.Sprintf("[POST /host][%d] addHostCreated  %+v", 201, o.Payload)
}

func (o *AddHostCreated) GetPayload() *models.Host {
	return o.Payload
}

func (o *AddHostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddHostBadRequest creates a AddHostBadRequest with default headers values
func NewAddHostBadRequest() *AddHostBadRequest {
	return &AddHostBadRequest{}
}

/*AddHostBadRequest handles this case with default header values.

bad request
*/
type AddHostBadRequest struct {
}

func (o *AddHostBadRequest) Error() string {
	return fmt.Sprintf("[POST /host][%d] addHostBadRequest ", 400)
}

func (o *AddHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddHostConflict creates a AddHostConflict with default headers values
func NewAddHostConflict() *AddHostConflict {
	return &AddHostConflict{}
}

/*AddHostConflict handles this case with default header values.

host already exist
*/
type AddHostConflict struct {
}

func (o *AddHostConflict) Error() string {
	return fmt.Sprintf("[POST /host][%d] addHostConflict ", 409)
}

func (o *AddHostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddHostInternalServerError creates a AddHostInternalServerError with default headers values
func NewAddHostInternalServerError() *AddHostInternalServerError {
	return &AddHostInternalServerError{}
}

/*AddHostInternalServerError handles this case with default header values.

internal server error
*/
type AddHostInternalServerError struct {
}

func (o *AddHostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /host][%d] addHostInternalServerError ", 500)
}

func (o *AddHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}