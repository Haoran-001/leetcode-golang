package main

import "sync"

type TreeNodeQueue struct {
	TreeNodeItems []*TreeNode
	lock          sync.RWMutex
}

func (queue *TreeNodeQueue) Enqueue(item *TreeNode) {
	queue.lock.Lock()
	queue.TreeNodeItems = append(queue.TreeNodeItems, item)
	queue.lock.Unlock()
}

func (queue *TreeNodeQueue) Dequeue() *TreeNode {
	queue.lock.Lock()
	item := queue.TreeNodeItems[0]
	queue.TreeNodeItems = queue.TreeNodeItems[1:]
	queue.lock.Unlock()
	return item
}

func (queue *TreeNodeQueue) Len() int {
	return len(queue.TreeNodeItems)
}

func (queue *TreeNodeQueue) Empty() bool {
	return queue.Len() == 0
}
