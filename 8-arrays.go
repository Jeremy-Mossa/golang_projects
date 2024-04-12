package main

import (
  "fmt"
)

func main() {

  var a [17]int
  fmt.Println("emp:", a)

  a[16] = 100
  fmt.Println("set:", a)
  fmt.Println("get:", a[16])

  fmt.Println("len:", len(a))

  b := [6]int{1,2,3,4,5,6}
  fmt.Println("dcl:", b)

  var twoD [7][8]int
  for i:= 0; i < 7; i++ {
    for j := 0; j < 8; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD)
}
