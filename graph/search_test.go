package graph_test

import (
	"graph"
	"testing"

	"github.com/stretchr/testify/require"
)

var visited []graph.Node

func visit(node graph.Node) {
	visited = append(visited, node)
}

func Test_BreadthFirst(t *testing.T) {
	root := buildFourNodeGraph()
	visited = []graph.Node{}

	graph.BreadthFirstSearch(root, visit)

	expectedNames := []string{"one", "two", "four", "three"}
	require.Equal(t, len(expectedNames), len(visited))
}

func Test_DepthFirstSearch(t *testing.T) {
	root := buildFourNodeGraph()
	visited = []graph.Node{}

	graph.DepthFirstSearch(root, visit)

	expectedNames := []string{"one", "two", "three", "four"}
	require.Equal(t, len(expectedNames), len(visited))
}

func buildFourNodeGraphWithLoops() graph.Node {
	root, node2 := buildTwoNodeGraph()
	edge2r := graph.NewSimpleEdge(4, root)
	node2.AddEdge(&edge2r)

	node3 := graph.NewSimpleNode("three")
	edge23 := graph.NewSimpleEdge(1, &node3)
	node2.AddEdge(&edge23)

	node4 := graph.NewSimpleNode("four")
	edge14 := graph.NewSimpleEdge(4, &node4)
	root.AddEdge(&edge14)

	edge42 := graph.NewSimpleEdge(6, node2)
	node4.AddEdge(&edge42)

	//doGraphvizOutputToFile(root, "loopy.txt")

	return root
}
func Test_DepthFirstSearchWithLoop(t *testing.T) {
	root := buildFourNodeGraphWithLoops()
	visited = []graph.Node{}

	graph.DepthFirstSearch(root, visit)

	expectedNames := []string{"one", "two", "three", "four"}
	require.Equal(t, len(expectedNames), len(visited))
}
