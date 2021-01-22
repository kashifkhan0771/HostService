package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/kashifkhan0771/HostService"
	"github.com/kashifkhan0771/HostService/gen/restapi/operations"
)

// Handler replaces swagger handler.
type Handler *operations.HostServiceAPI

// NewHandler overrides swagger api handlers.
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewHostServiceAPI(spec)

	// Host handlers
	handler.AddHostHandler = NewAddHost(rt)
	handler.GetHostHandler = NewGetHost(rt)
	handler.EditHostHandler = NewEditHost(rt)
	handler.DeleteHostHandler = NewDeleteHost(rt)

	return handler
}
