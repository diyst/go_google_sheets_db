package sql

import "reflect"

type column struct {
	Name       string
	Type       string
	Size       int
	PK         bool
	FK         bool
	References string
	Nullable   bool
	Check      Check
}

func validateColumnType(columnType string) bool {
	postgresTypes := map[string]string{
		"bigint":                      "bigint",
		"bigserial":                   "bigserial",
		"bit":                         "bit",
		"varbit":                      "bit varying",
		"bit varying":                 "bit varying",
		"boolean":                     "boolean",
		"bool":                        "boolean",
		"box":                         "box",
		"bytea":                       "bytea",
		"char":                        "character",
		"character":                   "character",
		"varchar":                     "character varying",
		"character varying":           "character varying",
		"cidr":                        "cidr",
		"circle":                      "circle",
		"date":                        "date",
		"float8":                      "double precision",
		"double precision":            "double precision",
		"inet":                        "inet",
		"int":                         "integer",
		"integer":                     "integer",
		"interval":                    "interval",
		"json":                        "json",
		"jsonb":                       "jsonb",
		"line":                        "line",
		"lseg":                        "lseg",
		"macaddr":                     "macaddr",
		"macaddr8":                    "macaddr8",
		"money":                       "money",
		"decimal":                     "numeric",
		"numeric":                     "numeric",
		"path":                        "path",
		"pg_lsn":                      "pg_lsn",
		"point":                       "point",
		"polygon":                     "polygon",
		"float4":                      "real",
		"real":                        "real",
		"smallint":                    "smallint",
		"int2":                        "smallint",
		"serial":                      "serial",
		"serial4":                     "serial",
		"text":                        "text",
		"time":                        "time",
		"time without time zone":      "time",
		"timetz":                      "time with time zone",
		"time with time zone":         "time with time zone",
		"timestamp":                   "timestamp",
		"timestamp without time zone": "timestamp",
		"timestamptz":                 "timestamp with time zone",
		"timestamp with time zone":    "timestamp with time zone",
		"tsquery":                     "tsquery",
		"tsvector":                    "tsvector",
		"txid_snapshot":               "txid_snapshot",
		"uuid":                        "uuid",
		"xml":                         "xml",
	}

	_, isValid := postgresTypes[columnType]

	return isValid
}

func ParsePostgresTypeToGoType(columnType string) reflect.Type {

	postgresTypes := map[string]reflect.Type{
		"bigint":                      reflect.TypeOf(int64(0)),
		"bigserial":                   reflect.TypeOf(int64(0)),
		"bit":                         reflect.TypeOf(int8(0)),
		"varbit":                      reflect.TypeOf(int8(0)),
		"boolean":                     reflect.TypeOf(false),
		"bool":                        reflect.TypeOf(false),
		"box":                         reflect.TypeOf(struct{}{}),
		"bytea":                       reflect.TypeOf([]byte{}),
		"char":                        reflect.TypeOf(""),
		"character":                   reflect.TypeOf(""),
		"varchar":                     reflect.TypeOf(""),
		"character varying":           reflect.TypeOf(""),
		"cidr":                        reflect.TypeOf(""),
		"circle":                      reflect.TypeOf(struct{}{}),
		"date":                        reflect.TypeOf(""),
		"float8":                      reflect.TypeOf(float64(0)),
		"double precision":            reflect.TypeOf(float64(0)),
		"inet":                        reflect.TypeOf(""),
		"int":                         reflect.TypeOf(int(0)),
		"integer":                     reflect.TypeOf(int(0)),
		"interval":                    reflect.TypeOf(struct{}{}),
		"json":                        reflect.TypeOf(struct{}{}),
		"jsonb":                       reflect.TypeOf(struct{}{}),
		"line":                        reflect.TypeOf(struct{}{}),
		"lseg":                        reflect.TypeOf(struct{}{}),
		"macaddr":                     reflect.TypeOf(""),
		"macaddr8":                    reflect.TypeOf(""),
		"money":                       reflect.TypeOf(float64(0)),
		"decimal":                     reflect.TypeOf(float64(0)),
		"numeric":                     reflect.TypeOf(float64(0)),
		"path":                        reflect.TypeOf(struct{}{}),
		"pg_lsn":                      reflect.TypeOf(""),
		"point":                       reflect.TypeOf(struct{}{}),
		"polygon":                     reflect.TypeOf(struct{}{}),
		"float4":                      reflect.TypeOf(float32(0)),
		"real":                        reflect.TypeOf(float32(0)),
		"smallint":                    reflect.TypeOf(int16(0)),
		"int2":                        reflect.TypeOf(int16(0)),
		"serial":                      reflect.TypeOf(int64(0)),
		"serial4":                     reflect.TypeOf(int64(0)),
		"text":                        reflect.TypeOf(""),
		"time":                        reflect.TypeOf(""),
		"time without time zone":      reflect.TypeOf(""),
		"timetz":                      reflect.TypeOf(struct{}{}),
		"time with time zone":         reflect.TypeOf(struct{}{}),
		"timestamp":                   reflect.TypeOf(""),
		"timestamp without time zone": reflect.TypeOf(""),
		"timestamptz":                 reflect.TypeOf(struct{}{}),
		"timestamp with time zone":    reflect.TypeOf(struct{}{}),
		"tsquery":                     reflect.TypeOf(struct{}{}),
		"tsvector":                    reflect.TypeOf(struct{}{}),
		"txid_snapshot":               reflect.TypeOf(""),
		"uuid":                        reflect.TypeOf(""),
		"xml":                         reflect.TypeOf(struct{}{}),
	}

	return postgresTypes[columnType]
}
