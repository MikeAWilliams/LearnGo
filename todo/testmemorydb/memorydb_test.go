package testmemorydb

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/memorydb"
)

func TestMemoryDB_AddItem_(t *testing.T) {
	testObject := memorydb.MemoryDB{}
	item, err := testObject.GetItem("testItem")
	require.NotEqual(t, nill, err)
}
