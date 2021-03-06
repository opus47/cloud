// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostPiecesIDOKCode is the HTTP code returned for type PostPiecesIDOK
const PostPiecesIDOKCode int = 200

/*PostPiecesIDOK Piece updated

swagger:response postPiecesIdOK
*/
type PostPiecesIDOK struct {
}

// NewPostPiecesIDOK creates PostPiecesIDOK with default headers values
func NewPostPiecesIDOK() *PostPiecesIDOK {

	return &PostPiecesIDOK{}
}

// WriteResponse to the client
func (o *PostPiecesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostPiecesIDForbiddenCode is the HTTP code returned for type PostPiecesIDForbidden
const PostPiecesIDForbiddenCode int = 403

/*PostPiecesIDForbidden Unauthorized

swagger:response postPiecesIdForbidden
*/
type PostPiecesIDForbidden struct {
}

// NewPostPiecesIDForbidden creates PostPiecesIDForbidden with default headers values
func NewPostPiecesIDForbidden() *PostPiecesIDForbidden {

	return &PostPiecesIDForbidden{}
}

// WriteResponse to the client
func (o *PostPiecesIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
