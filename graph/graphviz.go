package maw

import (
	"fmt"
	"io"
	"strconv"
)

func writeLine(line string, writer io.Writer) {
	writer.Write([]byte(line + "\n"))
}

func outputNode(root Node, writer io.Writer) {
	name := strconv.Quote(root.Name())
	for _, edge := range root.Forward() {
		forward := edge.Forward()
		writeLine(name+"->"+strconv.Quote(forward.Name())+fmt.Sprintf(" [label=%v];", edge.Value()), writer)
		outputNode(forward, writer)
	}
}

func OutputGraph(root Node, writer io.Writer) {
	writeLine("digraph thegrapyh {", writer)
	outputNode(root, writer)
	writeLine("}", writer)
}

type outputOperation struct {
	writer io.Writer
}

func (o *outputOperation) write(node Node) {
	name := strconv.Quote(node.Name())
	for _, edge := range node.Forward() {
		forward := edge.Forward()
		writeLine(name+"->"+strconv.Quote(forward.Name())+fmt.Sprintf(" [label=%v];", edge.Value()), o.writer)
	}
}

func OutputGraphDfs(root Node, writer io.Writer) {
	writeLine("digraph thegrapyh {", writer)

	writeObject := outputOperation{writer: writer}
	DepthFirstSearch(root, writeObject.write)

	writeLine("}", writer)
}
