// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMusiciansHandlerFunc turns a function with the right signature into a get musicians handler
type GetMusiciansHandlerFunc func(GetMusiciansParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMusiciansHandlerFunc) Handle(params GetMusiciansParams) middleware.Responder {
	return fn(params)
}

// GetMusiciansHandler interface for that can handle valid get musicians params
type GetMusiciansHandler interface {
	Handle(GetMusiciansParams) middleware.Responder
}

// NewGetMusicians creates a new http.Handler for the get musicians operation
func NewGetMusicians(ctx *middleware.Context, handler GetMusiciansHandler) *GetMusicians {
	return &GetMusicians{Context: ctx, Handler: handler}
}

/*GetMusicians swagger:route GET /musicians getMusicians

List musicians

List musicians

*/
type GetMusicians struct {
	Context *middleware.Context
	Handler GetMusiciansHandler
}

func (o *GetMusicians) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMusiciansParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
