package main

import "fmt"
import "strconv"
import "unicode/utf8"

func reverse(s string) string {
	o := make([]rune, utf8.RuneCountInString(s))
	i := len(o)

	for _, c := range s {
		i--
		o[i] = c
	}
	return string(o)
}

func isPalindrome(str string) bool {
	return str == reverse(str)
}

func main() {
	var largest int

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			num := i * j
			if num > largest && isPalindrome(strconv.Itoa(num)) {
				largest = num
			}
		}
	}

	fmt.Println(largest)
}
