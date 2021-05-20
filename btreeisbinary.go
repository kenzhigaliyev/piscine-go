// package main

package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
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
// 	return root
// }

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Data > root.Data {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Data < root.Data {
			return false
		}
	}
	return (BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right))
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(BTreeIsBinary(root))
// }
