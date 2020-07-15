package main

import (
	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/MikeAWilliams/LearnGo/tree/master/todo/httpdriver"
	"github.com/MikeAWilliams/LearnGo/tree/master/todo/memorydb"
)

func getDatabase() busineslogic.Database {
	db := memorydb.MemoryDB{}

	busineslogic.AddItem("testItem1", "the test item", &db)
	busineslogic.AddItem("testItem2", "the second test item", &db)

	return &db
}

func main() {
	db := getDatabase()
	httpdriver.Start(db)
}
