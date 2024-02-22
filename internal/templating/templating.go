// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package templating

import (
	"github.com/google/uuid"
	"os"
	"strings"
	"text/template"
)

var probeTemplate = template.New("probe").Funcs(map[string]any{
	"EventId": generateUuid,
	"TraceId": generateUuid,
})

func RenderFile(file string) (string, error) {
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

	return rendered.String(), nil
}

func generateUuid() string {
	var id, _ = uuid.NewUUID()
	return id.String()
}
