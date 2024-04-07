package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "fly-pvn",
}

func main() {
	rootCmd.AddCommand(configureCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
