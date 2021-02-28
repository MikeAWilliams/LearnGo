package maw

import (
	"container/heap"
	"errors"
)

type PriorityQueueItem struct {
	value    Node
	priority int
	index    int
}

func (i *PriorityQueueItem) Priority() int {
	return i.priority
}

func NewPriorityQueueItem(node Node, priority int) PriorityQueueItem {
	return PriorityQueueItem{value: node, priority: priority, index: -1}
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) UpdatePriorityForNode(node Node, newPriority int) error {
	for _, item := range *pq {
		if item.value == node {
			pq.update(item, newPriority)
			return nil
		}
	}
	return errors.New("We should have found the node")
}

func (pq *PriorityQueue) update(item *PriorityQueueItem, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
