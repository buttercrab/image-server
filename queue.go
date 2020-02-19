package main

import (
	"sync"
)

var popLock sync.Mutex
var pushLock sync.Mutex

type QueueNode struct {
	value int
	next  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	len  int
}

func (q *Queue) Push(x int) {
	pushLock.Lock()
	defer pushLock.Unlock()
	next := QueueNode{
		value: x,
		next:  nil,
	}
	if q.len == 0 {
		q.tail = &next
		q.head = &next
		q.len = 1
		return
	}
	q.tail.next = &next
	q.tail = &next
	q.len++
}

// assert len > 0 in this project
func (q *Queue) Pop() int {
	popLock.Lock()
	defer popLock.Unlock()
	x := q.head.value
	q.head = q.head.next
	q.len--
	return x
}
