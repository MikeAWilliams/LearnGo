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

func (db *MemoryDB) HasItem(title string) (bool, error) {
	resultPtr := db.FindItem(title)
	if nil == resultPtr {
		return false, nil
	}
	return true, nil
}

func (db *MemoryDB) GetAllItems() ([]busineslogic.TodoItem, error) {
	result := make([]busineslogic.TodoItem, len(db.items))
	copy(result, db.items)
	return result, nil
}

func (db *MemoryDB) AddItem(item busineslogic.TodoItem) error {
	existingItem := db.FindItem(item.Title)
	if nil != existingItem {
		return errors.New(fmt.Sprintf("Item with title %v already in db", item.Title))
	}
	db.items = append(db.items, item)
	return nil
}

func (db *MemoryDB) UpdateItem(item busineslogic.TodoItem) error {
	findPtr := db.FindItem(item.Title)
	if nil == findPtr {
		return errors.New("item not in db")
	}
	findPtr.Complete = item.Complete
	findPtr.Description = item.Description
	return nil
}

func (db *MemoryDB) DeleteItem(title string) error {
	return errors.New("not implemented")
}

// Helper functions

func (db *MemoryDB) FindItem(title string) *busineslogic.TodoItem {
	for index := range db.items {
		if db.items[index].Title == title {
			return &db.items[index]
		}
	}
	return nil
}
