package cmd

import (
	"bufio"
	"github.com/mizuki1412/go-core-kit/v2/cli"
	"github.com/mizuki1412/go-core-kit/v2/library/concurrentkit"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
)

func Test() {
	apiTestCmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			path := "D:/ycj/doris-data"
			files, _ := os.ReadDir(path)
			batch := 6
			for i := 0; i < len(files); i += batch {
				g := concurrentkit.NewGroup()
				for j := 0; j < batch; j++ {
					k := i + j
					if k < len(files) {
						g.Add(func() {
							handle(path, files[k])
							err := os.Remove(path + "/" + files[k].Name())
							if err != nil {
								log.Println(err.Error())
							}
						}, true)
					}
				}
				g.Process()
			}
		},
	}
	//apiTestCmd.Flags().StringP("file", "f", "", "json配置文件路径")
	//apiTestCmd.Flags().StringP("dest", "d", "", "结果输出路径")
	cli.AddChildCMD(apiTestCmd)
}

func handle(path string, f os.DirEntry) {
	log.Println(f.Name())
	file, err := os.Open(path + "/" + f.Name())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	file2, err := os.OpenFile(path+"/a-"+f.Name(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	count := 0
	linef := ""
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		line = line[:58] + "'" + line[58:77] + "'" + line[77:]
		line = strings.TrimSpace(line) + ";\n"
		linef += line
		count++
		if count > 1000 {
			_, err = file2.WriteString(linef)
			if err != nil {
				panic(err)
			}
			count = 0
			linef = ""
		}
	}
	if count > 0 {
		_, err = file2.Write([]byte(linef))
		if err != nil {
			panic(err)
		}
	}
}
