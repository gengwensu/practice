/*** given an array of number of tickets wanted for each person i, []int{t1,t2,...} and
the position p of a person, calculate the ticket waiting time for this person.
Constraints:
1. each transaction takes 1 second time
2. each person can only purchase 1 ticket at a time
3. if a person wants more tickets, the person will need to wait at the end of line
***/
package main

import (
	"fmt"
)

func main() {
	tickets:=[]int{6,2,3,5,7}
	position:=3
	fmt.Printf("tickets %v, position %d, waiting time %d\n",tickets,position,waitingTime(tickets,position))
	ticket1:=[]int{1,1,1,1}
	position1:=1
	fmt.Printf("tickets %v, position %d, waiting time %d\n",ticket1,position1,waitingTime(ticket1,position1))
	ticket2:=[]int{6,2,7,5,3}
	position2:=5
	fmt.Printf("tickets %v, position %d, waiting time %d\n",ticket2,position2,waitingTime(ticket2,position2))
}

func waitingTime(tickets []int, p int) int {
	work:=make([]int,len(tickets))
	copy(work,tickets)
	ans:=0
	sum:=0
	for i:=0;i<p;i++ { // finish the first round and it takes p seconds
		work[i]--
	}
	target:=work[p-1] // number of tickets desired for p
	for _,n:=range work {
		if n<target { // time saved for persons leave early
			sum+=target-n
		}
	}
	ans=len(tickets)*target+p-sum
	return ans
}

