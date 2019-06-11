// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// CreateSdxReader is a Reader for the CreateSdx structure.
type CreateSdxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSdxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSdxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSdxOK creates a CreateSdxOK with default headers values
func NewCreateSdxOK() *CreateSdxOK {
	return &CreateSdxOK{}
}

/*CreateSdxOK handles this case with default header values.

successful operation
*/
type CreateSdxOK struct {
	Payload *model.SdxClusterResponse
}

func (o *CreateSdxOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}][%d] createSdxOK  %+v", 200, o.Payload)
}

func (o *CreateSdxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
