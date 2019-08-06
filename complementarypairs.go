/**
Given an arr of integers and a target number, return the total number of complemntary pairs
Example:
a=[1,5,7,-1,5] target=6, output: 6
pairs are: {0,1},{0,4},{1,0},{2,3}{3,2},{4,0}
**/
package main

import "fmt"

func main() {
	// input := []int{1, 5, 7, -1, 5}
	// target := 6
	// input := []int{1, 3, 7, -1, 3}
	// target := 6
	input := []int{1, 1, 1, 1}
	target := 2

	fmt.Printf("Number of pairs - %d\n", countComplementaryPairs(input, target))
}

func countComplementaryPairs(arr []int, sum int) int {
	ans := [][]int{}
	hm := map[int][]int{} //key - a[i], value - arr of {i,j...} a[i]=a[j]
	group := []int{}      //arr without dups
	for i, n := range arr {
		if _, ok := hm[n]; !ok {
			group = append(group, n)
			hm[n] = []int{i}
		} else {
			tmp := hm[n]
			tmp = append(tmp, i) //collect index
			hm[n] = tmp
		}
	}

	for _, n := range group {
		compl := sum - n
		if _, ok := hm[compl]; ok { //a hit
			for _, i := range hm[n] { //matching up pairs
				for _, j := range hm[compl] {
					if compl != n || i != j {
						ans = append(ans, []int{i, j})
					}
				}
			}
		}
	}
	fmt.Printf("pairs - %v\n", ans)
	return len(ans)
}
