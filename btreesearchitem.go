// package main

package piscine

// import (
// 	"fmt"
// )

type TreeNode struct {
	Parent, Left, Right *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	if root.Left != nil {
		root.Left.Parent = root
	}
	if root.Right != nil {
		root.Right.Parent = root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	// if root == nil {
	// 	return root
	// }
	// if root.Data == elem {
	// 	return root
	// }

	// if root.Data > elem {
	// 	return BTreeSearchItem(root.Left, elem)
	// }
	// return BTreeSearchItem(root.Right, elem)

	if root.Parent != nil {
		if root.Parent.Data == elem {
			return root.Parent
		}
	}

	if root.Left != nil {
		if root.Left.Data == elem {
			return root.Left
		}
	}

	if root.Right != nil {
		if root.Right.Data == elem {
			return root.Right
		}
	}
	return root
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	selected := BTreeSearchItem(root, "7")
// 	fmt.Print("Item selected -> ")
// 	if selected != nil {
// 		fmt.Println(selected.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Parent of selected item -> ")
// 	if selected.Parent != nil {
// 		fmt.Println(selected.Parent.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Left child of selected item -> ")
// 	if selected.Left != nil {
// 		fmt.Println(selected.Left.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Right child of selected item -> ")
// 	if selected.Right != nil {
// 		fmt.Println(selected.Right.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}
// }
