package maw_test

import (
	"container/heap"
	"maw"
	"testing"

	"github.com/stretchr/testify/require"
)

func getItemWithPriority(priority int) *maw.PriorityQueueItem {
	result := maw.NewPriorityQueueItem(nil, priority)
	return &result
}

// go test -run Test_PriorityQueue_* .
func Test_PriorityQueue_PushAssending(t *testing.T) {
	testObject := maw.PriorityQueue{}

	for i := 1; i <= 10; i++ {
		heap.Push(&testObject, getItemWithPriority(i))
	}
	heap.Init(&testObject)
	require.Equal(t, 10, testObject.Len())

	for i := 1; i <= 10; i++ {
		thisItem := heap.Pop(&testObject).(*maw.PriorityQueueItem)
		require.Equal(t, i, thisItem.Priority())
	}
	require.Equal(t, 0, testObject.Len())
}
func Test_PriorityQueue_PushDecending(t *testing.T) {
	testObject := maw.PriorityQueue{}

	for i := 10; i >= 1; i-- {
		heap.Push(&testObject, getItemWithPriority(i))
	}
	heap.Init(&testObject)
	require.Equal(t, 10, testObject.Len())

	for i := 1; i <= 10; i++ {
		thisItem := heap.Pop(&testObject).(*maw.PriorityQueueItem)
		require.Equal(t, i, thisItem.Priority())
	}
	require.Equal(t, 0, testObject.Len())
}
