package main

import (
  "fmt"
)

func vals() (int, string, bool, float32) {
  return 17, "seventeen", true, 17.17
}

func main() {

  var pln = fmt.Println
  a, b, c, d := vals()
  pln(a)
  pln(b)
  pln(c)
  pln(d)
  pln()

  _, _, bool1, _ := vals()
  pln(bool1)

}
