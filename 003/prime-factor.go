// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"
import "math"

const NUMBER = 600851475143

func eratosthenes(max int, c chan int) {
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

		c <- p

		if i ==  max {
			break
		}
	}

	close(c)
}

func main() {
	max := int(math.Sqrt(NUMBER))

	c := make(chan int)

	go eratosthenes(max, c)

	var factor int
	for prime := range c {
		if NUMBER % prime == 0 {
			factor = prime
		}
	}

	fmt.Println(factor)
}
