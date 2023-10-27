package main

import "fmt"

func FibonacciIterative(n uint64) uint64 {
	if n <= 1 {
		return n
	}

	var oneBack, twoBack, current uint64 = 1, 0, 0

	for i := uint64(2); i <= n; i++ {
		current = oneBack + twoBack
		twoBack = oneBack
		oneBack = current
	}

	return current
}


func FibonacciRecursive(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}


var memo = make(map[uint64]uint64)

// Memoized FibonacciRecursiveMemo function
func FibonacciRecursiveMemo(n uint64) uint64 {
	// if value is present in the map, return it
	if val, exists := memo[n]; exists {
		return val
	}

	// base cases
	if n <= 1 {
		return n
	}

	// compute and store in map before returning
	memo[n] = FibonacciRecursiveMemo(n-1) + FibonacciRecursiveMemo(n-2)
	return memo[n]
}


func main() {

	fmt.Println(FibonacciIterative(1000000))
	//fmt.Println(FibonacciRecursive(1000000))
	fmt.Println(FibonacciRecursiveMemo(1000000))
}