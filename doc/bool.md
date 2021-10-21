# Bool value type

`Bool` is wrapper of go builtin `bool` type

## Operators

| Operator | Syntax                  | A      | B      | Result | Equivalent |
|----------|-------------------------|--------|--------|--------|------------|
| `and`    | `{"and": ["@A", "@B"]}` | `Bool` | `Bool` | `Bool` | `A and B`  |
| `or`     | `{"or": ["@A", "@B"]}`  | `Bool` | `Bool` | `Bool` | `A or B`   |
| `not`    | `{"not": "@A"}`         | `Bool` |        | `Bool` | `not A`    |