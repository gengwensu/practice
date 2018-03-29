/***
a="apple",b="ape"
a="apple",b="aep"
a="laptop",b="lpp" ===> true s2="lpa"===>false
***/
package main

import (
	"fmt"
)

func main() {
	a,b:="laptop","lpa"
	fmt.Println(matching(a,b))
}

func matching(s1,s2 string) bool{
	if len(s1)<len(s2){
		return false
	}
	i,j:=len(s2)-1,len(s1)-1 
	for i>=0 && j>=0 {
		//fmt.Printf("i %d, j %d, s1 %s, s2 %s\n",i,j,s1,s2)
		if s1[j]==s2[i]{
			i--
			s1=s1[:j]
			j--
		} else {
			j--
		}
	}
	if i==-1 {
		return true
	} else if i>-1 && j==-1 {
		return false
	}
	return false
}