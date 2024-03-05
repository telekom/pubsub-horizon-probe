// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package utils

import "net/http"

// HttpClient is an interface to allow mocking the http.Client or replacing it with a different implementation
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
