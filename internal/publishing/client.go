// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package publishing

import (
	"fmt"
	"github.com/telekom/pubsub-horizon-probe/internal/errors"
	"github.com/telekom/pubsub-horizon-probe/internal/messaging"
	"net/http"
	"strings"
	"time"

	"github.com/telekom/pubsub-horizon-probe/internal/auth"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"github.com/telekom/pubsub-horizon-probe/internal/templating"
	"github.com/telekom/pubsub-horizon-probe/internal/utils"
)

var Client utils.HttpClient

func init() {
	Client = &http.Client{Timeout: 10 * time.Second}
}

// Publish publishes an event produced to the given url by rendering the given templateFile.
// This returned string will be the event-id if the publishing request succeeds.
func Publish(publishingConfig *config.PublishingConfig, templateFile string) (string, error) {
	var message = new(messaging.Message)
	renderedFile, err := templating.RenderFile(templateFile, message)
	if err != nil {
		return "", err
	}

	var request, _ = http.NewRequest(http.MethodPost, publishingConfig.Endpoint, strings.NewReader(renderedFile))
	request.Header.Set("Content-Type", "application/json")

	var authConfig = &publishingConfig.Oidc
	token, err := auth.RetrieveToken(authConfig.Url, authConfig.ClientId, authConfig.ClientSecret)
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := Client.Do(request)
	if err != nil {
		return "", err
	}

	if response.StatusCode == http.StatusAccepted {
		return "", errors.ErrNoSubscribers
	}

	return message.Id, nil
}
