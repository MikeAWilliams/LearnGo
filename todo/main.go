package main

import (
	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
)

func main() {
	toPrint := busineslogic.TodoItem{"The thing", "You better get this done", false}
	busineslogic.PrintTodoItem(toPrint)
}
