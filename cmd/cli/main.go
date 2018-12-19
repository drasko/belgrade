//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"log"

	"github.com/drasko/belgrade/cli"
	"github.com/spf13/cobra"
)

func main() {
	// Root
	var rootCmd = &cobra.Command{
		Use: "belgrade",
	}

	// API commands
	newCmd := cli.NewNewCmd()
	serveCmd := cli.NewServeCmd()

	// Root Commands
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
