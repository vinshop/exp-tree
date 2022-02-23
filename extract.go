package exp_tree

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrNilValue = errors.New("nil value")

func extract(s interface{}, path []string) (interface{}, error) {
	if s == nil {
		return nil, ErrNilValue
	}
	if len(path) == 0 {
		return s, nil
	}
	ref := reflect.ValueOf(s)
	now := ref.FieldByName(path[0])
	if !now.IsValid() {
		return nil, fmt.Errorf("field %v does not exist", path[0])
	}
	if len(path) > 1 {
		return extract(now.Interface(), path[1:])
	}
	return now.Interface(), nil
}
