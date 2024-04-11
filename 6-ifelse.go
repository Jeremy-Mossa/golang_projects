package main

import (
  "fmt"
)

func main() {

  for n := range 10 {
    if (n+1)%2 == 0 {
      fmt.Println(n+1, "is even")
    } else {
      fmt.Println(n+1, "is odd")
    }
  }
}
