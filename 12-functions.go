package main

import (
  "fmt"
)

func plus(a int, b int) int {
  return a + b
}

func plusPlus(a, b, c int) int {
  return a + b + c
}

func main() {

  var pln = fmt.Println

  res := plus(1111, 3333)
  pln("1111 + 3333 =", res)

  res2 := plusPlus(11110, 33330, 88567)
  pln("11110 + 33330 + 88567=", res2)

}
