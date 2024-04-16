package main

import (
  "fmt"
  "math"
)

type geometry interface {
  // interfaces are named collections of method signatures
  area() float64
  perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) perim() float64 {
  return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

func measure(g geometry) {
  pln(g)
  pln(g.area())
  pln(g.perim())
}

var pln = fmt.Println

func main() {
  r := rect{width: 33, height: 3.12}
  c := circle{radius: 99.63}

  measure(r)
  measure(c)
}
