// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudCloudinstancesPutParams creates a new PcloudCloudinstancesPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesPutParams() *PcloudCloudinstancesPutParams {
	return &PcloudCloudinstancesPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesPutParamsWithTimeout creates a new PcloudCloudinstancesPutParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesPutParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesPutParams {
	return &PcloudCloudinstancesPutParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesPutParamsWithContext creates a new PcloudCloudinstancesPutParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesPutParamsWithContext(ctx context.Context) *PcloudCloudinstancesPutParams {
	return &PcloudCloudinstancesPutParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesPutParamsWithHTTPClient creates a new PcloudCloudinstancesPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesPutParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesPutParams {
	return &PcloudCloudinstancesPutParams{
		HTTPClient: client,
	}
}

/* PcloudCloudinstancesPutParams contains all the parameters to send to the API endpoint
   for the pcloud cloudinstances put operation.

   Typically these are written to a http.Request.
*/
type PcloudCloudinstancesPutParams struct {

	/* Body.

	   Parameters for updating a Power Cloud Instance
	*/
	Body *models.CloudInstanceUpdate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesPutParams) WithDefaults() *PcloudCloudinstancesPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) WithContext(ctx context.Context) *PcloudCloudinstancesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) WithBody(body *models.CloudInstanceUpdate) *PcloudCloudinstancesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) SetBody(body *models.CloudInstanceUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances put params
func (o *PcloudCloudinstancesPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}