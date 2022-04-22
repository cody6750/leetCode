package datastructures

import (
	"log"
)

/* Tree
Description:
- A tree is a graph that has no cycles
- Non linear data structure
- Made up of nodes or verticies and edges without having any cycle.
- Tree with no nodes are empty trees or nil
- Tree can have only one parent, but multiple children
- Leaf nodes are nodes with no children
- Sibling nodes are nodes with the same parent
- Height : Number of edges from deepest node to node
- Depth : Numbert of edges from node to deepest node
- Binary Tree : Maxiumum of two children
- Binary Search Tree : Nodes are ordered by their values, making it efficient for search and inserting data.
  Key of any node is >= the value of its left child and <= the value of its right cild
- Two main traversal strategies: BFS and DFS
- BFS: Visits node of the same level before going to ndoes of the next depth level
- DFS: Visitrs nodes in order from root down to leaf recursivley
- DFS has three travesal sub-types, each refer to the order the nodes are traversed:
	1. In-order: Left, root, right
	2. Pre-order: root, Left, Right
	3. Post-Order Left, Right, root
- Binary trees become unbalanced when nodes are inserted and removed. To keep search and insert times logirthmic, keep the binary trees balanced.
- A height balanced tree ensures that the left and right subtrees of every node have a height difference of not more than one.
- Rotations are used to balance a binary tree
- Self balancing trees automaticall ykeep thei rmaximum height difference as close to one as possible.const
- Types of self-balancing Binary trees
	1. Red-Black trees
	2. AVL trees
*/

type Tree struct {
	height int
	root   *Node
}
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (t *Tree) Insert(v int) {
	if t.root == nil {
		t.root = &Node{Value: v}
		return
	} else {
		t.root.Insert(v)
	}
}

func (n *Node) Insert(v int) {
	if n.Left == nil && n.Value >= v {
		n.Left = &Node{Value: v}
		return
	}
	if n.Right == nil && n.Value < v {
		n.Right = &Node{Value: v}
		return
	}
	if n.Value >= v {
		n.Left.Insert(v)
	}
	if n.Value < v {
		n.Right.Insert(v)
	}
}

//  5
// 2  6
//2  3  7
//1   4  8

func (n *Node) Print() {
	if n.Left != nil {
		n.Left.Print()
	}

	log.Printf("%v", n)
	if n.Right != nil {
		n.Right.Print()
	}
}

func (t *Tree) Print() {
	if t.root != nil {
		t.root.Print()
	}
}

func (t *Tree) Delete(v int) {
	if t.root == nil {
		return
	}
	stringify(t.root, 0)
	log.Println("=================")
	t.root.Delete(v)
	stringify(t.root, 0)
}

func (n *Node) Max() *Node {
	if n == nil {
		return nil
	}

	Max := n
	stack := []*Node{}
	stack = append(stack, n)

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Value > Max.Value {
			Max = curr
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	return Max
}

func (n *Node) Min() *Node {
	if n == nil {
		return n
	}
	min := n
	stack := []*Node{}
	stack = append(stack, n)
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Value < min.Value {
			min = curr
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return min
}

func (n *Node) Delete(v int) *Node {
	if nil == n {
		return n
	}
	if n.Value > v {
		n.Left = n.Left.Delete(v)
	}
	if n.Value < v {
		n.Right = n.Right.Delete(v)
	}
	if n.Value == v {
		if n.Left != nil && n.Right != nil {
			// Delete using successor
			// min := n.Right.Min()
			// n.Right = n.Right.Delete(min.Value)
			// n.Value = min.Value

			//Delete using predessesor
			max := n.Left.Max()
			n.Left = n.Left.Delete(max.Value)
			n.Value = max.Value
		} else if n.Left != nil {
			n = n.Left
		} else if n.Right != nil {
			n = n.Right
		} else {
			n = nil

		}
	}
	return n
}
func (t *Tree) BFS(v int) bool {
	if t.root != nil {
		return t.root.BFS(v)
	}
	return false
}

func (n *Node) BFS(v int) bool {
	if n == nil {
		return false
	}
	queue := []*Node{}
	queue = append(queue, n)

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		if v == curr.Value {
			return true
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return false
}

func (t *Tree) DFSRecursive(v int) bool {
	if t.root != nil {
		return t.root.DFSRecursive(v)
	}
	return false
}

func (n *Node) DFSRecursive(v int) bool {
	var result bool
	log.Printf("comparing %v to %v", n.Value, v)
	if n.Value == v {
		return true
	}
	if n.Left != nil && !result {
		result = n.Left.DFSRecursive(v)
	}
	if n.Right != nil && !result {
		result = n.Right.DFSRecursive(v)
	}
	return result
}
func (t *Tree) DFSIterative(v int) bool {
	if t.root != nil {
		return t.root.DFSIterative(v)
	}
	return false
}
func (n *Node) DFSIterative(v int) bool {
	if n == nil {
		return false
	}
	stack := []*Node{}
	stack = append(stack, n)

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Value == v {
			return true
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

	}
	return false
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Left, level)
		log.Printf(format+"%d\n", n.Value)
		stringify(n.Right, level)
	}
}
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}
	n.Left.InOrderTraverse()
	log.Print(n.Value)
	n.Right.InOrderTraverse()

}

func (n *Node) PreOrderTraverse() {
	if n == nil {
		return
	}
	log.Print(n.Value)
	n.Left.PreOrderTraverse()
	n.Right.PreOrderTraverse()
}

func (n *Node) PostOrderTraverse() {
	if n == nil {
		return
	}
	n.Left.PostOrderTraverse()
	n.Right.PostOrderTraverse()
	log.Print(n.Value)
}

// Item with the min value
func MinDepth() {}

// Item with the max value
func MaxDepth() {}

func testTree() {
	T := initTree()
	log.Print(T.DFSIterative(10))
	stringify(T.root, 0)
}

func initTree() *Tree {
	T := &Tree{}
	T.Insert(5)
	T.Insert(6)

	T.Insert(2)
	T.Insert(3)
	T.Insert(4)
	T.Insert(1)
	T.Insert(9)
	T.Insert(7)
	T.Insert(8)
	T.Insert(0)
	return T
}

//  5
// 2  6
//2 3   7
//1   4   8
