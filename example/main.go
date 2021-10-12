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
	res, err := tree.Calculate(et.Variables{
		"a": et.True, // or et.Var(true)
	})
	fmt.Println(res) // true
}
