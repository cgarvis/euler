package main

import "fmt"

const MAX = 4e6 // 4,000,000

func main() {
	sum := 0

	for previous, current := 1, 2; current < MAX; previous, current = current, previous+current {
		if current%2 == 0 {
			sum += current
		}
	}

	fmt.Println(sum)
}
