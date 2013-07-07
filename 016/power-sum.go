// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
// What is the sum of the digits of the number 2^1000?

package main

import "fmt"
import "math"
import "strconv"

func main() {
	power := math.Pow(2, 1000)
	power_str := strconv.FormatFloat(power, 'f', 0, 64)

	sum := 0
	for i := 0; i < len(power_str); i++ {
		digit, _ := strconv.Atoi(power_str[i:i+1])
		sum += digit
	}
	fmt.Println(sum)
}
