package memorydb

import (
	"errors"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/businesslogic"
)

type MemoryDB struct{}

func (db *MemoryDB) GetItem(title string) (businesslogic.TodoItem, error) {
	return businesslogic.TodoItem{}, errors.New("Not Implemented")
}
