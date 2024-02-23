// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package config

type Configuration struct {
	LogLevel   string           `mapstructure:"logLevel"`
	Publishing PublishingConfig `mapstructure:"publishing"`
}

type PublishingConfig struct {
	Oidc          OidcConfiguration `mapstructure:"oidc"`
	Endpoint      string            `mapstructure:"endpoint"`
	TraceIdHeader string            `mapstructure:"traceIdHeader"`
}

type OidcConfiguration struct {
	Url          string `mapstructure:"url"`
	ClientId     string `mapstructure:"clientId"`
	ClientSecret string `mapstructure:"clientSecret"`
}
