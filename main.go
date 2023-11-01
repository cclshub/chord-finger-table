package main

import (
	"fmt"
	"math"
)

var nodes = []*Node{}

// init Nodes manually for now
// Initialize the nodes without specifying successors.
func initNodes() {
	Node1 := &Node{id: 10, keys: []int{}, fTable: []*Finger{}}
	Node2 := &Node{id: 20, keys: []int{}, fTable: []*Finger{}}
	Node3 := &Node{id: 30, keys: []int{}, fTable: []*Finger{}}
	Node4 := &Node{id: 40, keys: []int{}, fTable: []*Finger{}}
	Node5 := &Node{id: 50, keys: []int{}, fTable: []*Finger{}}

	// Set the successors to create the Chord ring.
	Node1.successor = Node2
	Node2.successor = Node3
	Node3.successor = Node4
	Node4.successor = Node5
	Node5.successor = Node1

	nodes = append(nodes, Node1, Node2, Node3, Node4, Node5)

	// populate keys
	for _, n := range nodes {
		n.populateKeys()
	}
	for _, n := range nodes {
		populateFingerTable(n)
		n.populatefTable()
	}
}

/*
we set the no. of bits as m = 6 for now
this means that the keys range from 0 - 64
we use the populateKeys fn to populate the keys slice for each Node
*/
func (n *Node) populateKeys() {
	myId := n.id
	predecessorID := n.predecessor.id
	if myId > predecessorID {
		for i := predecessorID + 1; i <= myId; i++ {
			n.keys = append(n.keys, i)
		}
	} else {
		for i := 0; i <= myId; i++ {
			n.keys = append(n.keys, i)
		}
		for i := predecessorID + 1; i <= int(math.Pow(2, float64(m))); i++ {
			n.keys = append(n.keys, i)
		}
	}
}

func (n *Node) populatefTable() {
	fTable := n.fTable
	for i, finger := range fTable {
		k := finger.key
		for _, node := range nodes {
			keys := node.keys
			for _, nkey := range keys {
				if k == nkey {
					n.fTable[i].node = node
				}
			}
		}
	}
}

func main() {
	initNodes()
	// for _, node := range nodes {
	// 	// fmt.Printf("The newly populated node is: %s \n", node.fTable)
	// }

	// time.Sleep(1 * time.Second)

	// pass a key and get the node in which it is stored in
	ownerof64 := nodes[3].findSuccessor(64)
	fmt.Printf("ownerof21 is: %s", ownerof64)
}
