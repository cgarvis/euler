// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
// how many letters would be used?

package main

import "fmt"


func int_to_words(i int) string {
	digits := make([]string, 0)
	digits = append(digits, "", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine")
	teens := make([]string, 0)
	teens = append(teens, "ten", "eleven", "twelve", "thirteen", "forteen", "fifthteen", "sixteen", "seventeen", "eightteen", "nineteen")
	tens := make([]string, 0)
	tens = append(tens, "", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety")

	words := ""

	if i < 10 {
		words += digits[i]
	} else if i < 20 {
		words += teens[i%10]
	} else if i < 100 {
		words += tens[i/10]
		words += int_to_words(i%10)
	} else if i < 1000 {
		words += digits[i/100] + "hundred"
		words += int_to_words(i%100)
	} else if i < 10000 {
		words += digits[i/1000] + "thousand"
		words += int_to_words(i%1000)
	}

	return words
}

func main() {
	count := 0

	for i := 1; i < 1001; i++ {
		words := int_to_words(i)
		count += len(words)
	}

	fmt.Println(count)
}
