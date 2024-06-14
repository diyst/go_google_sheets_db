package parser

import "strings"

var pgFuncResp = map[string]string{
	"PG_TRY_ADVISORY_LOCK": "t",
	"PG_ADVISORY_UNLOCK":   "f",
}

func QueryIsUsingPostgresFunc(statement string) bool {
	for key := range pgFuncResp {
		if strings.Contains(statement, key) {
			return true
		}
	}
	return false
}