// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

package main

import "fmt"

const max = 2000000 // 2,000,000

func eratosthenes(max int) []int {
	nums := make([]int, max)

	p := 2 // start with the first prime

	for {
		i := p - 1

		// mark multiples of p as not primes
		for i += p; i < max; i += p {
			nums[i] = -1
		}

		// find first unmarked noumber greater than p
		for i = p; i < max; i++ {
			if nums[i] != -1 {
				p = i + 1
				break
			}
		}


		if i ==  max {
			break
		}
	}

	// filter out all marked numbers
	primes := make([]int, max)
	j := 0
	for i := range nums {
		if nums[i] == 0 {
			primes[j] = i + 1
			j++
		}
	}

	return primes[:j]

}


func main() {
	primes := eratosthenes(max)

	sum := -1 // 1 is not prime in the problem
	for _, prime := range primes {
		sum += prime
	}

	fmt.Println(sum)
}
