// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPublicBlueprintReader is a Reader for the PostPublicBlueprint structure.
type PostPublicBlueprintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublicBlueprintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPublicBlueprintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPublicBlueprintOK creates a PostPublicBlueprintOK with default headers values
func NewPostPublicBlueprintOK() *PostPublicBlueprintOK {
	return &PostPublicBlueprintOK{}
}

/*PostPublicBlueprintOK handles this case with default header values.

successful operation
*/
type PostPublicBlueprintOK struct {
	Payload *models_cloudbreak.BlueprintResponse
}

func (o *PostPublicBlueprintOK) Error() string {
	return fmt.Sprintf("[POST /v1/blueprints/account][%d] postPublicBlueprintOK  %+v", 200, o.Payload)
}

func (o *PostPublicBlueprintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.BlueprintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}