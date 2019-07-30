// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewDeleteSdxByCrnParams creates a new DeleteSdxByCrnParams object
// with the default values initialized.
func NewDeleteSdxByCrnParams() *DeleteSdxByCrnParams {
	var ()
	return &DeleteSdxByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSdxByCrnParamsWithTimeout creates a new DeleteSdxByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSdxByCrnParamsWithTimeout(timeout time.Duration) *DeleteSdxByCrnParams {
	var ()
	return &DeleteSdxByCrnParams{

		timeout: timeout,
	}
}

// NewDeleteSdxByCrnParamsWithContext creates a new DeleteSdxByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSdxByCrnParamsWithContext(ctx context.Context) *DeleteSdxByCrnParams {
	var ()
	return &DeleteSdxByCrnParams{

		Context: ctx,
	}
}

// NewDeleteSdxByCrnParamsWithHTTPClient creates a new DeleteSdxByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSdxByCrnParamsWithHTTPClient(client *http.Client) *DeleteSdxByCrnParams {
	var ()
	return &DeleteSdxByCrnParams{
		HTTPClient: client,
	}
}

/*DeleteSdxByCrnParams contains all the parameters to send to the API endpoint
for the delete sdx by crn operation typically these are written to a http.Request
*/
type DeleteSdxByCrnParams struct {

	/*ClusterCrn*/
	ClusterCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) WithTimeout(timeout time.Duration) *DeleteSdxByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) WithContext(ctx context.Context) *DeleteSdxByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) WithHTTPClient(client *http.Client) *DeleteSdxByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterCrn adds the clusterCrn to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) WithClusterCrn(clusterCrn string) *DeleteSdxByCrnParams {
	o.SetClusterCrn(clusterCrn)
	return o
}

// SetClusterCrn adds the clusterCrn to the delete sdx by crn params
func (o *DeleteSdxByCrnParams) SetClusterCrn(clusterCrn string) {
	o.ClusterCrn = clusterCrn
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSdxByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterCrn
	if err := r.SetPathParam("clusterCrn", o.ClusterCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
