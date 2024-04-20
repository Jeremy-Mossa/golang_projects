package main

import (
  "fmt"
)

var pln = fmt.Println

func main() {
  fortune := make(chan string)
  luckyNumber := make(chan int)

  go func() { fortune <- "play the lottery with int below sent from a goroutine ... yolo" }()  
  go func() { luckyNumber <- 24 }()  

  msg := <-fortune
  jackpot := <-luckyNumber
  pln(msg)
  pln(jackpot)
}
