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

  for index, runeValue := range s {
    fmt.Printf("%#U starts with %d\n", runeValue, index)
  }

  pln("\nUsing DecodeRuneInString")
  for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts with %d\n", runeValue, i)
    w = width
  
    examineRune(runeValue)
  }
}

func examineRune(r rune) {
  if r == 't' {
    pln("found tee")
  } else if r == 'ส' {
    pln("found soo suua")
  }
}
