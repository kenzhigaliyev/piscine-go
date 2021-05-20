// package main

package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Parent, Left, Right *TreeNode
// 	Data                string
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	if root.Left != nil {
// 		root.Left.Parent = root
// 	}
// 	if root.Right != nil {
// 		root.Right.Parent = root
// 	}
// 	return root
// }

// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	if root.Data == elem {
// 		return root
// 	}

// 	if root.Data > elem {
// 		return BTreeSearchItem(root.Left, elem)
// 	}
// 	return BTreeSearchItem(root.Right, elem)
// }

// func BTreeMin(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	if root.Left != nil {
// 		return BTreeMin(root.Left)
// 	}
// 	return root
// }

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			number := root.Right
			root = nil
			return number
		} else if root.Right == nil {
			number := root.Left
			root = nil
			return number
		}
		number := BTreeMin(root.Right)

		root.Data = number.Data
		root.Right = BTreeDeleteNode(root.Right, number)
	}
	return root
}

// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root.Left != nil {
// 		BTreeApplyInorder(root.Left, f)
// 	}
// 	f(root.Data)
// 	if root.Right != nil {
// 		BTreeApplyInorder(root.Right, f)
// 	}
// }

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	node := BTreeSearchItem(root, "4")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }
