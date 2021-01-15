package main

import (
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/kashifkhan0771/HostService"
	"github.com/kashifkhan0771/HostService/config"
	"github.com/kashifkhan0771/HostService/gen/restapi"
	"github.com/kashifkhan0771/HostService/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if serverErr := server.Shutdown(); serverErr != nil {
			panic(serverErr)
		}
	}()

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
