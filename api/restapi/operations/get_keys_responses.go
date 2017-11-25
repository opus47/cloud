// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opus47/cloud/api/models"
)

// GetKeysOKCode is the HTTP code returned for type GetKeysOK
const GetKeysOKCode int = 200

/*GetKeysOK List of keys

swagger:response getKeysOK
*/
type GetKeysOK struct {

	/*
	  In: Body
	*/
	Payload models.GetKeysOKBody `json:"body,omitempty"`
}

// NewGetKeysOK creates GetKeysOK with default headers values
func NewGetKeysOK() *GetKeysOK {
	return &GetKeysOK{}
}

// WithPayload adds the payload to the get keys o k response
func (o *GetKeysOK) WithPayload(payload models.GetKeysOKBody) *GetKeysOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get keys o k response
func (o *GetKeysOK) SetPayload(payload models.GetKeysOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetKeysOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetKeysOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}