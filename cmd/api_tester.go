package cmd

import (
	"github.com/spf13/cobra"
)

func ApiTest() *cobra.Command {
	apiTestCmd := &cobra.Command{
		Use: "api-test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	apiTestCmd.Flags().StringP("file", "f", "", "json配置文件路径")
	return apiTestCmd
}
