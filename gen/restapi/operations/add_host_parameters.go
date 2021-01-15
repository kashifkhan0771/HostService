// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kashifkhan0771/HostService/gen/models"
)

// NewAddHostParams creates a new AddHostParams object
// no default values defined in spec.
func NewAddHostParams() AddHostParams {

	return AddHostParams{}
}

// AddHostParams contains all the bound params for the add host operation
// typically these are obtained from a http.Request
//
// swagger:parameters addHost
type AddHostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*host's detail
	  In: body
	*/
	Host *models.Host
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddHostParams() beforehand.
func (o *AddHostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Host
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("host", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Host = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}