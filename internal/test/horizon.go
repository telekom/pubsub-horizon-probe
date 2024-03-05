// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"net/http"
)

type MockHorizonClient struct{}

func (m *MockHorizonClient) Do(req *http.Request) (*http.Response, error) {
	var response = http.Response{
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}
