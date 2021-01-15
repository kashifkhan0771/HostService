// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteHostHandlerFunc turns a function with the right signature into a delete host handler
type DeleteHostHandlerFunc func(DeleteHostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHostHandlerFunc) Handle(params DeleteHostParams) middleware.Responder {
	return fn(params)
}

// DeleteHostHandler interface for that can handle valid delete host params
type DeleteHostHandler interface {
	Handle(DeleteHostParams) middleware.Responder
}

// NewDeleteHost creates a new http.Handler for the delete host operation
func NewDeleteHost(ctx *middleware.Context, handler DeleteHostHandler) *DeleteHost {
	return &DeleteHost{Context: ctx, Handler: handler}
}

/*DeleteHost swagger:route DELETE /host/{ID} deleteHost

DeleteHost delete host API

*/
type DeleteHost struct {
	Context *middleware.Context
	Handler DeleteHostHandler
}

func (o *DeleteHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteHostParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}