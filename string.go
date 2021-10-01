package exp_tree

import "fmt"

type String string

var ErrCastString = func(v NodeValue) error { return fmt.Errorf("%v is not stringean", v) }

var stringComputeMap = ComputeMap{}

func validateString(values ...NodeValue) error {
	for _, v := range values {
		_, ok := v.(String)
		if !ok {
			return ErrCastInt64(v)
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
