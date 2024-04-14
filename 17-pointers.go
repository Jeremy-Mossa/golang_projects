package main

import (
  "fmt"
)

var pln = fmt.Println

func zeroval(ival int){
  ival = 0
}
func zeroptr(iptr *int){
  *iptr = 0
}
func main() {
  i := 1 
  pln("initial:", i)

  zeroval(i)
  pln("zeroval:", i)
  zeroptr(&i)
  pln("zeroptr:", i)

  pln("pointer:", &i)
}
