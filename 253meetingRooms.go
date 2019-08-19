/* [LeetCode]253 Meeting Rooms II
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

Example 1:

Input: [[0, 30],[5, 10],[15, 20]]
Output: 2

Example 2:

Input: [[7,10],[2,4]]
Output: 1 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []Interval{{0, 30}, {5, 10}, {15, 20}}
	// input := []Interval{{7, 10}, {2, 4}}
	fmt.Printf("%v\n", minMeetingRooms(input))
}

type Interval struct {
	Start, End int
}
type ByStart []Interval

func (h ByStart) Len() int           { return len(h) }
func (h ByStart) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByStart) Less(i, j int) bool { return h[i].Start < h[j].Start }
func minMeetingRooms(intervals []Interval) int {
	sort.Sort(ByStart(intervals))
	hm := map[int]int{}
	for _, ivl := range intervals {
		hm[ivl.Start]++
		hm[ivl.End]--
	}
	n := len(intervals)
	max := 0
	runningSum := 0
	for i := intervals[0].Start; i <= intervals[n-1].End; i++ {
		if room, ok := hm[i]; ok {
			//fmt.Printf("hm[%d] - %d\n", i, room)
			runningSum += room
		}
		//fmt.Printf("runningSum - %d\n", runningSum)
		if runningSum > max {
			max = runningSum
		}
		//fmt.Printf("max - %d\n", max)
	}
	return max
}
