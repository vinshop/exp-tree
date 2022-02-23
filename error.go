package exp_tree

import (
	"errors"
	"fmt"
)

var ErrNotArray = func(v interface{}) error {
	return fmt.Errorf("ErrNotArray: %v is not an array", v)
}

var ErrNotBool = func(v interface{}) error {
	return fmt.Errorf("ErrNotBool: %v is not boolean", v)
}

var ErrNotNumber = func(v interface{}) error {
	return fmt.Errorf("ErrNotNumber: %v is not a number", v)
}

var ErrNotString = func(v interface{}) error {
	return fmt.Errorf("ErrNotString: %v is not a string", v)
}

var ErrNotNested = func(v interface{}) error {
	return fmt.Errorf("ErrNotString: %v is not a nested struct", v)
}

var ErrVarNotFound = func(key string) error {
	return fmt.Errorf("ErrVarNotFound: variable %v not found", key)
}

var ErrParseTree = errors.New("ErrParseTree: error when parse tree")

var ErrOpMustBeUnique = errors.New("ErrOpMustBeUnique: operation must be map with single element")

var ErrCalcTree = errors.New("ErrCalcTree: error when calculate")

var ErrOperatorNotSupported = func(op Operator, value Value) error {
	return fmt.Errorf(`ErrOperatorNotSupported: operator "%v" not supported for value %v`, op, value)
}
