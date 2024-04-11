package main

import (
  "fmt"
)

func main() {

  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }
  
  for j := 0; j < 3; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("loop")
    break
  }

  for n := range 27 {
    if (n+1)%3 == 0 {
      fmt.Println(n+1)
    }
    continue
  }
}
