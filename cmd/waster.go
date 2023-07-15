package cmd

import (
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/spf13/cobra"
	"waster/mod_waster"
)

func init() {
	rootCmd.AddCommand(wasterCMD)
	wasterCMD.Flags().Int("delay", 0, "可选：循环计算时的延迟ms，用于控制cpu")
	wasterCMD.Flags().Int("core", 1, "可选：核心数")
	wasterCMD.Flags().Int("mem", 20, "内存增加的%，必填")
	wasterCMD.Flags().Int("ssize", 0, "增加的GB")
	wasterCMD.Flags().String("spath", "", "生成的filename")
}

var wasterCMD = &cobra.Command{
	Use: "waster",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		mod_waster.Run()
		select {}
	},
}
