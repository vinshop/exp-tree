package exp_tree

import (
	"fmt"
	"reflect"
	"strconv"
)

type Variables map[Variable]Value

func (v Variables) Get(key Variable) (Value, error) {
	value, ok := v[key]
	if !ok {
		return nil, ErrVarNotFound(string(key))
	}
	return value, nil
}

func Var(value ...interface{}) Value {
	if len(value) == 1 {
		return varOne(value[0])
	}
	return varOne(value)
}

func varOne(value interface{}) Value {
	t := reflect.TypeOf(value).Kind()
	val := reflect.ValueOf(value)
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
	case reflect.Slice, reflect.Array:
		arr := make(Array, 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			arr = append(arr, Var(val.Index(i).Interface()))
		}
		return arr
	}
	return nil
}
