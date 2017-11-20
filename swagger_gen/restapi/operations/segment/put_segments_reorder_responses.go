// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/checkr/flagr/swagger_gen/models"
)

// PutSegmentsReorderOKCode is the HTTP code returned for type PutSegmentsReorderOK
const PutSegmentsReorderOKCode int = 200

/*PutSegmentsReorderOK segments reordered

swagger:response putSegmentsReorderOK
*/
type PutSegmentsReorderOK struct {
}

// NewPutSegmentsReorderOK creates PutSegmentsReorderOK with default headers values
func NewPutSegmentsReorderOK() *PutSegmentsReorderOK {
	return &PutSegmentsReorderOK{}
}

// WriteResponse to the client
func (o *PutSegmentsReorderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*PutSegmentsReorderDefault generic error response

swagger:response putSegmentsReorderDefault
*/
type PutSegmentsReorderDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutSegmentsReorderDefault creates PutSegmentsReorderDefault with default headers values
func NewPutSegmentsReorderDefault(code int) *PutSegmentsReorderDefault {
	if code <= 0 {
		code = 500
	}

	return &PutSegmentsReorderDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put segments reorder default response
func (o *PutSegmentsReorderDefault) WithStatusCode(code int) *PutSegmentsReorderDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put segments reorder default response
func (o *PutSegmentsReorderDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put segments reorder default response
func (o *PutSegmentsReorderDefault) WithPayload(payload *models.Error) *PutSegmentsReorderDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put segments reorder default response
func (o *PutSegmentsReorderDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSegmentsReorderDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
