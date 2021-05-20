// package main
package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Parent, Right, Left *TreeNode
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

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	counter1 := BTreeLevelCount(root.Left)
	counter2 := BTreeLevelCount(root.Right)
	if counter1 > counter2 {
		return counter1 + 1
	}
	return counter2 + 1
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(BTreeLevelCount(root))
// }
