// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opus47/cloud/api/models"
)

// GetComposersOKCode is the HTTP code returned for type GetComposersOK
const GetComposersOKCode int = 200

/*GetComposersOK List of composers

swagger:response getComposersOK
*/
type GetComposersOK struct {

	/*
	  In: Body
	*/
	Payload models.GetComposersOKBody `json:"body,omitempty"`
}

// NewGetComposersOK creates GetComposersOK with default headers values
func NewGetComposersOK() *GetComposersOK {
	return &GetComposersOK{}
}

// WithPayload adds the payload to the get composers o k response
func (o *GetComposersOK) WithPayload(payload models.GetComposersOKBody) *GetComposersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get composers o k response
func (o *GetComposersOK) SetPayload(payload models.GetComposersOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetComposersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetComposersOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}