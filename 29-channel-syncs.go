package main

import (
  "fmt"
  "time"
)

var pln = fmt.Println

func worker(done chan bool) {
  pln("working ...")
  time.Sleep(time.Second)
  pln("done")

  done <- true
}

func main() {
  done := make(chan bool, 1)
  go worker(done)

  <-done
}
