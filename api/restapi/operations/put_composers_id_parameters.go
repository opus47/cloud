// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/opus47/cloud/api/models"
)

// NewPutComposersIDParams creates a new PutComposersIDParams object
// no default values defined in spec.
func NewPutComposersIDParams() PutComposersIDParams {

	return PutComposersIDParams{}
}

// PutComposersIDParams contains all the bound params for the put composers ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutComposersID
type PutComposersIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The composer to add or update
	  Required: true
	  In: body
	*/
	Data *models.Composer
	/*
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutComposersIDParams() beforehand.
func (o *PutComposersIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Composer
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PutComposersIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}
