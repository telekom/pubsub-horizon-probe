// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new configuration",
	Run: func(cmd *cobra.Command, args []string) {
		_ = config.Current
		if err := viper.SafeWriteConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
				log.Error().Msg("Config file already exists")
			} else {
				log.Fatal().Err(err).Msg("Failed to write config file")
			}
			return
		}
		log.Info().Msg("Config file created")
	},
}
