package exp_tree

func calc(op Operator, t Node) (Value, error) {
	switch t.t() {
	case NValue:
		value := t.(Value)
		return value.f(op).calc(value)
	case NOperation:
		op := t.(Op)
		return calc(op.op, op.args)
	case NGroup:
		group := t.(Group)
		arr := make(Array, 0, len(group))
		for _, n := range group {
			value, err := calc(None, n)
			if err != nil {
				return nil, err
			}
			arr = append(arr, value)
		}
		return calc(op, arr)
	default:
		panic("err")
	}
}
