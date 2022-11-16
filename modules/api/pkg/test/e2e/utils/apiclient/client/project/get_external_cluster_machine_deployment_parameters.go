// Code generated by go-swagger; DO NOT EDIT.

package project

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
)

// NewGetExternalClusterMachineDeploymentParams creates a new GetExternalClusterMachineDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalClusterMachineDeploymentParams() *GetExternalClusterMachineDeploymentParams {
	return &GetExternalClusterMachineDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalClusterMachineDeploymentParamsWithTimeout creates a new GetExternalClusterMachineDeploymentParams object
// with the ability to set a timeout on a request.
func NewGetExternalClusterMachineDeploymentParamsWithTimeout(timeout time.Duration) *GetExternalClusterMachineDeploymentParams {
	return &GetExternalClusterMachineDeploymentParams{
		timeout: timeout,
	}
}

// NewGetExternalClusterMachineDeploymentParamsWithContext creates a new GetExternalClusterMachineDeploymentParams object
// with the ability to set a context for a request.
func NewGetExternalClusterMachineDeploymentParamsWithContext(ctx context.Context) *GetExternalClusterMachineDeploymentParams {
	return &GetExternalClusterMachineDeploymentParams{
		Context: ctx,
	}
}

// NewGetExternalClusterMachineDeploymentParamsWithHTTPClient creates a new GetExternalClusterMachineDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalClusterMachineDeploymentParamsWithHTTPClient(client *http.Client) *GetExternalClusterMachineDeploymentParams {
	return &GetExternalClusterMachineDeploymentParams{
		HTTPClient: client,
	}
}

/*
GetExternalClusterMachineDeploymentParams contains all the parameters to send to the API endpoint

	for the get external cluster machine deployment operation.

	Typically these are written to a http.Request.
*/
type GetExternalClusterMachineDeploymentParams struct {

	// ClusterID.
	ClusterID string

	// MachinedeploymentID.
	MachineDeploymentID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get external cluster machine deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalClusterMachineDeploymentParams) WithDefaults() *GetExternalClusterMachineDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get external cluster machine deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalClusterMachineDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) WithTimeout(timeout time.Duration) *GetExternalClusterMachineDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) WithContext(ctx context.Context) *GetExternalClusterMachineDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) WithHTTPClient(client *http.Client) *GetExternalClusterMachineDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) WithClusterID(clusterID string) *GetExternalClusterMachineDeploymentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithMachineDeploymentID adds the machinedeploymentID to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) WithMachineDeploymentID(machinedeploymentID string) *GetExternalClusterMachineDeploymentParams {
	o.SetMachineDeploymentID(machinedeploymentID)
	return o
}

// SetMachineDeploymentID adds the machinedeploymentId to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) SetMachineDeploymentID(machinedeploymentID string) {
	o.MachineDeploymentID = machinedeploymentID
}

// WithProjectID adds the projectID to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) WithProjectID(projectID string) *GetExternalClusterMachineDeploymentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get external cluster machine deployment params
func (o *GetExternalClusterMachineDeploymentParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalClusterMachineDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param machinedeployment_id
	if err := r.SetPathParam("machinedeployment_id", o.MachineDeploymentID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}