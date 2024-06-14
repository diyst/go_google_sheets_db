package core

import (
	"go_google_sheets_db/internal/sql"
	"reflect"
	"testing"

	"google.golang.org/api/sheets/v4"
)


func TestGetTableColumns(t *testing.T) {
	// Test case: empty sheet
	sheet := &sheets.Sheet{
		Data: []*sheets.GridData{},
	}
	expected := []sql.Column{}
	actual := getTableColumns(sheet)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	sheet = &sheets.Sheet{
		Data: convertRowDataToGridData([]*sheets.RowData{
			{
				Values: []*sheets.CellData{
					{FormattedValue: "column1"},
				},
			},
		}),
	}
	expected = []sql.Column{
		{Name: "column1", Index: "A"},
	}
	actual = getTableColumns(sheet)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test case: one row with multiple columns
	sheet = &sheets.Sheet{
		Data: convertRowDataToGridData([]*sheets.RowData{
			{
				Values: []*sheets.CellData{
					{FormattedValue: "column1"},
					{FormattedValue: "column2"},
					{FormattedValue: "column3"},
				},
			},
		}),
	}
	expected = []sql.Column{
		{Name: "column1", Index: "A"},
		{Name: "column2", Index: "B"},
		{Name: "column3", Index: "C"},
	}
	actual = getTableColumns(sheet)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	// Test case: multiple rows with multiple columns
	sheet = &sheets.Sheet{
		Data: convertRowDataToGridData([]*sheets.RowData{
			{
				Values: []*sheets.CellData{
					{FormattedValue: "column1"},
					{FormattedValue: "column2"},
				},
			},
			{
				Values: []*sheets.CellData{
					{FormattedValue: "column3"},
					{FormattedValue: "column4"},
					{FormattedValue: "column5"},
				},
			},
		}),
	}
	expected = []sql.Column{
		{Name: "column1", Index: "A"},
		{Name: "column2", Index: "B"},
	}
	actual = getTableColumns(sheet)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func convertRowDataToGridData(rowData []*sheets.RowData) []*sheets.GridData {
	gridData := make([]*sheets.GridData, len(rowData))
	for i, row := range rowData {
		gridData[i] = &sheets.GridData{RowData: []*sheets.RowData{row}}
	}
	return gridData
}
