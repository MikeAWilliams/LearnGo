package trie_test

import (
	"maw/trie"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrie_SearchNoInsert(t *testing.T) {
	testObject := trie.NewTrie()
	require.False(t, testObject.Search("a"))
}

func TestTrie_InsertSingleCharacters(t *testing.T) {
	testObject := trie.NewTrie()
	testObject.Insert("a")
	testObject.Insert("b")
	require.True(t, testObject.Search("a"))
	require.True(t, testObject.Search("b"))
}

func TestTrie_InsertMultiCharacter(t *testing.T) {
	testObject := trie.NewTrie()
	testObject.Insert("ab")
	require.True(t, testObject.Search("ab"))
	require.False(t, testObject.Search("ac"))

	testObject.Insert("theword")
	require.True(t, testObject.Search("theword"))
}

func find(word string, slice []string) bool {
	for _, sliceW := range slice {
		if sliceW == word {
			return true
		}
	}
	return false
}

func TestTrie_AutoComplete_WordNotIn(t *testing.T) {
	testObject := trie.NewTrie()
	testObject.Insert("aaa")
	testObject.Insert("aba")
	testObject.Insert("aca")

	empty := testObject.AutoComplete("ba")
	require.Empty(t, empty)
}

func TestTrie_AutoComplete(t *testing.T) {
	testObject := trie.NewTrie()
	testObject.Insert("aaa")
	testObject.Insert("aba")
	testObject.Insert("aca")

	result := testObject.AutoComplete("a")

	require.Equal(t, 3, len(result))
	require.True(t, find("aaa", result))
	require.True(t, find("aba", result))
	require.True(t, find("aca", result))
}
