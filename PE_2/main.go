/*
Project Euler Problem #2 - Go Refactor

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

*/

package main

import (
	"fmt"
	//"math/big"
	"time"
)

func main() {
	// Initilize benchmark timer
	start := time.Now()

	var sol int
	//interate over soltuion x times
	for i := 0; i < 1000000; i++ {
		//Call even Fibs function
		sol = evenFibs(4000000)
	}

	// Print solution
	fmt.Printf("The sum of the even numbers in the Fibbonaci sequence is: %v\n", sol)

	// Stop timer and print outcome
	elapsed := time.Since(start)
	fmt.Printf("Solution 25 took %s", elapsed)
}

func evenFibs(x int) int {
	solution := 0
	a := 0
	b := 1
	for b < x {
		a = a + b
		a, b = b, a

		if b%2 == 0 {
			solution += b
		}
	}
	return solution
}
