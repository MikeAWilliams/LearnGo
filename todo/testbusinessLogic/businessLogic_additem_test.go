package testBusinessLogic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
)

type dbDouble struct {
	lastAdded busineslogic.TodoItem
}

func (db *dbDouble) GetItem(title string) (busineslogic.TodoItem, error) {
	if db.lastAdded.Title == title {
		return db.lastAdded, nil
	}
	return busineslogic.TodoItem{}, errors.New("The item is not in the db")
}

func (db *dbDouble) HasItem(title string) (bool, error) {
	if db.lastAdded.Title == title {
		return true, nil
	}
	return false, nil
}

func (db *dbDouble) GetAllItems() ([]busineslogic.TodoItem, error) {
	return []busineslogic.TodoItem{}, errors.New("Not implemented")
}

func (db *dbDouble) AddItem(item busineslogic.TodoItem) error {
	if db.lastAdded.Title == item.Title {
		return errors.New("the item is there")
	}
	db.lastAdded = item
	return nil
}

func (db *dbDouble) UpdateItem(item busineslogic.TodoItem) error {
	return errors.New("Not implemented")
}

func TestBusinessLogic_AddItem_ItemNotInDB(t *testing.T) {
	db := dbDouble{}

	wasAdded, item, err := busineslogic.AddItem("Laundry", "Clean laundry", &db)
	require.Equal(t, nil, err)
	require.True(t, wasAdded)
	require.Equal(t, "Laundry", item.Title)
	require.Equal(t, "Clean laundry", item.Description)
	require.False(t, item.Complete)

	require.Equal(t, item, db.lastAdded)
}

func TestBusinessLogic_AddItem_ItemInDB(t *testing.T) {
	toAdd := busineslogic.TodoItem{"Laundry", "Clean laundry", true}
	db := dbDouble{toAdd}

	wasAdded, item, err := busineslogic.AddItem(toAdd.Title, toAdd.Description, &db)
	require.Equal(t, nil, err)
	require.False(t, wasAdded)
	require.Equal(t, toAdd, item)
}

type dbErr struct {
	methodCalls   int
	hasItemErr    error
	getItemErr    error
	addItemErr    error
	updateItemErr error
	hasItem       bool
}

func (db *dbErr) GetItem(title string) (busineslogic.TodoItem, error) {
	db.methodCalls++
	return busineslogic.TodoItem{title, "item desc", false}, db.getItemErr
}

func (db *dbErr) HasItem(title string) (bool, error) {
	db.methodCalls++
	return db.hasItem, db.hasItemErr
}

func (db *dbErr) GetAllItems() ([]busineslogic.TodoItem, error) {
	db.methodCalls++
	return []busineslogic.TodoItem{}, errors.New("Not implemented")
}

func (db *dbErr) AddItem(item busineslogic.TodoItem) error {
	db.methodCalls++
	return db.addItemErr
}

func (db *dbErr) UpdateItem(item busineslogic.TodoItem) error {
	db.methodCalls++
	return db.updateItemErr
}

func TestBusinessLogic_AddItem_HasItemIsError(t *testing.T) {
	db := dbErr{0, errors.New("Error"), nil, nil, nil, false}

	wasAdded, _, err := busineslogic.AddItem("Laundry", "Clean laundry", &db)
	require.Equal(t, db.hasItemErr, err)
	require.Equal(t, 1, db.methodCalls)
	require.False(t, wasAdded)
}

func TestBusinessLogic_AddItem_GetItemIsError(t *testing.T) {
	db := dbErr{0, nil, errors.New("Error"), nil, nil, true}

	wasAdded, _, err := busineslogic.AddItem("Laundry", "Clean laundry", &db)
	require.Equal(t, db.getItemErr, err)
	require.Equal(t, 2, db.methodCalls)
	require.False(t, wasAdded)
}

func TestBusinessLogic_AddItem_AddItemIsError(t *testing.T) {
	db := dbErr{0, nil, nil, errors.New("Error"), nil, false}

	wasAdded, _, err := busineslogic.AddItem("Laundry", "Clean laundry", &db)
	require.Equal(t, db.addItemErr, err)
	require.Equal(t, 2, db.methodCalls)
	require.False(t, wasAdded)
}
