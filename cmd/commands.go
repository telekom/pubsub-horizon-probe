// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	expandArgs()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}

func expandArgs() {
	var args = os.Args[1:]
	for i, v := range args {
		args[i] = os.ExpandEnv(v)
	}
	os.Args = append([]string{os.Args[0]}, args...)
}
