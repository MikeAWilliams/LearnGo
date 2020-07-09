package memorydb

import (
	"errors"
	"fmt"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
)

type MemoryDB struct {
	items []busineslogic.TodoItem
}

// Database interface functions

func (db *MemoryDB) GetItem(title string) (busineslogic.TodoItem, error) {
	for _, item := range db.items {
		if item.Title == title {
			return item, nil
		}
	}
	return busineslogic.TodoItem{}, errors.New(fmt.Sprintf("Item with title %v not found", title))
}

func (db *MemoryDB) AddItem(item busineslogic.TodoItem) error {
	db.items = append(db.items, item)
	return nil
}
