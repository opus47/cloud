// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutPerformancesIDOKCode is the HTTP code returned for type PutPerformancesIDOK
const PutPerformancesIDOKCode int = 200

/*PutPerformancesIDOK Performance updated

swagger:response putPerformancesIdOK
*/
type PutPerformancesIDOK struct {
}

// NewPutPerformancesIDOK creates PutPerformancesIDOK with default headers values
func NewPutPerformancesIDOK() *PutPerformancesIDOK {
	return &PutPerformancesIDOK{}
}

// WriteResponse to the client
func (o *PutPerformancesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutPerformancesIDForbiddenCode is the HTTP code returned for type PutPerformancesIDForbidden
const PutPerformancesIDForbiddenCode int = 403

/*PutPerformancesIDForbidden Unauthorized

swagger:response putPerformancesIdForbidden
*/
type PutPerformancesIDForbidden struct {
}

// NewPutPerformancesIDForbidden creates PutPerformancesIDForbidden with default headers values
func NewPutPerformancesIDForbidden() *PutPerformancesIDForbidden {
	return &PutPerformancesIDForbidden{}
}

// WriteResponse to the client
func (o *PutPerformancesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
