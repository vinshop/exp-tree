package exp_tree


type Nested struct {
	value interface{}
}

func (n Nested) Value() interface{} {
	return n.value
}

func (n Nested) Type() NodeType {
	return NNested
}

func (n Nested) F(op Operator) *Math {
	return nil
}

