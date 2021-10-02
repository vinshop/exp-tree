package exp_tree

import (
	"fmt"
)

type String string

var ErrCastString = func(v NodeValue) error { return fmt.Errorf("%v is not stringe", v) }

var stringComputeMap = ComputeMap{
	In:  stringIn,
	Eq:  stringEq,
	Gt:  stringGt,
	Gte: stringGte,
	Lt:  stringLt,
	Lte: stringLte,
}

var stringIn ComputeFunc = func(values ...NodeValue) NodeValue {
	s := values[0].(String)
	if len(values) == 1 {
		return True
	}
	for i := 1; i < len(values); i++ {
		if s == values[i].(String) {
			return True
		}
	}
	return False
}

var stringEq ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i].(String) != values[i-1].(String) {
			return False
		}
	}
	return True
}

var stringGt ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(String) <= values[i].(String) {
			return False
		}
	}
	return True
}

var stringGte ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(String) < values[i].(String) {
			return False
		}
	}
	return True
}

var stringLt ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(String) >= values[i].(String) {
			return False
		}
	}
	return True
}

var stringLte ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(String) > values[i].(String) {
			return False
		}
	}
	return True
}

func validateString(values ...NodeValue) error {
	for _, v := range values {
		_, ok := v.(String)
		if !ok {
			return ErrCastString(v)
		}
	}
	return nil
}

func (b String) Validate(value ...NodeValue) error {
	return validateString(value...)
}

func (String) ComputeMap() ComputeMap {
	return stringComputeMap
}

func (b String) Byte() []byte {
	return []byte(fmt.Sprintf(`"%v"`, b))
}
