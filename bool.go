package exp_tree

import "fmt"

type Bool bool

var True = Bool(true)
var False = Bool(false)

var ErrNotBool = func(v interface{}) error {
	return fmt.Errorf("%v is not boolean", v)
}

var isBools ValidateFunc = func(value Value) error {
	if err := isArray(value); err != nil {
		return err
	}
	for _, v := range value.(Array) {
		if err := isBool(v); err != nil {
			return err
		}
	}
	return nil
}

var isBool ValidateFunc = func(value Value) error {
	if _, ok := value.(Bool); !ok {
		return ErrNotBool(value)
	}
	return nil
}

func (b Bool) t() NodeType {
	return NValue
}

var boolAnd = isBools.With(func(values Value) Value {
	for _, v := range values.(Array) {
		if v == False {
			return False
		}
	}
	return True
})

var boolOr = isBools.With(func(values Value) Value {
	for _, v := range values.(Array) {
		if v == True {
			return True
		}
	}
	return False
})

var boolMp = map[Operator]*Math{
	None: Keep,
	And:  boolAnd,
	Or:   boolOr,
}

func (b Bool) f(op Operator) *Math {
	return boolMp[op]
}
