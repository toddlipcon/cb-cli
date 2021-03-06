// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_kubernetes

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCreateKubernetesConfigInWorkspaceParams creates a new CreateKubernetesConfigInWorkspaceParams object
// with the default values initialized.
func NewCreateKubernetesConfigInWorkspaceParams() *CreateKubernetesConfigInWorkspaceParams {
	var ()
	return &CreateKubernetesConfigInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKubernetesConfigInWorkspaceParamsWithTimeout creates a new CreateKubernetesConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateKubernetesConfigInWorkspaceParamsWithTimeout(timeout time.Duration) *CreateKubernetesConfigInWorkspaceParams {
	var ()
	return &CreateKubernetesConfigInWorkspaceParams{

		timeout: timeout,
	}
}

// NewCreateKubernetesConfigInWorkspaceParamsWithContext creates a new CreateKubernetesConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateKubernetesConfigInWorkspaceParamsWithContext(ctx context.Context) *CreateKubernetesConfigInWorkspaceParams {
	var ()
	return &CreateKubernetesConfigInWorkspaceParams{

		Context: ctx,
	}
}

// NewCreateKubernetesConfigInWorkspaceParamsWithHTTPClient creates a new CreateKubernetesConfigInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateKubernetesConfigInWorkspaceParamsWithHTTPClient(client *http.Client) *CreateKubernetesConfigInWorkspaceParams {
	var ()
	return &CreateKubernetesConfigInWorkspaceParams{
		HTTPClient: client,
	}
}

/*CreateKubernetesConfigInWorkspaceParams contains all the parameters to send to the API endpoint
for the create kubernetes config in workspace operation typically these are written to a http.Request
*/
type CreateKubernetesConfigInWorkspaceParams struct {

	/*Body*/
	Body *model.KubernetesV4Request
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) WithTimeout(timeout time.Duration) *CreateKubernetesConfigInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) WithContext(ctx context.Context) *CreateKubernetesConfigInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) WithHTTPClient(client *http.Client) *CreateKubernetesConfigInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) WithBody(body *model.KubernetesV4Request) *CreateKubernetesConfigInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) SetBody(body *model.KubernetesV4Request) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) WithWorkspaceID(workspaceID int64) *CreateKubernetesConfigInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the create kubernetes config in workspace params
func (o *CreateKubernetesConfigInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKubernetesConfigInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
