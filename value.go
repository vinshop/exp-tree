package exp_tree

type NodeValue interface {
	Validate(values ...NodeValue) error // validate before compute
	ComputeMap() ComputeMap
	Byte() []byte
}

type NodeValueWithName interface {
	Name() string
}
