// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/opus47/cloud/api/models"
)

// GetMusiciansOKCode is the HTTP code returned for type GetMusiciansOK
const GetMusiciansOKCode int = 200

/*GetMusiciansOK List of musicians

swagger:response getMusiciansOK
*/
type GetMusiciansOK struct {

	/*
	  In: Body
	*/
	Payload models.GetMusiciansOKBody `json:"body,omitempty"`
}

// NewGetMusiciansOK creates GetMusiciansOK with default headers values
func NewGetMusiciansOK() *GetMusiciansOK {
	return &GetMusiciansOK{}
}

// WithPayload adds the payload to the get musicians o k response
func (o *GetMusiciansOK) WithPayload(payload models.GetMusiciansOKBody) *GetMusiciansOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get musicians o k response
func (o *GetMusiciansOK) SetPayload(payload models.GetMusiciansOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMusiciansOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetMusiciansOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
