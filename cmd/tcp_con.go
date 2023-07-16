package cmd

import (
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/spf13/cobra"
	"net"
	"time"
)

func init() {
	rootCmd.AddCommand(tcpConnectCMD)
	tcpConnectCMD.Flags().String("addr", "", "ip:port")
}

var tcpConnectCMD = &cobra.Command{
	Use: "tcp-con",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		d := net.Dialer{Timeout: time.Duration(30) * time.Second}
		addr := configkit.GetStringD("addr")
		_, err := d.Dial("tcp", addr)
		if err != nil {
			println("connect error:", addr, err.Error())
		} else {
			println("connect success")
		}
	},
}
