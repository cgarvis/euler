// A perfect number is a number for which the sum of its proper divisors is exactly equal
// to the number. For example, the sum of the proper divisors of 28 would be
// 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
//
// A number n is called deficient if the sum of its proper divisors is less than n and it
// is called abundant if this sum exceeds n.
//
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that
// can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can
// be shown that all integers greater than 28123 can be written as the sum of two abundant
// numbers. However, this upper limit cannot be reduced any further by analysis even though
// it is known that the greatest number that cannot be expressed as the sum of two abundant
// numbers is less than this limit.
//
// Find the sum of all the positive integers which cannot be written as the sum of
// two abundant numbers.

package main

import "fmt"
import "math"

func divisors(n int) []int {
	divisors := make([]int, 0)
	sqrt := int(math.Sqrt(float64(n)))
	divisors = append(divisors, 1)

	for i := 2; i <= sqrt; i ++ {
		if n % i == 0 {
			divisors = append(divisors, i, n/i)
		}
	}

	// handle perfect square
	if sqrt * sqrt == n {
		divisors = divisors[0:len(divisors) -1]
	}

	return divisors
}

func is_abundant(n int) bool {
	divisors := divisors(n)

	sum := 0
	for _, divisor := range divisors {
		sum += divisor
	}

	return sum > n
}

func main() {
	abundants := make([]int, 0)

	for i := 1; i < 28123; i++ {
		if is_abundant(i) {
			abundants = append(abundants, i)
		}
	}

	nums := make([]int, 28123)
	limit := len(abundants)

	for i := 0; i < limit; i++ {
		x := abundants[i]
		for j := i; j < limit; j++ {
			y := abundants[j]
			if x+y < 28123 {
				nums[x+y] = 1
			}
		}
	}

	sum := 0
	for n, v := range nums {
		if v == 0 {
			sum += n
		}
	}

	fmt.Println(sum)
}
