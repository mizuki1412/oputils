package cmd

import (
	"github.com/mizuki1412/go-core-kit/v2/cli"
	"github.com/mizuki1412/go-core-kit/v2/library/filekit"
	"github.com/mizuki1412/go-core-kit/v2/library/jsonkit"
	"github.com/mizuki1412/go-core-kit/v2/service/excelkit"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
	"log"
)

func Json2Excel() {
	json2excelCMD := &cobra.Command{
		Use: "json2excel",
		Run: func(cmd *cobra.Command, args []string) {
			str, _ := filekit.ReadString(viper.GetString("in"))
			var arr []map[string]any
			jsonkit.ParseObj(str, &arr)
			if len(arr) == 0 {
				log.Fatal("data null")
			}

			f := excelize.NewFile()
			titleStyle, _ := f.NewStyle(&excelize.Style{
				Alignment: &excelize.Alignment{
					Horizontal: "center",
					Vertical:   "center",
				},
				Font: &excelize.Font{
					Size: 15,
				},
			})
			keyStyle, _ := f.NewStyle(&excelize.Style{
				Alignment: &excelize.Alignment{
					Horizontal: "center",
					Vertical:   "center",
					WrapText:   true,
				},
				Font: &excelize.Font{
					Size: 12,
				},
				Border: excelkit.BorderStyleDefault(),
			})
			//valueStyle, _ := f.NewStyle(&excelize.Style{
			//	Alignment: &excelize.Alignment{
			//		//Horizontal: "right",
			//		Vertical: "center",
			//		WrapText: true,
			//	},
			//	Font: &excelize.Font{
			//		Size: 12,
			//	},
			//	Border: excelkit.BorderStyleDefault(),
			//})
			sheet := f.GetSheetName(0)
			//_ = f.MergeCell(sheet, "A1", baseConversion(len(arr[0])-1)+"1")
			_ = f.SetCellStyle(sheet, "A1", baseConversion(len(arr[0])-1)+"1", titleStyle)
			//_ = f.SetCellValue(f.GetSheetName(0), "A1", param.Title)
			// key title
			ki := 0
			var keys []string
			for k := range arr[0] {
				keys = append(keys, k)
				cell := baseConversion(ki) + "1"
				_ = f.SetCellStyle(sheet, cell, cell, keyStyle)
				_ = f.SetCellValue(sheet, cell, k)
				//if v.Width > 0 {
				//	err = f.SetColWidth(param.Sheet, baseConversion(v.Index), baseConversion(v.Index), v.Width)
				//}
				ki++
			}
			// data
			for i, data := range arr {
				index := i + 2
				// 每个cell加style
				for j, k := range keys {
					cell := baseConversion(j) + cast.ToString(index)
					_ = f.SetCellStyle(sheet, cell, cell, keyStyle)
					_ = f.SetCellValue(sheet, cell, cast.ToString(data[k]))
				}
			}
			f.SaveAs(viper.GetString("out"))
		},
	}
	json2excelCMD.Flags().String("in", "", "json文件路径")
	json2excelCMD.Flags().String("out", "", "导出路径")
	cli.AddChildCMD(json2excelCMD)
}

var base = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func baseConversion(number int) string {
	var ret string
	var m []rune
	length := len(base)
	flag := false
	for number >= 0 {
		if flag {
			break
		}
		index := number % length
		m = append(m, base[index])
		temp := number / length
		if temp == 0 {
			flag = true
		}
		number = temp - 1

	}
	for i := len(m) - 1; i >= 0; i-- {
		ret += string(m[i])
	}
	return ret
}
