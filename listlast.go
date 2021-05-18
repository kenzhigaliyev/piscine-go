// package main

package piscine

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

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	position := l.Head
	for position != nil {
		if position.Next == nil {
			return position.Data
		}
		position = position.Next
	}
	return position.Data
}

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	ListPushBack(link, "three")
// 	ListPushBack(link, 3)
// 	ListPushBack(link, "1")

// 	fmt.Println(ListLast(link))
// 	fmt.Println(ListLast(link2))
// }
