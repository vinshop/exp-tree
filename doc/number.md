# Number value type

`Number` is wrapper all number type (`int`, `int32`, `int64`, `float32`, ...) as `float64` builtin type

## Operator

| Operator | Syntax                | A        | B               | Result   | Equivalent |
|----------|-----------------------|----------|-----------------|----------|------------|
| `Sum`    | `{"sum":["@A","@B"]}` | `Number` | `Number`        | `Number` | `A + B`    |
| `Mul`    | `{"mul":["@A","@B"]}` | `Number` | `Number`        | `Number` | `A * B`    |
| `Gt`     | `{"gt":["@A","@B"]}`  | `Number` | `Number`        | `Bool`   | `A > B`    |
| `Gte`    | `{"gte":["@A","@B"]}` | `Number` | `Number`        | `Bool`   | `A >= B`   |
| `Lt`     | `{"lt":["@A","@B"]}`  | `Number` | `Number`        | `Bool`   | `A < B`    |
| `Lte`    | `{"lte":["@A","@B"]}` | `Number` | `Number`        | `Bool`   | `A <= B`   |
| `Div`    | `{"div":["@A","@B"]}` | `Number` | `Number`        | `Number` | `A / B`    |
| `In`     | `{"in":["@A","@B"]}`  | `Number` | `Array(Number)` | `Bool`   | `A in (B)` |
| `Eq`     | `{"eq":["@A","@B"]}`  | `Number` | `Number`        | `Bool`   | `A == B`   |
