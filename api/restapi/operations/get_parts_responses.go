// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/opus47/cloud/api/models"
)

// GetPartsOKCode is the HTTP code returned for type GetPartsOK
const GetPartsOKCode int = 200

/*GetPartsOK List of parts

swagger:response getPartsOK
*/
type GetPartsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Part `json:"body,omitempty"`
}

// NewGetPartsOK creates GetPartsOK with default headers values
func NewGetPartsOK() *GetPartsOK {

	return &GetPartsOK{}
}

// WithPayload adds the payload to the get parts o k response
func (o *GetPartsOK) WithPayload(payload []*models.Part) *GetPartsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get parts o k response
func (o *GetPartsOK) SetPayload(payload []*models.Part) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPartsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Part, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
