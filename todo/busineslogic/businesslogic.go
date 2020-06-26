package busineslogic

import "fmt"

type TodoItem struct {
	Title       string
	Description string
	Complete    bool
}

func (item TodoItem) String() string {
	return fmt.Sprintf("Title:%q Description:%q Complete:%v", item.Title, item.Description, item.Complete)
}

type Database interface {
	GetItem(title string) (TodoItem, error)
	HasItem(title string) (bool, error)
	GetAllItems() ([]TodoItem, error)
	AddItem(item TodoItem) error
}

func AddItem(title string, description string, db Database) (bool, TodoItem, error) {
	hasItem, hasErr := db.HasItem(title)
	if nil != hasErr {
		return false, TodoItem{}, hasErr
	}

	if hasItem {
		existingItem, getErr := db.GetItem(title)
		if nil != getErr {
			return false, TodoItem{}, getErr
		}
		return false, existingItem, nil
	}

	toAdd := TodoItem{title, description, false}
	addErr := db.AddItem(toAdd)
	if nil != addErr {
		return false, TodoItem{}, addErr
	}

	return true, toAdd, nil
}

func PrintTodoItem(item TodoItem) {
	fmt.Print(item)
}
