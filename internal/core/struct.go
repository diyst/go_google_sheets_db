package core

import (
	"go_google_sheets_db/internal/queriables"
)

var tables []*queriables.Table

func SetTables(gsTables []*queriables.Table) {
	gsTables = tables
}
