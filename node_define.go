package exp_tree

type NodeType int

const (
	nodeTypeOp NodeType = iota
	nodeTypeGroup
	nodeTypeVariable
	nodeTypeValue
)

type OpType string

const (
	None OpType = ""
	And  OpType = "and"
	Or   OpType = "or"
	Sum  OpType = "sum"
	Mul  OpType = "mul"
	Gt   OpType = "gt"  // greater than
	Gte  OpType = "gte" // greater than equal
	Lt   OpType = "lt"  // less than
	Lte  OpType = "lte" // less than equal
	Eq   OpType = "eq"  // equal
	Not  OpType = "not"
	Xor  OpType = "xor"
)

type Node interface {
	Type() NodeType
}
