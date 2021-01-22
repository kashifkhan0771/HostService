// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kashifkhan0771/HostService/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name HostService --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.HostServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HostServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddHostHandler == nil {
		api.AddHostHandler = operations.AddHostHandlerFunc(func(params operations.AddHostParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddHost has not yet been implemented")
		})
	}
	if api.DeleteHostHandler == nil {
		api.DeleteHostHandler = operations.DeleteHostHandlerFunc(func(params operations.DeleteHostParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteHost has not yet been implemented")
		})
	}
	if api.EditHostHandler == nil {
		api.EditHostHandler = operations.EditHostHandlerFunc(func(params operations.EditHostParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.EditHost has not yet been implemented")
		})
	}
	if api.GetHostHandler == nil {
		api.GetHostHandler = operations.GetHostHandlerFunc(func(params operations.GetHostParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHost has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
