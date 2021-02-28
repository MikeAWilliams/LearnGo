package graph_test

import (
	"graph"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Queue_IsEmptyAtStart(t *testing.T) {
	testObject := graph.Queue{}
	require.True(t, testObject.IsEmpty())
}

func Test_Queue_PushPop(t *testing.T) {
	items := []string{"one", "two", "three"}
	testObject := graph.Queue{}
	for _, item := range items {
		testObject.Add(item)
	}

	for _, item := range items {
		require.False(t, testObject.IsEmpty())
		queueItem := testObject.Remove()
		require.Equal(t, item, queueItem)
	}
	require.True(t, testObject.IsEmpty())
}
