package exp_tree

type Number float64

func (n Number) F(op Operator) *Math {
	switch op {
	case None:
		return Keep
	default:
		return numberMap[op]
	}
}

func (n Number) Type() NodeType {
	return NValue
}

var numberMap = map[Operator]*Math{
	Sum: numberSum,
	Mul: numberMul,
	Gt:  numberGt,
	Gte: numberGte,
	Lt:  numberLt,
	Lte: numberLte,
	Div: numberDiv,
	In:  numberIn,
}

var numberSum = isNumberArr.With(func(values Value) Value {
	res := Number(0)
	for _, v := range values.(Array) {
		res += v.(Number)
	}
	return res
})

var numberMul = isNumberArr.With(func(values Value) Value {
	res := Number(1)
	for _, v := range values.(Array) {
		res *= v.(Number)
	}
	return res
})

var numberGt = isNumberArr.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) >= values[i].(Number) {
			return False
		}
	}
	return True
})

var numberGte = isNumberArr.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) > values[i].(Number) {
			return False
		}
	}
	return True
})

var numberLt = isNumberArr.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) <= values[i].(Number) {
			return False
		}
	}
	return True
})

var numberLte = isNumberArr.With(func(value Value) Value {
	values := value.(Array)
	for i := 1; i < len(values); i++ {
		if values[i-1].(Number) < values[i].(Number) {
			return False
		}
	}
	return True
})

var numberDiv = isNumberArr.With(func(value Value) Value {
	values := value.(Array)
	res := values[0].(Number)
	for _, v := range values[1:] {
		res /= v.(Number)
	}
	return res
})

var numberIn = &Math{
	v: func(value Value) error {
		if err := isArr(value); err != nil {
			return err
		}
		values := value.(Array)
		if err := isNumber(values[0]); err != nil {
			return err
		}
		for _, v := range values[1:] {
			if err := isNumberArr(v); err != nil {
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
