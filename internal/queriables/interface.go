package queriables

import "github.com/jackc/pgproto3/v2"

type Queriable interface {
	GetDescription() pgproto3.RowDescription
	Query() (pgproto3.DataRow, error)
}
