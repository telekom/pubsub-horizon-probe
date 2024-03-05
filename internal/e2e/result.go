// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"time"
)

type Result struct {
	Sent     time.Time
	Received time.Time
}

func (r *Result) GetLatencyMs() int64 {
	return r.GetLatency().Milliseconds()
}

func (r *Result) GetLatency() time.Duration {
	return r.Received.Sub(r.Sent)
}

func (r *Result) IsComplete() bool {
	return !r.Sent.IsZero() && !r.Received.IsZero()
}
