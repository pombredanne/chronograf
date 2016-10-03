package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*PostSourcesIDKapacitorsCreated Successfully created kapacitor source

swagger:response postSourcesIdKapacitorsCreated
*/
type PostSourcesIDKapacitorsCreated struct {
	/*Location of the newly created kapacitor resource.
	  Required: true
	*/
	Location string `json:"Location"`

	// In: body
	Payload *models.Kapacitor `json:"body,omitempty"`
}

// NewPostSourcesIDKapacitorsCreated creates PostSourcesIDKapacitorsCreated with default headers values
func NewPostSourcesIDKapacitorsCreated() *PostSourcesIDKapacitorsCreated {
	return &PostSourcesIDKapacitorsCreated{}
}

// WithLocation adds the location to the post sources Id kapacitors created response
func (o *PostSourcesIDKapacitorsCreated) WithLocation(location string) *PostSourcesIDKapacitorsCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post sources Id kapacitors created response
func (o *PostSourcesIDKapacitorsCreated) SetLocation(location string) {
	o.Location = location
}

// WithPayload adds the payload to the post sources Id kapacitors created response
func (o *PostSourcesIDKapacitorsCreated) WithPayload(payload *models.Kapacitor) *PostSourcesIDKapacitorsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post sources Id kapacitors created response
func (o *PostSourcesIDKapacitorsCreated) SetPayload(payload *models.Kapacitor) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSourcesIDKapacitorsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location
	rw.Header().Add("Location", fmt.Sprintf("%v", o.Location))

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostSourcesIDKapacitorsDefault A processing or an unexpected error.

swagger:response postSourcesIdKapacitorsDefault
*/
type PostSourcesIDKapacitorsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostSourcesIDKapacitorsDefault creates PostSourcesIDKapacitorsDefault with default headers values
func NewPostSourcesIDKapacitorsDefault(code int) *PostSourcesIDKapacitorsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostSourcesIDKapacitorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post sources ID kapacitors default response
func (o *PostSourcesIDKapacitorsDefault) WithStatusCode(code int) *PostSourcesIDKapacitorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post sources ID kapacitors default response
func (o *PostSourcesIDKapacitorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post sources ID kapacitors default response
func (o *PostSourcesIDKapacitorsDefault) WithPayload(payload *models.Error) *PostSourcesIDKapacitorsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post sources ID kapacitors default response
func (o *PostSourcesIDKapacitorsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSourcesIDKapacitorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
