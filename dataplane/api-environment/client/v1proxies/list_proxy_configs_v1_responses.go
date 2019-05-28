// Code generated by go-swagger; DO NOT EDIT.

package v1proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// ListProxyConfigsV1Reader is a Reader for the ListProxyConfigsV1 structure.
type ListProxyConfigsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProxyConfigsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListProxyConfigsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListProxyConfigsV1OK creates a ListProxyConfigsV1OK with default headers values
func NewListProxyConfigsV1OK() *ListProxyConfigsV1OK {
	return &ListProxyConfigsV1OK{}
}

/*ListProxyConfigsV1OK handles this case with default header values.

successful operation
*/
type ListProxyConfigsV1OK struct {
	Payload *model.ProxyResponses
}

func (o *ListProxyConfigsV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/proxies][%d] listProxyConfigsV1OK  %+v", 200, o.Payload)
}

func (o *ListProxyConfigsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyResponses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
