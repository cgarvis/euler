// Euler discovered the remarkable quadratic formula:
//
// n² + n + 41
//
// It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
// However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly
// when n = 41, 41² + 41 + 41 is clearly divisible by 41.
//
// The incredible formula  n² - 79n + 1601 was discovered, which produces 80 primes for the
// consecutive values n = 0 to 79. The product of the coefficients,-79 and 1601, is -126479.
//
// Considering quadratics of the form:
//
// n² + an + b, where |a| < 1000 and |b| < 1000
//
// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |4| = 4
//
// Find the product of the coefficients, a and b, for the quadratic expression that produces
// the maximum number of primes for consecutive values of n, starting with n = 0.

package main

import "fmt"
import "math"

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func is_prime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func quadratic(n, a, b int) int {
	return n*n + a*n + b
}

func consecutive_quadratic_primes(a, b int) int {
	var n int

	for n = 0; true; n++ {
		if is_prime(quadratic(n, a, b)) == false {
			break
		}
	}

	return n
}

func main() {
	largest_n := 0
	largest_product := 0

	for a := -1000; a < 1000; a++ {
		for b := -1000; b < 1000; b++ {
			n := consecutive_quadratic_primes(a, b)

			if n > largest_n {
				largest_n = n
				largest_product = a * b
			}
		}
	}

	fmt.Println(largest_product)
}
