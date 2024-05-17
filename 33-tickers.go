package main

import (
  "fmt"
  "time"
)

var pln = fmt.Println

func main() {
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)

  go func () {
    for {
      select {
      case <-done:
        return
      case t := <-ticker.C:
        pln("Tick at", t)
      }
    }
  }()

  time.Sleep(1600 * time.Millisecond)
  ticker.Stop()
  done <- true
  pln("Ticker stopped")
}
