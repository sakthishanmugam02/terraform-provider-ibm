// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSubnetsSubnetIDPublicGatewayParams creates a new DeleteSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized.
func NewDeleteSubnetsSubnetIDPublicGatewayParams() *DeleteSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &DeleteSubnetsSubnetIDPublicGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubnetsSubnetIDPublicGatewayParamsWithTimeout creates a new DeleteSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubnetsSubnetIDPublicGatewayParamsWithTimeout(timeout time.Duration) *DeleteSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &DeleteSubnetsSubnetIDPublicGatewayParams{

		timeout: timeout,
	}
}

// NewDeleteSubnetsSubnetIDPublicGatewayParamsWithContext creates a new DeleteSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubnetsSubnetIDPublicGatewayParamsWithContext(ctx context.Context) *DeleteSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &DeleteSubnetsSubnetIDPublicGatewayParams{

		Context: ctx,
	}
}

// NewDeleteSubnetsSubnetIDPublicGatewayParamsWithHTTPClient creates a new DeleteSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubnetsSubnetIDPublicGatewayParamsWithHTTPClient(client *http.Client) *DeleteSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &DeleteSubnetsSubnetIDPublicGatewayParams{
		HTTPClient: client,
	}
}

/*DeleteSubnetsSubnetIDPublicGatewayParams contains all the parameters to send to the API endpoint
for the delete subnets subnet ID public gateway operation typically these are written to a http.Request
*/
type DeleteSubnetsSubnetIDPublicGatewayParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*SubnetID
	  The subnet identifier

	*/
	SubnetID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WithTimeout(timeout time.Duration) *DeleteSubnetsSubnetIDPublicGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WithContext(ctx context.Context) *DeleteSubnetsSubnetIDPublicGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WithHTTPClient(client *http.Client) *DeleteSubnetsSubnetIDPublicGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WithGeneration(generation int64) *DeleteSubnetsSubnetIDPublicGatewayParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithSubnetID adds the subnetID to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WithSubnetID(subnetID string) *DeleteSubnetsSubnetIDPublicGatewayParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) SetSubnetID(subnetID string) {
	o.SubnetID = subnetID
}

// WithVersion adds the version to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WithVersion(version string) *DeleteSubnetsSubnetIDPublicGatewayParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete subnets subnet ID public gateway params
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubnetsSubnetIDPublicGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param subnet_id
	if err := r.SetPathParam("subnet_id", o.SubnetID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}