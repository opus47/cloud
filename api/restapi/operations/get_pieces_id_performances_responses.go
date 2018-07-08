// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/opus47/cloud/api/models"
)

// GetPiecesIDPerformancesOKCode is the HTTP code returned for type GetPiecesIDPerformancesOK
const GetPiecesIDPerformancesOKCode int = 200

/*GetPiecesIDPerformancesOK Piece information

swagger:response getPiecesIdPerformancesOK
*/
type GetPiecesIDPerformancesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Performance `json:"body,omitempty"`
}

// NewGetPiecesIDPerformancesOK creates GetPiecesIDPerformancesOK with default headers values
func NewGetPiecesIDPerformancesOK() *GetPiecesIDPerformancesOK {

	return &GetPiecesIDPerformancesOK{}
}

// WithPayload adds the payload to the get pieces Id performances o k response
func (o *GetPiecesIDPerformancesOK) WithPayload(payload []*models.Performance) *GetPiecesIDPerformancesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pieces Id performances o k response
func (o *GetPiecesIDPerformancesOK) SetPayload(payload []*models.Performance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPiecesIDPerformancesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Performance, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetPiecesIDPerformancesBadRequestCode is the HTTP code returned for type GetPiecesIDPerformancesBadRequest
const GetPiecesIDPerformancesBadRequestCode int = 400

/*GetPiecesIDPerformancesBadRequest Bad Request

swagger:response getPiecesIdPerformancesBadRequest
*/
type GetPiecesIDPerformancesBadRequest struct {
}

// NewGetPiecesIDPerformancesBadRequest creates GetPiecesIDPerformancesBadRequest with default headers values
func NewGetPiecesIDPerformancesBadRequest() *GetPiecesIDPerformancesBadRequest {

	return &GetPiecesIDPerformancesBadRequest{}
}

// WriteResponse to the client
func (o *GetPiecesIDPerformancesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetPiecesIDPerformancesInternalServerErrorCode is the HTTP code returned for type GetPiecesIDPerformancesInternalServerError
const GetPiecesIDPerformancesInternalServerErrorCode int = 500

/*GetPiecesIDPerformancesInternalServerError Internal Error

swagger:response getPiecesIdPerformancesInternalServerError
*/
type GetPiecesIDPerformancesInternalServerError struct {
}

// NewGetPiecesIDPerformancesInternalServerError creates GetPiecesIDPerformancesInternalServerError with default headers values
func NewGetPiecesIDPerformancesInternalServerError() *GetPiecesIDPerformancesInternalServerError {

	return &GetPiecesIDPerformancesInternalServerError{}
}

// WriteResponse to the client
func (o *GetPiecesIDPerformancesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
