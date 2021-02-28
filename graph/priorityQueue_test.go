package graph_test

import (
	"container/heap"
	"graph"
	"testing"

	"github.com/stretchr/testify/require"
)

func getItemWithPriority(priority int) *graph.PriorityQueueItem {
	result := graph.NewPriorityQueueItem(nil, priority)
	return &result
}

// go test -run Test_PriorityQueue_* .
func Test_PriorityQueue_PushAssending(t *testing.T) {
	testObject := graph.PriorityQueue{}

	for i := 1; i <= 10; i++ {
		heap.Push(&testObject, getItemWithPriority(i))
	}
	heap.Init(&testObject)
	require.Equal(t, 10, testObject.Len())

	for i := 1; i <= 10; i++ {
		thisItem := heap.Pop(&testObject).(*graph.PriorityQueueItem)
		require.Equal(t, i, thisItem.Priority())
	}
	require.Equal(t, 0, testObject.Len())
}
func Test_PriorityQueue_PushDecending(t *testing.T) {
	testObject := graph.PriorityQueue{}

	for i := 10; i >= 1; i-- {
		heap.Push(&testObject, getItemWithPriority(i))
	}
	heap.Init(&testObject)
	require.Equal(t, 10, testObject.Len())

	for i := 1; i <= 10; i++ {
		thisItem := heap.Pop(&testObject).(*graph.PriorityQueueItem)
		require.Equal(t, i, thisItem.Priority())
	}
	require.Equal(t, 0, testObject.Len())
}
