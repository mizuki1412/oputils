package cmd

import (
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
	"github.com/spf13/cobra"
	"net"
	"time"
)

func TcpConCmd() *cobra.Command {
	tcpConnectCMD := &cobra.Command{
		Use: "tcp-con",
		Run: func(cmd *cobra.Command, args []string) {
			d := net.Dialer{Timeout: time.Duration(30) * time.Second}
			addr := configkit.GetString("addr")
			_, err := d.Dial("tcp", addr)
			if err != nil {
				println("connect error:", addr, err.Error())
			} else {
				println("connect success")
			}
		},
	}
	tcpConnectCMD.Flags().String("addr", "", "ip:port")
	return tcpConnectCMD
}
