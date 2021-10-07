package main

import (
	"fmt"

	et "github.com/vinshop/exp-tree"
)

func main() {
	tree, err := et.ParseTree(`{"and":["@a",{"lt":[1,2]}]}`)
	if err != nil {
		panic(err)
	}
	treeJSON, err := tree.JSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(treeJSON)
	res, err := tree.Calculate(et.Variables{
		"a": et.True,
	})
	fmt.Println(res)
}
