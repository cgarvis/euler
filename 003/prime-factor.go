// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"
import "math"

const NUMBER = 600851475143
//const NUMBER = 1000

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
	max := int(math.Sqrt(NUMBER))
	primes := eratosthenes(max)

	for i := len(primes) - 1; i >= 0; i-- {
		factor := primes[i]
		if NUMBER % factor == 0 {
			fmt.Println(factor)
			break
		}
	}
}
