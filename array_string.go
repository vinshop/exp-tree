package exp_tree

import (
	"fmt"
)

type ArrayString []string

var arrayStringComputeMap = ComputeMap{
	AnyIn: arrayStringAnyIn,
	AllIn: arrayStringAllIn,
}

var arrayStringAnyIn ComputeFunc = func(values ...NodeValue) NodeValue {
	if len(values) <= 1 {
		return True
	}

	listVariables := values[0].(ArrayString)
	for _, variable := range listVariables {
		for i := 1; i < len(values); i++ {
			if variable == string(values[i].(String)) {
				return True
			}
		}
	}

	return False
}

var arrayStringAllIn ComputeFunc = func(values ...NodeValue) NodeValue {
	if len(values) <= 1 {
		return True
	}

	listVariables := values[0].(ArrayString)
	for _, variable := range listVariables {
		var isIn bool

		for i := 1; i < len(values); i++ {
			if variable == string(values[i].(String)) {
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

var ErrCastArrayString = func(v NodeValue) error { return fmt.Errorf("%v is not array string", v) }

func validateArrayString(values ...NodeValue) error {
	for i, _ := range values {
		if i == 0 {
			_, ok := values[i].(ArrayString)
			if !ok {
				return ErrCastArrayString(values[i])
			}
		} else {
			_, ok := values[i].(String)
			if !ok {
				return ErrCastString(values[i])
			}
		}
	}
	return nil
}

func (b ArrayString) Validate(value ...NodeValue) error {
	return validateArrayString(value...)
}

func (ArrayString) ComputeMap() ComputeMap {
	return arrayStringComputeMap
}

func (b ArrayString) Byte() []byte {
	return []byte(fmt.Sprint(b))
}
