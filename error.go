package exp_tree

import "fmt"

var ErrNotArray = func(v interface{}) error {
	return fmt.Errorf("%v is not an array", v)
}

var ErrNotBool = func(v interface{}) error {
	return fmt.Errorf("%v is not boolean", v)
}

var ErrNotNumber = func(v interface{}) error {
	return fmt.Errorf("%v is not a number", v)
}

var ErrNotString = func(v interface{}) error {
	return fmt.Errorf("%v is not a string", v)
}
