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
	GetAllItems() ([]TodoItem, error)
	AddItem(item TodoItem) (bool, TodoItem)
}

func AddItem(title string, description string, db Database) (bool, TodoItem) {
	return db.AddItem(TodoItem{title, description, false})
}

func PrintTodoItem(item TodoItem) {
	fmt.Print(item)
}
