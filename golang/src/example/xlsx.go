package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "F:/吴乐/cindy/提取pdf文本/提取备注/过滤后.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {

	}
	count := 0
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				if count < 1000 {
					count++
					text, _ := cell.String()
					fmt.Printf("%s\n", text)
				}
			}
		}
	}
}
