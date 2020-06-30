package testBusinessLogic

import (
	"errors"
	"testing"

	"github.com/MikeAWilliams/LearnGo/tree/master/todo/busineslogic"
	"github.com/stretchr/testify/require"
)

func TestBusinessLogic_GetItem_ItemInDb(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, nil, true}

	item, err := busineslogic.GetItem("testItem", &db)
	require.Nil(t, err)
	require.Equal(t, "testItem", item.Title)
	require.Equal(t, 2, db.methodCalls)
}

func TestBusinessLogic_GetItem_ItemNotInDb(t *testing.T) {
	db := dbErr{0, nil, nil, nil, nil, nil, false}

	_, err := busineslogic.GetItem("testItem", &db)
	require.NotNil(t, err)
	require.Equal(t, 1, db.methodCalls)
}

func TestBusinessLogic_GetItem_HasItemReturnsError(t *testing.T) {
	db := dbErr{0, errors.New("bad has item"), nil, nil, nil, nil, false}

	_, err := busineslogic.GetItem("testItem", &db)
	require.NotNil(t, err)
	require.Equal(t, 1, db.methodCalls)
}

func TestBusinessLogic_GetItem_GetItemReturnsError(t *testing.T) {
	db := dbErr{0, nil, errors.New("can't get item"), nil, nil, nil, true}

	_, err := busineslogic.GetItem("testItem", &db)
	require.NotNil(t, err)
	require.Equal(t, 2, db.methodCalls)
}
