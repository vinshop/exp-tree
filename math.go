package exp_tree

//ValidateFunc validate args before do the math
type ValidateFunc func(value Value) error

func (v ValidateFunc) With(f ComputeFunc) *Math {
	return &Math{
		v: v,
		f: f,
	}
}

//ComputeFunc do the math
type ComputeFunc func(value Value) Value

//Math Some random asian will do the math for us
type Math struct {
	v ValidateFunc
	f ComputeFunc
}

//calc Math do the math for us
func (m *Math) calc(value Value) (Value, error) {
	if m.v != nil {
		if err := m.v(value); err != nil {
			return nil, err
		}
	}
	return m.f(value), nil
}

//Keep when the mathematician bored, they leave the value intact
var Keep = &Math{
	v: func(_ Value) error {
		return nil
	},
	f: func(v Value) Value {
		return v
	},
}
