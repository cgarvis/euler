// A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,
//		a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"
import "math"


func is_pythagorean(a, b, c int) bool {
	aabb := math.Pow(float64(a),2) + math.Pow(float64(b),2)
	cc := math.Pow(float64(c),2)

	if aabb == cc {
		return true
	}
	return false
}

func main() {
	for a := 1; a < 997; a++ {
		for b := a + 1; b < 998; b++ {
			c := 1000 - a - b
			if is_pythagorean(a,b,c) {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
