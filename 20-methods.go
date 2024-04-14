package main

import (
  "fmt"
)

type rect struct {
  width, height int
}

func (r *rect) area() int {
  return r.width * r.height
}

func (r rect) perim() int {
  return 2*r.width + 2*r.height
}

var pln = fmt.Println

func main() {
  r := rect{width: 9, height: 4}  

  pln("area:", r.area())
  pln("perimeter:", r.perim())

  rp := &r
  pln("area:", rp.area())
  pln("perimeter:", rp.perim())
}
