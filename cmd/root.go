package cmd

import (
	"github.com/mizuki1412/go-core-kit/class"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/library/filekit"
	"github.com/mizuki1412/go-core-kit/library/timekit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/mizuki1412/go-core-kit/service/cronkit"
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"log"
	"runtime"
	"time"
)

func init() {
	rootCmd.Flags().Int("delay", 0, "可选：循环计算时的延迟ms，用于控制cpu")
	rootCmd.Flags().Int("core", 1, "可选：核心数")
	rootCmd.Flags().Int("mem", 20, "内存增加的%，必填")
	rootCmd.Flags().Int("ssize", 0, "增加的GB")
	rootCmd.Flags().String("spath", "", "生成的filename")
}

var memory []byte

var rootCmd = &cobra.Command{
	Use: "waster",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		num := runtime.NumCPU() //获取CPU核心数量
		//if num > 1 {
		runtime.GOMAXPROCS(configkit.GetInt("core", 1))
		//}
		log.Println("核心: ", num)

		spath := configkit.GetStringD("spath")
		ssize := configkit.GetInt("ssize", 0)
		if spath != "" && ssize > 0 {
			log.Println("开始生成文件", spath, ssize, "GB")
			_ = filekit.CheckFilePath(spath)
			d := make([]byte, 1024*1024*1024, 1024*1024*1024)
			for k := 0; k < ssize; k++ {
				_ = filekit.WriteFileAppend(spath, d)
			}
		}

		cronkit.AddFunc("@every 3s", func() {
			percent, _ := cpu.Percent(time.Second, false)
			memInfo, _ := mem.VirtualMemory()
			log.Println("cpu%:", class.NewDecimal(percent[0]).Round(2).Decimal.String(), ", mem%: ", class.NewDecimal(memInfo.UsedPercent).Round(2).Decimal.String())
		})
		cronkit.Scheduler().Start()

		memInfo, _ := mem.VirtualMemory()
		l := cast.ToInt(cast.ToFloat64(memInfo.Total) * cast.ToFloat64(configkit.GetInt("mem", 20)) / 100)
		memory = make([]byte, l, l)
		for i := 0; i < num; i++ {
			go func() {
				for {
					for i := 0; i < len(memory); i++ {
						memory[i] = 1
					}
					if configkit.GetInt("delay", 0) > 0 {
						timekit.Sleep(cast.ToInt64(configkit.GetInt("delay", 0)))
					}
				}
			}()
		}
		select {}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logkit.Fatal(err.Error())
	}
}
