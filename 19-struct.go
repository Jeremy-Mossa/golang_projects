package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func newPerson(name string) *person {
  // Person constructor function, $10 words
  p := person{name: name}
  p.age = 42
  return &p
}

var pln = fmt.Println

func main() {
  pln(person{"Plato", 2500})  
  pln(person{name: "Miyamoto Musashi", age: 440})
  pln(person{name: "Mo"})
  pln(&person{name: "Johhny Appleseed", age: 250})
  pln(newPerson("Bobbybobbob"))

  s := person{name: "Sam", age: 51}
  pln(s.name)

  sp := &s
  pln(sp.age)

  sp.age = 17
  pln(sp.age)

  dog := struct {
    name string
    isGood bool
  }{
    "Huan",
    true,
  }
  pln(dog)
}
