package core

import (
	"fmt"
	"go_google_sheets_db/internal/gsheets"
	"go_google_sheets_db/internal/queriables"
	"go_google_sheets_db/internal/sql"

	"google.golang.org/api/sheets/v4"
)

func Setup() {
	resp, err := gsheets.Client.Service.Spreadsheets.Get(gsheets.Client.SheetID).IncludeGridData(true).Do()

	if err != nil {
		panic(err)
	}

	setupTables(resp.Sheets)
}

func setupTables(tables []*sheets.Sheet) {

	var gsTables = []queriables.Table{}

	for _, table := range tables {
		gTable := queriables.Table{
			Name:    table.Properties.Title,
			Columns: getTableColumns(table),
		}

		gsTables = append(gsTables, gTable)
	}

}

func getTableColumns(sheet *sheets.Sheet) []sql.Column {

	if len(sheet.Data) == 0 {
		return []sql.Column{}
	}

	var columns []sql.Column

	for _, row := range sheet.Data[0].RowData {
		for i, cell := range row.Values {
			fmt.Println(cell.FormattedValue, " ")

			columns = append(columns, sql.Column{
				Name:  cell.FormattedValue,
				Index: numberToCOlumnIndex(i),
			})
		}
	}

	return columns
}

func numberToCOlumnIndex(i int) string {
	var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	return letters[i]
}
