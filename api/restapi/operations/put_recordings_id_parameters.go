// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutRecordingsIDParams creates a new PutRecordingsIDParams object
// no default values defined in spec.
func NewPutRecordingsIDParams() PutRecordingsIDParams {

	return PutRecordingsIDParams{}
}

// PutRecordingsIDParams contains all the bound params for the put recordings ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutRecordingsID
type PutRecordingsIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*recording file
	  Required: true
	  In: formData
	*/
	File io.ReadCloser
	/*
	  Required: true
	  In: path
	*/
	ID string
	/*movement number
	  Required: true
	  In: formData
	*/
	Movement int64
	/*performance name
	  Required: true
	  In: formData
	*/
	Performance string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutRecordingsIDParams() beforehand.
func (o *PutRecordingsIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err := o.bindFile(file, fileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdMovement, fdhkMovement, _ := fds.GetOK("movement")
	if err := o.bindMovement(fdMovement, fdhkMovement, route.Formats); err != nil {
		res = append(res, err)
	}

	fdPerformance, fdhkPerformance, _ := fds.GetOK("performance")
	if err := o.bindPerformance(fdPerformance, fdhkPerformance, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PutRecordingsIDParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *PutRecordingsIDParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

// bindMovement binds and validates parameter Movement from formData.
func (o *PutRecordingsIDParams) bindMovement(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("movement", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("movement", "formData", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("movement", "formData", "int64", raw)
	}
	o.Movement = value

	return nil
}

// bindPerformance binds and validates parameter Performance from formData.
func (o *PutRecordingsIDParams) bindPerformance(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("performance", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("performance", "formData", raw); err != nil {
		return err
	}

	o.Performance = raw

	return nil
}
