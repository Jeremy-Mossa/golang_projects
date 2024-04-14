package main

import (
  "fmt"
  "unicode/utf8"
)

var pln = fmt.Println

func main() {
  const s = "สวัสดี"
  pln("Len:", len(s))
  
  for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i]) 
  }
  pln()

  pln("utf8 rune count:", utf8.RuneCountInString(s))
}
