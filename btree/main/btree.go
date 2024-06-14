package main

//노드는 최대 M개 부터 M/2개 까지의 자식을 가질 수 있습니다.
//노드에는 최대 M−1개 부터 [M/2]−1개의 키가 포함될 수 있습니다.
//노드의 키가 x개라면 자식의 수는 x+1개 입니다.
//최소차수는 자식수의 하한값을 의미하며, 최소차수가 t라면 M=2t−1을 만족합니다. (최소차수 t가 2라면 3차 B트리이며, key의 하한은 1개입니다.)

type BTreeNode struct {
	//key 갯수, 가질 수 있는 child 갯수(key - 1)
	keys   []int
	parent *BTreeNode
	child  []*BTreeNode
	isLeaf bool
}

func newNode() *BTreeNode {
	return &BTreeNode{}
}

func NewRootNode(value int) *BTreeNode {
	return &BTreeNode{
		keys:   make([]int, value),
		parent: nil,
		child:  nil,
		isLeaf: false}
}

func (node *BTreeNode) add(value int) {
	//TODO: 노드에 value 쳐박기 구현
}

type BTree struct {
	root *BTreeNode
}

func NewBTree() *BTree {
	return &BTree{}
}

func (node *BTreeNode) KeySize() int {
	return len(node.keys)
}

func (node *BTree) Add(element int) {
	if node.root == nil {
		node.root = newNode()
		return
	}
	// root node가 이미 존재하는 경우, add 함수를 호출하여 element 추가
	node.root.add(element)
}
