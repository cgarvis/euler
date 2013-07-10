// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

package main

import "fmt"

func Generate(ch chan <- int) {
  for i := 2; ; i++ {
    ch <- i
  }
}

func Filter(in <- chan int, out chan <- int, prime int) {
  for {
    i := <- in
    if i%prime != 0 {
      out <- i
    }
  }
}

func main() {
  ch := make(chan int)

  go Generate(ch)

  var prime int
  for i := 0; i < 10001; i++ {
    prime = <- ch
    ch1 := make(chan int)
    go Filter(ch, ch1, prime)
    ch = ch1
  }

  fmt.Println(prime)
}
