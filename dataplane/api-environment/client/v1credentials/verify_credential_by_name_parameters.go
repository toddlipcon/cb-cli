// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

// NewVerifyCredentialByNameParams creates a new VerifyCredentialByNameParams object
// with the default values initialized.
func NewVerifyCredentialByNameParams() *VerifyCredentialByNameParams {
	var ()
	return &VerifyCredentialByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyCredentialByNameParamsWithTimeout creates a new VerifyCredentialByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerifyCredentialByNameParamsWithTimeout(timeout time.Duration) *VerifyCredentialByNameParams {
	var ()
	return &VerifyCredentialByNameParams{

		timeout: timeout,
	}
}

// NewVerifyCredentialByNameParamsWithContext creates a new VerifyCredentialByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerifyCredentialByNameParamsWithContext(ctx context.Context) *VerifyCredentialByNameParams {
	var ()
	return &VerifyCredentialByNameParams{

		Context: ctx,
	}
}

// NewVerifyCredentialByNameParamsWithHTTPClient creates a new VerifyCredentialByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerifyCredentialByNameParamsWithHTTPClient(client *http.Client) *VerifyCredentialByNameParams {
	var ()
	return &VerifyCredentialByNameParams{
		HTTPClient: client,
	}
}

/*VerifyCredentialByNameParams contains all the parameters to send to the API endpoint
for the verify credential by name operation typically these are written to a http.Request
*/
type VerifyCredentialByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verify credential by name params
func (o *VerifyCredentialByNameParams) WithTimeout(timeout time.Duration) *VerifyCredentialByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify credential by name params
func (o *VerifyCredentialByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify credential by name params
func (o *VerifyCredentialByNameParams) WithContext(ctx context.Context) *VerifyCredentialByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify credential by name params
func (o *VerifyCredentialByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify credential by name params
func (o *VerifyCredentialByNameParams) WithHTTPClient(client *http.Client) *VerifyCredentialByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify credential by name params
func (o *VerifyCredentialByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the verify credential by name params
func (o *VerifyCredentialByNameParams) WithName(name string) *VerifyCredentialByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the verify credential by name params
func (o *VerifyCredentialByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyCredentialByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
