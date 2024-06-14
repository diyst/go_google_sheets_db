package queriables

import (
	"fmt"

	"github.com/jackc/pgproto3/v2"
	"google.golang.org/api/sheets/v4"

	"go_google_sheets_db/internal/gsheets"
	"go_google_sheets_db/internal/sql"
)

type Table struct {
	Name    string
	Columns []sql.Column
}

func NewTable(name string, columns []sql.Column) Table {
	return Table{
		Name:    name,
		Columns: columns,
	}
}

func (t Table) Create() (pgproto3.RowDescription, error) {
	spreadsheet, err := gsheets.Client.Service.Spreadsheets.Get(gsheets.Client.SheetID).Do()

	if err != nil {
		return pgproto3.RowDescription{}, err
	}

	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == t.Name {
			return pgproto3.RowDescription{}, fmt.Errorf("Sheet %s already exists", t.Name)
		}
	}

	sheet := &sheets.Sheet{
		Properties: &sheets.SheetProperties{
			Title: t.Name,
		},
	}

	newSheet := &sheets.AddSheetRequest{
		Properties: sheet.Properties,
	}

	batchUpdateRequest := sheets.Request{
		AddSheet: newSheet,
	}

	resp, err := gsheets.Client.BatchUpdate(batchUpdateRequest)

	if err != nil {
		return pgproto3.RowDescription{}, err
	}

	if len(resp.UpdatedSpreadsheet.Sheets) == 0 {
		return pgproto3.RowDescription{}, fmt.Errorf("Failed to create sheet %s", t.Name)
	}

	updateRequest := &sheets.UpdateCellsRequest{
		Rows: []*sheets.RowData{
			{Values: []*sheets.CellData{}},
		},
		Fields: "userEnteredValue",
		Start: &sheets.GridCoordinate{
			SheetId:     resp.UpdatedSpreadsheet.Sheets[0].Properties.SheetId,
			RowIndex:    0,
			ColumnIndex: 0,
		},
	}

	for _, column := range t.Columns {
		updateRequest.Rows[0].Values = append(updateRequest.Rows[0].Values, &sheets.CellData{
			UserEnteredValue: &sheets.ExtendedValue{
				StringValue: &column.Name,
			},
		})
	}

	updateRequestBatch := sheets.Request{
		UpdateCells: updateRequest,
	}

	resp, err = gsheets.Client.BatchUpdate(updateRequestBatch)
	if err != nil {
		return pgproto3.RowDescription{}, err
	}

	if len(resp.UpdatedSpreadsheet.Sheets) == 0 {
		return pgproto3.RowDescription{}, fmt.Errorf("Failed to create sheet %s", t.Name)
	}

	return pgproto3.RowDescription{
		Fields: t.columnToFields(t.Columns),
	}, nil

}

func (t Table) GetDescription() pgproto3.RowDescription {

	return pgproto3.RowDescription{
		Fields: t.columnToFields(t.Columns),
	}
}
func (t Table) columnToFields(columns []sql.Column) []pgproto3.FieldDescription {
	fields := []pgproto3.FieldDescription{}

	for _, column := range columns {
		fields = append(fields, pgproto3.FieldDescription{
			Name: []byte(column.Name),
		})
	}

	return fields
}

func (t Table) Query() (pgproto3.DataRow, error) {
	return pgproto3.DataRow{}, nil
}
