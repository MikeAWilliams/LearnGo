package graph

import (
	"container/heap"
	"errors"
	"math"
)

func reverseEdges(reversePath []Edge) []Edge {
	path := make([]Edge, len(reversePath))
	for i, edge := range reversePath {
		j := len(reversePath) - 1 - i
		path[j] = edge
	}
	return path
}

func findForwardEdge(currentNode Node, previousNode Node) (Edge, bool) {
	for _, edge := range previousNode.Forward() {
		if edge.Forward() == currentNode {
			return edge, true
		}
	}
	return nil, false
}

func getEdgesFromDijkstraPredicessorMap(root Node, target Node, predicessors map[Node]Node) ([]Edge, error) {
	previousNode, ok := predicessors[target]
	if !ok {
		return nil, errors.New("the target cannot be reached from the source")
	}

	reversePath := []Edge{}
	currentNode := target
	for {
		nextEdge, edgeFound := findForwardEdge(currentNode, previousNode)
		if !edgeFound {
			return nil, errors.New("unable to find edge between two nodes")
		}
		reversePath = append(reversePath, nextEdge)

		currentNode = previousNode
		if currentNode == root {
			break
		}
		previousNode, ok = predicessors[previousNode]
		if !ok {
			return nil, errors.New("can't find the path")
		}
	}

	return reverseEdges(reversePath), nil
}

func DijkstraBetweenTwo(root Node, target Node) (int, []Edge, error) {
	distances, predicessors, err := Dijkstra(root, root)
	if nil != err {
		return 0, nil, err
	}

	totalDistance, ok := distances[target]
	if !ok {
		return 0, nil, errors.New("distance map does not contain target")
	}
	path, err := getEdgesFromDijkstraPredicessorMap(root, target, predicessors)
	if nil != err {
		return 0, nil, err
	}

	return totalDistance, path, nil
}

func setupDijkstra(root Node, source Node) (map[Node]int, map[Node]Node, PriorityQueue) {
	distanceToTargetFromNode := make(map[Node]int)
	distanceToTargetFromNode[source] = 0

	queue := PriorityQueue{}

	predicessorOfNode := make(map[Node]Node)
	DepthFirstSearch(root, func(node Node) {
		var newItem PriorityQueueItem
		if node != source {
			distanceToTargetFromNode[node] = -1
			predicessorOfNode[node] = nil
			newItem = NewPriorityQueueItem(node, math.MaxInt32)
		} else {
			newItem = NewPriorityQueueItem(node, 0)
		}
		heap.Push(&queue, &newItem)
	})
	heap.Init(&queue)
	return distanceToTargetFromNode, predicessorOfNode, queue
}

func Dijkstra(root Node, source Node) (map[Node]int, map[Node]Node, error) {
	distanceToTargetFromNode, predicessorOfNode, queue := setupDijkstra(root, source)

	for queue.Len() > 0 {
		minItem := heap.Pop(&queue).(*PriorityQueueItem)
		for _, edge := range minItem.value.Forward() {
			currentDistance := distanceToTargetFromNode[minItem.value]
			if -1 == currentDistance {
				continue
			}
			newDistance := currentDistance + edge.Value()
			forwardNode := edge.Forward()
			distanceToTarget := distanceToTargetFromNode[forwardNode]
			if -1 == distanceToTarget || newDistance < distanceToTarget {
				distanceToTargetFromNode[forwardNode] = newDistance
				predicessorOfNode[forwardNode] = minItem.value
				err := queue.UpdatePriorityForNode(forwardNode, newDistance)
				if nil != err {
					return nil, nil, err
				}
			}
		}
	}

	return distanceToTargetFromNode, predicessorOfNode, nil
}
