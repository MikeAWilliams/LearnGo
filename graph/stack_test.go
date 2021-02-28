package maw_test

import (
	"maw"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Stack_IsEmptyAtStart(t *testing.T) {
	testObject := maw.Stack{}
	require.True(t, testObject.IsEmpty())
}

func Test_Stack_PushPop(t *testing.T) {
	items := []string{"one", "two", "three"}
	testObject := maw.Stack{}
	for _, item := range items {
		testObject.Add(item)
	}

	for i := len(items) - 1; i >= 0; i-- {
		require.False(t, testObject.IsEmpty())
		item := testObject.Remove()
		require.Equal(t, items[i], item)
	}
	require.True(t, testObject.IsEmpty())
}
