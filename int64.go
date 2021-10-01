package exp_tree

import "fmt"

type Int64 int64

var ErrCastInt64 = func(v NodeValue) error { return fmt.Errorf("%v is not int64ean", v) }

var int64ComputeMap = ComputeMap{
	None: int64And,
	And:  int64And,
	Or:   int64Or,
	Eq:   int64Eq,
	Xor:  int64Xor,
	Sum:  int64Sum,
	Mul:  int64Mul,
	Gt:   int64Gt,
	Gte:  int64Gte,
	Lt:   int64Lt,
	Lte:  int64Lte,
}

var int64And ComputeFunc = func(values ...NodeValue) NodeValue {
	res := values[0].(Int64)
	for i := 1; i < len(values); i++ {
		res &= values[i].(Int64)
	}
	return res
}

var int64Or ComputeFunc = func(values ...NodeValue) NodeValue {
	var res Int64
	for _, v := range values {
		res &= v.(Int64)
	}
	return res
}

var int64Eq ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Int64) != values[i].(Int64) {
			return False
		}
	}
	return True
}

var int64Xor ComputeFunc = func(values ...NodeValue) NodeValue {
	var res Int64
	for _, v := range values {
		res ^= v.(Int64)
	}
	return res
}

var int64Sum ComputeFunc = func(values ...NodeValue) NodeValue {
	var res Int64
	for _, v := range values {
		res += v.(Int64)
	}
	return res
}

var int64Mul ComputeFunc = func(values ...NodeValue) NodeValue {
	res := Int64(1)
	for _, v := range values {
		res *= v.(Int64)
	}
	return res
}

var int64Gt ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Int64) <= values[i].(Int64) {
			return False
		}
	}
	return True
}

var int64Gte ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Int64) < values[i].(Int64) {
			return False
		}
	}
	return True
}

var int64Lt ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Int64) >= values[i].(Int64) {
			return False
		}
	}
	return True
}

var int64Lte ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Int64) > values[i].(Int64) {
			return False
		}
	}
	return True
}

func validateInt64(values ...NodeValue) error {
	for _, v := range values {
		_, ok := v.(Int64)
		if !ok {
			return ErrCastInt64(v)
		}
	}
	return nil
}

func (b Int64) Validate(value ...NodeValue) error {
	return validateInt64(value...)
}

func (Int64) ComputeMap() ComputeMap {
	return int64ComputeMap
}

func (b Int64) Byte() []byte {
	return []byte(fmt.Sprint(b))
}
