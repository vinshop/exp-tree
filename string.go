package exp_tree

//String store string value
type String string

func (s String) Type() NodeType {
	return NValue
}

func (s String) F(op Operator) *Math {
	return stringMap[op]
}

var stringMap = map[Operator]*Math{
	None: Keep,
	In:   stringIn,
	Eq:   stringEqual,
}

var stringEqual = isStringArr.With(func(value Value) Value {
	values := value.(Array)
	now := values[0]
	for _, v := range values[1:] {
		if v != now {
			return False
		}
	}
	return True
})

var stringIn = &Math{
	v: func(value Value) error {
		if err := isArr(value); err != nil {
			return err
		}
		values := value.(Array)
		if err := isString(values[0]); err != nil {
			return err
		}
		for _, v := range values[1:] {
			if err := isStringArr(v); err != nil {
				return err
			}
		}
		return nil
	},
	f: func(value Value) Value {
		values := value.(Array)
		s := values[0].(String)
		for _, v := range values[1:] {
			if v.(Array).toMap()[s] == 0 {
				return False
			}
		}
		return True
	},
}
