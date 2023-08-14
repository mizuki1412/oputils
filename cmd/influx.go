package cmd

import (
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/spf13/cobra"
	"waster/mod_influx"
)

func init() {
	rootCmd.AddCommand(influxCmd)
	influxCmd.Flags().String("bucket", "", "")
	influxCmd.Flags().String("org", "", "")
	influxCmd.Flags().String("token", "", "")
	influxCmd.Flags().String("url", "", "")
	influxCmd.Flags().String("flux", "", "语句文件路径")
	influxCmd.Flags().String("csv", "", "csv文件路径")
	influxCmd.Flags().String("dstart", "", "删除开始时间")
	influxCmd.Flags().String("dend", "", "删除结束时间")
	influxCmd.Flags().Int("type", 0, "0-print, 1-csv, 10-delete")
}

var influxCmd = &cobra.Command{
	Use: "influx",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		mod_influx.Exec()
	},
}
