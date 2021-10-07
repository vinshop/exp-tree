package exp_tree

import (
	"fmt"
)

type ArrayInt64 []int64

var arrayInt64ComputeMap = ComputeMap{
	AnyIn: arrayInt64AnyIn,
	AllIn: arrayInt64AllIn,
}

var arrayInt64AnyIn ComputeFunc = func(values ...NodeValue) NodeValue {
	if len(values) <= 1 {
		return True
	}

	listVariables := values[0].(ArrayInt64)
	for _, variable := range listVariables {
		for i := 1; i < len(values); i++ {
			if variable == int64(values[i].(Int64)) {
				return True
			}
		}
	}

	return False
}

var arrayInt64AllIn ComputeFunc = func(values ...NodeValue) NodeValue {
	if len(values) <= 1 {
		return True
	}

	listVariables := values[0].(ArrayInt64)
	for _, variable := range listVariables {
		var isIn bool

		for i := 1; i < len(values); i++ {
			if variable == int64(values[i].(Int64)) {
				isIn = true
				break
			}
		}

		if !isIn {
			return False
		}
	}

	return True
}

var ErrCastArrayInt64 = func(v NodeValue) error { return fmt.Errorf("%v is not array int64", v) }

func validateArrayInt64(values ...NodeValue) error {
	for i, _ := range values {
		if i == 0 {
			_, ok := values[i].(ArrayInt64)
			if !ok {
				return ErrCastArrayInt64(values[i])
			}
		} else {
			_, ok := values[i].(Int64)
			if !ok {
				return ErrCastString(values[i])
			}
		}
	}
	return nil
}

func (b ArrayInt64) Validate(value ...NodeValue) error {
	return validateArrayInt64(value...)
}

func (ArrayInt64) ComputeMap() ComputeMap {
	return arrayInt64ComputeMap
}

func (b ArrayInt64) Byte() []byte {
	return []byte(fmt.Sprint(b))
}
