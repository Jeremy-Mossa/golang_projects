package main

import (
  "fmt"
)

func main() {

  var pln = fmt.Println

  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  pln("sum:", sum)

  for i, num := range nums {
    if num == 3 {
      pln("index:", i)
    }
  }

  kvs := map[string]string{"malic acid": "apple", "potassium": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

  for _, v := range kvs {
    pln("value:", v)
  }

  for i, c := range "go" {
    pln(i, c)
  }
}
