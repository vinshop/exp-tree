package exp_tree

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Variables map[Variable]Value

func (v Variables) Get(key Variable) (Value, error) {
	path := strings.Split(string(key), ".")
	value, ok := v[Variable(path[0])]
	iValue := value.(interface{})
	if !ok {
		return nil, ErrVarNotFound(path[0])
	}
	var sub []string
	if len(path) > 1 {
		nested, ok := value.(Nested)
		if !ok {
			return nil, ErrNotNested(value)
		}
		iValue = nested.Value()
		sub = path[1:]
	}
	result, err := extract(iValue, sub)
	if err == ErrNilValue {
		return nil, ErrVarNotFound(string(key))
	}
	if _, ok := result.(Value); ok {
		return result.(Value), nil
	}
	return Var(result), nil
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
	case reflect.Struct:
		return Nested{value}
	}
	return nil
}
