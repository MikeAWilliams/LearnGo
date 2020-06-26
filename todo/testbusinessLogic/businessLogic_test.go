package testBusinessLogic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
)

type dbSpy struct {
	lastAdded busineslogic.TodoItem
}

func (db *dbSpy) GetItem(title string) (busineslogic.TodoItem, error) {
	if db.lastAdded.Title == title {
		return db.lastAdded, nil
	}
	return busineslogic.TodoItem{}, errors.New("The item is not in the db")
}
func (db *dbSpy) GetAllItems() ([]busineslogic.TodoItem, error) {
	return []busineslogic.TodoItem{}, errors.New("Not implemented")
}

func (db *dbSpy) AddItem(item busineslogic.TodoItem) (bool, busineslogic.TodoItem) {
	if db.lastAdded.Title == item.Title {
		return false, db.lastAdded
	}
	db.lastAdded = item
	return true, item
}

func TestMySpy_AddItem(t *testing.T) {
	spy := dbSpy{}
	toAdd := busineslogic.TodoItem{"title", "description", false}
	addWorked, added := spy.AddItem(toAdd)
	require.True(t, addWorked)
	require.Equal(t, toAdd, added)
}
