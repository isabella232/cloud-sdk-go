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

package runnerapi

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestResync(t *testing.T) {
	type args struct {
		params ResyncParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("invalid runner resync params",
				errors.New("api reference is required for the operation"),
				errors.New("id not specified and is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Fails due to unknown API response",
			args: args{params: ResyncParams{
				ID:     "2c221bd86b7f48959a59ee3128d5c5e8",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusForbidden,
					Body:       mock.NewStringBody(`{"error": "some forbidden error"}`),
				}}),
			}},
			err: errors.New(`{"error": "some forbidden error"}`),
		},
		{
			name: "Fails due to API error",
			args: args{params: ResyncParams{
				ID:     "2c221bd86b7f48959a59ee3128d5c5e8",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New("error with API"),
				}),
			}},
			err: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/runners/2c221bd86b7f48959a59ee3128d5c5e8/_resync",
				Err: errors.New("error with API"),
			},
		},
		{
			name: "Succeeds to resynchronize Kibana instance without errors",
			args: args{params: ResyncParams{
				ID:     "d324608c97154bdba2dff97511d40368",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Body:       mock.NewStringBody(`{}`),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/runners/d324608c97154bdba2dff97511d40368/_resync",
					},
				}),
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Resync(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}

func TestResyncAll(t *testing.T) {
	type args struct {
		params ResyncAllParams
	}
	tests := []struct {
		name string
		args args
		err  error
		want *models.ModelVersionIndexSynchronizationResults
	}{
		{
			name: "Fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("invalid runner resync all params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Fails due to unknown API response",
			args: args{params: ResyncAllParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					StatusCode: http.StatusForbidden,
					Body:       mock.NewStringBody(`{"error": "some forbidden error"}`),
				}}),
			}},
			err: errors.New(`{"error": "some forbidden error"}`),
		},
		{
			name: "Fails due to API error",
			args: args{params: ResyncAllParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New("error with API"),
				}),
			}},
			err: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/runners/_resync?skip_matching_version=true",
				Err: errors.New("error with API"),
			},
		},
		{
			name: "Succeeds to re-synchronize all Kibana instances without errors",
			args: args{params: ResyncAllParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusAccepted,
						Body:       mock.NewStringBody(`{}`),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/runners/_resync",
						Query: url.Values{
							"skip_matching_version": []string{"true"},
						},
					},
				}),
			}},
			want: &models.ModelVersionIndexSynchronizationResults{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResyncAll(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
