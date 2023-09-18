package cmd

import (
	"github.com/spf13/cobra"
	"oputils/mod_influx"
)

var influxCmd = &cobra.Command{
	Use: "influx",
	Run: func(cmd *cobra.Command, args []string) {
		mod_influx.Exec()
	},
}
