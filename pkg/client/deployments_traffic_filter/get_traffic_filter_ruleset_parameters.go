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

package deployments_traffic_filter

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

// NewGetTrafficFilterRulesetParams creates a new GetTrafficFilterRulesetParams object
// with the default values initialized.
func NewGetTrafficFilterRulesetParams() *GetTrafficFilterRulesetParams {
	var (
		includeAssociationsDefault = bool(false)
	)
	return &GetTrafficFilterRulesetParams{
		IncludeAssociations: &includeAssociationsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrafficFilterRulesetParamsWithTimeout creates a new GetTrafficFilterRulesetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTrafficFilterRulesetParamsWithTimeout(timeout time.Duration) *GetTrafficFilterRulesetParams {
	var (
		includeAssociationsDefault = bool(false)
	)
	return &GetTrafficFilterRulesetParams{
		IncludeAssociations: &includeAssociationsDefault,

		timeout: timeout,
	}
}

// NewGetTrafficFilterRulesetParamsWithContext creates a new GetTrafficFilterRulesetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTrafficFilterRulesetParamsWithContext(ctx context.Context) *GetTrafficFilterRulesetParams {
	var (
		includeAssociationsDefault = bool(false)
	)
	return &GetTrafficFilterRulesetParams{
		IncludeAssociations: &includeAssociationsDefault,

		Context: ctx,
	}
}

// NewGetTrafficFilterRulesetParamsWithHTTPClient creates a new GetTrafficFilterRulesetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTrafficFilterRulesetParamsWithHTTPClient(client *http.Client) *GetTrafficFilterRulesetParams {
	var (
		includeAssociationsDefault = bool(false)
	)
	return &GetTrafficFilterRulesetParams{
		IncludeAssociations: &includeAssociationsDefault,
		HTTPClient:          client,
	}
}

/*GetTrafficFilterRulesetParams contains all the parameters to send to the API endpoint
for the get traffic filter ruleset operation typically these are written to a http.Request
*/
type GetTrafficFilterRulesetParams struct {

	/*IncludeAssociations
	  Retrieves a list of resources that are associated to the specified ruleset.

	*/
	IncludeAssociations *bool
	/*RulesetID
	  The mandatory ruleset ID.

	*/
	RulesetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) WithTimeout(timeout time.Duration) *GetTrafficFilterRulesetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) WithContext(ctx context.Context) *GetTrafficFilterRulesetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) WithHTTPClient(client *http.Client) *GetTrafficFilterRulesetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeAssociations adds the includeAssociations to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) WithIncludeAssociations(includeAssociations *bool) *GetTrafficFilterRulesetParams {
	o.SetIncludeAssociations(includeAssociations)
	return o
}

// SetIncludeAssociations adds the includeAssociations to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) SetIncludeAssociations(includeAssociations *bool) {
	o.IncludeAssociations = includeAssociations
}

// WithRulesetID adds the rulesetID to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) WithRulesetID(rulesetID string) *GetTrafficFilterRulesetParams {
	o.SetRulesetID(rulesetID)
	return o
}

// SetRulesetID adds the rulesetId to the get traffic filter ruleset params
func (o *GetTrafficFilterRulesetParams) SetRulesetID(rulesetID string) {
	o.RulesetID = rulesetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrafficFilterRulesetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeAssociations != nil {

		// query param include_associations
		var qrIncludeAssociations bool
		if o.IncludeAssociations != nil {
			qrIncludeAssociations = *o.IncludeAssociations
		}
		qIncludeAssociations := swag.FormatBool(qrIncludeAssociations)
		if qIncludeAssociations != "" {
			if err := r.SetQueryParam("include_associations", qIncludeAssociations); err != nil {
				return err
			}
		}

	}

	// path param ruleset_id
	if err := r.SetPathParam("ruleset_id", o.RulesetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
