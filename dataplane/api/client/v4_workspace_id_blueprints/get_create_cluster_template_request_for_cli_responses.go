// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetCreateClusterTemplateRequestForCliReader is a Reader for the GetCreateClusterTemplateRequestForCli structure.
type GetCreateClusterTemplateRequestForCliReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCreateClusterTemplateRequestForCliReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCreateClusterTemplateRequestForCliOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCreateClusterTemplateRequestForCliOK creates a GetCreateClusterTemplateRequestForCliOK with default headers values
func NewGetCreateClusterTemplateRequestForCliOK() *GetCreateClusterTemplateRequestForCliOK {
	return &GetCreateClusterTemplateRequestForCliOK{}
}

/*GetCreateClusterTemplateRequestForCliOK handles this case with default header values.

successful operation
*/
type GetCreateClusterTemplateRequestForCliOK struct {
	Payload *model.CreateClusterTemplateRequest
}

func (o *GetCreateClusterTemplateRequestForCliOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/blueprints/cli_create][%d] getCreateClusterTemplateRequestForCliOK  %+v", 200, o.Payload)
}

func (o *GetCreateClusterTemplateRequestForCliOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CreateClusterTemplateRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
