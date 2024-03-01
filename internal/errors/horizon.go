// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package errors

import "errors"

var (
	ErrNoSubscribers  = errors.New("unable to produce event because there are no subscribers")
	ErrGatewayTimeout = NewTimeoutError("gateway timeout")
)
