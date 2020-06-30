package testBusinessLogic

import (
	"errors"
	"testing"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/stretchr/testify/require"
)

func TestBusinessLogic_DeleteItem_ItemInDb(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, nil, true}

	item, err := busineslogic.DeleteItem("LearnGo", &db)
	require.Nil(t, err)
	require.Equal(t, "LearnGo", item.Title)
	require.Equal(t, 3, db.methodCalls)
}

func TestBusinessLogic_DeleteItem_ItemNotInDb(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, nil, false}

	_, err := busineslogic.DeleteItem("LearnGo", &db)
	require.NotNil(t, err)
	require.Equal(t, 1, db.methodCalls)
}

func TestBusinessLogic_DeleteItem_HasItemErr(t *testing.T) {
	db := dbErr{0, errors.New("Error"), nil, nil, nil, nil, true}

	_, err := busineslogic.DeleteItem("LearnGo", &db)
	require.NotNil(t, err)
	require.Equal(t, 1, db.methodCalls)
}

func TestBusinessLogic_DeleteItem_GetItemErr(t *testing.T) {
	db := dbErr{0, nil, errors.New("Error"), nil, nil, nil, true}

	_, err := busineslogic.DeleteItem("LearnGo", &db)
	require.NotNil(t, err)
	require.Equal(t, 2, db.methodCalls)
}

func TestBusinessLogic_DeleteItem_DeleteItemErr(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, errors.New("error"), true}

	_, err := busineslogic.DeleteItem("LearnGo", &db)
	require.NotNil(t, err)
	require.Equal(t, 3, db.methodCalls)
}
