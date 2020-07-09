package testmemorydb

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/MikeAWilliams/LearnGo/tree/master/todo/memorydb"
)

func TestMemoryDB_AddItemGetItem_noError(t *testing.T) {
	testObject := memorydb.MemoryDB{}
	toAdd := busineslogic.TodoItem{"testItem", "the best item", false}

	add_err := testObject.AddItem(toAdd)
	require.Equal(t, nil, add_err)

	resultItem, err := testObject.GetItem("testItem")
	require.Equal(t, nil, err)
	require.Equal(t, toAdd, resultItem)
}

func TestMemoryDB_AddItem_multipleAdd(t *testing.T) {
	testObject := memorydb.MemoryDB{}
	toAdd := busineslogic.TodoItem{"testItem", "the best item", false}

	add_err1 := testObject.AddItem(toAdd)
	require.Equal(t, nil, add_err1)
	add_err2 := testObject.AddItem(toAdd)
	require.NotEqual(t, nil, add_err2)
}

func TestMemoryDB_GetItem_noAdd(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	_, err := testObject.GetItem("not in there")
	require.NotEqual(t, nil, err)
}
