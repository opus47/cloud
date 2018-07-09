// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewOpus47API creates a new Opus47 instance
func NewOpus47API(spec *loads.Document) *Opus47API {
	return &Opus47API{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		DeleteComposersIDHandler: DeleteComposersIDHandlerFunc(func(params DeleteComposersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteComposersID has not yet been implemented")
		}),
		DeleteKeysIDHandler: DeleteKeysIDHandlerFunc(func(params DeleteKeysIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteKeysID has not yet been implemented")
		}),
		DeleteMusiciansIDHandler: DeleteMusiciansIDHandlerFunc(func(params DeleteMusiciansIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteMusiciansID has not yet been implemented")
		}),
		DeletePartsIDHandler: DeletePartsIDHandlerFunc(func(params DeletePartsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeletePartsID has not yet been implemented")
		}),
		DeletePerformancesIDHandler: DeletePerformancesIDHandlerFunc(func(params DeletePerformancesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeletePerformancesID has not yet been implemented")
		}),
		DeletePiecesIDHandler: DeletePiecesIDHandlerFunc(func(params DeletePiecesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeletePiecesID has not yet been implemented")
		}),
		DeleteRecordingsIDHandler: DeleteRecordingsIDHandlerFunc(func(params DeleteRecordingsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteRecordingsID has not yet been implemented")
		}),
		GetComposersHandler: GetComposersHandlerFunc(func(params GetComposersParams) middleware.Responder {
			return middleware.NotImplemented("operation GetComposers has not yet been implemented")
		}),
		GetComposersIDHandler: GetComposersIDHandlerFunc(func(params GetComposersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetComposersID has not yet been implemented")
		}),
		GetKeysHandler: GetKeysHandlerFunc(func(params GetKeysParams) middleware.Responder {
			return middleware.NotImplemented("operation GetKeys has not yet been implemented")
		}),
		GetKeysIDHandler: GetKeysIDHandlerFunc(func(params GetKeysIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetKeysID has not yet been implemented")
		}),
		GetMusiciansHandler: GetMusiciansHandlerFunc(func(params GetMusiciansParams) middleware.Responder {
			return middleware.NotImplemented("operation GetMusicians has not yet been implemented")
		}),
		GetMusiciansIDHandler: GetMusiciansIDHandlerFunc(func(params GetMusiciansIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetMusiciansID has not yet been implemented")
		}),
		GetPartsHandler: GetPartsHandlerFunc(func(params GetPartsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetParts has not yet been implemented")
		}),
		GetPartsIDHandler: GetPartsIDHandlerFunc(func(params GetPartsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPartsID has not yet been implemented")
		}),
		GetPerformancesHandler: GetPerformancesHandlerFunc(func(params GetPerformancesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPerformances has not yet been implemented")
		}),
		GetPerformancesIDHandler: GetPerformancesIDHandlerFunc(func(params GetPerformancesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPerformancesID has not yet been implemented")
		}),
		GetPiecesHandler: GetPiecesHandlerFunc(func(params GetPiecesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPieces has not yet been implemented")
		}),
		GetPiecesIDHandler: GetPiecesIDHandlerFunc(func(params GetPiecesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPiecesID has not yet been implemented")
		}),
		GetPiecesIDPerformancesHandler: GetPiecesIDPerformancesHandlerFunc(func(params GetPiecesIDPerformancesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPiecesIDPerformances has not yet been implemented")
		}),
		GetPiecesSearchHandler: GetPiecesSearchHandlerFunc(func(params GetPiecesSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation GetPiecesSearch has not yet been implemented")
		}),
		GetRecordingsHandler: GetRecordingsHandlerFunc(func(params GetRecordingsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRecordings has not yet been implemented")
		}),
		GetRecordingsIDHandler: GetRecordingsIDHandlerFunc(func(params GetRecordingsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRecordingsID has not yet been implemented")
		}),
		PostPiecesIDHandler: PostPiecesIDHandlerFunc(func(params PostPiecesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PostPiecesID has not yet been implemented")
		}),
		PutComposersIDHandler: PutComposersIDHandlerFunc(func(params PutComposersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PutComposersID has not yet been implemented")
		}),
		PutKeysIDHandler: PutKeysIDHandlerFunc(func(params PutKeysIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PutKeysID has not yet been implemented")
		}),
		PutMusiciansIDHandler: PutMusiciansIDHandlerFunc(func(params PutMusiciansIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PutMusiciansID has not yet been implemented")
		}),
		PutPartsIDHandler: PutPartsIDHandlerFunc(func(params PutPartsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation PutPartsID has not yet been implemented")
		}),
		PutPerformancesHandler: PutPerformancesHandlerFunc(func(params PutPerformancesParams) middleware.Responder {
			return middleware.NotImplemented("operation PutPerformances has not yet been implemented")
		}),
		PutPiecesHandler: PutPiecesHandlerFunc(func(params PutPiecesParams) middleware.Responder {
			return middleware.NotImplemented("operation PutPieces has not yet been implemented")
		}),
		PutRecordingsHandler: PutRecordingsHandlerFunc(func(params PutRecordingsParams) middleware.Responder {
			return middleware.NotImplemented("operation PutRecordings has not yet been implemented")
		}),
	}
}

/*Opus47API the opus47 API */
type Opus47API struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// DeleteComposersIDHandler sets the operation handler for the delete composers ID operation
	DeleteComposersIDHandler DeleteComposersIDHandler
	// DeleteKeysIDHandler sets the operation handler for the delete keys ID operation
	DeleteKeysIDHandler DeleteKeysIDHandler
	// DeleteMusiciansIDHandler sets the operation handler for the delete musicians ID operation
	DeleteMusiciansIDHandler DeleteMusiciansIDHandler
	// DeletePartsIDHandler sets the operation handler for the delete parts ID operation
	DeletePartsIDHandler DeletePartsIDHandler
	// DeletePerformancesIDHandler sets the operation handler for the delete performances ID operation
	DeletePerformancesIDHandler DeletePerformancesIDHandler
	// DeletePiecesIDHandler sets the operation handler for the delete pieces ID operation
	DeletePiecesIDHandler DeletePiecesIDHandler
	// DeleteRecordingsIDHandler sets the operation handler for the delete recordings ID operation
	DeleteRecordingsIDHandler DeleteRecordingsIDHandler
	// GetComposersHandler sets the operation handler for the get composers operation
	GetComposersHandler GetComposersHandler
	// GetComposersIDHandler sets the operation handler for the get composers ID operation
	GetComposersIDHandler GetComposersIDHandler
	// GetKeysHandler sets the operation handler for the get keys operation
	GetKeysHandler GetKeysHandler
	// GetKeysIDHandler sets the operation handler for the get keys ID operation
	GetKeysIDHandler GetKeysIDHandler
	// GetMusiciansHandler sets the operation handler for the get musicians operation
	GetMusiciansHandler GetMusiciansHandler
	// GetMusiciansIDHandler sets the operation handler for the get musicians ID operation
	GetMusiciansIDHandler GetMusiciansIDHandler
	// GetPartsHandler sets the operation handler for the get parts operation
	GetPartsHandler GetPartsHandler
	// GetPartsIDHandler sets the operation handler for the get parts ID operation
	GetPartsIDHandler GetPartsIDHandler
	// GetPerformancesHandler sets the operation handler for the get performances operation
	GetPerformancesHandler GetPerformancesHandler
	// GetPerformancesIDHandler sets the operation handler for the get performances ID operation
	GetPerformancesIDHandler GetPerformancesIDHandler
	// GetPiecesHandler sets the operation handler for the get pieces operation
	GetPiecesHandler GetPiecesHandler
	// GetPiecesIDHandler sets the operation handler for the get pieces ID operation
	GetPiecesIDHandler GetPiecesIDHandler
	// GetPiecesIDPerformancesHandler sets the operation handler for the get pieces ID performances operation
	GetPiecesIDPerformancesHandler GetPiecesIDPerformancesHandler
	// GetPiecesSearchHandler sets the operation handler for the get pieces search operation
	GetPiecesSearchHandler GetPiecesSearchHandler
	// GetRecordingsHandler sets the operation handler for the get recordings operation
	GetRecordingsHandler GetRecordingsHandler
	// GetRecordingsIDHandler sets the operation handler for the get recordings ID operation
	GetRecordingsIDHandler GetRecordingsIDHandler
	// PostPiecesIDHandler sets the operation handler for the post pieces ID operation
	PostPiecesIDHandler PostPiecesIDHandler
	// PutComposersIDHandler sets the operation handler for the put composers ID operation
	PutComposersIDHandler PutComposersIDHandler
	// PutKeysIDHandler sets the operation handler for the put keys ID operation
	PutKeysIDHandler PutKeysIDHandler
	// PutMusiciansIDHandler sets the operation handler for the put musicians ID operation
	PutMusiciansIDHandler PutMusiciansIDHandler
	// PutPartsIDHandler sets the operation handler for the put parts ID operation
	PutPartsIDHandler PutPartsIDHandler
	// PutPerformancesHandler sets the operation handler for the put performances operation
	PutPerformancesHandler PutPerformancesHandler
	// PutPiecesHandler sets the operation handler for the put pieces operation
	PutPiecesHandler PutPiecesHandler
	// PutRecordingsHandler sets the operation handler for the put recordings operation
	PutRecordingsHandler PutRecordingsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *Opus47API) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *Opus47API) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *Opus47API) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *Opus47API) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *Opus47API) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *Opus47API) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *Opus47API) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the Opus47API
func (o *Opus47API) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DeleteComposersIDHandler == nil {
		unregistered = append(unregistered, "DeleteComposersIDHandler")
	}

	if o.DeleteKeysIDHandler == nil {
		unregistered = append(unregistered, "DeleteKeysIDHandler")
	}

	if o.DeleteMusiciansIDHandler == nil {
		unregistered = append(unregistered, "DeleteMusiciansIDHandler")
	}

	if o.DeletePartsIDHandler == nil {
		unregistered = append(unregistered, "DeletePartsIDHandler")
	}

	if o.DeletePerformancesIDHandler == nil {
		unregistered = append(unregistered, "DeletePerformancesIDHandler")
	}

	if o.DeletePiecesIDHandler == nil {
		unregistered = append(unregistered, "DeletePiecesIDHandler")
	}

	if o.DeleteRecordingsIDHandler == nil {
		unregistered = append(unregistered, "DeleteRecordingsIDHandler")
	}

	if o.GetComposersHandler == nil {
		unregistered = append(unregistered, "GetComposersHandler")
	}

	if o.GetComposersIDHandler == nil {
		unregistered = append(unregistered, "GetComposersIDHandler")
	}

	if o.GetKeysHandler == nil {
		unregistered = append(unregistered, "GetKeysHandler")
	}

	if o.GetKeysIDHandler == nil {
		unregistered = append(unregistered, "GetKeysIDHandler")
	}

	if o.GetMusiciansHandler == nil {
		unregistered = append(unregistered, "GetMusiciansHandler")
	}

	if o.GetMusiciansIDHandler == nil {
		unregistered = append(unregistered, "GetMusiciansIDHandler")
	}

	if o.GetPartsHandler == nil {
		unregistered = append(unregistered, "GetPartsHandler")
	}

	if o.GetPartsIDHandler == nil {
		unregistered = append(unregistered, "GetPartsIDHandler")
	}

	if o.GetPerformancesHandler == nil {
		unregistered = append(unregistered, "GetPerformancesHandler")
	}

	if o.GetPerformancesIDHandler == nil {
		unregistered = append(unregistered, "GetPerformancesIDHandler")
	}

	if o.GetPiecesHandler == nil {
		unregistered = append(unregistered, "GetPiecesHandler")
	}

	if o.GetPiecesIDHandler == nil {
		unregistered = append(unregistered, "GetPiecesIDHandler")
	}

	if o.GetPiecesIDPerformancesHandler == nil {
		unregistered = append(unregistered, "GetPiecesIDPerformancesHandler")
	}

	if o.GetPiecesSearchHandler == nil {
		unregistered = append(unregistered, "GetPiecesSearchHandler")
	}

	if o.GetRecordingsHandler == nil {
		unregistered = append(unregistered, "GetRecordingsHandler")
	}

	if o.GetRecordingsIDHandler == nil {
		unregistered = append(unregistered, "GetRecordingsIDHandler")
	}

	if o.PostPiecesIDHandler == nil {
		unregistered = append(unregistered, "PostPiecesIDHandler")
	}

	if o.PutComposersIDHandler == nil {
		unregistered = append(unregistered, "PutComposersIDHandler")
	}

	if o.PutKeysIDHandler == nil {
		unregistered = append(unregistered, "PutKeysIDHandler")
	}

	if o.PutMusiciansIDHandler == nil {
		unregistered = append(unregistered, "PutMusiciansIDHandler")
	}

	if o.PutPartsIDHandler == nil {
		unregistered = append(unregistered, "PutPartsIDHandler")
	}

	if o.PutPerformancesHandler == nil {
		unregistered = append(unregistered, "PutPerformancesHandler")
	}

	if o.PutPiecesHandler == nil {
		unregistered = append(unregistered, "PutPiecesHandler")
	}

	if o.PutRecordingsHandler == nil {
		unregistered = append(unregistered, "PutRecordingsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *Opus47API) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *Opus47API) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *Opus47API) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *Opus47API) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *Opus47API) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *Opus47API) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the opus47 API
func (o *Opus47API) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *Opus47API) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/composers/{id}"] = NewDeleteComposersID(o.context, o.DeleteComposersIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/keys/{id}"] = NewDeleteKeysID(o.context, o.DeleteKeysIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/musicians/{id}"] = NewDeleteMusiciansID(o.context, o.DeleteMusiciansIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/parts/{id}"] = NewDeletePartsID(o.context, o.DeletePartsIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/performances/{id}"] = NewDeletePerformancesID(o.context, o.DeletePerformancesIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/pieces/{id}"] = NewDeletePiecesID(o.context, o.DeletePiecesIDHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/recordings/{id}"] = NewDeleteRecordingsID(o.context, o.DeleteRecordingsIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/composers"] = NewGetComposers(o.context, o.GetComposersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/composers/{id}"] = NewGetComposersID(o.context, o.GetComposersIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/keys"] = NewGetKeys(o.context, o.GetKeysHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/keys/{id}"] = NewGetKeysID(o.context, o.GetKeysIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/musicians"] = NewGetMusicians(o.context, o.GetMusiciansHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/musicians/{id}"] = NewGetMusiciansID(o.context, o.GetMusiciansIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/parts"] = NewGetParts(o.context, o.GetPartsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/parts/{id}"] = NewGetPartsID(o.context, o.GetPartsIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/performances"] = NewGetPerformances(o.context, o.GetPerformancesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/performances/{id}"] = NewGetPerformancesID(o.context, o.GetPerformancesIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pieces"] = NewGetPieces(o.context, o.GetPiecesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pieces/{id}"] = NewGetPiecesID(o.context, o.GetPiecesIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pieces/{id}/performances"] = NewGetPiecesIDPerformances(o.context, o.GetPiecesIDPerformancesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pieces/search"] = NewGetPiecesSearch(o.context, o.GetPiecesSearchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/recordings"] = NewGetRecordings(o.context, o.GetRecordingsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/recordings/{id}"] = NewGetRecordingsID(o.context, o.GetRecordingsIDHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pieces/{id}"] = NewPostPiecesID(o.context, o.PostPiecesIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/composers/{id}"] = NewPutComposersID(o.context, o.PutComposersIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/keys/{id}"] = NewPutKeysID(o.context, o.PutKeysIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/musicians/{id}"] = NewPutMusiciansID(o.context, o.PutMusiciansIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/parts/{id}"] = NewPutPartsID(o.context, o.PutPartsIDHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/performances"] = NewPutPerformances(o.context, o.PutPerformancesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/pieces"] = NewPutPieces(o.context, o.PutPiecesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/recordings"] = NewPutRecordings(o.context, o.PutRecordingsHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *Opus47API) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *Opus47API) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *Opus47API) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *Opus47API) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
