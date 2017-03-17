package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetTemplatesAccountReader is a Reader for the GetTemplatesAccount structure.
type GetTemplatesAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetTemplatesAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTemplatesAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTemplatesAccountOK creates a GetTemplatesAccountOK with default headers values
func NewGetTemplatesAccountOK() *GetTemplatesAccountOK {
	return &GetTemplatesAccountOK{}
}

/*GetTemplatesAccountOK handles this case with default header values.

successful operation
*/
type GetTemplatesAccountOK struct {
	Payload []*models_cloudbreak.TemplateResponse
}

func (o *GetTemplatesAccountOK) Error() string {
	return fmt.Sprintf("[GET /templates/account][%d] getTemplatesAccountOK  %+v", 200, o.Payload)
}

func (o *GetTemplatesAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}