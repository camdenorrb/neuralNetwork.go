package neural

import "gonum.org/v1/gonum/mat"

type Network struct {
	headNode NetworkNode // The first node in the network (input layer)
}

type NetworkNode interface {
	forward(input mat.Matrix)

	result() mat.Matrix

	next() *NetworkNode
}

func (n *Network) forward(input mat.Matrix) {
	n.headNode.forward(input)
}

func (n *Network) resultOfLastNode() mat.Matrix {
	node := n.headNode
	next := node.next()
	for next != nil {
		node = *next
		next = node.next()
	}
	return node.result()
}
