// The following iterative sequence is defined for the set of positive integers:
//
// n -> n/2 (n is even)
// n -> 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following sequence:
//
// 13  40  20  10  5  16  8  4  2  1
//
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
// Although it has not been proved yet (Collatz Problem), it is thought that all starting
// numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?
//
// NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

const workers = 3

type Chain struct {
	start, length int
}

func collatz(start int) Chain{
	count := 0
	for i := start; i > 1; count++ {
		if i % 2 == 0 {
			i = i/2
		} else {
			i = 3*i + 1
		}
	}

	return Chain{start,count}
}

func main() {
	work := make(chan int)
	chains := make(chan Chain)

	go func() {
		for i := 2; i < 1e6; i++ {
			work <- i
		}
		close(work)
	}()

	go func() {
		for i := 0; i < workers; i++ {
			for start := range work {
				chains <- collatz(start)
			}
			chains <- Chain{}
		}
	}()

	var largest Chain
	n := workers

	for chain := range chains {
		if chain.start == 0 {
			n--
		} else if chain.length > largest.length {
			largest = chain
		}

		if n == 0 {
			fmt.Println(largest)
			return
		}
	}
}
