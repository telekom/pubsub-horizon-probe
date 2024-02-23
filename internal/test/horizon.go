// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/google/uuid"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"net/http"
)

type MockHorizonClient struct{}

func (m *MockHorizonClient) Do(req *http.Request) (*http.Response, error) {
	var traceId, _ = uuid.NewUUID()
	var response = http.Response{
		StatusCode: http.StatusCreated,
		Header: http.Header{
			config.Current.Publishing.TraceIdHeader: []string{traceId.String()},
		},
	}

	return &response, nil
}
