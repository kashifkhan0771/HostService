// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteHostReader is a Reader for the DeleteHost structure.
type DeleteHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteHostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteHostNoContent creates a DeleteHostNoContent with default headers values
func NewDeleteHostNoContent() *DeleteHostNoContent {
	return &DeleteHostNoContent{}
}

/*DeleteHostNoContent handles this case with default header values.

host deleted
*/
type DeleteHostNoContent struct {
}

func (o *DeleteHostNoContent) Error() string {
	return fmt.Sprintf("[DELETE /host/{ID}][%d] deleteHostNoContent ", 204)
}

func (o *DeleteHostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHostBadRequest creates a DeleteHostBadRequest with default headers values
func NewDeleteHostBadRequest() *DeleteHostBadRequest {
	return &DeleteHostBadRequest{}
}

/*DeleteHostBadRequest handles this case with default header values.

bad request
*/
type DeleteHostBadRequest struct {
}

func (o *DeleteHostBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /host/{ID}][%d] deleteHostBadRequest ", 400)
}

func (o *DeleteHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHostInternalServerError creates a DeleteHostInternalServerError with default headers values
func NewDeleteHostInternalServerError() *DeleteHostInternalServerError {
	return &DeleteHostInternalServerError{}
}

/*DeleteHostInternalServerError handles this case with default header values.

internal server error
*/
type DeleteHostInternalServerError struct {
}

func (o *DeleteHostInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /host/{ID}][%d] deleteHostInternalServerError ", 500)
}

func (o *DeleteHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
