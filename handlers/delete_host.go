package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/HostService"
	domainErr "github.com/kashifkhan0771/HostService/errors"
	"github.com/kashifkhan0771/HostService/gen/restapi/operations"
)

// NewDeleteHost function will delete the Host.
func NewDeleteHost(rt *runtime.Runtime) operations.DeleteHostHandler {
	return &deleteHost{rt: rt}
}

type deleteHost struct {
	rt *runtime.Runtime
}

// Handle the delete entry request.
func (d *deleteHost) Handle(params operations.DeleteHostParams) middleware.Responder {
	if err := d.rt.Service().DeleteHost(params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteHostBadRequest()
		default:
			return operations.NewDeleteHostInternalServerError()
		}
	}

	return operations.NewDeleteHostNoContent()
}
