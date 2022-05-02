package ds

import "sync"

func NewQueue() Container {
	return &queue{}
}

type queue struct {
	mu   sync.Mutex
	head *node
	tail *node
}

func (q *queue) Empty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.head == nil
}

func (q *queue) Push(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	next := &node{v, nil}

	if q.tail == nil {
		q.head = next
	} else {
		q.tail.next = next
	}
	q.tail = next
}

func (q *queue) Top() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil {
		return nil
	}
	return q.head.value
}

func (q *queue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil {
		return nil
	}
	v := q.head.value
	q.head = q.head.next
	return v
}
