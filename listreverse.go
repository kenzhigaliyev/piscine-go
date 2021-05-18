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

// func ListPushBack(l *List, data interface{}) {
// 	if l.Head == nil {
// 		l.Head = &NodeL{Data: data}
// 		return
// 	}
// 	val := l.Head
// 	for val.Next != nil {
// 		val = val.Next
// 	}
// 	val.Next = &NodeL{Data: data}
// }

package piscine

func ListPushFront(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = val
		return
	}
	val.Next = l.Head
	l.Head = val
}

func ListReverse(l *List) {
	reverseList := &List{}
	val := l.Head
	for val != nil {
		ListPushFront(reverseList, val.Data)
		val = val.Next
	}
	l.Head = reverseList.Head
	l.Tail = reverseList.Tail
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, 2)
// 	ListPushBack(link, 3)
// 	ListPushBack(link, 4)

// 	ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }
