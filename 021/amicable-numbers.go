// Let d(n) be defined as the sum of proper divisors of n (numbers less than n
// which divide evenly into n).
//
// If d(a) = b and d(b) = a, where a  b, then a and b are an amicable pair and each
// of a and b are called amicable numbers.
//
// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
// therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
//
// Evaluate the sum of all the amicable numbers under 10000.

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

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func d(n int) int {
	return sum(divisors(n))
}

func main() {
	sum := 0

	for i := 1; i < 10001; i++ {
		n := d(i)
		n2 := d(n)

		if i == n2 && i != n {
			sum += n
		}
	}

	fmt.Println(sum)
}
