// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutKeysIDOKCode is the HTTP code returned for type PutKeysIDOK
const PutKeysIDOKCode int = 200

/*PutKeysIDOK Key updated

swagger:response putKeysIdOK
*/
type PutKeysIDOK struct {
}

// NewPutKeysIDOK creates PutKeysIDOK with default headers values
func NewPutKeysIDOK() *PutKeysIDOK {

	return &PutKeysIDOK{}
}

// WriteResponse to the client
func (o *PutKeysIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutKeysIDForbiddenCode is the HTTP code returned for type PutKeysIDForbidden
const PutKeysIDForbiddenCode int = 403

/*PutKeysIDForbidden Unauthorized

swagger:response putKeysIdForbidden
*/
type PutKeysIDForbidden struct {
}

// NewPutKeysIDForbidden creates PutKeysIDForbidden with default headers values
func NewPutKeysIDForbidden() *PutKeysIDForbidden {

	return &PutKeysIDForbidden{}
}

// WriteResponse to the client
func (o *PutKeysIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
