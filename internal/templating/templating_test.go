// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package templating_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/telekom/pubsub-horizon-probe/internal/messaging"
	"github.com/telekom/pubsub-horizon-probe/internal/templating"
	"os"
	"path/filepath"
	"testing"
)

func TestRenderFile(t *testing.T) {
	var assertions = assert.New(t)
	var testFile = "./../../testdata/event.json"
	var unexpected string

	if bytes, err := os.ReadFile(testFile); err != nil {
		path, _ := filepath.Abs(testFile)
		panic(fmt.Sprintf("could not read file '%s'", path))
	} else {
		unexpected = string(bytes)
	}

	var message messaging.Message
	rendered, err := templating.RenderFile(testFile, &message)
	assertions.NoError(err, "unexpected error")
	assertions.NotEmpty(message.Id, "message id is empty")
	assertions.NotEqual(unexpected, rendered, "expected template to render correctly")
}
