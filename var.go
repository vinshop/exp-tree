package exp_tree

import (
	"fmt"
	"reflect"
	"strconv"
)

func Var(value interface{}) Value {
	t := reflect.TypeOf(value).Kind()
	switch t {
	case reflect.String:
		return String(value.(string))
	case reflect.Bool:
		return Bool(value.(bool))
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		v, _ := strconv.ParseFloat(fmt.Sprint(value), 64)
		return Number(v)
	}
	return nil
}
