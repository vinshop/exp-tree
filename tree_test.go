package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSON(t *testing.T) {
	q := `{"and":["@a",{"lt":[1,2]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	data, err := tree.JSON()
	assert.Nil(t, err)
	assert.Equal(t, q, data)
}

func TestJSON2(t *testing.T) {
	q :=`{"and":[{"in":["@district","001"]},{"in":["@province","01"]},{"in":["@gt_level",1,2]},{"lte":["@order.total_amount",1000000]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	data, err := tree.JSON()
	assert.Nil(t, err)
	assert.Equal(t, q, data)
}
