// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package clusters_elasticsearch

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
	"github.com/go-openapi/swag"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSetEsClusterInstancesAllSettingsOverridesParams creates a new SetEsClusterInstancesAllSettingsOverridesParams object
// with the default values initialized.
func NewSetEsClusterInstancesAllSettingsOverridesParams() *SetEsClusterInstancesAllSettingsOverridesParams {
	var (
		restartAfterUpdateDefault = bool(false)
	)
	return &SetEsClusterInstancesAllSettingsOverridesParams{
		RestartAfterUpdate: &restartAfterUpdateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSetEsClusterInstancesAllSettingsOverridesParamsWithTimeout creates a new SetEsClusterInstancesAllSettingsOverridesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetEsClusterInstancesAllSettingsOverridesParamsWithTimeout(timeout time.Duration) *SetEsClusterInstancesAllSettingsOverridesParams {
	var (
		restartAfterUpdateDefault = bool(false)
	)
	return &SetEsClusterInstancesAllSettingsOverridesParams{
		RestartAfterUpdate: &restartAfterUpdateDefault,

		timeout: timeout,
	}
}

// NewSetEsClusterInstancesAllSettingsOverridesParamsWithContext creates a new SetEsClusterInstancesAllSettingsOverridesParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetEsClusterInstancesAllSettingsOverridesParamsWithContext(ctx context.Context) *SetEsClusterInstancesAllSettingsOverridesParams {
	var (
		restartAfterUpdateDefault = bool(false)
	)
	return &SetEsClusterInstancesAllSettingsOverridesParams{
		RestartAfterUpdate: &restartAfterUpdateDefault,

		Context: ctx,
	}
}

// NewSetEsClusterInstancesAllSettingsOverridesParamsWithHTTPClient creates a new SetEsClusterInstancesAllSettingsOverridesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetEsClusterInstancesAllSettingsOverridesParamsWithHTTPClient(client *http.Client) *SetEsClusterInstancesAllSettingsOverridesParams {
	var (
		restartAfterUpdateDefault = bool(false)
	)
	return &SetEsClusterInstancesAllSettingsOverridesParams{
		RestartAfterUpdate: &restartAfterUpdateDefault,
		HTTPClient:         client,
	}
}

/*SetEsClusterInstancesAllSettingsOverridesParams contains all the parameters to send to the API endpoint
for the set es cluster instances all settings overrides operation typically these are written to a http.Request
*/
type SetEsClusterInstancesAllSettingsOverridesParams struct {

	/*Body
	  The settings to override for the instances

	*/
	Body *models.ElasticsearchClusterInstanceSettingsOverrides
	/*ClusterID
	  The Elasticsearch cluster identifier.

	*/
	ClusterID string
	/*RestartAfterUpdate
	  After overrides are applied, restarts the instances.

	*/
	RestartAfterUpdate *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WithTimeout(timeout time.Duration) *SetEsClusterInstancesAllSettingsOverridesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WithContext(ctx context.Context) *SetEsClusterInstancesAllSettingsOverridesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WithHTTPClient(client *http.Client) *SetEsClusterInstancesAllSettingsOverridesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WithBody(body *models.ElasticsearchClusterInstanceSettingsOverrides) *SetEsClusterInstancesAllSettingsOverridesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) SetBody(body *models.ElasticsearchClusterInstanceSettingsOverrides) {
	o.Body = body
}

// WithClusterID adds the clusterID to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WithClusterID(clusterID string) *SetEsClusterInstancesAllSettingsOverridesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithRestartAfterUpdate adds the restartAfterUpdate to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WithRestartAfterUpdate(restartAfterUpdate *bool) *SetEsClusterInstancesAllSettingsOverridesParams {
	o.SetRestartAfterUpdate(restartAfterUpdate)
	return o
}

// SetRestartAfterUpdate adds the restartAfterUpdate to the set es cluster instances all settings overrides params
func (o *SetEsClusterInstancesAllSettingsOverridesParams) SetRestartAfterUpdate(restartAfterUpdate *bool) {
	o.RestartAfterUpdate = restartAfterUpdate
}

// WriteToRequest writes these params to a swagger request
func (o *SetEsClusterInstancesAllSettingsOverridesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.RestartAfterUpdate != nil {

		// query param restart_after_update
		var qrRestartAfterUpdate bool
		if o.RestartAfterUpdate != nil {
			qrRestartAfterUpdate = *o.RestartAfterUpdate
		}
		qRestartAfterUpdate := swag.FormatBool(qrRestartAfterUpdate)
		if qRestartAfterUpdate != "" {
			if err := r.SetQueryParam("restart_after_update", qRestartAfterUpdate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
