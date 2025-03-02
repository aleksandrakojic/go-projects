package main

import (
	"fmt"
	"time"
)

// A LinkedList is a collection of linked nodes.
type LinkedList struct {
	next *LinkedList
	data string
}

// Next returns the next node in the linked list.
func (l *LinkedList) Next() *LinkedList { return l.next }

// Iterate iterates over a LinkedList or a channel of strings.
// It takes a function that takes the type of the argument and returns a type of Result.
func Iterate[M LinkedList | chan string, R string | *LinkedList](arg M, iter func(M) R) R {
    return iter(arg)
}

// cIter is a function that takes a channel of strings and returns a string.
// It waits for a message on the channel or times out after 1 second.
func cIter(c chan string) string {
    select {
    case msg := <-c:
        return msg
    case <-time.After(time.Second):
        return "nothing"
    }
}

// lIter is a function that takes a LinkedList and returns the next node in the list.
func lIter(l LinkedList) *LinkedList {
	if l.next == nil {
		return nil
	}
	return l.next
}

func main() {
	// Create a channel of strings.
	c := make(chan string, 5)
	c <- "ok"
	c <- "ok2"

	// Iterate over the channel.
	go func() {
		for msg := range Iterate(c, cIter) {
			fmt.Println(msg)
		}
	}()

	// Create a LinkedList.
	n1 := &LinkedList{data: "n1"}
	n2 := &LinkedList{data: "n2"}
	n3 := &LinkedList{data: "n3"}
	n1.next = n2
	n2.next = n3

	// Iterate over the LinkedList.
	for node := Iterate(*n1, lIter); node != nil; node = Iterate(*node, lIter) {
		fmt.Println(node.data)
	}
}
