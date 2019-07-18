// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteEnvironmentV1ByNameParams creates a new DeleteEnvironmentV1ByNameParams object
// with the default values initialized.
func NewDeleteEnvironmentV1ByNameParams() *DeleteEnvironmentV1ByNameParams {
	var ()
	return &DeleteEnvironmentV1ByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentV1ByNameParamsWithTimeout creates a new DeleteEnvironmentV1ByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvironmentV1ByNameParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentV1ByNameParams {
	var ()
	return &DeleteEnvironmentV1ByNameParams{

		timeout: timeout,
	}
}

// NewDeleteEnvironmentV1ByNameParamsWithContext creates a new DeleteEnvironmentV1ByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvironmentV1ByNameParamsWithContext(ctx context.Context) *DeleteEnvironmentV1ByNameParams {
	var ()
	return &DeleteEnvironmentV1ByNameParams{

		Context: ctx,
	}
}

// NewDeleteEnvironmentV1ByNameParamsWithHTTPClient creates a new DeleteEnvironmentV1ByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvironmentV1ByNameParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentV1ByNameParams {
	var ()
	return &DeleteEnvironmentV1ByNameParams{
		HTTPClient: client,
	}
}

/*DeleteEnvironmentV1ByNameParams contains all the parameters to send to the API endpoint
for the delete environment v1 by name operation typically these are written to a http.Request
*/
type DeleteEnvironmentV1ByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentV1ByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) WithContext(ctx context.Context) *DeleteEnvironmentV1ByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentV1ByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) WithName(name string) *DeleteEnvironmentV1ByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete environment v1 by name params
func (o *DeleteEnvironmentV1ByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentV1ByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
