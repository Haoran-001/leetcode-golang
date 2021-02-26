package main

import "sync"

type NodeQueue struct {
	NodeItems []*Node
	lock      sync.RWMutex
}

func (queue *NodeQueue) Enqueue(item *Node) {
	queue.lock.Lock()
	queue.NodeItems = append(queue.NodeItems, item)
	queue.lock.Unlock()
}

func (queue *NodeQueue) Dequeue() *Node {
	queue.lock.Lock()
	item := queue.NodeItems[0]
	queue.NodeItems = queue.NodeItems[1:]
	queue.lock.Unlock()
	return item
}

func (queue *NodeQueue) Len() int {
	return len(queue.NodeItems)
}

func (queue *NodeQueue) Empty() bool {
	return queue.Len() == 0
}
