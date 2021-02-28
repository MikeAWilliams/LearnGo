package maw_test

import (
	"maw"
	"testing"

	"github.com/stretchr/testify/require"
)

func buildStraitLineGraph(costFrom4to5 int) (maw.Node, maw.Node) {
	node1 := maw.NewSimpleNode("one")
	node2 := maw.NewSimpleNode("two")
	node3 := maw.NewSimpleNode("three")
	node4 := maw.NewSimpleNode("four")
	node5 := maw.NewSimpleNode("five")

	edge1 := maw.NewSimpleEdge(1, &node2)
	edge2 := maw.NewSimpleEdge(2, &node3)
	edge3 := maw.NewSimpleEdge(3, &node4)
	edge4 := maw.NewSimpleEdge(costFrom4to5, &node5)

	node1.AddEdge(&edge1)
	node2.AddEdge(&edge2)
	node3.AddEdge(&edge3)
	node4.AddEdge(&edge4)

	return &node1, &node5
}

func buildTwoPathGraph(costFrom4to5 int) (maw.Node, maw.Node) {
	node1, node5 := buildStraitLineGraph(costFrom4to5)

	extraN1 := maw.NewSimpleNode("extra one")
	extraN2 := maw.NewSimpleNode("extra two")
	extraN3 := maw.NewSimpleNode("extra three")

	extraEdge1 := maw.NewSimpleEdge(10, &extraN1)
	extraEdge2 := maw.NewSimpleEdge(20, &extraN2)
	extraEdge3 := maw.NewSimpleEdge(30, &extraN3)
	extraEdge4 := maw.NewSimpleEdge(40, node5)

	node1.AddEdge(&extraEdge1)
	extraN1.AddEdge(&extraEdge2)
	extraN2.AddEdge(&extraEdge3)
	extraN3.AddEdge(&extraEdge4)

	//doGraphvizOutputToFile(node1, "twoPath.txt")

	return node1, node5
}

func buildTwoPathGraphWithLoop() (maw.Node, maw.Node) {
	node1, node5 := buildTwoPathGraph(4)
	extraEdge := maw.NewSimpleEdge(0, node1)
	node5.AddEdge(&extraEdge)
	//doGraphvizOutputToFile(node1, "twoPathLoop.txt")
	return node1, node5
}

func testDijkstraOutput(
	t *testing.T,
	expectedPredicessorNames map[string]string,
	expectedDistance map[string]int,
	distanceMap map[maw.Node]int,
	pathMap map[maw.Node]maw.Node,
	err error) {

	require.Nil(t, err)
	require.Equal(t, len(expectedDistance), len(distanceMap))
	require.Equal(t, len(expectedPredicessorNames), len(pathMap))

	for destinationNode, predicessorNode := range pathMap {
		require.Equal(t, expectedPredicessorNames[destinationNode.Name()], predicessorNode.Name())
	}

	for node, distance := range distanceMap {
		require.Equal(t, expectedDistance[node.Name()], distance)
	}
}

func testDijkstraBetweenTwoNodesOutput(
	t *testing.T,
	expectedDistance int,
	expectedEdges []maw.Edge,
	distance int,
	path []maw.Edge,
	err error) {
	require.Nil(t, err)
	require.Equal(t, expectedDistance, distance)
	require.Equal(t, len(expectedEdges), len(path))

	for resultIndex, resultEdge := range path {
		require.Equal(t, expectedEdges[resultIndex], resultEdge)
	}
}

func Test_Dijkstra_TwoNodes(t *testing.T) {
	root, _ := buildTwoNodeGraph()

	expectedPredicessorNames := map[string]string{
		"two": "one",
	}

	expectedDistance := map[string]int{
		"one": 0,
		"two": 1,
	}

	distanceMap, pathMap, err := maw.Dijkstra(root, root)
	testDijkstraOutput(t, expectedPredicessorNames, expectedDistance, distanceMap, pathMap, err)
}

func Test_DijkstraBetweenTwo_TwoNodes(t *testing.T) {
	root, target := buildTwoNodeGraph()

	expectedEdges := []maw.Edge{}
	edge := root.Forward()[0]
	expectedEdges = append(expectedEdges, edge)

	totalDistance, path, err := maw.DijkstraBetweenTwo(root, target)
	testDijkstraBetweenTwoNodesOutput(t, 1, expectedEdges, totalDistance, path, err)
}

func Test_Dijkstra_SinglePath(t *testing.T) {
	root, _ := buildStraitLineGraph(4)

	expectedPredicessorNames := map[string]string{
		"two":   "one",
		"three": "two",
		"four":  "three",
		"five":  "four",
	}

	expectedDistance := map[string]int{
		"one":   0,
		"two":   1,
		"three": 3,
		"four":  6,
		"five":  10,
	}

	distanceMap, pathMap, err := maw.Dijkstra(root, root)
	testDijkstraOutput(t, expectedPredicessorNames, expectedDistance, distanceMap, pathMap, err)
}

func Test_DijkstraBetweenTwo_SinglePath(t *testing.T) {
	root, target := buildStraitLineGraph(4)

	expectedEdges := []maw.Edge{}
	edge := root.Forward()[0]
	expectedEdges = append(expectedEdges, edge)
	edge = edge.Forward().Forward()[0]
	expectedEdges = append(expectedEdges, edge)
	edge = edge.Forward().Forward()[0]
	expectedEdges = append(expectedEdges, edge)
	edge = edge.Forward().Forward()[0]
	expectedEdges = append(expectedEdges, edge)

	totalDistance, path, err := maw.DijkstraBetweenTwo(root, target)
	testDijkstraBetweenTwoNodesOutput(t, 10, expectedEdges, totalDistance, path, err)
}

func Test_Dijkstra_TwoPath(t *testing.T) {
	root, _ := buildTwoPathGraph(4)

	expectedPredicessorNames := map[string]string{
		"two":         "one",
		"three":       "two",
		"four":        "three",
		"five":        "four",
		"extra one":   "one",
		"extra two":   "extra one",
		"extra three": "extra two",
	}

	expectedDistance := map[string]int{
		"one":         0,
		"two":         1,
		"three":       3,
		"four":        6,
		"five":        10,
		"extra one":   10,
		"extra two":   30,
		"extra three": 60,
	}

	distanceMap, pathMap, err := maw.Dijkstra(root, root)
	testDijkstraOutput(t, expectedPredicessorNames, expectedDistance, distanceMap, pathMap, err)
}

func Test_Dijkstra_TwoPathWithLoop(t *testing.T) {
	root, _ := buildTwoPathGraphWithLoop()

	expectedPredicessorNames := map[string]string{
		"two":         "one",
		"three":       "two",
		"four":        "three",
		"five":        "four",
		"extra one":   "one",
		"extra two":   "extra one",
		"extra three": "extra two",
	}

	expectedDistance := map[string]int{
		"one":         0,
		"two":         1,
		"three":       3,
		"four":        6,
		"five":        10,
		"extra one":   10,
		"extra two":   30,
		"extra three": 60,
	}

	distanceMap, pathMap, err := maw.Dijkstra(root, root)
	testDijkstraOutput(t, expectedPredicessorNames, expectedDistance, distanceMap, pathMap, err)
}

func Test_Dijkstra_TwoPathHighCostAtEnd(t *testing.T) {
	root, _ := buildTwoPathGraph(400)

	expectedPredicessorNames := map[string]string{
		"two":         "one",
		"three":       "two",
		"four":        "three",
		"five":        "extra three",
		"extra one":   "one",
		"extra two":   "extra one",
		"extra three": "extra two",
	}

	expectedDistance := map[string]int{
		"one":         0,
		"two":         1,
		"three":       3,
		"four":        6,
		"five":        100,
		"extra one":   10,
		"extra two":   30,
		"extra three": 60,
	}

	distanceMap, pathMap, err := maw.Dijkstra(root, root)
	testDijkstraOutput(t, expectedPredicessorNames, expectedDistance, distanceMap, pathMap, err)
}
