// package main
package piscine

// import (
// 	"fmt"
// )

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

// func BTreeLevelCount(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	counter1 := BTreeLevelCount(root.Left)
// 	counter2 := BTreeLevelCount(root.Right)
// 	if counter1 > counter2 {
// 		return counter1 + 1
// 	}
// 	return counter2 + 1
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

func NewLine(root *TreeNode, f func(...interface{}) (int, error), val int) {
	if root == nil {
		return
	}
	if val == 0 {
		f(root.Data)
	}
	if val > 0 {
		NewLine(root.Left, f, val-1)
		NewLine(root.Right, f, val-1)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	levels := BTreeLevelCount(root)
	for i := 0; i < levels; i++ {
		NewLine(root, f, i)
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyByLevel(root, fmt.Println)
// }
