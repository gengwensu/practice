/*  A linked list is given such that each node contains an additional random pointer 
which could point to any node in the list or null. Return a deep copy of the list.

type ListNode {
	Val int
	Next *ListNode
	Random *ListNode
}

*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
	Random *ListNode
}

func main() {
	input := []*ListNode{}
	for i:=0;i<5;i++ {
		input=append(input,&ListNode{i+1,nil,nil})
	}
	root:=input[0]
	ptr:=root
	for i:=1;i<5;i++{
		ptr.Next=input[i]
		ptr=ptr.Next
	}
	ptr=root
	for i,_:=range input {
		if i%2==1 {
			ptr.Random=input[i] // point to self
		} else {
			ptr.Random=input[4-i]
		}
		ptr=ptr.Next
	}
	ans:=deepCopy(root)
	fmt.Printf("root->")
	for ptr=root;ptr!=nil;ptr=ptr.Next {
		fmt.Printf("%v->", ptr)
	}
	fmt.Println()
	fmt.Printf("ans->")
	for ptr=ans;ptr!=nil;ptr=ptr.Next {
		fmt.Printf("%v->", ptr)
	}
	fmt.Println()
}

func deepCopy(root *ListNode) *ListNode {
	if root==nil {
		return nil
	}
	hm:=map[*ListNode]*ListNode{}
	var result,prev *ListNode
	result=&ListNode{root.Val,nil,nil}
	prev=result
	hm[root]=result
	for ptr:=root.Next;ptr!=nil;ptr=ptr.Next {
		prev.Next=&ListNode{ptr.Val,nil,nil}
		hm[ptr]=prev.Next
		prev=prev.Next
	}

	for ptr,rptr:=root,result;ptr!=nil && rptr!= nil;ptr,rptr=ptr.Next,rptr.Next{
		rptr.Random=hm[ptr.Random]
	}
	return result
}


