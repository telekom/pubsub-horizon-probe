// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package errors

type TimeoutError struct {
	err string
}

func (e *TimeoutError) Error() string {
	return e.err
}

func (e *TimeoutError) Timeout() bool {
	return true
}

func (e *TimeoutError) Temporary() bool {
	return true
}

func NewTimeoutError(err string) error {
	return &TimeoutError{
		err: err,
	}
}
