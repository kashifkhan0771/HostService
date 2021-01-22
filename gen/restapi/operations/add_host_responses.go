// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kashifkhan0771/HostService/gen/models"
)

// AddHostCreatedCode is the HTTP code returned for type AddHostCreated
const AddHostCreatedCode int = 201

/*AddHostCreated host added

swagger:response addHostCreated
*/
type AddHostCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Host `json:"body,omitempty"`
}

// NewAddHostCreated creates AddHostCreated with default headers values
func NewAddHostCreated() *AddHostCreated {

	return &AddHostCreated{}
}

// WithPayload adds the payload to the add host created response
func (o *AddHostCreated) WithPayload(payload *models.Host) *AddHostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add host created response
func (o *AddHostCreated) SetPayload(payload *models.Host) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddHostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddHostBadRequestCode is the HTTP code returned for type AddHostBadRequest
const AddHostBadRequestCode int = 400

/*AddHostBadRequest bad request

swagger:response addHostBadRequest
*/
type AddHostBadRequest struct {
}

// NewAddHostBadRequest creates AddHostBadRequest with default headers values
func NewAddHostBadRequest() *AddHostBadRequest {

	return &AddHostBadRequest{}
}

// WriteResponse to the client
func (o *AddHostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddHostConflictCode is the HTTP code returned for type AddHostConflict
const AddHostConflictCode int = 409

/*AddHostConflict host already exist

swagger:response addHostConflict
*/
type AddHostConflict struct {
}

// NewAddHostConflict creates AddHostConflict with default headers values
func NewAddHostConflict() *AddHostConflict {

	return &AddHostConflict{}
}

// WriteResponse to the client
func (o *AddHostConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// AddHostInternalServerErrorCode is the HTTP code returned for type AddHostInternalServerError
const AddHostInternalServerErrorCode int = 500

/*AddHostInternalServerError internal server error

swagger:response addHostInternalServerError
*/
type AddHostInternalServerError struct {
}

// NewAddHostInternalServerError creates AddHostInternalServerError with default headers values
func NewAddHostInternalServerError() *AddHostInternalServerError {

	return &AddHostInternalServerError{}
}

// WriteResponse to the client
func (o *AddHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
