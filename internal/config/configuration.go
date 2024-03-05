// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package config

type Configuration struct {
	LogLevel   string           `mapstructure:"logLevel"`
	Publishing PublishingConfig `mapstructure:"publishing"`
	Consuming  ConsumerConfig   `mapstructure:"consuming"`
}

type PublishingConfig struct {
	Oidc     OidcConfiguration `mapstructure:"oidc"`
	Endpoint string            `mapstructure:"endpoint"`
}

type ConsumerConfig struct {
	Oidc     OidcConfiguration `mapstructure:"oidc"`
	Endpoint string            `mapstructure:"endpoint"`
}

type OidcConfiguration struct {
	Url          string `mapstructure:"url"`
	ClientId     string `mapstructure:"clientId"`
	ClientSecret string `mapstructure:"clientSecret"`
}
