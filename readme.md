# Vinshop expression tree

[![Status](https://github.com/vinshop/exp-tree/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vinshop/exp-tree/actions/workflows/go.yml)

Exp-tree is go library for parsing expression tree

## Installation
```sh
go get -u github.com/vinshop/exp-tree
```
## Quick start
Expression tree is in format:
```json
{
  "<operator>": [
    "<arg1>",
    "<arg2>",
    {
      "<opetator>": [
        "<arg>",
        "<arg>",
        "<@variable>"
      ]
    }
  ]
}
```
example
```json
{"and":["@a",{"lt":[1,2]}]}
```
is equivalent to `@a and (1 < 2)` with `a` is a variable
### Parse tree
```go
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
}
```
### Calculate
```go
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
		"a": et.True,
	})
	fmt.Println(res) // true
}
```
Current only support for String, Float64, Bool type

With Bool type, we already define True and False