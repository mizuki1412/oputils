package cmd

import (
	"github.com/mizuki1412/go-core-kit/cli"
	"github.com/spf13/cobra"
)

func ApiTest() {
	apiTestCmd := &cobra.Command{
		Use: "api-test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	apiTestCmd.Flags().StringP("file", "f", "", "json配置文件路径")
	cli.AddChildCMD(apiTestCmd)
}
