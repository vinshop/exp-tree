package exp_tree

type ComputeFunc func(values ...NodeValue) NodeValue
type ComputeMap map[OpType]ComputeFunc
