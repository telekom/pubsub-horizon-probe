// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/telekom/pubsub-horizon-probe/internal/e2e"
	"testing"
	"time"
)

func TestRecordAsSent(t *testing.T) {
	var rs = e2e.NewResultSet()
	var key = "testKey"

	rs.RecordAsSent(key)

	assert.NotNil(t, rs.GetResult(key).Sent, "expected sent time to be not-nil")
}

func TestRecordAsReceived(t *testing.T) {
	var rs = e2e.NewResultSet()
	var key = "testKey"

	rs.RecordAsReceived(key)

	assert.NotNil(t, rs.GetResult(key).Received, "expected received time to be not-nil")
}

func TestIsCompleteResultSet(t *testing.T) {
	var rs = e2e.NewResultSet()
	var key = "testKey"

	rs.RecordAsSent(key)
	rs.RecordAsReceived(key)

	assert.True(t, rs.IsComplete(), "expected resultset to be complete")
}

func TestHasRecorded(t *testing.T) {
	var rs = e2e.NewResultSet()
	var key = "testKey"
	var assertions = assert.New(t)

	assertions.False(rs.HasRecorded(key), "did not expect resultset to be aware of %s yet", key)

	rs.RecordAsSent(key)

	assertions.True(rs.HasRecorded(key), "expected resultset to be aware of %s", key)
}

func TestSurpasses(t *testing.T) {
	var rs = e2e.NewResultSet()
	var key = "testKey"
	var assertions = assert.New(t)

	rs.RecordAsSent(key)
	time.Sleep(2 * time.Second)
	rs.RecordAsReceived(key)

	assertions.True(rs.Surpasses(1*time.Second), "expected latency to be >1ms")
	assertions.False(rs.Surpasses(3*time.Second), "expected latency to be <3ms")
}
