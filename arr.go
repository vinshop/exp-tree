package exp_tree

import "fmt"

type Array []Value

var ToArray = func(values ...Value) Value {
	return Array(values)
}

func (a Array) toMap() map[Value]int {
	mp := make(map[Value]int)
	for _, v := range a {
		mp[v]++
	}
	return mp
}

var ErrNotArray = func(v interface{}) error {
	return fmt.Errorf("%v is not an array", v)
}

var isArray ValidateFunc = func(values Value) error {
	if _, ok := values.(Array); !ok {
		return ErrNotArray(values)
	}
	return nil
}

func (a Array) validate(op Operator, value ...Value) error {
	panic("implement me")
}

func (a Array) t() NodeType {
	return NValue
}

func (a Array) f(op Operator) *Math {
	return a[0].f(op)
}
