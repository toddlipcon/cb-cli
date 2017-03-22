package sssd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPostPublicSssdParams creates a new PostPublicSssdParams object
// with the default values initialized.
func NewPostPublicSssdParams() *PostPublicSssdParams {
	var ()
	return &PostPublicSssdParams{}
}

/*PostPublicSssdParams contains all the parameters to send to the API endpoint
for the post public sssd operation typically these are written to a http.Request
*/
type PostPublicSssdParams struct {

	/*Body*/
	Body *models_cloudbreak.SssdConfigRequest
}

// WithBody adds the body to the post public sssd params
func (o *PostPublicSssdParams) WithBody(body *models_cloudbreak.SssdConfigRequest) *PostPublicSssdParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicSssdParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.SssdConfigRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}