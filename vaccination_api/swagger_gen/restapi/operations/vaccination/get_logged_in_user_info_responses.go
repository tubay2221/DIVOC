// Code generated by go-swagger; DO NOT EDIT.

package vaccination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/api/swagger_gen/models"
)

// GetLoggedInUserInfoOKCode is the HTTP code returned for type GetLoggedInUserInfoOK
const GetLoggedInUserInfoOKCode int = 200

/*GetLoggedInUserInfoOK OK

swagger:response getLoggedInUserInfoOK
*/
type GetLoggedInUserInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserInfo `json:"body,omitempty"`
}

// NewGetLoggedInUserInfoOK creates GetLoggedInUserInfoOK with default headers values
func NewGetLoggedInUserInfoOK() *GetLoggedInUserInfoOK {

	return &GetLoggedInUserInfoOK{}
}

// WithPayload adds the payload to the get logged in user info o k response
func (o *GetLoggedInUserInfoOK) WithPayload(payload *models.UserInfo) *GetLoggedInUserInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get logged in user info o k response
func (o *GetLoggedInUserInfoOK) SetPayload(payload *models.UserInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLoggedInUserInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
