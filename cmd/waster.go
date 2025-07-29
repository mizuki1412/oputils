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
		},
	}
	wasterCMD.Flags().Int("cpu", 20, "设定要增加的cpu%")
	wasterCMD.Flags().Int("mem", 20, "设定的内存%")
	wasterCMD.Flags().Int("ssize", 0, "增加的GB")
	wasterCMD.Flags().String("spath", "", "生成的filename")
	return wasterCMD
}
