package graph_test

import (
	"fmt"
	"graph"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func buildTwoNodeGraph() (graph.Node, graph.Node) {
	node1 := graph.NewSimpleNode("one")
	node2 := graph.NewSimpleNode("two")
	edge := graph.NewSimpleEdge(1, &node2)
	node1.AddEdge(&edge)
	return &node1, &node2
}

func buildFourNodeGraph() graph.Node {
	root, node2 := buildTwoNodeGraph()
	node3 := graph.NewSimpleNode("three")
	edge23 := graph.NewSimpleEdge(1, &node3)
	node2.AddEdge(&edge23)

	node4 := graph.NewSimpleNode("four")
	edge14 := graph.NewSimpleEdge(4, &node4)
	root.AddEdge(&edge14)
	return root
}

func Test_BuildTwoNodeGraph(t *testing.T) {
	node1, node2 := buildTwoNodeGraph()

	forwardEdges := node1.Forward()
	require.Equal(t, 1, len(forwardEdges))
	require.Equal(t, 1, forwardEdges[0].Value())
	require.Equal(t, 0, len(node1.Backward()))
	require.Equal(t, "two", forwardEdges[0].Forward().Name())

	backwardNodes := node2.Backward()
	require.Equal(t, 1, len(backwardNodes))
	require.Equal(t, "one", backwardNodes[0].Name())
	require.Equal(t, 0, len(node2.Forward()))
}

type terminalWriter struct{}

func (t terminalWriter) Write(bytes []byte) (int, error) {
	fmt.Print(string(bytes))
	return len(bytes), nil
}

func doGraphvizOutputToTerminal(root graph.Node) {
	writer := terminalWriter{}
	graph.OutputGraph(root, writer)
}

func doGraphvizOutputToFile(root graph.Node, fileName string) error {
	file, err := os.Create(fileName)
	if nil != err {
		return err
	}
	defer file.Close()
	graph.OutputGraphDfs(root, file)
	return nil
}

//func Test_GraphvizOutput(t *testing.T) {
//root := buildFourNodeGraph()
//doGraphvizOutputToTerminal(root)
//doGraphvizOutputToFile(root, "stuff.txt")
//}
