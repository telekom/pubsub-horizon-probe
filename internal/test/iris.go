// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type MockIrisClient struct{}

func (m *MockIrisClient) Do(req *http.Request) (*http.Response, error) {
	var response = new(http.Response)

	var contentType = req.Header.Get("Content-Type")
	if contentType != "application/x-www-form-urlencoded" {
		response.StatusCode = http.StatusUnsupportedMediaType
		return response, nil
	}

	_, _, ok := req.BasicAuth()
	if !ok {
		response.StatusCode = http.StatusBadRequest
		return response, nil
	}

	var body = struct {
		AccessToken string `json:"access_token"`
	}{"foobar"}

	bodyBytes, _ := json.Marshal(body)
	response.StatusCode = http.StatusOK
	response.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	return response, nil
}
