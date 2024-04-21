package cmd

import (
	"github.com/spf13/cobra"
	"oputils/mod_waster"
)

func WasterCmd() *cobra.Command {
	wasterCMD := &cobra.Command{
		Use: "waster",
		Run: func(cmd *cobra.Command, args []string) {
			mod_waster.Run()
			select {}
		},
	}
	wasterCMD.Flags().Int("delay", 0, "可选：循环计算时的延迟ms，用于控制cpu")
	wasterCMD.Flags().Int("core", 1, "可选：核心数")
	wasterCMD.Flags().Int("mem", 20, "内存增加的%，必填")
	wasterCMD.Flags().Int("ssize", 0, "增加的GB")
	wasterCMD.Flags().String("spath", "", "生成的filename")
	return wasterCMD
}
