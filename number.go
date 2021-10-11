package exp_tree

import "fmt"

type Number float64

var ErrNotNumber = func(val interface{}) error {
	return fmt.Errorf("%v is not a number", val)
}

var ErrNotNumbers = func(val interface{}) error {
	return fmt.Errorf("%v is not a array of number", val)
}

var isNumbers ValidateFunc = func(value Value) error {
	if err := isArray(value); err != nil {
		return err
	}
	for _, v := range value.(Array) {
		if err := isNumber(v); err != nil {
			return err
		}
	}
	return nil
}

var isNumber ValidateFunc = func(v Value) error {
	if _, ok := v.(Number); !ok {
		return ErrNotNumber(v)
	}
	return nil
}

func (n Number) f(op Operator) *Math {
	switch op {
	case None:
		return Keep
	default:
		return float64Mp[op]
	}

}

func (n Number) t() NodeType {
	return NValue
}

var numberSum = isNumbers.With(func(values Value) Value {
	res := Number(0)
	for _, v := range values.(Array) {
		res += v.(Number)
	}
	return res
})

var numberMul = isNumbers.With(func(values Value) Value {
	res := Number(1)
	for _, v := range values.(Array) {
		res *= v.(Number)
	}
	return res
})

var numberGt = isNumbers.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) >= values[i].(Number) {
			return False
		}
	}
	return True
})

var numberGte = isNumbers.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) > values[i].(Number) {
			return False
		}
	}
	return True
})

var numberLt = isNumbers.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) <= values[i].(Number) {
			return False
		}
	}
	return True
})

var numberLte = isNumbers.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) < values[i].(Number) {
			return False
		}
	}
	return True
})

var numberDiv = isNumbers.With(func(value Value) Value {
	values := value.(Array)
	res := values[0].(Number)
	for _, v := range values[1:] {
		res /= v.(Number)
	}
	return res
})

var numberIn = &Math{
	v: func(value Value) error {
		if err := isArray(value); err != nil {
			return err
		}
		values := value.(Array)
		if err := isNumber(values[0]); err != nil {
			return err
		}
		for _, v := range values[1:] {
			if err := isNumbers(v); err != nil {
				return err
			}
		}
		return nil
	},
	f: func(value Value) Value {
		values := value.(Array)
		now := values[0].(Number)
		for _, arr := range values[1:] {
			if arr.(Array).toMap()[now] == 0 {
				return False
			}
		}
		return True
	},
}

var float64Mp = map[Operator]*Math{
	Sum: numberSum,
	Mul: numberMul,
	Gt:  numberGt,
	Gte: numberGte,
	Lt:  numberLt,
	Lte: numberLte,
	Div: numberDiv,
	In:  numberIn,
}
