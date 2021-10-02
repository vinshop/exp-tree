package exp_tree

import "fmt"

//Bool type
type Bool bool

const (
	True  Bool = true
	False Bool = false
)

var ErrCastBool = func(v NodeValue) error { return fmt.Errorf("%v is not boolean", v) }

var boolComputeMap = ComputeMap{
	None: boolAnd,
	And:  boolAnd,
	Or:   boolOr,
	Not:  boolNot,
	Eq:   boolEq,
	Xor:  boolXor,
}

var boolAnd ComputeFunc = func(values ...NodeValue) NodeValue {
	for _, v := range values {
		vBool := v.(Bool)
		if !vBool {
			return False
		}
	}
	return True
}

var boolOr ComputeFunc = func(values ...NodeValue) NodeValue {
	for _, v := range values {
		vBool := v.(Bool)
		if vBool {
			return True
		}
	}
	return False
}

var boolNot ComputeFunc = func(values ...NodeValue) NodeValue {
	res := boolAnd(values...)
	return !res.(Bool)
}

var boolEq ComputeFunc = func(values ...NodeValue) NodeValue {
	for i := 1; i < len(values); i++ {
		if values[i-1].(Bool) != values[i].(Bool) {
			return False
		}
	}
	return True
}

var boolXor ComputeFunc = func(values ...NodeValue) NodeValue {
	res := False
	for i := 1; i < len(values); i++ {
		res = res != values[i].(Bool)
	}
	return res
}

func validateBool(values ...NodeValue) error {
	for _, v := range values {
		_, ok := v.(Bool)
		if !ok {
			return ErrCastBool(v)
		}
	}
	return nil
}

func (b Bool) Validate(value ...NodeValue) error {
	return validateBool(value...)
}

func (Bool) ComputeMap() ComputeMap {
	return boolComputeMap
}

func (b Bool) Byte() []byte {
	return []byte(fmt.Sprint(b))
}
