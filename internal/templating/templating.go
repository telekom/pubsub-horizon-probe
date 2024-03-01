// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package templating

import (
	"github.com/google/uuid"
	"github.com/telekom/pubsub-horizon-probe/internal/messaging"
	"os"
	"strings"
	"text/template"
)

func RenderFile(file string, message *messaging.Message) (string, error) {
	var eventId, traceId = generateUuid(), generateUuid()
	var probeTemplate = template.New("probe").Funcs(map[string]any{
		"EventId": func() string { return eventId },
		"TraceId": func() string { return traceId },
	})

	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	parsed, err := probeTemplate.Parse(string(bytes))
	if err != nil {
		return "", err
	}

	var rendered = new(strings.Builder)
	if err := parsed.Execute(rendered, nil); err != nil {
		return "", err
	}

	message.Id = eventId
	return rendered.String(), nil
}

func generateUuid() string {
	var id, _ = uuid.NewRandom()
	return id.String()
}
