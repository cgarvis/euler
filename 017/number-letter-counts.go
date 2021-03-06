// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
// how many letters would be used?

package main

import "fmt"
import "flag"

var (
    digits = make([]string, 0)
    teens = make([]string, 0)
    tens = make([]string, 0)
)

func int_to_words(i int) string {
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
		if i % 100 != 0 {
            words += "and" + int_to_words(i%100)
        }
	} else if i < 10000 {
		words += digits[i/1000] + "thousand"
		words += int_to_words(i%1000)
	}

	return words
}

func main() {
	count := 0
	count_to := flag.Int("to", 1000, "number to count to")

	flag.Parse()

    digits = append(digits, "", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine")
    teens = append(teens, "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen")
    tens = append(tens, "", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety")

	for i := 1; i <= *count_to; i++ {
		words := int_to_words(i)
		count += len(words)
		fmt.Println(fmt.Sprintf("%s %d %d", words, len(words), count))
	}

	fmt.Println(count)
}
