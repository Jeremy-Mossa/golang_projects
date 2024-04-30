package main

import (
  "fmt"
  "time"
)

var pln = fmt.Println

func main() {
  c1 := make(chan string, 1)
    go func() {
      time.Sleep(2 * time.Second)
      c1 <- "result 1"
    }()

  c2 := make(chan string, 1)
    go func() {
      time.Sleep(2 * time.Second)
      c2 <- "result 2"
    }()
}
