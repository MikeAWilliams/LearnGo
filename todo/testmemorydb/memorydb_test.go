package testmemorydb

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/MikeAWilliams/LearnGo/tree/master/todo/memorydb"
)

func TestMemoryDB_AddItem_GetItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}
	toAdd := busineslogic.TodoItem{"testItem", "the best item", false}

	add_err := testObject.AddItem(toAdd)
	require.Equal(t, nil, add_err)

	resultItem, err := testObject.GetItem("testItem")
	require.Equal(t, nil, err)
	require.Equal(t, toAdd, resultItem)
}
