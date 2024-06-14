package parser

import (
	"fmt"
	"go_google_sheets_db/internal/queriables"
	"go_google_sheets_db/internal/sql"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Parse(statement string) (queriables.Queriable, error) {
	statement = strings.ToLower(statement)

	if strings.HasPrefix(statement, "create") {
		statement = strings.TrimLeft(strings.Replace(statement, "create", "", 1), " ")
		if strings.HasPrefix(statement, "table") {

			statement = strings.TrimLeft(strings.Replace(statement, "table", "", 1), " ")

			tableName := strings.Split(statement, " ")[0]

			statement = strings.Replace(statement, tableName, "", 1)

			statement = strings.TrimSuffix(statement, ");")

			statement = strings.TrimLeft(statement, " ")

			statement = strings.TrimRight(statement, " ")

			statement = strings.TrimPrefix(statement, "(")

			columns := strings.Split(statement, ",")

			tableColumns, err := parseColumns(columns)

			if err != nil {
				return nil, err
			}

			return queriables.Table{
				Name:    tableName,
				Columns: tableColumns,
			}, nil
		}

		return nil, nil
	}

	if strings.HasPrefix(statement, "select") {

		if QueryIsUsingPostgresFunc(statement) {
			return nil, nil
		}

		return nil, nil
	}

	return nil, nil
}

func parseColumns(columns []string) ([]sql.Column, error) {
	var columnsParsed []sql.Column

	var onlyNumberRegex = regexp.MustCompile("^[0-9]+$")

	for _, columnStr := range columns {
		columnStr = strings.ToLower(strings.TrimSpace(columnStr))
		parts := strings.FieldsFunc(columnStr, func(r rune) bool {
			return r == ' ' || r == '(' || r == ')'
		})

		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid column format: %s", columnStr)
		}

		name := parts[0]

		columnType := parts[1]

		if !validateColumnType(columnType) {
			return nil, fmt.Errorf("invalid column type: %s", columnType)
		}

		newColumn := sql.Column{
			Name:     name,
			Type:     columnType,
			Nullable: true,
		}

		if len(parts) == 2 {
			columnsParsed = append(columnsParsed, newColumn)
			continue
		}

		currentIndex := 2

		match := onlyNumberRegex.MatchString(parts[currentIndex])
		if match {
			size, err := strconv.Atoi(parts[2])
			if err != nil {
				log.Printf("Invalid column size: %s", parts[2])
				return nil, fmt.Errorf("invalid column size: %s", parts[2])
			}

			newColumn.Size = size
			currentIndex++

			if len(parts) == currentIndex {
				columnsParsed = append(columnsParsed, newColumn)
				continue
			}
		}

		if strings.Contains(columnStr, "primary key") {
			newColumn.PK = true
		}

		if strings.Contains(columnStr, "not null") {
			newColumn.Nullable = false
		} else if strings.Contains(columnStr, "null") {
			newColumn.Nullable = true
		} else {
			newColumn.Nullable = true
		}

		if strings.Contains(columnStr, "check") {
			newColumn.Check = createCheck(columnStr, newColumn.Type)
		}

		columnsParsed = append(columnsParsed, newColumn)

	}

	log.Printf("Parsed %d columns", len(columnsParsed))
	return columnsParsed, nil
}

func createCheck(columnStr string, columnType string) sql.Check {
	check := strings.Split(strings.Split(columnStr, "check")[1], ")")[0]

	check = strings.Replace(check, "(", "", 1)

	check = strings.TrimLeft(check, " ")

	checkDef := strings.Split(check, " ")

	checkOperator := checkDef[1]

	checkThreshold := checkDef[2]

	checkInstance := sql.NewCheck(checkOperator, checkThreshold, ParsePostgresTypeToGoType(columnType))

	return checkInstance
}
