package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/HostService"
	domainErr "github.com/kashifkhan0771/HostService/errors"
	"github.com/kashifkhan0771/HostService/gen/restapi/operations"
)

// NewEditHost function is for edit Host.
func NewEditHost(rt *runtime.Runtime) operations.EditHostHandler {
	return &editHost{rt: rt}
}

type editHost struct {
	rt *runtime.Runtime
}

// Handle the put Host request.
func (d *editHost) Handle(params operations.EditHostParams) middleware.Responder {
	host, _ := d.rt.Service().RetrieveHost(params.ID)
	if err := d.rt.Service().UpdateHost(host, params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteHostBadRequest()
		default:
			return operations.NewDeleteHostInternalServerError()
		}
	}

	return operations.NewEditHostOK()
}
