package main

import (
  "fmt"
)

var pln = fmt.Println

func sum(nums ...int) {
  total := 0

  for _, num := range nums {
    total += num
  }
  pln(nums, total)
  
}

func main() {

  sum(1, 2)
  sum(1, 2, 3, 4, 5, 6, 7, 8)

  s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
  sum(s...)

}
