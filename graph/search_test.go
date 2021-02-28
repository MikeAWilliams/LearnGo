package maw_test

import (
	"maw"
	"testing"

	"github.com/stretchr/testify/require"
)

var visited []maw.Node

func visit(node maw.Node) {
	visited = append(visited, node)
}

func Test_BreadthFirst(t *testing.T) {
	root := buildFourNodeGraph()
	visited = []maw.Node{}

	maw.BreadthFirstSearch(root, visit)

	expectedNames := []string{"one", "two", "four", "three"}
	require.Equal(t, len(expectedNames), len(visited))
}

func Test_DepthFirstSearch(t *testing.T) {
	root := buildFourNodeGraph()
	visited = []maw.Node{}

	maw.DepthFirstSearch(root, visit)

	expectedNames := []string{"one", "two", "three", "four"}
	require.Equal(t, len(expectedNames), len(visited))
}

func buildFourNodeGraphWithLoops() maw.Node {
	root, node2 := buildTwoNodeGraph()
	edge2r := maw.NewSimpleEdge(4, root)
	node2.AddEdge(&edge2r)

	node3 := maw.NewSimpleNode("three")
	edge23 := maw.NewSimpleEdge(1, &node3)
	node2.AddEdge(&edge23)

	node4 := maw.NewSimpleNode("four")
	edge14 := maw.NewSimpleEdge(4, &node4)
	root.AddEdge(&edge14)

	edge42 := maw.NewSimpleEdge(6, node2)
	node4.AddEdge(&edge42)

	//doGraphvizOutputToFile(root, "loopy.txt")

	return root
}
func Test_DepthFirstSearchWithLoop(t *testing.T) {
	root := buildFourNodeGraphWithLoops()
	visited = []maw.Node{}

	maw.DepthFirstSearch(root, visit)

	expectedNames := []string{"one", "two", "three", "four"}
	require.Equal(t, len(expectedNames), len(visited))
}
