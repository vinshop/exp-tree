package exp_tree

func calcValue(op Operator, value Value) (Value, error) {
	math := value.F(op)
	if math == nil {
		return nil, ErrOperatorNotSupported(op, value)
	}
	return math.calc(value)
}

func calc(op Operator, t Node, vars Variables) (Value, error) {
	switch t.Type() {
	case NVariable:
		return vars.Get(t.(Variable))
	case NValue:
		value := t.(Value)
		return calcValue(op, value)
	case NOperation:
		op := t.(*Operation)
		res, err := calc(op.op, op.args, vars)
		op.result = res
		return res, err
	case NGroup:
		group := t.(Group)
		arr := make(Array, 0, len(group))
		for _, n := range group {
			value, err := calc(None, n, vars)
			if err != nil {
				return nil, err
			}
			arr = append(arr, value)
		}
		return calc(op, arr, vars)
	default:
		return nil, ErrCalcTree
	}
}
