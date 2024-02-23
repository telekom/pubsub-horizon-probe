// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var Current = LoadConfiguration()

func LoadConfiguration() *Configuration {
	setDefault()
	var config = readConfiguration()
	applyLogLevel(config.LogLevel)
	return config
}

func ReloadConfiguration() {
	Current = LoadConfiguration()
}

func setDefault() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("probe")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("logLevel", "info")

	viper.SetDefault("publishing.endpoint", "https://horizon.example.com/events")
	viper.SetDefault("publishing.oidc.clientId", "client-id")
	viper.SetDefault("publishing.oidc.clientSecret", "client-secret")
	viper.SetDefault("publishing.oidc.url", "https://oidc.example.com/")
	viper.SetDefault("publishing.traceIdHeader", "X-B3-Traceid")
}

func readConfiguration() *Configuration {
	if err := viper.ReadInConfig(); err != nil {
		var configErr viper.ConfigFileNotFoundError
		if !errors.As(err, &configErr) {
			log.Panic().Err(err).Msg("Failed to read configuration")
		}
	}

	viper.AutomaticEnv()

	var config Configuration
	if err := viper.Unmarshal(&config); err != nil {
		log.Panic().Err(err).Msg("Failed to unmarshal configuration")
	}

	return &config
}

func applyLogLevel(level string) {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to parse log level")
	}

	log.Logger = log.Logger.Level(logLevel).Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
