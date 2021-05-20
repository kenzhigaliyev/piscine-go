// package main

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

package piscine

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {

		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(root.Left.Data)
// 	fmt.Println(root.Data)
// 	fmt.Println(root.Right.Left.Data)
// 	fmt.Println(root.Right.Data)

// }
