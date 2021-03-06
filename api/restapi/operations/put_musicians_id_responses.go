// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutMusiciansIDOKCode is the HTTP code returned for type PutMusiciansIDOK
const PutMusiciansIDOKCode int = 200

/*PutMusiciansIDOK Musician updated

swagger:response putMusiciansIdOK
*/
type PutMusiciansIDOK struct {
}

// NewPutMusiciansIDOK creates PutMusiciansIDOK with default headers values
func NewPutMusiciansIDOK() *PutMusiciansIDOK {

	return &PutMusiciansIDOK{}
}

// WriteResponse to the client
func (o *PutMusiciansIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutMusiciansIDForbiddenCode is the HTTP code returned for type PutMusiciansIDForbidden
const PutMusiciansIDForbiddenCode int = 403

/*PutMusiciansIDForbidden Unauthorized

swagger:response putMusiciansIdForbidden
*/
type PutMusiciansIDForbidden struct {
}

// NewPutMusiciansIDForbidden creates PutMusiciansIDForbidden with default headers values
func NewPutMusiciansIDForbidden() *PutMusiciansIDForbidden {

	return &PutMusiciansIDForbidden{}
}

// WriteResponse to the client
func (o *PutMusiciansIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
