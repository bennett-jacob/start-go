package cmd

import (
	"github.com/bennett-jacob/start-go/pkg/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var getSettingsCmd = &cobra.Command{
	Use:   "get-settings",
	Short: "Print all settings",
	Long:  "Print all settings loaded from the environment",
	Run: func(cmd *cobra.Command, args []string) {
		settings, err := config.Load()
		if err != nil {
			panic(err)
		}

		spew.Dump(settings)
	},
}
