package exp_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_JSON(t *testing.T) {
	q := `{"and":[":con",{"lt":[1,2]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)
	treeJSON, err := tree.JSON()
	assert.Nil(t, err)
	assert.Equal(t, q, treeJSON)
}

func TestTree_Caculate_NoVariable(t *testing.T) {
	q := `{"and": [":con", {"lt": [1, 2]}]}`
	tree, err := ParseTree(q)

	assert.Nil(t, err)

	_, err = tree.Caculate(nil)
	assert.Equal(t, ErrVariableNotFound("con"), err)
}

func TestTree_Caculate(t *testing.T) {
	q := `{"and": [":con", {"lt": [1, 2]}]}`
	tree, err := ParseTree(q)

	assert.Nil(t, err)

	res, err := tree.Caculate(Variables{
		"con": True,
	})

	assert.Nil(t, err)
	assert.Equal(t, True, res)
}

func TestTree_Caculate2(t *testing.T) {
	q := `{"eq":[{"sum":[1,2,3,4,5]},{"sum":[{"mul":[2,5]},5]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)

	res, err := tree.Caculate(nil)

	assert.Nil(t, err)
	assert.Equal(t, True, res)
}
