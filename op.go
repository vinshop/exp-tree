package exp_tree

type Operator string

const (
	None Operator = ""
	And  Operator = "and"
	Or   Operator = "or"
	Sum  Operator = "sum"
	Mul  Operator = "mul"
	Div  Operator = "div"
	Lte  Operator = "lte"
	Gte  Operator = "gte"
	Lt   Operator = "lt"
	Gt   Operator = "gt"
	In   Operator = "in"
)
