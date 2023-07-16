package cmd

import (
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/spf13/cobra"
)

func init() {

}

var rootCmd = &cobra.Command{
	Use: "main",
	Run: func(cmd *cobra.Command, args []string) {
		//initkit.BindFlags(cmd)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logkit.Fatal(err.Error())
	}
}
