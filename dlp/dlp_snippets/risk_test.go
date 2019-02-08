// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

const (
	riskTopicName        = "dlp-risk-test-topic"
	riskSubscriptionName = "dlp-risk-test-sub"
)

func TestRiskNumerical(t *testing.T) {
	testutil.SystemTest(t)
	testutil.Retry(t, 20, 2*time.Second, func(r *testutil.R) {
		buf := new(bytes.Buffer)
		riskNumerical(buf, client, projectID, "bigquery-public-data", riskTopicName, riskSubscriptionName, "nhtsa_traffic_fatalities", "accident_2015", "state_number")
		wants := []string{"Created job", "Value range", "Value at"}
		got := buf.String()
		for _, want := range wants {
			if !strings.Contains(got, want) {
				r.Errorf("riskNumerical got %s, want substring %q", got, want)
			}
		}
	})
}

func TestRiskCategorical(t *testing.T) {
	testutil.SystemTest(t)
	testutil.Retry(t, 20, 2*time.Second, func(r *testutil.R) {
		buf := new(bytes.Buffer)
		riskCategorical(buf, client, projectID, "bigquery-public-data", riskTopicName, riskSubscriptionName, "nhtsa_traffic_fatalities", "accident_2015", "state_number")
		wants := []string{"Created job", "Histogram bucket", "Most common value occurs"}
		got := buf.String()
		for _, want := range wants {
			if !strings.Contains(got, want) {
				r.Errorf("riskCategorical got %s, want substring %q", got, want)
			}
		}
	})
}

func TestRiskKAnonymity(t *testing.T) {
	testutil.SystemTest(t)
	testutil.Retry(t, 20, 2*time.Second, func(r *testutil.R) {
		buf := new(bytes.Buffer)
		riskKAnonymity(buf, client, projectID, "bigquery-public-data", riskTopicName, riskSubscriptionName, "nhtsa_traffic_fatalities", "accident_2015", "state_number", "county")
		wants := []string{"Created job", "Histogram bucket", "Size range"}
		got := buf.String()
		for _, want := range wants {
			if !strings.Contains(got, want) {
				r.Errorf("riskKAnonymity got %s, want substring %q", got, want)
			}
		}
	})
}

func TestRiskLDiversity(t *testing.T) {
	testutil.SystemTest(t)
	testutil.Retry(t, 20, 2*time.Second, func(r *testutil.R) {
		buf := new(bytes.Buffer)
		riskLDiversity(buf, client, projectID, "bigquery-public-data", riskTopicName, riskSubscriptionName, "nhtsa_traffic_fatalities", "accident_2015", "city", "state_number", "county")
		wants := []string{"Created job", "Histogram bucket", "Size range"}
		got := buf.String()
		for _, want := range wants {
			if !strings.Contains(got, want) {
				r.Errorf("riskLDiversity got %s, want substring %q", got, want)
			}
		}
	})
}

func TestRiskKMap(t *testing.T) {
	testutil.SystemTest(t)
	testutil.Retry(t, 20, 2*time.Second, func(r *testutil.R) {
		buf := new(bytes.Buffer)
		riskKMap(buf, client, projectID, "bigquery-public-data", riskTopicName, riskSubscriptionName, "san_francisco", "bikeshare_trips", "US", "zip_code")
		wants := []string{"Created job", "Histogram bucket", "Anonymity range"}
		got := buf.String()
		for _, want := range wants {
			if !strings.Contains(got, want) {
				r.Errorf("riskKMap got %s, want substring %q", got, want)
			}
		}
	})
}
