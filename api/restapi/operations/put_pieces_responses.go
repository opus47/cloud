// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/opus47/cloud/api/models"
)

// PutPiecesOKCode is the HTTP code returned for type PutPiecesOK
const PutPiecesOKCode int = 200

/*PutPiecesOK Piece updated

swagger:response putPiecesOK
*/
type PutPiecesOK struct {
}

// NewPutPiecesOK creates PutPiecesOK with default headers values
func NewPutPiecesOK() *PutPiecesOK {

	return &PutPiecesOK{}
}

// WriteResponse to the client
func (o *PutPiecesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutPiecesBadRequestCode is the HTTP code returned for type PutPiecesBadRequest
const PutPiecesBadRequestCode int = 400

/*PutPiecesBadRequest Bad Request

swagger:response putPiecesBadRequest
*/
type PutPiecesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPutPiecesBadRequest creates PutPiecesBadRequest with default headers values
func NewPutPiecesBadRequest() *PutPiecesBadRequest {

	return &PutPiecesBadRequest{}
}

// WithPayload adds the payload to the put pieces bad request response
func (o *PutPiecesBadRequest) WithPayload(payload *models.ErrorMessage) *PutPiecesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put pieces bad request response
func (o *PutPiecesBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPiecesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutPiecesForbiddenCode is the HTTP code returned for type PutPiecesForbidden
const PutPiecesForbiddenCode int = 403

/*PutPiecesForbidden Unauthorized

swagger:response putPiecesForbidden
*/
type PutPiecesForbidden struct {
}

// NewPutPiecesForbidden creates PutPiecesForbidden with default headers values
func NewPutPiecesForbidden() *PutPiecesForbidden {

	return &PutPiecesForbidden{}
}

// WriteResponse to the client
func (o *PutPiecesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PutPiecesInternalServerErrorCode is the HTTP code returned for type PutPiecesInternalServerError
const PutPiecesInternalServerErrorCode int = 500

/*PutPiecesInternalServerError Internal Error

swagger:response putPiecesInternalServerError
*/
type PutPiecesInternalServerError struct {
}

// NewPutPiecesInternalServerError creates PutPiecesInternalServerError with default headers values
func NewPutPiecesInternalServerError() *PutPiecesInternalServerError {

	return &PutPiecesInternalServerError{}
}

// WriteResponse to the client
func (o *PutPiecesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
