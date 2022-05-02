package exp_tree

import (
	"fmt"
	"github.com/vinshop/exp-tree/ds"
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Operator struct {
	Priority  int
	Direction Direction
}

var OM = map[string]*Operator{
	"^": {4, Right},
	"*": {3, Left},
	"/": {3, Left},
	"+": {2, Left},
	"-": {2, Left},
}

func isNumber(e string) (float64, bool) {
	num, err := strconv.ParseFloat(e, 64)
	if err != nil {
		return 0, false
	}
	return num, true
}

func isOp(e string) (*Operator, bool) {
	op, ok := OM[e]
	if !ok {
		return nil, false
	}
	return op, true
}

func isFun(s string) bool {
	return false
}

func Parse(exp string) {
	tokens := strings.Split(exp, " ")
	output := ds.NewQueue()
	op := ds.NewStack()
	for _, token := range tokens {
		if _, ok := isNumber(token); ok {
			output.Push(token)
			continue
		}
		if _, ok := isOp(token); ok {
			for !op.Empty() {
				top := op.Top().(string)
				if top == "(" {
					break
				}
				if OM[top].Priority > OM[token].Priority || (OM[token].Direction == Left && OM[token].Priority == OM[top].Priority) {
					output.Push(top)
					op.Pop()
					continue
				}
				break
			}
			op.Push(token)
			continue
		}
		if token == "(" {
			op.Push("(")
			continue
		}
		if token == ")" {
			for {
				if op.Empty() {
					panic("err")
				}
				top := op.Top().(string)
				if top == "(" {
					break
				}
				output.Push(top)
				op.Pop()
			}
			op.Pop()
			continue
		}

	}
	for !op.Empty() {
		output.Push(op.Pop())
	}
	s := strings.Builder{}
	for !output.Empty() {
		s.WriteString(output.Pop().(string))
	}
	fmt.Println(s.String())
}
