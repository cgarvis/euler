// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
// there are exactly 6 routes to the bottom right corner.
//
// How many such routes are there through a 20×20 grid?

package main

import "fmt"
import "math/big"

func main() {
	// http://en.wikipedia.org/wiki/Combination
	// http://en.wikipedia.org/wiki/Binomial_coefficient

	n := big.NewInt(40)
	k := big.NewInt(20)


	fmt.Println(choose(n, k))
}

//     n!
// -----------
//  k! (n-k)!
func choose(n, k *big.Int) *big.Int {
	return big.NewInt(0)
}

func factorial(n *big.Int) (result *big.Int) {
    //fmt.Println("n = ", n)
    b := big.NewInt(0)
    c := big.NewInt(1)

    if n.Cmp(b) == -1 {
        result = big.NewInt(1)
    }
    if n.Cmp(b) == 0 {
        result = big.NewInt(1)
    } else {
        // return n * factorial(n - 1);
        fmt.Println("n = ", n)
        result = n.Mul(n, factorial(n.Sub(n, c)))
    }
    return result
}
