package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/HostService"
	"github.com/kashifkhan0771/HostService/gen/restapi/operations"
	"github.com/kashifkhan0771/HostService/models"
)

// NewAddHost function will add new host.
func NewAddHost(rt *runtime.Runtime) operations.AddHostHandler {
	return &addHost{rt: rt}
}

type addHost struct {
	rt *runtime.Runtime
}

// Handle the add student request.
func (d *addHost) Handle(params operations.AddHostParams) middleware.Responder {
	host := models.Host{
		ID:       params.Host.ID,
		IP:       params.Host.IP,
		Name:     params.Host.Name,
		Metadata: params.Host.MetaData,
		Status:   params.Host.Status,
	}

	id, err := d.rt.Service().AddHost(&host)
	if err != nil {
		log().Errorf("failed to add student: %s", err)

		return operations.NewAddHostBadRequest()
	}

	params.Host.ID = id

	return operations.NewAddHostCreated().WithPayload(params.Host)
}
