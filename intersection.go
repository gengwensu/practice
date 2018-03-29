/***
a={1,2,3},b={2,3,4}
print out {2,3}
***/

package main

import (
	"fmt"
)

func main() {
	a:=[]int{1,2,3}
	b:=[]int{2,3,4}
	fmt.Println(inter(a,b))
}

func inter(a1,a2 []int) []int {
	ans:=[]int{}
	hm:=map[int]int{}
	for _,num:=range a1{
		hm[num]++
	}
	for _,num:=range a2{
		if _,ok:=hm[num];ok{
			ans=append(ans,num)
		}
	}
	return ans
}