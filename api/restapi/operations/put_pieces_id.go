// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutPiecesIDHandlerFunc turns a function with the right signature into a put pieces ID handler
type PutPiecesIDHandlerFunc func(PutPiecesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPiecesIDHandlerFunc) Handle(params PutPiecesIDParams) middleware.Responder {
	return fn(params)
}

// PutPiecesIDHandler interface for that can handle valid put pieces ID params
type PutPiecesIDHandler interface {
	Handle(PutPiecesIDParams) middleware.Responder
}

// NewPutPiecesID creates a new http.Handler for the put pieces ID operation
func NewPutPiecesID(ctx *middleware.Context, handler PutPiecesIDHandler) *PutPiecesID {
	return &PutPiecesID{Context: ctx, Handler: handler}
}

/*PutPiecesID swagger:route PUT /pieces/{id} putPiecesId

Add or update a piece

Add or update a piece

*/
type PutPiecesID struct {
	Context *middleware.Context
	Handler PutPiecesIDHandler
}

func (o *PutPiecesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutPiecesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
