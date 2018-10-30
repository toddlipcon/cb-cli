// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_ldapconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListLdapsByWorkspaceParams creates a new ListLdapsByWorkspaceParams object
// with the default values initialized.
func NewListLdapsByWorkspaceParams() *ListLdapsByWorkspaceParams {
	var ()
	return &ListLdapsByWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListLdapsByWorkspaceParamsWithTimeout creates a new ListLdapsByWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListLdapsByWorkspaceParamsWithTimeout(timeout time.Duration) *ListLdapsByWorkspaceParams {
	var ()
	return &ListLdapsByWorkspaceParams{

		timeout: timeout,
	}
}

// NewListLdapsByWorkspaceParamsWithContext creates a new ListLdapsByWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListLdapsByWorkspaceParamsWithContext(ctx context.Context) *ListLdapsByWorkspaceParams {
	var ()
	return &ListLdapsByWorkspaceParams{

		Context: ctx,
	}
}

// NewListLdapsByWorkspaceParamsWithHTTPClient creates a new ListLdapsByWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListLdapsByWorkspaceParamsWithHTTPClient(client *http.Client) *ListLdapsByWorkspaceParams {
	var ()
	return &ListLdapsByWorkspaceParams{
		HTTPClient: client,
	}
}

/*ListLdapsByWorkspaceParams contains all the parameters to send to the API endpoint
for the list ldaps by workspace operation typically these are written to a http.Request
*/
type ListLdapsByWorkspaceParams struct {

	/*AttachGlobal*/
	AttachGlobal *bool
	/*Environment*/
	Environment *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) WithTimeout(timeout time.Duration) *ListLdapsByWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) WithContext(ctx context.Context) *ListLdapsByWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) WithHTTPClient(client *http.Client) *ListLdapsByWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachGlobal adds the attachGlobal to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) WithAttachGlobal(attachGlobal *bool) *ListLdapsByWorkspaceParams {
	o.SetAttachGlobal(attachGlobal)
	return o
}

// SetAttachGlobal adds the attachGlobal to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) SetAttachGlobal(attachGlobal *bool) {
	o.AttachGlobal = attachGlobal
}

// WithEnvironment adds the environment to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) WithEnvironment(environment *string) *ListLdapsByWorkspaceParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithWorkspaceID adds the workspaceID to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) WithWorkspaceID(workspaceID int64) *ListLdapsByWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the list ldaps by workspace params
func (o *ListLdapsByWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListLdapsByWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AttachGlobal != nil {

		// query param attachGlobal
		var qrAttachGlobal bool
		if o.AttachGlobal != nil {
			qrAttachGlobal = *o.AttachGlobal
		}
		qAttachGlobal := swag.FormatBool(qrAttachGlobal)
		if qAttachGlobal != "" {
			if err := r.SetQueryParam("attachGlobal", qAttachGlobal); err != nil {
				return err
			}
		}

	}

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

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
