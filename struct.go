package exp_tree

import (
	"reflect"
	"strings"
)

func Extract(v interface{}, query string) Value {
	return Var(extract(v, query))
}

func extract(v interface{}, key string) interface{} {
	if key == "" {
		return v
	}
	keys := strings.Split(key, ".")

	for _, key = range keys {
		t := reflect.ValueOf(v)
		switch t.Kind() {
		case reflect.Struct:
			str, ok := t.Type().FieldByName(key)
			if !ok {
				return nil
			}
			v = t.FieldByIndex(str.Index).Interface()
		case reflect.Map:
			val := t.MapIndex(reflect.ValueOf(key))
			if val.Kind() == reflect.Invalid {
				return nil
			}
			v = val.Interface()
		default:
			return nil
		}
	}
	return v
}
