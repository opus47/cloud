// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetComposersIDHandlerFunc turns a function with the right signature into a get composers ID handler
type GetComposersIDHandlerFunc func(GetComposersIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetComposersIDHandlerFunc) Handle(params GetComposersIDParams) middleware.Responder {
	return fn(params)
}

// GetComposersIDHandler interface for that can handle valid get composers ID params
type GetComposersIDHandler interface {
	Handle(GetComposersIDParams) middleware.Responder
}

// NewGetComposersID creates a new http.Handler for the get composers ID operation
func NewGetComposersID(ctx *middleware.Context, handler GetComposersIDHandler) *GetComposersID {
	return &GetComposersID{Context: ctx, Handler: handler}
}

/*GetComposersID swagger:route GET /composers/{id} getComposersId

Get composer information

Get composer information

*/
type GetComposersID struct {
	Context *middleware.Context
	Handler GetComposersIDHandler
}

func (o *GetComposersID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetComposersIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
