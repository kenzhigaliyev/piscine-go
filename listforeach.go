// package main

// import (
// 	"fmt"
// )

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

package piscine

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}

	val := l.Head

	for val.Next != nil {
		val = val.Next
	}
	val.Next = &NodeL{Data: data}
}

func ListForEach(l *List, f func(*NodeL)) {
	val := l.Head
	for val != nil {
		f(val)
		val = val.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "1")
// 	ListPushBack(link, "2")
// 	ListPushBack(link, "3")
// 	ListPushBack(link, "5")

// 	ListForEach(link, Add2_node)

// 	it := link.Head
// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}
// }
