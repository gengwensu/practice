/* [LeetCode]252 Meeting Rooms
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), determine if a person could attend all meetings.

For example,
Given [[0, 30],[5, 10],[15, 20]],
return false. */
package main

import (
	"fmt"
	"sort"
)

func main() {
	//input := []Interval{{0, 30}, {5, 10}, {15, 20}}
	input := []Interval{{0, 5}, {5, 10}, {15, 20}}
	fmt.Printf("%v\n", canAttendMeetings(input))
}

type Interval struct {
	Start, End int
}
type ByStart []Interval

func (h ByStart) Len() int           { return len(h) }
func (h ByStart) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByStart) Less(i, j int) bool { return h[i].Start < h[j].Start }
func canAttendMeetings(intervals []Interval) bool {
	sort.Sort(ByStart(intervals))
	prev := Interval{-1, -1}
	for _, ivl := range intervals {
		if ivl.Start < prev.End { //overlap
			return false
		}
		prev = ivl
	}
	return true
}
