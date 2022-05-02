package ds

import "sync"

type stack struct {
	mu  sync.Mutex
	top *node
}

func NewStack() Container {
	return &stack{}
}

func (s *stack) Empty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.top == nil
}

func (s *stack) Push(e interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	next := &node{e, nil}
	next.next = s.top
	s.top = next
}

func (s *stack) Top() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.top == nil {
		return nil
	}
	return s.top.value
}
func (s *stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.top == nil {
		return nil
	}
	v := s.top.value
	s.top = s.top.next
	return v
}
