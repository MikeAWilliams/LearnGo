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
	resultPtr := db.FindItem(title)
	if nil == resultPtr {
		return busineslogic.TodoItem{}, errors.New(fmt.Sprintf("Item with title %v not found", title))
	}
	return *resultPtr, nil
}

func (db *MemoryDB) AddItem(item busineslogic.TodoItem) error {
	existingItem := db.FindItem(item.Title)
	if nil != existingItem {
		return errors.New(fmt.Sprintf("Item with title %v already in db", item.Title))
	}
	db.items = append(db.items, item)
	return nil
}

// Helper functions

func (db *MemoryDB) FindItem(title string) *busineslogic.TodoItem {
	for _, item := range db.items {
		if item.Title == title {
			return &item
		}
	}
	return nil
}
