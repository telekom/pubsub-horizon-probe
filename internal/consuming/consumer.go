// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package consuming

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/telekom/pubsub-horizon-probe/internal/errors"
	"github.com/telekom/pubsub-horizon-probe/internal/utils"

	"github.com/telekom/pubsub-horizon-probe/internal/auth"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
)

var Client utils.HttpClient = &http.Client{Timeout: 65 * time.Second}

type Consumer struct {
	Events chan ConsumedEvent
	config *config.ConsumerConfig
}

// NewConsumer creates a new Consumer instance with the given consumerConfig.
// It opens a connection, initializes the Events channel, and returns the Consumer instance.
// If an error occurs while opening the connection, it returns nil and the error.
func NewConsumer(consumerConfig *config.ConsumerConfig) *Consumer {
	return &Consumer{
		Events: make(chan ConsumedEvent),
		config: consumerConfig,
	}
}

// Start starts reading from the stream and sends the events to the Events channel.
// Calling this function is blocking and should only be done in a goroutine.
func (c *Consumer) Start() error {
	connection, err := openConnection(c.config)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(connection.Body)

	for {
		select {

		case <-connection.Request.Context().Done():
			if connection.StatusCode == http.StatusGatewayTimeout {
				return errors.NewTimeoutError("gateway timeout")
			}

		default:
			line, err := reader.ReadBytes('\n')
			if err != nil {
				return err
			}

			var event ConsumedEvent
			if err := json.Unmarshal(line, &event); err != nil {
				panic(err)
			}

			c.Events <- event
		}
	}
}

// openConnection opens a connection to the specified URL using the provided consumerConfig.
// It retrieves an access token using the OIDC configuration from the consumerConfig,
// creates an HTTP request with the token in the Authorization header,
// and sends the request to the URL.
// It returns the HTTP response or an error if one occurs.
func openConnection(consumerConfig *config.ConsumerConfig) (*http.Response, error) {
	var authConfig = consumerConfig.Oidc
	var url, clientId, clientSecret = authConfig.Url, authConfig.ClientId, authConfig.ClientSecret
	token, err := auth.RetrieveToken(url, clientId, clientSecret)
	if err != nil {
		return nil, err
	}

	var request, _ = http.NewRequest(http.MethodGet, consumerConfig.Endpoint, nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	request.Header.Add("Accept", "application/stream+json")

	return Client.Do(request)
}
