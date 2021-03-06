package main

import "sync"

type IntQueue struct {
	items []int
	lock  sync.RWMutex
}

// 创建队列
func (q *IntQueue) New() *IntQueue {
	q.items = []int{}
	return q
}

// 入队列
func (q *IntQueue) Enqueue(t int) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

// 出队列
func (q *IntQueue) Dequeue() *int {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// 获取队列的第一个元素, 不移除
func (q *IntQueue) Front() *int {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// 判空
func (q *IntQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 获取队列的长度
func (q *IntQueue) Size() int {
	return len(q.items)
}
