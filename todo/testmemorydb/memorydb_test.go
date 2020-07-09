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
	require.Nil(t, add_err)

	resultItem, err := testObject.GetItem("testItem")
	require.Nil(t, err)
	require.Equal(t, toAdd, resultItem)
}

func TestMemoryDB_AddItem_multipleAdd(t *testing.T) {
	testObject := memorydb.MemoryDB{}
	toAdd := busineslogic.TodoItem{"testItem", "the best item", false}

	add_err1 := testObject.AddItem(toAdd)
	require.Nil(t, add_err1)
	add_err2 := testObject.AddItem(toAdd)
	require.NotNil(t, add_err2)
}

func TestMemoryDB_GetItem_noAdd(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	_, err := testObject.GetItem("not in there")
	require.NotNil(t, err)
}

func TestMemoryDB_HasItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	isThere1, err1 := testObject.HasItem("not in there")
	require.Nil(t, err1)
	require.False(t, isThere1)

	toAdd := busineslogic.TodoItem{"testItem", "the best item", false}
	testObject.AddItem(toAdd)

	isThere2, err2 := testObject.HasItem(toAdd.Title)
	require.Nil(t, err2)
	require.True(t, isThere2)
}

func TestMemoryDB_GetAllItems(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	expectEmpty, err1 := testObject.GetAllItems()
	require.Nil(t, err1)
	require.Equal(t, 0, len(expectEmpty))

	toAdd1 := busineslogic.TodoItem{"testItem", "the best item", false}
	testObject.AddItem(toAdd1)
	toAdd2 := busineslogic.TodoItem{"testItem 2", "the second best item", false}
	testObject.AddItem(toAdd2)

	expectSize2, err2 := testObject.GetAllItems()
	require.Nil(t, err2)
	require.Equal(t, 2, len(expectSize2))
}
