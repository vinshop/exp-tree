package exp_tree

//Array store array of value
type Array []Value

func (a Array) Type() NodeType {
	return NValue
}

func (a Array) F(op Operator) *Math {
	return a[0].F(op)
}

//toMap convert Array to with key is Value and value is number of Value in Array
func (a Array) toMap() map[Value]int {
	mp := make(map[Value]int)
	for _, v := range a {
		mp[v]++
	}
	return mp
}
