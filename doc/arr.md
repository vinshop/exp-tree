# Array value type

`Array` is slice of `Value` type

## Operators

| Operator | Syntax                     | A       | B       | Result | Equivalent               |
|----------|----------------------------|---------|---------|--------|--------------------------|
| `in`     | `{"in": ["@A", "@B"]}`     | `Array` | `Array` | `Bool` | `all of A in B`          |
| `one_in` | `{"one_in": ["@A", "@B"]}` | `Array` | `Array` | `Bool` | `at least one of A in B` |

