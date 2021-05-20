// package main
package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Right, Left, Parent *TreeNode
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

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		return BTreeMin(root.Left)
	}
	return root
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	min := BTreeMin(root)
// 	fmt.Println(min.Data)
// }
