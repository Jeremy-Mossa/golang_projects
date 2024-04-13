package main

import (
  "fmt"
)

func main() {

  var pln = fmt.Println

  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  pln("sum:", sum)

  for i, num := range nums {
    if num == 3 {
      pln("index:", i)
    }
  }

}
