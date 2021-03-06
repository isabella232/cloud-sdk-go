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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsageStats Information about the current usage.
//
// swagger:model UsageStats
type UsageStats struct {

	// The total amount of memory available to your system
	// Required: true
	TotalConnectedMemoryTotal *int32 `json:"total_connected_memory_total"`
}

// Validate validates this usage stats
func (m *UsageStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotalConnectedMemoryTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsageStats) validateTotalConnectedMemoryTotal(formats strfmt.Registry) error {

	if err := validate.Required("total_connected_memory_total", "body", m.TotalConnectedMemoryTotal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsageStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsageStats) UnmarshalBinary(b []byte) error {
	var res UsageStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
