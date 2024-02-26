// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/telekom/pubsub-horizon-probe/internal/config"
	"github.com/telekom/pubsub-horizon-probe/internal/consuming"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start probing",
	Run: func(cmd *cobra.Command, args []string) {
		consumer := consuming.NewConsumer(&config.Current.Consuming)

		go func() {
			for {
				if err := consumer.Start(); err != nil {
					if os.IsTimeout(err) {
						log.Debug().Msg("Connection timed out. Reconnecting...")
						continue
					}

					if errors.Is(err, io.EOF) {
						log.Debug().Msg("Received end of stream (EOF). Reconnecting...")
						continue
					}
					log.Fatal().Err(err).Msg("Error while consuming events")
				}
			}
		}()

		for {
			event, ok := <-consumer.Events
			if !ok {
				log.Fatal().Msg("Consumer channel closed")
			} else {
				fmt.Printf("Received Event: %+v\n", event)
			}
		}
	},
}
