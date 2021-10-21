# Vinshop expression tree

[![Status](https://github.com/vinshop/exp-tree/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vinshop/exp-tree/actions/workflows/go.yml)

Exp-tree is go library for parsing expression tree

## Installation

```sh
go get -u github.com/vinshop/exp-tree
```

## Quick start

### Format

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

### Variable

Variables use format `@<string>`, and will be replaced as `Value` from `Variables` when
call `Tree.Calculate(v Variables)`

Support for `String`, `Number` ( as float64), `Bool`, `Array` type

With `Bool` type, we already define `True` and `False`
You could use `et.Var(value)` to auto convert value into corresponding type

### Data type
[Bool](doc/bool.md)

[Number](doc/number.md)

[Array](doc/arr.md)

### Operator
[Operator](doc/operator.md)

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
		"a": et.True, // or et.Var(true)
	})
	fmt.Println(res) // true
}
```
