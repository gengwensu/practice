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

type interval struct {
	start int
	end   int
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
	list1 := []interval{{2, 3}, {5, 10}, {16, 20}, {24, 25}}
	list2 := []interval{{1, 5}, {10, 18}, {20, 23}}
	result := mergeInterval(list1, list2)
	fmt.Printf("result is %v\n", result)
}

func mergeInterval(l1, l2 []interval) []interval {
	result := []interval{}
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			tmp := e1
			if e1.start > e2.end ||
				e1.end < e2.start {
				continue
			}
			if e1.start < e2.start {
				tmp.start = e2.start
			}
			if e1.end > e2.end {
				tmp.end = e2.end
			}
			result = append(result, tmp)
		}
	}
	return result
}
