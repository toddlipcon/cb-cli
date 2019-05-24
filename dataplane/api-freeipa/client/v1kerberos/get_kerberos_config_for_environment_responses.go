// Code generated by go-swagger; DO NOT EDIT.

package v1kerberos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GetKerberosConfigForEnvironmentReader is a Reader for the GetKerberosConfigForEnvironment structure.
type GetKerberosConfigForEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKerberosConfigForEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKerberosConfigForEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKerberosConfigForEnvironmentOK creates a GetKerberosConfigForEnvironmentOK with default headers values
func NewGetKerberosConfigForEnvironmentOK() *GetKerberosConfigForEnvironmentOK {
	return &GetKerberosConfigForEnvironmentOK{}
}

/*GetKerberosConfigForEnvironmentOK handles this case with default header values.

successful operation
*/
type GetKerberosConfigForEnvironmentOK struct {
	Payload *model.DescribeKerberosConfigV1Response
}

func (o *GetKerberosConfigForEnvironmentOK) Error() string {
	return fmt.Sprintf("[GET /v1/kerberos][%d] getKerberosConfigForEnvironmentOK  %+v", 200, o.Payload)
}

func (o *GetKerberosConfigForEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DescribeKerberosConfigV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
