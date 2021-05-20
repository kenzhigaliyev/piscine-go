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

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	replacement := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		replacement.Parent.Left = rplc
	} else {
		replacement.Parent.Right = rplc
	}
	replacement.Parent = node.Parent
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
// 	node := BTreeSearchItem(root, "1")
// 	replacement := &TreeNode{Data: "3"}
// 	root = BTreeTransplant(root, node, replacement)
// 	BTreeApplyInorder(root, fmt.Println)
// }
