package exp_tree

import (
	"fmt"
)

type ArrayFloat64 []float64

var arrayFloat64ComputeMap = ComputeMap{
	AnyIn: arrayFloat64AnyIn,
	AllIn: arrayFloat64AllIn,
}

var arrayFloat64AnyIn ComputeFunc = func(values ...NodeValue) NodeValue {
	if len(values) <= 1 {
		return True
	}

	listVariables := values[0].(ArrayFloat64)
	for _, variable := range listVariables {
		for i := 1; i < len(values); i++ {
			if variable == float64(values[i].(Float64)) {
				return True
			}
		}
	}

	return False
}

var arrayFloat64AllIn ComputeFunc = func(values ...NodeValue) NodeValue {
	if len(values) <= 1 {
		return True
	}

	listVariables := values[0].(ArrayFloat64)
	for _, variable := range listVariables {
		var isIn bool

		for i := 1; i < len(values); i++ {
			if variable == float64(values[i].(Float64)) {
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

var ErrCastArrayFloat64 = func(v NodeValue) error { return fmt.Errorf("%v is not array float64", v) }

func validateArrayFloat64(values ...NodeValue) error {
	for i, _ := range values {
		if i == 0 {
			_, ok := values[i].(ArrayFloat64)
			if !ok {
				return ErrCastArrayFloat64(values[i])
			}
		} else {
			_, ok := values[i].(Float64)
			if !ok {
				return ErrCastString(values[i])
			}
		}
	}
	return nil
}

func (b ArrayFloat64) Validate(value ...NodeValue) error {
	return validateArrayFloat64(value...)
}

func (ArrayFloat64) ComputeMap() ComputeMap {
	return arrayFloat64ComputeMap
}

func (b ArrayFloat64) Byte() []byte {
	return []byte(fmt.Sprint(b))
}
