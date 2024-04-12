package main

import (
  "fmt"
  "slices"
)

func main() {

  var s []string
  fmt.Println("uninitialized:", s, s==nil, len(s)==0)

  s = make([]string, 3)
  fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

  s[0] = "a" 
  s[1] = "b"
  s[2] = "c"

  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

  fmt.Println("len:", len(s))

  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  l := s[1:5]
  fmt.Println("sl1:", l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[3:]
  fmt.Println("sl3:", l)

  t := []string{"x", "y", "z"}
  t2 := []string{"x", "y", "z"}
  fmt.Println(t); fmt.Println(t2)

  if slices.Equal(t, t2) {
    fmt.Println("t == t2")
  }
  
  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("twoD", twoD)
}
