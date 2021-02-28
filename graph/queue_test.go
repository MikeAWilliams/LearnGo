package maw_test

import (
	"maw"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Queue_IsEmptyAtStart(t *testing.T) {
	testObject := maw.Queue{}
	require.True(t, testObject.IsEmpty())
}

func Test_Queue_PushPop(t *testing.T) {
	items := []string{"one", "two", "three"}
	testObject := maw.Queue{}
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
