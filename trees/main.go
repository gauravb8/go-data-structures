package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node == nil || (node.Left == nil && node.Right == nil)
}

func calcLeafSum(root *TreeNode, sumType string) int {
	if isLeaf(root) {
		return 0
	}

	if sumType == "left" {
		if root.Left != nil && isLeaf(root.Left) {
			fmt.Println("Adding left node:", root.Left.Val)
			return root.Left.Val + calcLeafSum(root.Right, sumType)
		}
	} else {
		if root.Right != nil && isLeaf(root.Right) {
			fmt.Println("Subtracting right node:", root.Right.Val)
			return calcLeafSum(root.Left, sumType) + root.Right.Val
		}
	}

	return calcLeafSum(root.Left, sumType) + calcLeafSum(root.Right, sumType)
}

func levelOrderTraversal(root *TreeNode) {
	q := Queue{}

	q.Push(root)

	for !q.IsEmpty() {
		node := q.pop()
		q.push(node.left)
		q.push(node.right)
		fmt.Println(node.Val)
	}

}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Println(calcLeafSum(t, "left") - calcLeafSum(t, "right"))

}
