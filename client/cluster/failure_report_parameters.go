package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// NewFailureReportParams creates a new FailureReportParams object
// with the default values initialized.
func NewFailureReportParams() *FailureReportParams {
	var ()
	return &FailureReportParams{}
}

/*FailureReportParams contains all the parameters to send to the API endpoint
for the failure report operation typically these are written to a http.Request
*/
type FailureReportParams struct {

	/*Body*/
	Body *models.FailureReport
	/*ID*/
	ID int64
}

// WithBody adds the body to the failure report params
func (o *FailureReportParams) WithBody(body *models.FailureReport) *FailureReportParams {
	o.Body = body
	return o
}

// WithID adds the id to the failure report params
func (o *FailureReportParams) WithID(id int64) *FailureReportParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *FailureReportParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.FailureReport)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
