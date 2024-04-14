package main

import (
  "fmt"
)

var pln = fmt.Println

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}

func main() {
  pln(fact(12))

  var fib func(n int) int
  fib = func(n int) int {
    if n >= 0 && n < 2 {
      return n
    }
    return fib(n-1) + fib(n-2)
  }
  for i := 0; i < 26; i++ {
    pln(i, fib(i))
  }
}
