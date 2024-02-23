// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package publishing_test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/telekom/pubsub-horizon-probe/internal/auth"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"github.com/telekom/pubsub-horizon-probe/internal/publishing"
	"github.com/telekom/pubsub-horizon-probe/internal/test"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("./../..")
	config.ReloadConfiguration()

	auth.Client = new(test.MockIrisClient)
	publishing.Client = new(test.MockHorizonClient)

	m.Run()
}

func TestPublish(t *testing.T) {
	var assertions = assert.New(t)

	traceId, err := publishing.Publish(&config.Current.Publishing, "./../../testdata/event.json")
	assertions.NoError(err, "unexpected error")
	assertions.NotEmpty(traceId)
}
