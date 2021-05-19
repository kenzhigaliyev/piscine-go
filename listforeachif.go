// package main

package piscine

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

func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) < 0
	case string, rune:
		return false
	}
	return false
}

func IsNotNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	case string, rune:
		return true
	}
	return true
}

// func ListPushBack(l *List, data interface{}) {
// 	if l.Head == nil {
// 		l.Head = &NodeL{Data: data}
// 		return
// 	}
// 	val := &NodeL{Data: data}
// 	for val.Next != nil {
// 		val = val.Next
// 	}
// 	val.Next = &NodeL{Data: data}
// }

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	val := l.Head
	for val != nil {
		if cond(val) {
			f(val)
		}
		val = val.Next
	}
}

// func PrintElem(node *NodeL) {
// 	fmt.Println(node.Data)
// }

// func StringToInt(node *NodeL) {
// 	node.Data = 2
// }

// func PrintList(l *List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, "->")
// 		it = it.Next
// 	}
// 	fmt.Print("nil", "\n")
// }

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, "hello")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "there")
// 	ListPushBack(link, 23)
// 	ListPushBack(link, "!")
// 	ListPushBack(link, 54)

// 	PrintList(link)

// 	fmt.Println("--------function applied--------")
// 	ListForEachIf(link, PrintElem, IsPositive_node)

// 	ListForEachIf(link, StringToInt, IsNotNumeric_node)

// 	fmt.Println("--------function applied--------")
// 	PrintList(link)

// 	fmt.Println()
// }
