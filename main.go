package main

import (
	"github.com/mizuki1412/go-core-kit/v2/cli"
	"github.com/spf13/cobra"
	"oputils/cmd"
)

func main() {
	cli.RootCMD(&cobra.Command{
		Use: "main",
		Run: func(cmd *cobra.Command, args []string) {
		},
	})
	//cmd.TcpConCmd()
	//cmd.WasterCmd()
	//cmd.Test()
	cmd.Json2Excel()
	cli.Execute()
}
