// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10
// without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers
// from 1 to 20?

package main

import "fmt"

func is_divisible_to(number, x int) bool {
	for i := x; i > 0; i-- {
		if number % i != 0 {
			return false
		}
	}
	return true
}

func divisible_to(x int) int {
	if x < 1 { return 0 }
	if x == 1 { return 1 }

	n := 0
	step := divisible_to(x-1)

	found := false
	for !found {
		n += step
		found = is_divisible_to(n, x)
	}

	return n
}

func main() {
	fmt.Println(divisible_to(20))
}
