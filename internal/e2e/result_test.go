// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"github.com/stretchr/testify/assert" // replace with your actual project path
	"github.com/telekom/pubsub-horizon-probe/internal/e2e"
	"testing"
	"time"
)

func TestGetLatencyMs(t *testing.T) {
	var assertions = assert.New(t)
	var r = &e2e.Result{
		Sent:     time.Now(),
		Received: time.Now().Add(10 * time.Millisecond),
	}

	var latency = r.GetLatencyMs()
	assertions.Equal(int64(10), latency, "latency does not match expectation")
}

func TestIsComplete(t *testing.T) {
	var assertions = assert.New(t)
	var r = &e2e.Result{
		Sent:     time.Now(),
		Received: time.Now(),
	}

	var isComplete = r.IsComplete()
	assertions.True(isComplete, "expected result to be complete")
}
