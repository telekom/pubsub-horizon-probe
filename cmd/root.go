// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "Probe is a tool for performing Horizon e2e tests.",
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(initCmd, startCmd)
}
