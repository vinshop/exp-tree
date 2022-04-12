package exp_tree

//Bool store bool value
type Bool bool

func (b Bool) F(op Operator) *Math {
	return boolMp[op]
}

func (b Bool) Type() NodeType {
	return NValue
}

func (b Bool) Variables() Variables {
	return nil
}

const True = Bool(true)
const False = Bool(false)

var boolAnd = isBoolArr.With(func(values Value) Value {
	for _, v := range values.(Array) {
		if v == False {
			return False
		}
	}
	return True
})

var boolOr = isBoolArr.With(func(values Value) Value {
	for _, v := range values.(Array) {
		if v == True {
			return True
		}
	}
	return False
})

var boolNot = isBool.With(func(value Value) Value {
	return !value.(Bool)
})

var boolMp = map[Operator]*Math{
	None: Keep,
	And:  boolAnd,
	Or:   boolOr,
	Not:  boolNot,
}
