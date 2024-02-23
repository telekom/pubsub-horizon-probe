// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package auth_test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/telekom/pubsub-horizon-probe/internal/auth"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"github.com/telekom/pubsub-horizon-probe/internal/test"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("./../..")
	config.ReloadConfiguration()

	auth.Client = new(test.MockIrisClient)

	m.Run()
}

func TestRetrieveToken(t *testing.T) {
	var oidc = config.Current.Publishing.Oidc
	var url, clientId, clientSecret = oidc.Url, oidc.ClientId, oidc.ClientSecret
	var assertions = assert.New(t)

	token, err := auth.RetrieveToken(url, clientId, clientSecret)
	assertions.NoError(err, "unexpected error")
	assertions.NotEmpty(token, "token is empty")
}
