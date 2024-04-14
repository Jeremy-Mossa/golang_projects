package main

import (
  "fmt"
)

var pln = fmt.Println

func intSeq() func() int {
  i := 0
   return func() int{
      i++
      return i
   }
}

func main() {
  nextInt := intSeq()    

  pln(nextInt())
  pln(nextInt())
  pln(nextInt())
  pln(nextInt())

  newInt := intSeq()
  pln(newInt())
}
