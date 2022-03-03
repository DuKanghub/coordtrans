package save

import (
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

func Save2Excel(fileName string, data []string) error {
	xlsx := excelize.NewFile()
	// 设置表头，即列标题
	xlsx.SetCellValue("Sheet1", "A1", "序号")
	xlsx.SetCellValue("Sheet1", "B1", "原经度")
	xlsx.SetCellValue("Sheet1", "C1", "原纬度")
	xlsx.SetCellValue("Sheet1", "D1", "新经度")
	xlsx.SetCellValue("Sheet1", "E1", "新纬度")
	for i, row := range data {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), strconv.Itoa(i+1))
		arr := strings.Split(row, ",")
		if len(arr) == 4 {
			xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), arr[0])
			xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), arr[1])
			xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), arr[2])
			xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), arr[3])
		}
	}
	return xlsx.SaveAs(fileName)
}

func ReadExcel(fileName string) ([]string, error) {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}
	var data []string
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) == 2 {
			data = append(data, row[0]+","+row[1])
		}
	}
	return data, nil
}
