/*
Merge lists of integer intervals
Within interval, integers are non-decreasing
Within list, intervals are increasing, non-overlapping, non-touching
1st list: [2, 3], [5, 10], [16, 20], [24, 25]
2nd list: [1, 5], [10, 18], [20, 23]
  AND Result: [{1 23} {24 25}]
1st list: [3, 5], [10, 18]
2nd list: [0, 1], [6, 10]
  AND Result: [{0 1} {3 5} {6 18}]
*/
package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func main() {
	list1 := []Interval{{3, 5}, {10, 18}}
	list2 := []Interval{{0, 1}, {6, 10}}
	list3 := []Interval{{2, 3}, {5, 10}, {16, 20}, {24, 25}}
	list4 := []Interval{{1, 5}, {10, 18}, {20, 23}}
	//result := mergeInterval(list1, list2)
	fmt.Printf("result of merge list1 and list2 is %v\n", mergeInterval(list1, list2))
	fmt.Printf("result of merge list3 and list4 is %v\n", mergeInterval(list3, list4))
}

func mergeInterval(l1, l2 []Interval) []Interval {
	ans := []Interval{}
	current := Interval{0, 0}
	i1, i2 := 0, 0
	for i1 < len(l1) || i2 < len(l2) {
		var work Interval
		if i1 < len(l1) && i2 < len(l2) {
			switch {
			case l1[i1].Start < l2[i2].Start:
				work = l1[i1]
				i1++

			case l1[i1].Start == l2[i2].Start: //merge into one
				tend := l1[i1].End
				if l1[i1].End < l2[i2].End {
					tend = l2[i2].End
				}
				work = Interval{l1[i1].Start, tend}
				i1++
				i2++

			case l1[i1].Start > l2[i2].Start:
				work = l2[i2]
				i2++
			}
		} else if i1 == len(l1) && i2 < len(l2) {
			work = l2[i2]
			i2++
		} else if i2 == len(l2) && i1 < len(l1) {
			work = l1[i1]
			i1++
		}

		if work.Start <= current.End {
			if work.End > current.End { // merge
				current.End = work.End
			}
		} else {
			if current.End > 0 {
				ans = append(ans, current)
			}
			current.Start = work.Start
			current.End = work.End
		}
	}
	ans = append(ans, current)

	return ans
}
