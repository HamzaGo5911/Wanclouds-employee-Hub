// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name WancloudsEmpolyeeHub --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.WancloudsEmpolyeeHubAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WancloudsEmpolyeeHubAPI) http.Handler {
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

	if api.AddEmployeeHandler == nil {
		api.AddEmployeeHandler = operations.AddEmployeeHandlerFunc(func(params operations.AddEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddEmployee has not yet been implemented")
		})
	}
	if api.AddHomeHandler == nil {
		api.AddHomeHandler = operations.AddHomeHandlerFunc(func(params operations.AddHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddHome has not yet been implemented")
		})
	}
	if api.AddOfficeHandler == nil {
		api.AddOfficeHandler = operations.AddOfficeHandlerFunc(func(params operations.AddOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddOffice has not yet been implemented")
		})
	}
	if api.DeleteEmployeeHandler == nil {
		api.DeleteEmployeeHandler = operations.DeleteEmployeeHandlerFunc(func(params operations.DeleteEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteEmployee has not yet been implemented")
		})
	}
	if api.DeleteHomeHandler == nil {
		api.DeleteHomeHandler = operations.DeleteHomeHandlerFunc(func(params operations.DeleteHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteHome has not yet been implemented")
		})
	}
	if api.DeleteOfficeHandler == nil {
		api.DeleteOfficeHandler = operations.DeleteOfficeHandlerFunc(func(params operations.DeleteOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteOffice has not yet been implemented")
		})
	}
	if api.GetEmployeeHandler == nil {
		api.GetEmployeeHandler = operations.GetEmployeeHandlerFunc(func(params operations.GetEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetEmployee has not yet been implemented")
		})
	}
	if api.GetHomeHandler == nil {
		api.GetHomeHandler = operations.GetHomeHandlerFunc(func(params operations.GetHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHome has not yet been implemented")
		})
	}
	if api.GetOfficeHandler == nil {
		api.GetOfficeHandler = operations.GetOfficeHandlerFunc(func(params operations.GetOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetOffice has not yet been implemented")
		})
	}
	if api.ListHomesHandler == nil {
		api.ListHomesHandler = operations.ListHomesHandlerFunc(func(params operations.ListHomesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListHomes has not yet been implemented")
		})
	}
	if api.ListOfficesHandler == nil {
		api.ListOfficesHandler = operations.ListOfficesHandlerFunc(func(params operations.ListOfficesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListOffices has not yet been implemented")
		})
	}
	if api.UpdateEmployeeHandler == nil {
		api.UpdateEmployeeHandler = operations.UpdateEmployeeHandlerFunc(func(params operations.UpdateEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateEmployee has not yet been implemented")
		})
	}
	if api.UpdateHomeHandler == nil {
		api.UpdateHomeHandler = operations.UpdateHomeHandlerFunc(func(params operations.UpdateHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateHome has not yet been implemented")
		})
	}
	if api.UpdateOfficeHandler == nil {
		api.UpdateOfficeHandler = operations.UpdateOfficeHandlerFunc(func(params operations.UpdateOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateOffice has not yet been implemented")
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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
