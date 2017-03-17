package usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GenerateReader is a Reader for the Generate structure.
type GenerateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GenerateReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewGenerateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGenerateDefault creates a GenerateDefault with default headers values
func NewGenerateDefault(code int) *GenerateDefault {
	return &GenerateDefault{
		_statusCode: code,
	}
}

/*GenerateDefault handles this case with default header values.

successful operation
*/
type GenerateDefault struct {
	_statusCode int
}

// Code gets the status code for the generate default response
func (o *GenerateDefault) Code() int {
	return o._statusCode
}

func (o *GenerateDefault) Error() string {
	return fmt.Sprintf("[GET /usages/generate][%d] generate default ", o._statusCode)
}

func (o *GenerateDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}