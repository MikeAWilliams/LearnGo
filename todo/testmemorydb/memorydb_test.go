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

func TestMemoryDB_UpdateItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	testItem := busineslogic.TodoItem{"testItem", "the best item", false}
	errNotThere := testObject.UpdateItem(testItem)
	require.NotNil(t, errNotThere)

	testObject.AddItem(testItem)

	testItem.Description = "not as good as I though"
	testItem.Complete = true
	shouldBeNil := testObject.UpdateItem(testItem)
	require.Nil(t, shouldBeNil)

	resultItem, _ := testObject.GetItem(testItem.Title)
	require.Equal(t, testItem.Description, resultItem.Description)
	require.Equal(t, testItem.Complete, resultItem.Complete)
}

func TestMemoryDB_DeleteItem_NotThere(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	errNotThere := testObject.DeleteItem("not in there")
	require.NotNil(t, errNotThere)
}

func TestMemoryDB_DeleteItem_SingleItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	testItem := busineslogic.TodoItem{"testItem", "the best item", false}
	testObject.AddItem(testItem)
	isThere, _ := testObject.HasItem(testItem.Title)
	require.True(t, isThere)

	errorShouldBeNil := testObject.DeleteItem(testItem.Title)
	require.Nil(t, errorShouldBeNil)

	isThereNow, _ := testObject.HasItem(testItem.Title)
	require.False(t, isThereNow)
}

func TestMemoryDB_DeleteItem_FirstItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	item1 := busineslogic.TodoItem{"testItem1", "the best item", false}
	testObject.AddItem(item1)
	item2 := busineslogic.TodoItem{"testItem2", "the best item", false}
	testObject.AddItem(item2)
	item3 := busineslogic.TodoItem{"testItem3", "the best item", false}
	testObject.AddItem(item3)

	errorShouldBeNil := testObject.DeleteItem(item1.Title)
	require.Nil(t, errorShouldBeNil)

	item1IsHere, _ := testObject.HasItem(item1.Title)
	require.False(t, item1IsHere)
	item2IsHere, _ := testObject.HasItem(item2.Title)
	require.True(t, item2IsHere)
	item3IsHere, _ := testObject.HasItem(item3.Title)
	require.True(t, item3IsHere)
}

func TestMemoryDB_DeleteItem_SecondItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	item1 := busineslogic.TodoItem{"testItem1", "the best item", false}
	testObject.AddItem(item1)
	item2 := busineslogic.TodoItem{"testItem2", "the best item", false}
	testObject.AddItem(item2)
	item3 := busineslogic.TodoItem{"testItem3", "the best item", false}
	testObject.AddItem(item3)

	errorShouldBeNil := testObject.DeleteItem(item2.Title)
	require.Nil(t, errorShouldBeNil)

	item1IsHere, _ := testObject.HasItem(item1.Title)
	require.True(t, item1IsHere)
	item2IsHere, _ := testObject.HasItem(item2.Title)
	require.False(t, item2IsHere)
	item3IsHere, _ := testObject.HasItem(item3.Title)
	require.True(t, item3IsHere)
}

func TestMemoryDB_DeleteItem_LastItem(t *testing.T) {
	testObject := memorydb.MemoryDB{}

	item1 := busineslogic.TodoItem{"testItem1", "the best item", false}
	testObject.AddItem(item1)
	item2 := busineslogic.TodoItem{"testItem2", "the best item", false}
	testObject.AddItem(item2)
	item3 := busineslogic.TodoItem{"testItem3", "the best item", false}
	testObject.AddItem(item3)

	errorShouldBeNil := testObject.DeleteItem(item3.Title)
	require.Nil(t, errorShouldBeNil)

	item1IsHere, _ := testObject.HasItem(item1.Title)
	require.True(t, item1IsHere)
	item2IsHere, _ := testObject.HasItem(item2.Title)
	require.True(t, item2IsHere)
	item3IsHere, _ := testObject.HasItem(item3.Title)
	require.False(t, item3IsHere)
}
