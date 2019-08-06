/**
O(log(n)) Algo. to calculate Fibonacci sequences to 6 digits
**/
package main

import "fmt"

const mod = 1000000 // to 6 digits
func main() {

	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("|   N   \t|\tfib(N)\t|\n")
	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("| 50 \t\t|\t%7d\t|\n", fib(50))
	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("| 100 \t\t|\t%7d\t|\n", fib(100))
	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("| 1000 \t\t|\t%7d\t|\n", fib(1000))
	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("| 10000 \t|\t%7d\t|\n", fib(10000))
	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("| 100000 \t|\t%7d\t|\n", fib(100000))
	fmt.Printf("|---------------|---------------|\n")
	fmt.Printf("| 1000000 \t|\t%7d\t|\n", fib(1000000))
	fmt.Printf("|---------------|---------------|\n")
}

func fib(n int) int {
	F := [][]int{{1, 1}, {1, 0}}
	if n == 0 {
		return 0
	}
	power(F, n-1)
	return F[0][0]
}

func power(F [][]int, n int) {
	if n == 0 || n == 1 {
		return
	}
	M := [][]int{{1, 1}, {1, 0}}

	power(F, n/2)
	multiply(F, F)
	if n%2 != 0 {
		multiply(F, M)
	}
}

func multiply(F, M [][]int) {
	x := F[0][0]*M[0][0]%mod + F[0][1]*M[1][0]%mod
	y := F[0][0]*M[0][1]%mod + F[0][1]*M[1][1]%mod
	z := F[1][0]*M[0][0]%mod + F[1][1]*M[1][0]%mod
	w := F[1][0]*M[0][1]%mod + F[1][1]*M[1][1]%mod

	F[0][0] = x % mod
	F[0][1] = y % mod
	F[1][0] = z % mod
	F[1][1] = w % mod
}
