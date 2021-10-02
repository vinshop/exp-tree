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

func TestTree_Calculate_NoVariable(t *testing.T) {
	q := `{"and": [":con", {"lt": [1, 2]}]}`
	tree, err := ParseTree(q)

	assert.Nil(t, err)

	_, err = tree.Calculate(nil)
	assert.Equal(t, ErrVariableNotFound("con"), err)
}

func TestTree_Calculate(t *testing.T) {
	q := `{"and": [":con", {"lt": [1, 2]}]}`
	tree, err := ParseTree(q)

	assert.Nil(t, err)

	res, err := tree.Calculate(Variables{
		"con": True,
	})

	assert.Nil(t, err)
	assert.Equal(t, True, res)
}

func TestTree_Calculate2(t *testing.T) {
	q := `{"eq":[{"sum":[1,2,3,4,5]},{"sum":[{"mul":[2,5]},5]}]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)

	res, err := tree.Calculate(nil)

	assert.Nil(t, err)
	assert.Equal(t, True, res)
}

func TestTree_CalculateString(t *testing.T) {
	q := `{"in": [":val", "A", "B", "C"]}`
	tree, err := ParseTree(q)
	assert.Nil(t, err)

	_, err = tree.Calculate(nil)
	assert.Equal(t, ErrVariableNotFound(Variable("val")), err)

	res, err := tree.Calculate(Variables{
		"val": String("A"),
	})
	assert.Nil(t, err)
	assert.Equal(t, True, res)
	res, err = tree.Calculate(Variables{
		"val": String("D"),
	})
	assert.Nil(t, err)
	assert.Equal(t, False, res)

}
