// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opus47/cloud/api/models"
)

// GetPiecesIDOKCode is the HTTP code returned for type GetPiecesIDOK
const GetPiecesIDOKCode int = 200

/*GetPiecesIDOK Piece information

swagger:response getPiecesIdOK
*/
type GetPiecesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Piece `json:"body,omitempty"`
}

// NewGetPiecesIDOK creates GetPiecesIDOK with default headers values
func NewGetPiecesIDOK() *GetPiecesIDOK {
	return &GetPiecesIDOK{}
}

// WithPayload adds the payload to the get pieces Id o k response
func (o *GetPiecesIDOK) WithPayload(payload *models.Piece) *GetPiecesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pieces Id o k response
func (o *GetPiecesIDOK) SetPayload(payload *models.Piece) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPiecesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
