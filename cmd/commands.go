// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
