// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMusiciansIDParams creates a new DeleteMusiciansIDParams object
// with the default values initialized.
func NewDeleteMusiciansIDParams() DeleteMusiciansIDParams {
	var ()
	return DeleteMusiciansIDParams{}
}

// DeleteMusiciansIDParams contains all the bound params for the delete musicians ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteMusiciansID
type DeleteMusiciansIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*API access authorization in JWT access token format
	  Required: true
	  In: header
	*/
	Authorization string
	/*
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteMusiciansIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("authorization")], true, route.Formats); err != nil {
		res = append(res, err)
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

func (o *DeleteMusiciansIDParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("authorization", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("authorization", "header", raw); err != nil {
		return err
	}

	o.Authorization = raw

	return nil
}

func (o *DeleteMusiciansIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}
