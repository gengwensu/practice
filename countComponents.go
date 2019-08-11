/* Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes),
write a function to find the number of connected components in an undirected graph.
Example 1:
     0          3
     |          |
     1 --- 2    4

Given n = 5 and edges = [[0, 1], [1, 2], [3, 4]], return 2.
Example 2:
     0           4
     |           |
     1 --- 2 --- 3

Given n = 5 and edges = [[0, 1], [1, 2], [2, 3], [3, 4]], return 1.
 Note:
You can assume that no duplicate edges will appear in edges. Since all edges are undirected,
[0, 1] is the same as [1, 0] and thus will not appear together in edges. */

package main

import (
	"fmt"
)

func main() {
	n := 5
     // edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
     edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Printf("n - %d, connectedComponents - %d\n", n, countComponents(n, edges))
}

var visited []bool
var adjNodes [][]int

func countComponents(N int, edges [][]int) int {
	visited = make([]bool, N)
	adjNodes = make([][]int, N)
	for _, e := range edges { //build adjcentNode array for each node
		t0 := adjNodes[e[0]]
		t0 = append(t0, e[1])
		adjNodes[e[0]] = t0
		t1 := adjNodes[e[1]]
		t1 = append(t1, e[0])
		adjNodes[e[1]] = t1
	}
	count := 0
	for i := 0; i < N; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}
	return count
}

func dfs(node int) {
	if visited[node] {
		return
	}
	visited[node] = true
	for _, n := range adjNodes[node] {
		if !visited[n] {
			dfs(n)
		}
	}
}
