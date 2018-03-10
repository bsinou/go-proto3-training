package xlsx

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func AddHeaders(sheet *xlsx.Sheet, headerLabels []string) {
	var row *xlsx.Row
	var cell *xlsx.Cell
	headerStyle := createHeaderStyle()

	row = sheet.AddRow()
	for _, label := range headerLabels {
		cell = row.AddCell()
		cell.Value = label
		cell.SetStyle(headerStyle)
	}
}

func AddBodyRow(sheet *xlsx.Sheet, values []interface{}) {
	var row *xlsx.Row
	var cell *xlsx.Cell

	bodyStyle := createBodyStyle()
	row = sheet.AddRow()
	for _, label := range values {
		cell = row.AddCell()
		cell.SetStyle(bodyStyle)

		switch label := label.(type) {
		case string:
			cell.Value = label
		default:
			fmt.Println("Unexpected type...")
		}
	}
}

func CreateSingleSheetXlsxFile(sheetName string) (*xlsx.Sheet, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return nil, err
	}
	return sheet, nil
}

func createHeaderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	font := *xlsx.NewFont(14, "Verdana")
	font.Bold = true
	font.Color = "FFFFFFFF"
	style.Font = font
	// fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
	fill := *xlsx.NewFill("solid", "00134E6C", "00FFFFFF")
	style.Fill = fill
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Border = border
	style.ApplyBorder = true
	style.ApplyFill = true
	return style
}

func createBodyStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	font := *xlsx.NewFont(12, "Verdana")
	style.Font = font
	border := *xlsx.NewBorder("thin", "thin", "none", "none")
	style.Border = border
	style.ApplyBorder = true
	return style
}
