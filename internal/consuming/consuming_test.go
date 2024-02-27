// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package consuming

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/telekom/pubsub-horizon-probe/internal/auth"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"github.com/telekom/pubsub-horizon-probe/internal/test"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

var client *Consumer

func TestMain(m *testing.M) {
	viper.AddConfigPath("./../..")
	config.ReloadConfiguration()

	auth.Client = new(test.MockIrisClient)

	client = NewConsumer(&config.ConsumerConfig{
		Endpoint: "http://localhost:8080/sse",
	})

	server := test.NewMockServer()
	server.Start()

	code := m.Run()

	server.Stop()
	os.Exit(code)
}

func TestConsumer_Start(t *testing.T) {
	var assertions = assert.New(t)
	var count atomic.Int32

	go client.Start()
	go func() {
		for {
			_ = <-client.Events
			count.Add(1)
			if count.Load() == 3 {
				client.Stop()
				return
			}
		}
	}()

	var success = assertions.Eventually(func() bool { return count.Load() == 3 }, 30*time.Second, time.Second)
	assertions.True(success, "timeout exceeded")
}
