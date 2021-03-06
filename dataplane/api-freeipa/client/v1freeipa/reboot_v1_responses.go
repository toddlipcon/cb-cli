// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RebootV1Reader is a Reader for the RebootV1 structure.
type RebootV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebootV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRebootV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRebootV1Default creates a RebootV1Default with default headers values
func NewRebootV1Default(code int) *RebootV1Default {
	return &RebootV1Default{
		_statusCode: code,
	}
}

/*RebootV1Default handles this case with default header values.

successful operation
*/
type RebootV1Default struct {
	_statusCode int
}

// Code gets the status code for the reboot v1 default response
func (o *RebootV1Default) Code() int {
	return o._statusCode
}

func (o *RebootV1Default) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/reboot][%d] rebootV1 default ", o._statusCode)
}

func (o *RebootV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
