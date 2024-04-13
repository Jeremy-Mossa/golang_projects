package main

import (
  "fmt"
  "maps"
)

func main() {

  var pln = fmt.Println

  m := make(map[string]int)

  m["prime6th"] = 13
  m["prime7th"] = 17

  pln("map m:", m)

  v1 := m["prime6th"]
  pln("v1:", v1)

  v3 := m["prime8th"]
  pln("v3:", v3)

  pln("len:", len(m))

  delete(m, "prime6th")
  pln("map m:", m)
  
  
  clear(m)
  pln("map m:", m)
  
  _, prs := m["prime7th"]
  pln("prs:", prs)

  n := map[string]float32{"foo": 1.41, "bar": 3.14}
  pln("map n:", n)

  n2 := map[string]float32{"foo": 1.41, "bar": 3.14}
  
  foobar := maps.Equal(n, n2)
  pln(foobar)
  
}
