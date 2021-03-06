// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteMusiciansIDHandlerFunc turns a function with the right signature into a delete musicians ID handler
type DeleteMusiciansIDHandlerFunc func(DeleteMusiciansIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMusiciansIDHandlerFunc) Handle(params DeleteMusiciansIDParams) middleware.Responder {
	return fn(params)
}

// DeleteMusiciansIDHandler interface for that can handle valid delete musicians ID params
type DeleteMusiciansIDHandler interface {
	Handle(DeleteMusiciansIDParams) middleware.Responder
}

// NewDeleteMusiciansID creates a new http.Handler for the delete musicians ID operation
func NewDeleteMusiciansID(ctx *middleware.Context, handler DeleteMusiciansIDHandler) *DeleteMusiciansID {
	return &DeleteMusiciansID{Context: ctx, Handler: handler}
}

/*DeleteMusiciansID swagger:route DELETE /musicians/{id} deleteMusiciansId

Delete a musician

Delete a musician

*/
type DeleteMusiciansID struct {
	Context *middleware.Context
	Handler DeleteMusiciansIDHandler
}

func (o *DeleteMusiciansID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMusiciansIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
