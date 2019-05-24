// Code generated by go-swagger; DO NOT EDIT.

package v1kerberos

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

// NewDeleteKerberosConfigForEnvironmentParams creates a new DeleteKerberosConfigForEnvironmentParams object
// with the default values initialized.
func NewDeleteKerberosConfigForEnvironmentParams() *DeleteKerberosConfigForEnvironmentParams {
	var ()
	return &DeleteKerberosConfigForEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKerberosConfigForEnvironmentParamsWithTimeout creates a new DeleteKerberosConfigForEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteKerberosConfigForEnvironmentParamsWithTimeout(timeout time.Duration) *DeleteKerberosConfigForEnvironmentParams {
	var ()
	return &DeleteKerberosConfigForEnvironmentParams{

		timeout: timeout,
	}
}

// NewDeleteKerberosConfigForEnvironmentParamsWithContext creates a new DeleteKerberosConfigForEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteKerberosConfigForEnvironmentParamsWithContext(ctx context.Context) *DeleteKerberosConfigForEnvironmentParams {
	var ()
	return &DeleteKerberosConfigForEnvironmentParams{

		Context: ctx,
	}
}

// NewDeleteKerberosConfigForEnvironmentParamsWithHTTPClient creates a new DeleteKerberosConfigForEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteKerberosConfigForEnvironmentParamsWithHTTPClient(client *http.Client) *DeleteKerberosConfigForEnvironmentParams {
	var ()
	return &DeleteKerberosConfigForEnvironmentParams{
		HTTPClient: client,
	}
}

/*DeleteKerberosConfigForEnvironmentParams contains all the parameters to send to the API endpoint
for the delete kerberos config for environment operation typically these are written to a http.Request
*/
type DeleteKerberosConfigForEnvironmentParams struct {

	/*Environment*/
	Environment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) WithTimeout(timeout time.Duration) *DeleteKerberosConfigForEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) WithContext(ctx context.Context) *DeleteKerberosConfigForEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) WithHTTPClient(client *http.Client) *DeleteKerberosConfigForEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) WithEnvironment(environment *string) *DeleteKerberosConfigForEnvironmentParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the delete kerberos config for environment params
func (o *DeleteKerberosConfigForEnvironmentParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKerberosConfigForEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
