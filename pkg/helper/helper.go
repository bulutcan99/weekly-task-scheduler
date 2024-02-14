package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strconv"
)

func MinValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ConvertToFloat64(value any) float64 {
	if value == nil {
		return 0
	}
	typeOf := reflect.TypeOf(value).String()

	switch typeOf {
	case "float64":
		return value.(float64)
	case "float32":
		return float64(value.(float32))
	case "int":
		return float64(value.(int))
	case "int32":
		return float64(value.(int32))
	case "int64":
		return float64(value.(int64))
	case "string":
		f, _ := strconv.ParseFloat(value.(string), 64)
		return f
	case "primitive.Decimal128":
		dec := value.(primitive.Decimal128)
		str := dec.String()
		f, _ := strconv.ParseFloat(str, 64)
		return f
	default:
		return 0
	}
}
func ConvertToInt(value interface{}) int {
	if value == nil {
		return 0
	}

	typeOf := reflect.TypeOf(value).String()
	switch typeOf {
	case "float64":
		return int(value.(float64))
	case "float32":
		return int(value.(float32))
	case "int":
		return value.(int)
	case "int32":
		return int(value.(int32))
	case "int64":
		return int(value.(int64))
	case "string":
		i, _ := strconv.Atoi(value.(string))
		return i
	case "primitive.Decimal128":
		dec := value.(primitive.Decimal128)
		str := dec.String()
		i, _ := strconv.Atoi(str)
		return i
	default:
		return 0
	}
}
