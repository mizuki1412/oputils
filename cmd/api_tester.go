package cmd

import (
	"github.com/mizuki1412/go-core-kit/v2/cli"
	"github.com/spf13/cobra"
	"oputils/mod_api"
)

func ApiTest() {
	apiTestCmd := &cobra.Command{
		Use: "api-test",
		Run: func(cmd *cobra.Command, args []string) {
			mod_api.Run()
		},
	}
	apiTestCmd.Flags().StringP("file", "f", "", "json配置文件路径")
	apiTestCmd.Flags().StringP("dest", "d", "", "结果输出路径")
	cli.AddChildCMD(apiTestCmd)
}
