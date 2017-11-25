// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPiecesIDHandlerFunc turns a function with the right signature into a get pieces ID handler
type GetPiecesIDHandlerFunc func(GetPiecesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPiecesIDHandlerFunc) Handle(params GetPiecesIDParams) middleware.Responder {
	return fn(params)
}

// GetPiecesIDHandler interface for that can handle valid get pieces ID params
type GetPiecesIDHandler interface {
	Handle(GetPiecesIDParams) middleware.Responder
}

// NewGetPiecesID creates a new http.Handler for the get pieces ID operation
func NewGetPiecesID(ctx *middleware.Context, handler GetPiecesIDHandler) *GetPiecesID {
	return &GetPiecesID{Context: ctx, Handler: handler}
}

/*GetPiecesID swagger:route GET /pieces/{id} getPiecesId

Get piece information

Get pieces information

*/
type GetPiecesID struct {
	Context *middleware.Context
	Handler GetPiecesIDHandler
}

func (o *GetPiecesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPiecesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}