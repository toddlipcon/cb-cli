// Code generated by go-swagger; DO NOT EDIT.

package v3workspaces

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

// NewRemoveWorkspaceUsersParams creates a new RemoveWorkspaceUsersParams object
// with the default values initialized.
func NewRemoveWorkspaceUsersParams() *RemoveWorkspaceUsersParams {
	var ()
	return &RemoveWorkspaceUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveWorkspaceUsersParamsWithTimeout creates a new RemoveWorkspaceUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveWorkspaceUsersParamsWithTimeout(timeout time.Duration) *RemoveWorkspaceUsersParams {
	var ()
	return &RemoveWorkspaceUsersParams{

		timeout: timeout,
	}
}

// NewRemoveWorkspaceUsersParamsWithContext creates a new RemoveWorkspaceUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveWorkspaceUsersParamsWithContext(ctx context.Context) *RemoveWorkspaceUsersParams {
	var ()
	return &RemoveWorkspaceUsersParams{

		Context: ctx,
	}
}

// NewRemoveWorkspaceUsersParamsWithHTTPClient creates a new RemoveWorkspaceUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveWorkspaceUsersParamsWithHTTPClient(client *http.Client) *RemoveWorkspaceUsersParams {
	var ()
	return &RemoveWorkspaceUsersParams{
		HTTPClient: client,
	}
}

/*RemoveWorkspaceUsersParams contains all the parameters to send to the API endpoint
for the remove workspace users operation typically these are written to a http.Request
*/
type RemoveWorkspaceUsersParams struct {

	/*Body*/
	Body []string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) WithTimeout(timeout time.Duration) *RemoveWorkspaceUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) WithContext(ctx context.Context) *RemoveWorkspaceUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) WithHTTPClient(client *http.Client) *RemoveWorkspaceUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) WithBody(body []string) *RemoveWorkspaceUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) SetBody(body []string) {
	o.Body = body
}

// WithName adds the name to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) WithName(name string) *RemoveWorkspaceUsersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the remove workspace users params
func (o *RemoveWorkspaceUsersParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveWorkspaceUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}