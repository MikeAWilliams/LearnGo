package graph

import "errors"

type NodeOperation func(n Node)

func search(root Node, operation NodeOperation, container SearchContainer) error {
	visited := make(map[Node]bool)
	visited[root] = true
	container.Add(root)

	for !container.IsEmpty() {
		item := container.Remove()
		thisNode, ok := item.(Node)
		if !ok {
			return errors.New("Got an unxpected type in the search container")
		}
		operation(thisNode)
		for _, edge := range thisNode.Forward() {
			forwardNode := edge.Forward()
			if _, ok := visited[forwardNode]; !ok {
				visited[forwardNode] = true
				container.Add(forwardNode)
			}
		}
	}
	return nil
}

func BreadthFirstSearch(root Node, operation NodeOperation) error {
	container := Queue{}
	return search(root, operation, &container)
}

func DepthFirstSearch(root Node, operation NodeOperation) error {
	container := Stack{}
	return search(root, operation, &container)
}
