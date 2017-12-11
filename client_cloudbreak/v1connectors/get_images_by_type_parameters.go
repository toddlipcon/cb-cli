// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetImagesByTypeParams creates a new GetImagesByTypeParams object
// with the default values initialized.
func NewGetImagesByTypeParams() *GetImagesByTypeParams {
	var ()
	return &GetImagesByTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesByTypeParamsWithTimeout creates a new GetImagesByTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesByTypeParamsWithTimeout(timeout time.Duration) *GetImagesByTypeParams {
	var ()
	return &GetImagesByTypeParams{

		timeout: timeout,
	}
}

// NewGetImagesByTypeParamsWithContext creates a new GetImagesByTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesByTypeParamsWithContext(ctx context.Context) *GetImagesByTypeParams {
	var ()
	return &GetImagesByTypeParams{

		Context: ctx,
	}
}

// NewGetImagesByTypeParamsWithHTTPClient creates a new GetImagesByTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImagesByTypeParamsWithHTTPClient(client *http.Client) *GetImagesByTypeParams {
	var ()
	return &GetImagesByTypeParams{
		HTTPClient: client,
	}
}

/*GetImagesByTypeParams contains all the parameters to send to the API endpoint
for the get images by type operation typically these are written to a http.Request
*/
type GetImagesByTypeParams struct {

	/*Type*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get images by type params
func (o *GetImagesByTypeParams) WithTimeout(timeout time.Duration) *GetImagesByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images by type params
func (o *GetImagesByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images by type params
func (o *GetImagesByTypeParams) WithContext(ctx context.Context) *GetImagesByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images by type params
func (o *GetImagesByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images by type params
func (o *GetImagesByTypeParams) WithHTTPClient(client *http.Client) *GetImagesByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images by type params
func (o *GetImagesByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get images by type params
func (o *GetImagesByTypeParams) WithType(typeVar string) *GetImagesByTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get images by type params
func (o *GetImagesByTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}