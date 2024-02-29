// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/telekom/pubsub-horizon-probe/internal/e2e"
	"os"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start probing",
	Run: func(cmd *cobra.Command, args []string) {
		var testCase = e2e.NewTestCase(3, 30*time.Second, "./testdata/event.json")
		if testCase.Start() {
			log.Info().Msg("Test succeeded")
			os.Exit(0)
		}
		log.Error().Msg("Test didn't succeed")
		os.Exit(1)
	},
}

func init() {
	startCmd.Flags().IntP("message-count", "c", 3, "the amount of messaged sent")
	startCmd.Flags().DurationP("timeout", "t", 30*time.Second, "the timeout after which the test is considered failed")
	startCmd.Flags().String("template", "", "the template file used to generate events")

	_ = startCmd.MarkFlagFilename("template")
	_ = startCmd.MarkFlagRequired("template")
}
