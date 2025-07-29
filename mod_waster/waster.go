package mod_waster

import (
	"github.com/mizuki1412/go-core-kit/v2/class"
	"github.com/mizuki1412/go-core-kit/v2/library/c"
	"github.com/mizuki1412/go-core-kit/v2/library/filekit"
	"github.com/mizuki1412/go-core-kit/v2/library/timekit"
	"github.com/mizuki1412/go-core-kit/v2/service/configkit"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math"
	"runtime"
	"time"
)

func Run() {
	//log.Println(runtime.NumCPU())
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	spath := configkit.GetString("spath")
	ssize := configkit.GetInt("ssize", 0)
	if spath != "" && ssize > 0 {
		log.Println("开始生成文件", spath, ssize, "GB")
		_ = filekit.CheckFilePath(spath)
		d := make([]byte, 1024*1024*1024, 1024*1024*1024)
		for k := 0; k < ssize; k++ {
			_ = filekit.WriteFileAppend(spath, d)
		}
	}
	// cpu调整
	cpuR := configkit.GetInt("cpu")
	if cpuR > 0 {
		if cpuR > 100 {
			log.Println("cpu不可大于100")
			return
		}
		// 按核数开协程
		for i := 0; i < cores; i++ {
			c.RecoverGoFuncWrapper(func() {
				sampleInterval := 500
				// 采样间隔 500ms * 占比
				busyTime := time.Duration(sampleInterval*cpuR/100) * time.Millisecond
				sleepTime := time.Duration(sampleInterval)*time.Millisecond - busyTime
				for {
					lastTime := time.Now()
					workUntil := lastTime.Add(busyTime)
					for time.Now().Before(workUntil) {
						// 空循环可能会被优化
						_ = math.Sqrt(float64(time.Now().UnixNano()))
					}
					// 睡眠阶段
					if sleepTime > 0 {
						time.Sleep(sleepTime)
					}
				}
			})
		}
	}

	for {
		percent, _ := cpu.Percent(time.Second, false)
		memInfo, _ := mem.VirtualMemory()
		m0 := configkit.GetInt("mem")
		mm := class.NewDecimal(memInfo.UsedPercent)
		// 内存调整
		if m0 > 0 {
			m := class.NewDecimal(m0)
			if mm.Float64() <= m.Float64() {
				freeMemory()
				allocateMemory(uint64(m.Sub(mm).Div(class.NewDecimal(100)).Mul(class.NewDecimal(memInfo.Total)).Decimal.IntPart() + 1))
			} else if memoryHog != nil && mm.Float64()-m.Float64() >= 10 {
				// 如果实际超过10%
				freeMemory()
			}
		}
		log.Println("cpu%:", class.NewDecimal(percent[0]).Round(2).Decimal.String(), ", mem%: ", mm.Round(2).Decimal.String())
		timekit.Sleep(2000)
	}
}

var memoryHog []byte

func allocateMemory(amount uint64) {
	// Allocate new memory and keep it referenced
	newMemory := make([]byte, amount)
	for i := range newMemory {
		newMemory[i] = byte(i % 256)
	}
	memoryHog = append(memoryHog, newMemory...)
}

func freeMemory() {
	// Release memory by clearing the slice
	memoryHog = nil
	// Force garbage collection
	runtime.GC()
	// Give GC some time to work
	time.Sleep(1 * time.Second)
}
