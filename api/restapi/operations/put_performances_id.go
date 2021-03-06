// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutPerformancesIDHandlerFunc turns a function with the right signature into a put performances ID handler
type PutPerformancesIDHandlerFunc func(PutPerformancesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPerformancesIDHandlerFunc) Handle(params PutPerformancesIDParams) middleware.Responder {
	return fn(params)
}

// PutPerformancesIDHandler interface for that can handle valid put performances ID params
type PutPerformancesIDHandler interface {
	Handle(PutPerformancesIDParams) middleware.Responder
}

// NewPutPerformancesID creates a new http.Handler for the put performances ID operation
func NewPutPerformancesID(ctx *middleware.Context, handler PutPerformancesIDHandler) *PutPerformancesID {
	return &PutPerformancesID{Context: ctx, Handler: handler}
}

/*PutPerformancesID swagger:route PUT /performances/{id} putPerformancesId

Add or update a performance

Add or update a performance

*/
type PutPerformancesID struct {
	Context *middleware.Context
	Handler PutPerformancesIDHandler
}

func (o *PutPerformancesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutPerformancesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
