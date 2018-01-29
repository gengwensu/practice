/*
Merge lists of integer intervals
Within interval, integers are non-decreasing
Within list, intervals are increasing, non-overlapping, non-touching
1st list: [2, 3], [5, 10], [16, 20], [24, 25]
2nd list: [1, 5], [10, 18], [20, 23]
  AND Result: [2, 3], [5, 5], [10, 10], [16, 18], [20, 20]
1st list: [3, 5], [10, 18]
2nd list: [0, 1], [6, 10]
  AND Result: [10, 10]
*/
package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func main() {
	// list1:=[]interval{
	//     {3,5},
	//     {10,18},
	// }
	// list2:=[]interval{
	//     {0,1},
	//     {6,10},
	// }
	list1 := []Interval{{2, 3}, {5, 10}, {16, 20}, {24, 25}}
	list2 := []Interval{{1, 5}, {10, 18}, {20, 23}}
	result := mergeInterval(list1, list2)
	fmt.Printf("result is %v\n", result)
}

func mergeInterval(l1, l2 []Interval) []Interval {
	list := []Interval{}
	i1, i2 := 0, 0
	for i1 < len(l1) && i2 < len(l2) {
		switch {
		case l1[i1].Start < l2[i2].Start:
			list = append(list, l1[i1])
			i1++

		case l1[i1].Start == l2[i2].Start: //merge into one
			tend := l1[i1].End
			if l1[i1].End < l2[i2].End {
				tend = l2[i2].End
			}
			list = append(list, Interval{l1[i1].Start, tend})
			i1++
			i2++

		case l1[i1].Start > l2[i2].Start:
			list = append(list, l2[i2])
			i2++
		}
	}

	if i1 == len(l1) && i2 < len(l2) {
		for ; i2 < len(l2); i2++ {
			list = append(list, l2[i2])
		}
	} else if i2 == len(l2) && i1 < len(l1) {
		for ; i1 < len(l1); i1++ {
			list = append(list, l1[i1])
		}
	}

	ans := []Interval{}
	current := Interval{0, 0}
	for i, intvl := range list {
		if i == 0 {
			current.Start = list[0].Start
			current.End = list[0].End
		} else {
			if intvl.Start <= current.End {
				if intvl.End > current.End { // merge
					current.End = intvl.End
				}
			} else {
				temp := current
				ans = append(ans, temp)
				current.Start = intvl.Start
				current.End = intvl.End
			}
		}
	}
	ans = append(ans, current)

	return ans
}
