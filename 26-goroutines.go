package main

import (
  "fmt"
  "time"
)

var pln = fmt.Println

func f(from string) {
  for i := 0; i < 3; i++ {
    pln(from, ":", i)
  }
}


func main() {
  f("direct")

  go f("goroutine")

  go func (msg string) {
    pln(msg)
  }("going")

  time.Sleep(time.Second)
  pln("done")
}
