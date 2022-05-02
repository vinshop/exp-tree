package exp_tree

import "testing"

func TestParse(t *testing.T) {
	input := "3 + 4 * 2 / ( 1 − 5 ) ^ 2 ^ 3"
	Parse(input)
}
