package sql

import (
	"fmt"
	"reflect"
	"strconv"
)

type Check struct {
	Condition string
	Threshold string
	Type      reflect.Type
}

func NewCheck(condition string, threshold string, T any) Check {
	return Check{
		Condition: condition,
		Threshold: threshold,
		Type:      reflect.TypeOf(T),
	}
}

var stringType = reflect.TypeOf("")
var intType = reflect.TypeOf(int(0))
var int64Type = reflect.TypeOf(int64(0))
var float64Type = reflect.TypeOf(float64(0))

func (c Check) isValid(value any) bool {

	stringValue := fmt.Sprintf("%v", value)
	stringThreshold := fmt.Sprintf("%v", c.Threshold)

	intValue, _ := strconv.Atoi(stringValue)
	intThreshold, _ := strconv.Atoi(stringThreshold)

	floatValue, _ := strconv.ParseFloat(stringValue, 64)
	floatThreshold, _ := strconv.ParseFloat(stringThreshold, 64)

	switch reflect.TypeOf(c.Threshold) {
	case stringType:
		return validString(stringValue, stringThreshold, c.Condition)
	case intType:
	case int64Type:
		return validNumber(int64(intValue), int64(intThreshold), c.Condition)
	case float64Type:
		return validFloat(floatValue, floatThreshold, c.Condition)
	}

	return false
}

func validString(value string, threshold string, condition string) bool {
	switch condition {
	case "<":
		return value < threshold
	case "<=":
		return value <= threshold
	case ">":
		return value > threshold
	case ">=":
		return value >= threshold
	case "=":
		return value == threshold
	case "<>":
		return value != threshold
	default:
		return false
	}
}

func validFloat(value float64, threshold float64, condition string) bool {
	switch condition {
	case "<":
		return value < threshold
	case "<=":
		return value <= threshold
	case ">":
		return value > threshold
	case ">=":
		return value >= threshold
	case "=":
		return value == threshold
	case "<>":
		return value != threshold
	default:
		return false
	}
}

func validNumber(value int64, threshold int64, condition string) bool {
	switch condition {
	case "<":
		return value < threshold
	case "<=":
		return value <= threshold
	case ">":
		return value > threshold
	case ">=":
		return value >= threshold
	case "=":
		return value == threshold
	case "<>":
		return value != threshold
	default:
		return false
	}
}
