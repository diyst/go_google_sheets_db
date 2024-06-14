package sql

type Column struct {
	Name       string
	Type       string
	Size       int
	PK         bool
	FK         bool
	References string
	Nullable   bool
	Check      Check
	Index      string
}
