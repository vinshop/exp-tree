package exp_tree

import "fmt"

//Float64 type
type Float64 float64

var ErrCastFloat64 = func(v NodeValue) error { return fmt.Errorf("%v is not float64ean", v) }

var float64ComputeMap = ComputeMap{
	None: float64Sum,
	Eq:   float64Eq,
	Sum:  float64Sum,
	Mul:  float64Mul,
	Gt:   float64Gt,
	Gte:  float64Gte,
	Lt:   float64Lt,
	Lte:  float64Lte,
	In:   float64In,
}

var float64Eq ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Float64) != values[i].(Float64) {
			return False
		}
	}
	return True
}

var float64Sum ComputeFunc = func(values ...NodeValue) NodeValue {
	var res Float64
	for _, v := range values {
		res += v.(Float64)
	}
	return res
}

var float64Mul ComputeFunc = func(values ...NodeValue) NodeValue {
	res := Float64(1)
	for _, v := range values {
		res *= v.(Float64)
	}
	return res
}

var float64Gt ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Float64) <= values[i].(Float64) {
			return False
		}
	}
	return True
}

var float64Gte ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Float64) < values[i].(Float64) {
			return False
		}
	}
	return True
}

var float64Lt ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Float64) >= values[i].(Float64) {
			return False
		}
	}
	return True
}

var float64Lte ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Float64) > values[i].(Float64) {
			return False
		}
	}
	return True
}

var float64In ComputeFunc = func(values ...NodeValue) NodeValue {
	v := values[0].(Float64)
	if len(values) == 1 {
		return True
	}
	for i := 1; i < len(values); i++ {
		if values[i].(Float64) == v {
			return True
		}
	}
	return False
}

func validateFloat64(values ...NodeValue) error {
	for _, v := range values {
		_, ok := v.(Float64)
		if !ok {
			return ErrCastFloat64(v)
		}
	}
	return nil
}

func (b Float64) Validate(value ...NodeValue) error {
	return validateFloat64(value...)
}

func (Float64) ComputeMap() ComputeMap {
	return float64ComputeMap
}

func (b Float64) Byte() []byte {
	return []byte(fmt.Sprint(b))
}
