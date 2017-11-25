// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteKeysIDOKCode is the HTTP code returned for type DeleteKeysIDOK
const DeleteKeysIDOKCode int = 200

/*DeleteKeysIDOK Key deleted

swagger:response deleteKeysIdOK
*/
type DeleteKeysIDOK struct {
}

// NewDeleteKeysIDOK creates DeleteKeysIDOK with default headers values
func NewDeleteKeysIDOK() *DeleteKeysIDOK {
	return &DeleteKeysIDOK{}
}

// WriteResponse to the client
func (o *DeleteKeysIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteKeysIDForbiddenCode is the HTTP code returned for type DeleteKeysIDForbidden
const DeleteKeysIDForbiddenCode int = 403

/*DeleteKeysIDForbidden Unauthorized

swagger:response deleteKeysIdForbidden
*/
type DeleteKeysIDForbidden struct {
}

// NewDeleteKeysIDForbidden creates DeleteKeysIDForbidden with default headers values
func NewDeleteKeysIDForbidden() *DeleteKeysIDForbidden {
	return &DeleteKeysIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteKeysIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}