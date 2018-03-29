/***
The median of a dataset of integers is the midpoint value of the dataset for which an equal number of integers are less than and greater than the value. To find the median, you must first sort your dataset of integers in non-decreasing order, then:

    If your dataset contains an odd number of elements, the median is the middle element of the sorted sample. In the sorted dataset , is the median.
    If your dataset contains an even number of elements, the median is the average of the two middle elements of the sorted sample. In the sorted dataset , is the median.

Given an input stream of integers, you must perform the following task for each integer:

    Add the integer to a running list of integers.
    Find the median of the updated list (i.e., for the first element through the element).
    Print the list's updated median on a new line. The printed value must be a double-precision number scaled to decimal place (i.e., format).

Input Format

The first line contains a single integer, , denoting the number of integers in the data stream.
Each line of the subsequent lines contains an integer, , to be added to your list.

Constraints

Output Format

After each new integer is added to the list, print the list's updated median on a new line as a single double-precision number scaled to decimal place (i.e., format).

Sample Input

6
12
4
5
3
8
7

Sample Output

12.0
8.0
5.0
4.5
5.0
6.0

***/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	_, _ = fmt.Scan(&n)
	minh, maxh := &MinHeap{}, &MaxHeap{}
	heap.Init(minh)
	heap.Init(maxh)
	median := float32(0.0)
	for ; n > 0; n-- {
		var num int
		_, _ = fmt.Scan(&num)

		//fmt.Printf("num %d, maxh.Len %d, minh.Len %d\n",num,maxh.Len(),minh.Len())
		if float32(num) <= median {
			heap.Push(maxh, num)
			//fmt.Printf("push to maxh,num %d, maxh %v, median %f\n",num,maxh,median)
		} else if float32(num) > median {
			heap.Push(minh, num)
			//fmt.Printf("push to minh, num %d, minh %v, median %f\n",num,minh,median)
		}
		//fmt.Printf("before balacing,num %d, maxh %v, minh %v, prev median %f\n",num,maxh,minh,median)
		if maxh.Len()-minh.Len() > 1 {
			top := heap.Pop(maxh)
			heap.Push(minh, top)
		} else if maxh.Len()-minh.Len() < (-1) {
			top := heap.Pop(minh)
			heap.Push(maxh, top)
		}

		if maxh.Len() == minh.Len() {
			min, max := *minh, *maxh
			median = float32(max[0]+min[0]) / 2.0
			//fmt.Printf("even, maxh %v, minh %v, median %f\n",max,min,median)
		} else if maxh.Len() > minh.Len() {
			max := *maxh
			median = float32(max[0])
		} else {
			min := *minh
			median = float32(min[0])
		}

		fmt.Printf("%.1f\n", median)
	}
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
