// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeletePartsIDOKCode is the HTTP code returned for type DeletePartsIDOK
const DeletePartsIDOKCode int = 200

/*DeletePartsIDOK Part deleted

swagger:response deletePartsIdOK
*/
type DeletePartsIDOK struct {
}

// NewDeletePartsIDOK creates DeletePartsIDOK with default headers values
func NewDeletePartsIDOK() *DeletePartsIDOK {
	return &DeletePartsIDOK{}
}

// WriteResponse to the client
func (o *DeletePartsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeletePartsIDForbiddenCode is the HTTP code returned for type DeletePartsIDForbidden
const DeletePartsIDForbiddenCode int = 403

/*DeletePartsIDForbidden Unauthorized

swagger:response deletePartsIdForbidden
*/
type DeletePartsIDForbidden struct {
}

// NewDeletePartsIDForbidden creates DeletePartsIDForbidden with default headers values
func NewDeletePartsIDForbidden() *DeletePartsIDForbidden {
	return &DeletePartsIDForbidden{}
}

// WriteResponse to the client
func (o *DeletePartsIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
