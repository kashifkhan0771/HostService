// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddHostHandlerFunc turns a function with the right signature into a add host handler
type AddHostHandlerFunc func(AddHostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddHostHandlerFunc) Handle(params AddHostParams) middleware.Responder {
	return fn(params)
}

// AddHostHandler interface for that can handle valid add host params
type AddHostHandler interface {
	Handle(AddHostParams) middleware.Responder
}

// NewAddHost creates a new http.Handler for the add host operation
func NewAddHost(ctx *middleware.Context, handler AddHostHandler) *AddHost {
	return &AddHost{Context: ctx, Handler: handler}
}

/*AddHost swagger:route POST /host addHost

AddHost add host API

*/
type AddHost struct {
	Context *middleware.Context
	Handler AddHostHandler
}

func (o *AddHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddHostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
