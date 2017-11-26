// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/opus47/cloud/api/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../openapi.yml

func configureFlags(api *operations.Opus47API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Opus47API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DeleteComposersIDHandler = operations.DeleteComposersIDHandlerFunc(func(params operations.DeleteComposersIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteComposersID has not yet been implemented")
	})
	api.DeleteKeysIDHandler = operations.DeleteKeysIDHandlerFunc(func(params operations.DeleteKeysIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteKeysID has not yet been implemented")
	})
	api.DeleteMusiciansIDHandler = operations.DeleteMusiciansIDHandlerFunc(func(params operations.DeleteMusiciansIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteMusiciansID has not yet been implemented")
	})
	api.DeletePartsIDHandler = operations.DeletePartsIDHandlerFunc(func(params operations.DeletePartsIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeletePartsID has not yet been implemented")
	})
	api.DeletePerformancesIDHandler = operations.DeletePerformancesIDHandlerFunc(func(params operations.DeletePerformancesIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeletePerformancesID has not yet been implemented")
	})
	api.DeletePiecesIDHandler = operations.DeletePiecesIDHandlerFunc(func(params operations.DeletePiecesIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeletePiecesID has not yet been implemented")
	})
	api.DeleteRecordingsIDHandler = operations.DeleteRecordingsIDHandlerFunc(func(params operations.DeleteRecordingsIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteRecordingsID has not yet been implemented")
	})
	api.GetComposersHandler = operations.GetComposersHandlerFunc(func(params operations.GetComposersParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetComposers has not yet been implemented")
	})
	api.GetComposersIDHandler = operations.GetComposersIDHandlerFunc(func(params operations.GetComposersIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetComposersID has not yet been implemented")
	})
	api.GetKeysHandler = operations.GetKeysHandlerFunc(func(params operations.GetKeysParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetKeys has not yet been implemented")
	})
	api.GetKeysIDHandler = operations.GetKeysIDHandlerFunc(func(params operations.GetKeysIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetKeysID has not yet been implemented")
	})
	api.GetMusiciansHandler = operations.GetMusiciansHandlerFunc(func(params operations.GetMusiciansParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetMusicians has not yet been implemented")
	})
	api.GetMusiciansIDHandler = operations.GetMusiciansIDHandlerFunc(func(params operations.GetMusiciansIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetMusiciansID has not yet been implemented")
	})
	api.GetPartsHandler = operations.GetPartsHandlerFunc(func(params operations.GetPartsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetParts has not yet been implemented")
	})
	api.GetPartsIDHandler = operations.GetPartsIDHandlerFunc(func(params operations.GetPartsIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetPartsID has not yet been implemented")
	})
	api.GetPerformancesHandler = operations.GetPerformancesHandlerFunc(func(params operations.GetPerformancesParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetPerformances has not yet been implemented")
	})
	api.GetPerformancesIDHandler = operations.GetPerformancesIDHandlerFunc(func(params operations.GetPerformancesIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetPerformancesID has not yet been implemented")
	})
	api.GetPiecesHandler = operations.GetPiecesHandlerFunc(func(params operations.GetPiecesParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetPieces has not yet been implemented")
	})

	api.GetPiecesIDHandler = operations.GetPiecesIDHandlerFunc(handleGetPiecesId)

	api.GetPiecesSearchHandler = operations.GetPiecesSearchHandlerFunc(handleGetPiecesSearch)

	api.GetRecordingsHandler = operations.GetRecordingsHandlerFunc(func(params operations.GetRecordingsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetRecordings has not yet been implemented")
	})
	api.GetRecordingsIDHandler = operations.GetRecordingsIDHandlerFunc(func(params operations.GetRecordingsIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetRecordingsID has not yet been implemented")
	})
	api.PutComposersIDHandler = operations.PutComposersIDHandlerFunc(func(params operations.PutComposersIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutComposersID has not yet been implemented")
	})
	api.PutKeysIDHandler = operations.PutKeysIDHandlerFunc(func(params operations.PutKeysIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutKeysID has not yet been implemented")
	})
	api.PutMusiciansIDHandler = operations.PutMusiciansIDHandlerFunc(func(params operations.PutMusiciansIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutMusiciansID has not yet been implemented")
	})
	api.PutPartsIDHandler = operations.PutPartsIDHandlerFunc(func(params operations.PutPartsIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutPartsID has not yet been implemented")
	})
	api.PutPerformancesIDHandler = operations.PutPerformancesIDHandlerFunc(func(params operations.PutPerformancesIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutPerformancesID has not yet been implemented")
	})
	api.PutPiecesIDHandler = operations.PutPiecesIDHandlerFunc(func(params operations.PutPiecesIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutPiecesID has not yet been implemented")
	})
	api.PutRecordingsIDHandler = operations.PutRecordingsIDHandlerFunc(func(params operations.PutRecordingsIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutRecordingsID has not yet been implemented")
	})

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
func configureServer(s *graceful.Server, scheme, addr string) {
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
