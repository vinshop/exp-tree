package ds

type Container interface {
	Empty() bool
	Push(v interface{})
	Top() interface{}
	Pop() interface{}
}
