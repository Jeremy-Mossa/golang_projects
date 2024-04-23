package main

import (
  "fmt"
)

var pln = fmt.Println

func antenna(call_sign chan<- string, msg string) {
  call_sign <- msg
}


func transmitter(call_sign <-chan string, commands chan<- string) {
  msg := <-call_sign
  commands <- msg
}

func main() {
  call_sign := make(chan string, 1)
  commands := make(chan string, 1)
  antenna(call_sign, "come in, over")
  transmitter(call_sign, commands)
  pln(<-commands)
}
