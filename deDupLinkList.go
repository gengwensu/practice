/*
deDupLinkList: remove duplicates from an unsorted linked list
input format: 	1st int - number of data points
				2nd to n - data points
				example: 6 1 2 2 3 3 4
				6 data points - 1 2 2 3 3 4
*/
package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func main() {
	var n int
	var head *node

	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var d int
		_, _ = fmt.Scan(&d)

		head = insert(head, d)
	}
	fmt.Println("Input linklist is:")
	display(head)
	head = deDupLinkList(head)
	fmt.Println("Deduped linklist is:")
	display(head)
}

func deDupLinkList(head *node) *node {
	m := map[int]int{} //create an empty map
	var n *node = head
	var prev *node = nil
	for n != nil {
		if _, ok := m[n.data]; ok {
			prev.next = n.next
		} else {
			m[n.data]++
			prev = n
		}
		n = n.next
	}
	return head
}

func insert(head *node, data int) *node {
	p := &node{data, nil}
	if head == nil {
		head = p
	} else if head.next == nil {
		head.next = p
	} else {
		s := head
		for s.next != nil {
			s = s.next
		}
		s.next = p
	}
	return head
}

func display(head *node) {
	s := head
	for s != nil {
		fmt.Printf("%d ", s.data)
		s = s.next
	}
}
