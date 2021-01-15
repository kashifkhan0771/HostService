package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/HostService"
	domainErr "github.com/kashifkhan0771/HostService/errors"
	"github.com/kashifkhan0771/HostService/gen/models"
	"github.com/kashifkhan0771/HostService/gen/restapi/operations"
)

// NewGetHost handles a request for retrieving Host.
func NewGetHost(rt *runtime.Runtime) operations.GetHostHandler {
	return &getHost{rt: rt}
}

type getHost struct {
	rt *runtime.Runtime
}

// Handle the get Host request.
func (d *getHost) Handle(params operations.GetHostParams) middleware.Responder {
	host, err := d.rt.Service().RetrieveHost(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetHostNotFound()
		default:
			return operations.NewGetHostInternalServerError()
		}
	}

	return operations.NewGetHostOK().WithPayload(&models.Host{
		ID:       host.ID,
		IP:       host.IP,
		Name:     host.Name,
		MetaData: host.Metadata,
		Status:   host.Status,
	})
}
