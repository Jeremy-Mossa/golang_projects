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

  for n := range 10 {
    if (n+1)%2 == 0 {
      fmt.Println(n+1, "is even")
    } else if n%3 == 0{
        fmt.Println(n, "is divisible by 3 and therefore odd")
    } else {
      fmt.Println(n+1, "is odd")
    }
  }

  for n := range 100 {
    if (n+1)%7 == 0 {
      fmt.Println(n+1, "is divisible by 7")
    }
  }
}
