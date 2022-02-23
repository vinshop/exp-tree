package exp_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTree(t *testing.T) {
	q := `{"and":[{"gt":[{"sum":[1,2,3]},5]},{"in":["hello",["hello","world"]]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	fmt.Printf("%#v\n", tree)
	res, err := tree.Calculate(nil)
	assert.Nil(t, err)
	assert.Equal(t, True, res)
}

func TestParseTreeNested(t *testing.T) {
	q := `{"and": "@nested.Bool"}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	resp, err := tree.Calculate(Variables{
		"nested": Var(struct {
			Bool bool
		}{
			Bool: true,
		}),
	})
	assert.Nil(t, err)
	assert.Equal(t, True, resp)
}

func TestParseTree2(t *testing.T) {
	q := `{"and": [{"in": ["@district", ["001"]]}, {"in": ["@province", ["01"]]}, {"in": ["@gt_level", [1, 2]]}, {"lte": ["@order.TotalAmount", 1000000]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	resp, err := tree.Calculate(Variables{
		"district": Var("001"),
		"province": Var("01"),
		"gt_level": Var(1),
		"order": Var(struct {
			TotalAmount int
		}{
			TotalAmount: 100000,
		}),
		//"order.total_amount": Var(10000),
	})
	assert.Nil(t, err)
	assert.Equal(t, True, resp)
}
