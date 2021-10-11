package exp_tree

var isArr ValidateFunc = func(values Value) error {
	if _, ok := values.(Array); !ok {
		return ErrNotArray(values)
	}
	return nil
}

var isBoolArr ValidateFunc = func(value Value) error {
	if err := isArr(value); err != nil {
		return err
	}
	for _, v := range value.(Array) {
		if err := isBool(v); err != nil {
			return err
		}
	}
	return nil
}

var isBool ValidateFunc = func(value Value) error {
	if _, ok := value.(Bool); !ok {
		return ErrNotBool(value)
	}
	return nil
}

var isNumberArr ValidateFunc = func(value Value) error {
	if err := isArr(value); err != nil {
		return err
	}
	for _, v := range value.(Array) {
		if err := isNumber(v); err != nil {
			return err
		}
	}
	return nil
}

var isNumber ValidateFunc = func(v Value) error {
	if _, ok := v.(Number); !ok {
		return ErrNotNumber(v)
	}
	return nil
}

var isString ValidateFunc = func(v Value) error {
	if _, ok := v.(String); !ok {
		return ErrNotString(v)
	}
	return nil
}

var isStringArr ValidateFunc = func(value Value) error {
	if err := isArr(value); err != nil {
		return err
	}
	for _, v := range value.(Array) {
		if err := isString(v); err != nil {
			return err
		}
	}
	return nil
}
