package cmdutil

import (
	"log"

	"github.com/spf13/cobra"
)

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MustGetString(cmd *cobra.Command, key string) string {
	value, err := cmd.Flags().GetString(key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
