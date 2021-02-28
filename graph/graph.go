package graph

type Node interface {
	Name() string
	Forward() []Edge
	Backward() []Node
	AddEdge(Edge)
	AddBackward(Node)
}

type Edge interface {
	Forward() Node
	Value() int
}

type SimpleNode struct {
	name     string
	forward  []Edge
	backward []Node
}

func NewSimpleNode(name string) SimpleNode {
	return SimpleNode{name: name, forward: []Edge{}, backward: []Node{}}
}

type SimpleEdge struct {
	value   int
	forward Node
}

func NewSimpleEdge(value int, forward Node) SimpleEdge {
	return SimpleEdge{value: value, forward: forward}
}

func (n *SimpleNode) Name() string {
	return n.name
}

func (n *SimpleNode) Forward() []Edge {
	return n.forward
}

func (n *SimpleNode) Backward() []Node {
	return n.backward
}

func (n *SimpleNode) AddEdge(newEdge Edge) {
	n.forward = append(n.forward, newEdge)
	newEdge.Forward().AddBackward(n)
}

func (n *SimpleNode) AddBackward(newNode Node) {
	n.backward = append(n.backward, newNode)
}

func (e *SimpleEdge) Forward() Node {
	return e.forward
}

func (e *SimpleEdge) Value() int {
	return e.value
}
