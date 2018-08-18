package main

//refer to https://www.jianshu.com/p/456af5480cee for tree traverse
//先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)
//中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
//后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)
import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) print() {
	fmt.Print(node.Value, "->")
}

func (node *TreeNode) setvalue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node.")
		return
	}
	node.Value = value
}

func print2(node TreeNode) {
	fmt.Print(node.Value)
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func normalMethod_setValtoNil() {
	var root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Right.Right = createNode(2)

	root.print()
	root.setvalue(4)
	fmt.Println()
	print2(root)
	fmt.Println()

	pRoot := &root
	pRoot.print()
	pRoot.Right.Left.setvalue(100)
	pRoot.Right.Left.print()

	var zeroRoot TreeNode
	fmt.Println(zeroRoot)
	zeroRoot.setvalue(100)
	fmt.Println(zeroRoot)
	fmt.Println("=====================")
	var nilRoot *TreeNode
	nilRoot.setvalue(100)
	fmt.Println(nilRoot)
	nilRoot = &zeroRoot
	nilRoot.setvalue(200)
	fmt.Println(nilRoot)

	//	fmt.Println(nilRoot.Value)
}

func normalMethod2() {
	var root TreeNode
	nodes := []TreeNode{{Value: 3}, {}, {6, nil, &root}}
	fmt.Println(nodes)
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.print()
	node.Right.Traverse()

}

type myTreeNode struct {
	mynode *TreeNode
}

func (node *myTreeNode) PostOrder_Traverse() {
	if node == nil || node.mynode == nil {
		return
	}
	left := myTreeNode{node.mynode.Left}
	left.PostOrder_Traverse()
	right := myTreeNode{node.mynode.Right}
	right.PostOrder_Traverse()
	node.mynode.print()
}

func TryBrowse() {
	var root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Right.Right = createNode(2)

	root.print()
	root.setvalue(4)

	fmt.Println("root.Value: \n======", root.Value, "==============\n")
	fmt.Println("           (          )          ")
	fmt.Println("====  ", root.Left.Value, "      ", root.Right.Value, "=====\n")
	fmt.Println("          (  )       (  )     ")
	fmt.Println("=", root.Left.Left, root.Left.Right, " ", root.Right.Left.Value, "", root.Right.Right.Value, "=\n")
	fmt.Println("======Traverse========")
	root.Traverse()
	fmt.Println()
	mynode := myTreeNode{&root}
	fmt.Println("======PostOrder_Traverse========")
	mynode.PostOrder_Traverse()
	fmt.Println()
}

func main() {
	//	normalMethod_setValtoNil()
	TryBrowse()
	//	var node = new(TreeNode)
	//	var newnode = node.createNode(1)
	//	newnode.Left = newnode.createNode(2)
	//	newnode.Right = newnode.createNode(3)
	//	fmt.Println(newnode, newnode.Left, newnode.Right)

	//	normalMethod2()
}
