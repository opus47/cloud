// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMusiciansIDHandlerFunc turns a function with the right signature into a get musicians ID handler
type GetMusiciansIDHandlerFunc func(GetMusiciansIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMusiciansIDHandlerFunc) Handle(params GetMusiciansIDParams) middleware.Responder {
	return fn(params)
}

// GetMusiciansIDHandler interface for that can handle valid get musicians ID params
type GetMusiciansIDHandler interface {
	Handle(GetMusiciansIDParams) middleware.Responder
}

// NewGetMusiciansID creates a new http.Handler for the get musicians ID operation
func NewGetMusiciansID(ctx *middleware.Context, handler GetMusiciansIDHandler) *GetMusiciansID {
	return &GetMusiciansID{Context: ctx, Handler: handler}
}

/*GetMusiciansID swagger:route GET /musicians/{id} getMusiciansId

Get musician information

Get musician information

*/
type GetMusiciansID struct {
	Context *middleware.Context
	Handler GetMusiciansIDHandler
}

func (o *GetMusiciansID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMusiciansIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
