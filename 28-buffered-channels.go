package main

import (
  "fmt"
)

var pln = fmt.Println

func main() {
  fortune := make(chan string)
  luckyNumber := make(chan int, 6)

  go func() { fortune <- "play the lottery with numbers below ... yolo" }()  
  msg := <-fortune

  luckyNumber <- 3
  luckyNumber <- 12
  luckyNumber <- 24
  luckyNumber <- 33
  luckyNumber <- 39
  luckyNumber <- 41

  pln(msg)
  for i := 0; i < 6; i++ {
    pln(<-luckyNumber)
  }
}
