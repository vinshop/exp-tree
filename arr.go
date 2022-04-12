package exp_tree

//Array store array of value
type Array []Value

func (a Array) Type() NodeType {
	return NValue
}

func (Array) Variables() Variables {
	return nil
}

func (a Array) F(op Operator) *Math {

	switch op {
	case In:
		if err := isArr(a[0]); err != nil {
			return a[0].F(op)
		}
		fallthrough
	default:
		m := arrMap[op]
		if m == nil {
			return a[0].F(op)
		}
		return m
	}
}

var arrMap = map[Operator]*Math{
	In: isArr.With(
		func(value Value) Value {
			values := value.(Array)
			mp := values[0].(Array).toMap()
			for _, v := range values[1:] {
				vMp := v.(Array).toMap()
				for k := range mp {
					if vMp[k] == 0 {
						return False
					}
				}
			}
			return True
		}),
	OneIn: isArr.With(
		func(value Value) Value {
			values := value.(Array)
			mp := values[0].(Array).toMap()
			for _, v := range values[1:] {
				ok := False
				vMp := v.(Array).toMap()
				for k := range mp {
					if vMp[k] != 0 {
						ok = True
						break
					}
				}
				if !ok {
					return False
				}
			}
			return True
		},
	),
}

//toMap convert Array to with key is Value and value is number of Value in Array
func (a Array) toMap() map[Value]int {
	mp := make(map[Value]int)
	for _, v := range a {
		mp[v]++
	}
	return mp
}
