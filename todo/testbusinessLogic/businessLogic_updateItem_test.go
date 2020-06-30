package testBusinessLogic

import (
	"errors"
	"testing"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/stretchr/testify/require"
)

func TestBusinessLogic_UpdateItem_ItemInDb(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, true}

	newItem := busineslogic.TodoItem{"LearnGo", "write some code", true}
	item, err := busineslogic.UpdateItem(newItem, &db)
	require.Nil(t, err)
	require.Equal(t, "LearnGo", item.Title)
	require.Equal(t, 3, db.methodCalls)
}

func TestBusinessLogic_UpdateItem_ItemNotInDb(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, false}

	newItem := busineslogic.TodoItem{"LearnGo", "write some code", true}
	_, err := busineslogic.UpdateItem(newItem, &db)
	require.NotNil(t, err)
	require.Equal(t, 1, db.methodCalls)
}

func TestBusinessLogic_UpdateItem_HasItemErr(t *testing.T) {
	db := dbErr{0, errors.New("Error"), nil, nil, nil, true}

	newItem := busineslogic.TodoItem{"LearnGo", "write some code", true}
	_, err := busineslogic.UpdateItem(newItem, &db)
	require.NotNil(t, err)
	require.Equal(t, 1, db.methodCalls)
}

func TestBusinessLogic_UpdateItem_GetItemErr(t *testing.T) {
	db := dbErr{0, nil, errors.New("Error"), nil, nil, true}

	newItem := busineslogic.TodoItem{"LearnGo", "write some code", true}
	_, err := busineslogic.UpdateItem(newItem, &db)
	require.NotNil(t, err)
	require.Equal(t, 2, db.methodCalls)
}

func TestBusinessLogic_UpdateItem_UpdateItemErr(t *testing.T) {
	db := dbErr{0, nil, nil, nil, errors.New("error"), true}

	newItem := busineslogic.TodoItem{"LearnGo", "write some code", true}
	_, err := busineslogic.UpdateItem(newItem, &db)
	require.NotNil(t, err)
	require.Equal(t, 3, db.methodCalls)
}
